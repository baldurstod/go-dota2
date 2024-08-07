package dota2

import (
	"errors"
)

type Hero struct {
	template *HeroTemplate
	items    map[string]*Item
}

func newHero(template *HeroTemplate) *Hero {
	return &Hero{
		template: template,
		items:    make(map[string]*Item),
	}
}

func (h *Hero) GetEntity() string {
	return h.template.entity
}

func (h *Hero) EquipItem(index string, replaceExisting bool) (*Item, error) {
	var template *ItemTemplate
	var err error
	if template, err = GetItemTemplate(index); err != nil {
		return nil, err
	}
	return h.equipItem(template, replaceExisting)
}

func (h *Hero) equipItem(template *ItemTemplate, replaceExisting bool) (*Item, error) {
	item := newItem(template)

	if !item.IsUsedByHero(h.template.entity) {
		return nil, errors.New("item is not equipable by this hero")
	}

	if replaceExisting {
		h.items[item.GetItemSlot()] = item
	} else {
		if _, ok := h.items[item.GetItemSlot()]; !ok {
			h.items[item.GetItemSlot()] = item
		} else {
			return nil, errors.New("slot already occupied")
		}
	}

	return item, nil
}

func (h *Hero) EquipBundle(index string, replaceExisting bool) ([]*Item, error) {
	var template *ItemTemplate
	var err error
	if template, err = GetItemTemplate(index); err != nil {
		return nil, err
	}

	return h.equipBundle(template, replaceExisting)
}

func (h *Hero) equipBundle(template *ItemTemplate, replaceExisting bool) ([]*Item, error) {
	var items []*Item
	if len(template.Bundle) == 0 {
		return nil, errors.New("not a bundle")
	}

	var t *ItemTemplate
	var err error
	var item *Item
	for name := range template.Bundle {
		if t, err = GetItemTemplateByName(name); err != nil {
			return nil, err
		}

		if item, err = h.equipItem(t, replaceExisting); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// Get hero model depending on the equipped items
func (h *Hero) GetModel() string {
	model := h.template.model

	items := h.GetItems()
	for _, item := range items {
		for _, modifier := range item.GetAssetModifiers() {
			if modifier.Type == MODIFIER_ENTITY_MODEL && modifier.Asset == h.template.entity {
				model = modifier.Modifier
			}
		}
	}
	/*
		"asset_modifier"
		{
			"type"		"entity_model"
			"asset"		"npc_dota_hero_crystal_maiden"
			"modifier"		"models/heroes/crystal_maiden_persona/crystal_maiden_persona.vmdl"
		}
	*/

	return model
}

func (h *Hero) GetSkin() int {
	for _, item := range h.items {
		for _, modifier := range item.GetAssetModifiers() {
			if modifier.Type == MODIFIER_MODEL_SKIN {
				return modifier.Skin
			}
		}
	}

	return 0
}

func (h *Hero) GetModelScale() float32 {
	for _, item := range h.items {
		for _, modifier := range item.GetAssetModifiers() {
			if modifier.Type == MODIFIER_HERO_SCALE {
				return modifier.ModelScale
			}
		}
	}

	return 1.0
}

func (h *Hero) GetItems() []*Item {
	//itemsPerPersona := map[int][]*Item{0: make([]*Item, 0)}
	persona := 0
	if item, ok := h.items["persona_selector"]; ok {
		if id := item.GetPersonaId(); id >= 0 {
			persona = id
		}
	}

	var exist bool
	var slot ItemSlot
	slots := make(map[string]*Item)

	ret := make([]*Item, 0, 5)
	items, exist := itemsPerHero[h.template.entity]
	if !exist {
		return ret
	}

	// First, we add equipped items
	for slot, item := range h.items {
		if slot != "bundle" {
			ret = append(ret, item)
			slots[item.template.ItemSlot] = item
		}
	}

	// Second, we add bundle items, if any
	if item, ok := h.items["bundle"]; ok {
		var itemTemplate *ItemTemplate
		var err error
		for name := range item.template.Bundle {

			if itemTemplate, err = GetItemTemplateByName(name); err != nil {
				continue
			}

			if _, exist = h.items[itemTemplate.ItemSlot]; exist {
				continue
			}

			if slot, exist = h.template.itemSlots[itemTemplate.ItemSlot]; !exist {
				continue
			}

			if !slot.IsPersonaSlot(persona) {
				continue
			}

			item := newItem(itemTemplate)
			ret = append(ret, item)
			slots[itemTemplate.ItemSlot] = item
		}

	}

	for _, itemTemplate := range items {
		if !itemTemplate.BaseItem {
			continue
		}

		if _, exist = slots[itemTemplate.ItemSlot]; exist {
			continue
		}

		if slot, exist = h.template.itemSlots[itemTemplate.ItemSlot]; !exist {
			continue
		}

		if !slot.IsPersonaSlot(persona) {
			continue
		}

		ret = append(ret, newItem(itemTemplate))

	}

	return ret
}

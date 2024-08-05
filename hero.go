package dota2

import "errors"

type Hero struct {
	template *HeroTemplate
	items    map[string]*item
}

func newHero(template *HeroTemplate) *Hero {
	return &Hero{
		template: template,
		items:    make(map[string]*item),
	}
}

func (h *Hero) GetEntity() string {
	return h.template.entity
}

func (h *Hero) EquipItem(index string) error {
	var item *item
	var err error
	if item, err = CreateItem(index); err != nil {
		return err
	}

	if !item.IsUsedByHero(h.template.entity) {
		return errors.New("item is not equipable by this hero")
	}

	h.items[item.GetItemSlot()] = item

	return nil
}

// Get hero model for the selected persona. base hero = 0
func (h *Hero) GetModel() string {
	model := h.template.model

	for _, item := range h.items {
		for _, modifier := range item.GetAssetModifiers(0) {
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

func (h *Hero) GetItems() []*item {
	//itemsPerPersona := map[int][]*item{0: make([]*item, 0)}
	persona := 0
	if item, ok := h.items["persona_selector"]; ok {
		if id := item.GetPersonaId(); id >= 0 {
			persona = id
		}
	}

	var exist bool

	ret := make([]*item, 0, 5)
	items, exist := itemsPerHero[h.template.entity]
	if !exist {
		return ret
	}

	for _, item := range h.items {
		ret = append(ret, item)
	}

	var slot ItemSlot
	for _, itemTemplate := range items {
		if !itemTemplate.BaseItem {
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

		ret = append(ret, newItem(itemTemplate))

	}

	return ret
}

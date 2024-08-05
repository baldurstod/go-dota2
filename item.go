package dota2

type Item struct {
	template *ItemTemplate
	style    int
}

func newItem(template *ItemTemplate) *Item {
	return &Item{
		template: template,
	}
}

func (item *Item) GetIndex() string {
	return item.template.Index
}

func (item *Item) GetName() string {
	return item.template.Name
}

func (item *Item) GetItemSlot() string {
	return item.template.ItemSlot
}

func (item *Item) GetModelPlayer() string {
	return item.template.ModelPlayer
}

func (item *Item) IsUsedByHero(hero string) bool {
	return item.template.IsUsedByHero(hero)
}

func (item *Item) GetAssetModifiers(style int) []*AssetModifier {
	return item.template.GetAssetModifiers(style)
}

func (item *Item) GetPersonaId() int {
	return item.template.GetPersonaId()
}

/*

func (h *Hero) GetEntity() string {
	return h.template.entity
}

func (h *Hero) EquipItem(index string) error {
	var item *Item
	var err error
	if item, err = GetItem(index); err != nil {
		return err
	}

	if !item.IsUsedByHero(h.template.entity) {
		return errors.New("item is not equipable by this hero")
	}

	h.items[item.ItemSlot] = item

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

	return model
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

	ret := make([]*Item, 0, 5)
	items, exist := itemsPerHero[h.template.entity]
	if !exist {
		return ret
	}

	for _, item := range h.items {
		ret = append(ret, item)
	}

	var slot ItemSlot
	for _, item := range items {
		if !item.BaseItem {
			continue
		}

		if _, exist = h.items[item.ItemSlot]; exist {
			continue
		}

		if slot, exist = h.template.itemSlots[item.ItemSlot]; !exist {
			continue
		}

		if !slot.IsPersonaSlot(persona) {
			continue
		}

		ret = append(ret, item)

	}

	return ret
}
*/

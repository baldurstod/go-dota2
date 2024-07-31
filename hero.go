package dota2

type Hero struct {
	template *HeroTemplate
	Items    map[string]*Item
}

func newHero(template *HeroTemplate) *Hero {
	return &Hero{
		template: template,
		Items:    make(map[string]*Item),
	}
}

func (h *Hero) GetEntity() string {
	return h.template.entity
}

// Get hero model for the selected persona. base hero = 0
func (h *Hero) GetModel() string {
	model := h.template.model
	for _, item := range h.GetItems() {
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

func (h *Hero) GetItems() []*Item {
	var exist bool

	ret := make([]*Item, 0, 5)
	//items, exist := itemsPerHero[h.template.entity]
	if !exist {
		return ret
	}
	/*
		var slot ItemSlot
		for _, item := range items {
			if slot, exist = h.ItemSlots[item.ItemSlot]; !exist {
				continue
			}

			if item.BaseItem {
				ret = append(ret, item)
			}
		}
	*/

	return ret
}

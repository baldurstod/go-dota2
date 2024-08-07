package dota2

import (
	"errors"
	"strconv"
	"strings"

	"github.com/baldurstod/vdf"
)

type ItemTemplate struct {
	Index          string          `json:"index,omitempty"`
	Name           string          `json:"name,omitempty"`
	ItemName       string          `json:"item_name,omitempty"`
	ItemClass      string          `json:"item_class,omitempty"`
	ItemTypeName   string          `json:"item_type_name,omitempty"`
	ItemSlot       string          `json:"item_slot,omitempty"`
	ModelPlayer    string          `json:"model_player,omitempty"`
	ImageInventory string          `json:"image_inventory,omitempty"`
	Prefab         string          `json:"prefab,omitempty"`
	BaseItem       bool            `json:"base_item,omitempty"`
	UsedByHeroes   map[string]bool `json:"used_by_heroes,omitempty"`
	Bundle         map[string]bool `json:"bundle,omitempty"`
	Visuals        *Visuals        `json:"visuals,omitempty"`
}

func newItemTemplate(index string) *ItemTemplate {
	return &ItemTemplate{
		Index:        index,
		UsedByHeroes: make(map[string]bool),
		Bundle:       make(map[string]bool),
	}
}

func (i *ItemTemplate) CreateItem() *Item {
	return &Item{
		template: i,
	}
}

func (i *ItemTemplate) initFromData(data *vdf.KeyValue) error {
	var err error
	var prefab *ItemTemplate

	if index, err := data.GetString("prefab"); err == nil {
		prefab, err = GetPrefab(index)
		if err != nil {
			return errors.New("unknonw prefab " + index)
		}

	}

	if i.Name, err = data.GetString("name"); err != nil {
		if prefab != nil {
			i.Name = prefab.Name
		}
	}

	i.ModelPlayer, _ = data.GetString("model_player")
	i.ImageInventory, _ = data.GetString("image_inventory")

	if i.ItemClass, err = data.GetString("item_class"); err != nil {
		if prefab != nil {
			i.ItemClass = prefab.ItemClass
		}
	}

	if i.ItemName, err = data.GetString("item_name"); err != nil {
		if prefab != nil {
			i.ItemName = prefab.ItemName
		}
	}

	if i.ItemTypeName, err = data.GetString("item_type_name"); err != nil {
		if prefab != nil {
			i.ItemTypeName = prefab.ItemTypeName
		}
	}

	if i.ItemSlot, err = data.GetString("item_slot"); err != nil {
		if prefab != nil {
			i.ItemSlot = prefab.ItemSlot
		}
	}

	if i.BaseItem, err = data.GetBool("baseitem"); err != nil {
		if prefab != nil {
			i.BaseItem = prefab.BaseItem
		}
	}

	if usedByHeroes, err := data.Get("used_by_heroes"); err == nil {
		for _, hero := range usedByHeroes.GetChilds() {
			if b, _ := hero.ToBool(); b {
				i.UsedByHeroes[hero.Key] = true
			}
		}
	}

	if bundle, err := data.Get("bundle"); err == nil {
		i.ItemSlot = "bundle"
		for _, name := range bundle.GetChilds() {
			if b, _ := name.ToBool(); b {
				i.Bundle[name.Key] = true
			}
		}
	}

	if visuals, err := data.Get("visuals"); err == nil {
		i.Visuals = NewVisuals()
		err = i.Visuals.initFromData(visuals)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *ItemTemplate) GetAssetModifiers(style int) []*AssetModifier {
	modifiers := make([]*AssetModifier, 0)

	if i.Visuals != nil {
		for _, modifier := range i.Visuals.AssetModifiers {
			if modifier.Style == -1 || modifier.Style == style {
				modifiers = append(modifiers, modifier)

			}
		}
	}

	return modifiers
}

/*
func (i *ItemTemplate) MarshalJSON() ([]byte, error) {
	ret := make(map[string]interface{})

	ret["index"] = i.Index
	ret["name"] = i.Name
	ret["item_name"] = i.ItemName
	ret["item_class"] = i.ItemClass
	ret["item_type_name"] = i.ItemTypeName
	ret["item_slot"] = i.ItemSlot
	if i.BaseItem {
		ret["base_item"] = i.BaseItem
	}

	if i.ModelPlayer != "" {
		ret["model_player"] = i.ModelPlayer
	}

	if len(i.UsedByHeroes) > 0 {
		usedByHeroes := make(map[string]interface{})

		for hero := range i.UsedByHeroes {
			usedByHeroes[hero] = true
		}

		ret["used_by_heroes"] = usedByHeroes
	}

	if i.Visuals != nil {
		ret["visuals"] = i.Visuals
	}

	return json.Marshal(ret)
}*/

func (i *ItemTemplate) IsUsedByHero(hero string) bool {
	_, ok := i.UsedByHeroes[hero]
	return ok
}

func (i *ItemTemplate) IsPersonaItem(id int) bool {
	return strings.Contains(i.ItemSlot, "persona_"+strconv.Itoa(id))
}

// Return the persona id for items having a "persona_selector" slot
func (i *ItemTemplate) GetPersonaId() int {
	id := -1

	for _, modifier := range i.GetAssetModifiers(0) {
		if modifier.Type == MODIFIER_PERSONA {
			id = modifier.Persona
		}
	}
	return id
}

func (i *ItemTemplate) GetSkin(style int) int {
	skin := 0

	if i.Visuals != nil {
		skin = i.Visuals.getSkin(style)
	}
	/*

		for _, modifier := range i.GetAssetModifiers(0) {
			if modifier.Type == MODIFIER_PERSONA {
				skin = modifier.Persona
			}
		}
	*/
	return skin
}

/*
"prefab"		"ancient"
"prefab"		"announcer"
"prefab"		"blink_effect"
"prefab"		"bundle"
"prefab"		"courier"
"prefab"		"courier_effect"
"prefab"		"courier_wearable"
"prefab"		"cursor_pack"
"prefab"		"death_effect"
"prefab"		"default_item"
"prefab"		"direcreeps"
"prefab"		"diresiegecreeps"
"prefab"		"diretowers"
"prefab"		"dynamic_recipe"
"prefab"		"emblem"
"prefab"		"emoticon_tool"
"prefab"		"head_effect"
"prefab"		"hud_skin"
"prefab"		"key"
"prefab"		"league"
"prefab"		"loading_screen"
"prefab"		"map_effect"
"prefab"		"misc"
"prefab"		"music"
"prefab"		"pennant"
"prefab"		"player_card"
"prefab"		"radiantcreeps"
"prefab"		"radiantsiegecreeps"
"prefab"		"radianttowers"
"prefab"		"retired_treasure_chest"
"prefab"		"roshan"
"prefab"		"showcase_decoration"
"prefab"		"socket_gem"
"prefab"		"sticker"
"prefab"		"sticker_capsule"
"prefab"		"streak_effect"
"prefab"		"summons"
"prefab"		"taunt"
"prefab"		"teleport_effect"
"prefab"		"terrain"
"prefab"		"tool"
"prefab"		"tormentor"
"prefab"		"treasure_chest"
"prefab"		"versus_screen"
"prefab"		"ward"
"prefab"		"wearable"
"prefab"		"weather"
*/

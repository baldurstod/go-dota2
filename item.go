package dota2

import (
	"encoding/json"
	"errors"

	"github.com/baldurstod/vdf"
)

type Item struct {
	Index        string
	Name         string
	ItemName     string
	ItemClass    string
	ItemTypeName string
	ItemSlot     string

	ModelPlayer string
	Prefab      string
}

func NewItem(index string) *Item {
	return &Item{
		Index: index,
	}
}

func (i *Item) initFromData(data *vdf.KeyValue) error {
	var err error
	var prefab *Item

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

	return nil
}

func (i *Item) MarshalJSON() ([]byte, error) {
	ret := make(map[string]interface{})

	ret["Index"] = i.Index
	ret["Name"] = i.Name
	ret["ItemName"] = i.ItemName
	ret["ItemClass"] = i.ItemClass
	ret["ItemTypeName"] = i.ItemTypeName
	ret["ItemSlot"] = i.ItemSlot

	if i.ModelPlayer != "" {
		ret["ModelPlayer"] = i.ModelPlayer
	}

	return json.Marshal(ret)
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

package dota2

import (
	"errors"

	"github.com/baldurstod/vdf"
)

func createItems() map[string]*ItemTemplate {
	return map[string]*ItemTemplate{}
}

var items = createItems()
var itemsByName = createItems()
var itemsPerHero = func() map[string][]*ItemTemplate { return make(map[string][]*ItemTemplate) }()
var prefabs = createItems()

func InitItems(buf []byte) error {
	vdf := vdf.VDF{}
	root := vdf.Parse(buf)
	itemsGame, err := root.Get("items_game")
	if err != nil {
		return err
	}

	err = initItems(itemsGame)
	if err != nil {
		return err
	}
	return nil
}

func initItems(datas *vdf.KeyValue) error {
	prefabsVdf, err := datas.Get("prefabs")
	if err != nil {
		return err
	}

	for _, prefabVdf := range prefabsVdf.GetChilds() {
		prefab, err := addItemTemplate(prefabVdf)
		if err != nil {
			return err
		}

		prefabs[prefab.Index] = prefab
	}

	itemsVdf, err := datas.Get("items")
	if err != nil {
		return err
	}

	for _, itemVdf := range itemsVdf.GetChilds() {
		item, err := addItemTemplate(itemVdf)
		if err != nil {
			return err
		}

		items[item.Index] = item
		itemsByName[item.Name] = item
	}
	return nil
}

func addItemTemplate(datas *vdf.KeyValue) (*ItemTemplate, error) {

	item := newItemTemplate(datas.Key)
	err := item.initFromData(datas)
	if err != nil {
		return nil, err
	}

	for npc, used := range item.UsedByHeroes {
		if used {
			arr, exist := itemsPerHero[npc]
			if !exist {
				arr = make([]*ItemTemplate, 0, 100)
			}
			itemsPerHero[npc] = append(arr, item)
		}
	}

	return item, nil
}

func CreateItem(index string) (*Item, error) {
	t, ok := items[index]
	if !ok {
		return nil, errors.New("item not found " + index)
	}

	return newItem(t), nil
}

func GetPrefab(index string) (*ItemTemplate, error) {
	h, ok := prefabs[index]
	if !ok {
		return nil, errors.New("prefab not found " + index)
	}
	return h, nil
}

func GetItemTemplates() map[string]*ItemTemplate {
	return items
}

func GetItemTemplate(index string) (*ItemTemplate, error) {
	if template, ok := items[index]; ok {
		return template, nil
	}

	return nil, errors.New("item not found " + index)
}

func GetItemTemplateByName(name string) (*ItemTemplate, error) {
	if template, ok := itemsByName[name]; ok {
		return template, nil
	}

	return nil, errors.New("item not found " + name)
}

/*
func GetBaseItems(hero string) []*Item {
	i := make([]*Item, 0, 10)
	for _, item := range items {
		if item.BaseItem && item.IsUsedByHero(hero) {
			i = append(i, item)
		}
	}
	return i
}
*/

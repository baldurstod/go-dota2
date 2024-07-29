package dota2

import (
	"errors"

	"github.com/baldurstod/vdf"
)

func createItems() map[string]*Item {
	return map[string]*Item{}
}

var items = createItems()
var itemsPerHero = func() map[string][]*Item { return make(map[string][]*Item) }()
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
		prefab, err := addItem(prefabVdf)
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
		item, err := addItem(itemVdf)
		if err != nil {
			return err
		}

		items[item.Index] = item
	}
	return nil
}

func addItem(datas *vdf.KeyValue) (*Item, error) {

	item := NewItem(datas.Key)
	err := item.initFromData(datas)
	if err != nil {
		return nil, err
	}

	for npc, used := range item.UsedByHeroes {
		if used {
			arr, exist := itemsPerHero[npc]
			if !exist {
				arr = make([]*Item, 0, 100)
			}
			itemsPerHero[npc] = append(arr, item)
		}
	}

	return item, nil
}

func GetItem(index string) (*Item, error) {
	h, ok := items[index]
	if !ok {
		return nil, errors.New("item not found " + index)
	}
	return h, nil
}

func GetPrefab(index string) (*Item, error) {
	h, ok := prefabs[index]
	if !ok {
		return nil, errors.New("prefab not found " + index)
	}
	return h, nil
}

func GetItems() map[string]*Item {
	return items
}

func GetBaseItems(hero string) []*Item {
	i := make([]*Item, 0, 10)
	for _, item := range items {
		if item.BaseItem && item.IsUsedByHero(hero) {
			i = append(i, item)
		}
	}
	return i
}

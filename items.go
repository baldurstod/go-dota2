package dota2

import (
	"errors"
	"log"

	"github.com/baldurstod/vdf"
)

func createItems() map[string]*Item {
	return map[string]*Item{}
}

var items = createItems()
var prefabs = createItems()

func InitItems(buf []byte) error {
	vdf := vdf.VDF{}
	root := vdf.Parse(buf)
	//log.Println(root.Get("DOTAHeroes"))
	itemsGame, err := root.Get("items_game")
	if err != nil {
		return err
	}

	log.Println(itemsGame)
	/*items, err := itemsGame.Get("items")*/
	err = initItems(itemsGame)
	if err != nil {
		return err
	}
	/*
		if err != nil {
			return err
		}

		for _, hero := range heroes.GetChilds() {
			if strings.HasPrefix(hero.Key, "npc_") {
				addHero(hero)
			}
		}

		//log.Println(heroes)
	/*/
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

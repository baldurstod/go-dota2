package dota2

import (
	"errors"
	"log"
	"strconv"

	"github.com/baldurstod/vdf"
)

func createItems() map[int]*Item {
	return map[int]*Item{}
}

var items = createItems()

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
	items, err := datas.Get("items")
	if err != nil {
		return err
	}

	for _, item := range items.GetChilds() {
		_, err = addItem(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func addItem(datas *vdf.KeyValue) (*Item, error) {
	index, err := strconv.Atoi(datas.Key)
	if err != nil {
		return nil, err
	}

	item := NewItem(index)
	err = item.initFromData(datas)
	if err != nil {
		return nil, err
	}

	items[index] = item

	return item, nil
}

func GetIndex(index int) (*Item, error) {
	h, ok := items[index]
	if !ok {
		return nil, errors.New("item not found " + strconv.Itoa(index))
	}
	return h, nil
}

func GetItems() map[int]*Item {
	return items
}

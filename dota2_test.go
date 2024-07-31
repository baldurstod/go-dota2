package dota2_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/baldurstod/go-dota2"
)

func initHeroes() error {
	buf, err := os.ReadFile(varFolder + "npc_heroes.txt")
	if err != nil {
		return err
	}
	err = dota2.InitHeroes(buf)
	if err != nil {
		return err
	}
	return nil
}

func initItems() error {
	buf, err := os.ReadFile(varFolder + "items_game.txt")
	if err != nil {
		return err
	}
	err = dota2.InitItems(buf)
	if err != nil {
		return err
	}

	return nil
}

func initAll() error {
	if err := initHeroes(); err != nil {
		return err
	}
	if err := initItems(); err != nil {
		return err
	}
	return nil
}

func TestHeroes(t *testing.T) {
	log.Println("start heroes")

	err := initHeroes()
	if err != nil {
		t.Error(err)
		return
	}
	log.Println("end heroes")

	//j, _ := json.MarshalIndent(dota2.GetHeroes(), "", "\t")
	//log.Println(string(j[:]))
	//os.WriteFile(path.Join(varFolder, "heroes.json"), j, 0666)

}

func TestItems(t *testing.T) {
	if err := initItems(); err != nil {
		t.Error(err)
		return
	}

	//j, _ := json.MarshalIndent(dota2.GetItems(), "", "\t")
	j, _ := json.MarshalIndent(dota2.GetBaseItems("npc_dota_hero_dark_willow"), "", "\t")
	log.Println(string(j[:]))
	//os.WriteFile(path.Join(varFolder, "items.json"), j, 0666)
}

func TestAssetModifiers(t *testing.T) {
	if err := initItems(); err != nil {
		t.Error(err)
		return
	}

	//j, _ := json.MarshalIndent(dota2.GetBaseItems("npc_dota_hero_dark_willow"), "", "\t")
	//log.Println(string(j[:]))
	item, err := dota2.GetItem("5156")
	if err != nil {
		t.Error(err)
		return
	}
	if item != nil {
		log.Println(item.GetAssetModifiers(1))
	}
}

func TestHeroItems(t *testing.T) {
	if err := initAll(); err != nil {
		t.Error(err)
		return
	}

	h, err := dota2.GetHero("npc_dota_hero_crystal_maiden")
	if err != nil {
		t.Error(err)
		return
	}

	//h.EquipItem("19205")

	items := h.GetItems()
	j, _ := json.MarshalIndent(items, "", "\t")
	log.Println(string(j[:]))

}

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
	if err := initAll(); err != nil {
		t.Error(err)
		return
	}

	var h *dota2.Hero
	var err error
	if h, err = dota2.GetHero("npc_dota_hero_dark_willow"); err != nil {
		t.Error(err)
		return
	}
	for _, item := range h.GetItems() {
		log.Println(item.GetName())
	}
}

func TestAssetModifiers(t *testing.T) {
	if err := initItems(); err != nil {
		t.Error(err)
		return
	}

	//j, _ := json.MarshalIndent(dota2.GetBaseItems("npc_dota_hero_dark_willow"), "", "\t")
	//log.Println(string(j[:]))
	item, err := dota2.CreateItem("5156")
	if err != nil {
		t.Error(err)
		return
	}
	if item != nil {
		log.Println(item.GetAssetModifiers(0))
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

func TestWrongHero(t *testing.T) {
	if err := initAll(); err != nil {
		t.Error(err)
		return
	}

	h, err := dota2.GetHero("npc_dota_hero_antimage")
	if err != nil {
		t.Error(err)
		return
	}

	if err = h.EquipItem("19205"); err == nil { // try to equip Conduit of the Blueheart
		t.Error("EquipItem should return an error")
		return
	}
}

func TestHeroModel(t *testing.T) {
	if err := initAll(); err != nil {
		t.Error(err)
		return
	}

	h, err := dota2.GetHero("npc_dota_hero_crystal_maiden")
	if err != nil {
		t.Error(err)
		return
	}

	m1 := h.GetModel()
	h.EquipItem("19205") // cm persona
	m2 := h.GetModel()

	if m1 != "models/heroes/crystal_maiden/crystal_maiden.vmdl" {
		t.Error("wrong hero model")
		return
	}

	if m1 != " models/heroes/crystal_maiden_persona/crystal_maiden_persona.vmdl" {
		t.Error("wrong hero model")
		return
	}

	log.Println(m1, m2)
}

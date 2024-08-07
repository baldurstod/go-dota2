package dota2_test

import (
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
	printHero(h)
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
		item.Style = 0
		log.Println(item.GetAssetModifiers())
		item.Style = 1
		log.Println(item.GetAssetModifiers())
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

	printHero(h)
	h.EquipItem("19205", true)
	printHero(h)
}

func printHero(h *dota2.Hero) {
	log.Println(h.GetModel())
	for _, item := range h.GetItems() {
		log.Println(item.GetName(), item.GetSkin())
	}
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

	if _, err = h.EquipItem("19205", true); err == nil { // try to equip Conduit of the Blueheart
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
	h.EquipItem("19205", true) // cm persona
	m2 := h.GetModel()

	if m1 != "models/heroes/crystal_maiden/crystal_maiden.vmdl" {
		t.Error("wrong hero model")
		return
	}

	if m2 != "models/heroes/crystal_maiden_persona/crystal_maiden_persona.vmdl" {
		t.Error("wrong hero model")
		return
	}

	log.Println(m1, m2)
}

func TestItemSkin(t *testing.T) {
	if err := initAll(); err != nil {
		t.Error(err)
		return
	}

	h, err := dota2.GetHero("npc_dota_hero_ogre_magi")
	if err != nil {
		t.Error(err)
		return
	}

	item, err := h.EquipItem("13670", true)
	if err != nil {
		t.Error(err)
		return
	}
	printHero(h)
	item.Style = 1
	printHero(h)
}

func TestEquipBundle(t *testing.T) {
	if err := initAll(); err != nil {
		t.Error(err)
		return
	}

	h, err := dota2.GetHero("npc_dota_hero_razor")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = h.EquipItem("23100", true)
	if err != nil {
		t.Error(err)
		return
	}
	printHero(h)
}

package dota2_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/baldurstod/go-dota2"
)

func DisabledTestHeroes(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("start heroes")

	buf, err := os.ReadFile(varFolder + "npc_heroes.txt")
	if err != nil {
		t.Error(err)
		return
	}
	log.Println("start parse")
	dota2.InitHeroes(buf)
	log.Println("end heroes")

	//j, _ := json.MarshalIndent(dota2.GetHeroes(), "", "\t")
	//log.Println(string(j[:]))
	//os.WriteFile(path.Join(varFolder, "heroes.json"), j, 0666)

}

func TestItems(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	buf, err := os.ReadFile(varFolder + "items_game.txt")
	if err != nil {
		t.Error(err)
		return
	}
	err = dota2.InitItems(buf)
	if err != nil {
		t.Error(err)
		return
	}

	//j, _ := json.MarshalIndent(dota2.GetItems(), "", "\t")
	j, _ := json.MarshalIndent(dota2.GetBaseItems("npc_dota_hero_dark_willow"), "", "\t")
	log.Println(string(j[:]))
	//os.WriteFile(path.Join(varFolder, "items.json"), j, 0666)
}

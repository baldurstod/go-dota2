package dota2_test

import (
	"log"
	"os"
	"testing"

	"encoding/json"

	"github.com/baldurstod/go-dota2"
)

func TestHeroes(t *testing.T) {
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

	j, _ := json.MarshalIndent(dota2.GetHeroes(), "", "\t")
	log.Println(string(j[:]))
	//os.WriteFile(path.Join(varFolder, "heroes.json"), j, 0666)

}

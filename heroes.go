package dota2

import (
	"errors"
	"strings"

	"github.com/baldurstod/vdf"
)

var heroes = func() map[string]*HeroTemplate { return map[string]*HeroTemplate{} }()

func InitHeroes(buf []byte) error {
	vdf := vdf.VDF{}
	root := vdf.Parse(buf)
	//log.Println(root.Get("DOTAHeroes"))
	heroes, err := root.Get("DOTAHeroes")

	if err != nil {
		return err
	}

	for _, hero := range heroes.GetChilds() {
		if strings.HasPrefix(hero.Key, "npc_") {
			addHeroTemplate(hero)
		}
	}

	//log.Println(heroes)

	return nil
}

func addHeroTemplate(datas *vdf.KeyValue) (*HeroTemplate, error) {
	hero := newHeroTemplate(datas.Key)
	hero.initFromData(datas)

	heroes[hero.entity] = hero

	return hero, nil
}

func GetHero(entity string) (*Hero, error) {
	template, ok := heroes[entity]
	if !ok {
		return nil, errors.New("hero not found " + entity)
	}

	return newHero(template), nil
}

/*
func GetHeroes() map[string]*HeroTemplate {
	return heroes
}
*/

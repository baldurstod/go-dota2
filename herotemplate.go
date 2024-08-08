package dota2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/baldurstod/vdf"
)

type HeroTemplate struct {
	entity      string
	name        string
	heroID      int
	heroOrderID int
	model       string
	personas    []Persona
	itemSlots   map[string]ItemSlot
}

func newHeroTemplate(entity string) *HeroTemplate {
	return &HeroTemplate{
		entity:    entity,
		personas:  make([]Persona, 0),
		itemSlots: make(map[string]ItemSlot),
	}
}

func (h *HeroTemplate) initFromData(data *vdf.KeyValue) error {
	var err error
	if h.model, err = data.GetString("Model"); err != nil {
		return err
	}

	if h.heroID, err = data.GetInt("HeroID"); err != nil {
		return err
	}

	if h.heroOrderID, err = data.GetInt("HeroOrderID"); err != nil {
		return err
	}

	if personas, err := data.Get("Persona"); err == nil {
		for _, p := range personas.GetChilds() {
			persona := Persona{}
			err = persona.initFromData(p)
			if err != nil {
				return fmt.Errorf("error while initializing persona: <%w>", err)
			}
			h.personas = append(h.personas, persona)
		}
	}

	if slots, err := data.Get("ItemSlots"); err == nil {
		for _, p := range slots.GetChilds() {
			itemSlot := ItemSlot{}
			err = itemSlot.initFromData(p)
			if err != nil {
				return fmt.Errorf("error while initializing persona: <%w>", err)
			}
			h.itemSlots[itemSlot.SlotName] = itemSlot
		}
	}

	h.itemSlots["bundle"] = ItemSlot{SlotIndex: -1, SlotName: "bundle", SlotText: "Bundle"}

	return nil
}

func (h *HeroTemplate) String() string {
	var sb strings.Builder

	sb.WriteString("Model: " + h.model + "\n")
	sb.WriteString("HeroID " + strconv.Itoa(h.heroID) + "\n")

	for _, p := range h.personas {
		sb.WriteString("Persona " + p.String())

	}
	//Personas

	return sb.String()
}

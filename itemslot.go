package dota2

import (
	"strconv"
	"strings"

	"github.com/baldurstod/vdf"
)

type ItemSlot struct {
	SlotIndex int
	SlotName  string
	SlotText  string
}

func (is *ItemSlot) initFromData(data *vdf.KeyValue) error {
	var err error

	if is.SlotIndex, err = data.GetInt("SlotIndex"); err != nil {
		return err
	}

	if is.SlotName, err = data.GetString("SlotName"); err != nil {
		return err
	}

	if is.SlotText, err = data.GetString("SlotText"); err != nil {
		return err
	}

	return nil
}

func (is *ItemSlot) String() string {
	var sb strings.Builder

	sb.WriteString("SlotIndex: " + strconv.Itoa(is.SlotIndex) + "\n")
	sb.WriteString("SlotName: " + is.SlotName + "\n")
	sb.WriteString("SlotText: " + is.SlotText + "\n")

	return sb.String()
}

/*
	"0"
	{
		"SlotIndex"		"0"
		"SlotName"		"weapon"
		"SlotText"		"#LoadoutSlot_Weapon"
		"TextureWidth"		"256"
		"TextureHeight"		"256"
		"MaxPolygonsLOD0"		"2500"
		"MaxPolygonsLOD1"		"1000"
	}
*/

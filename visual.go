package dota2

import (
	"github.com/baldurstod/vdf"
)

type Visual struct {
	Type     string
	Modifier string
	Asset    string
}

func (v *Visual) initFromData(data *vdf.KeyValue) error {
	v.Type, _ = data.GetString("type")
	v.Modifier, _ = data.GetString("modifier")
	v.Asset, _ = data.GetString("asset")
	return nil
}

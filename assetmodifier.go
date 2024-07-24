package dota2

import (
	"github.com/baldurstod/vdf"
)

type AssetModifier struct {
	Type     string
	Modifier string
	Asset    string
}

func (am *AssetModifier) initFromData(data *vdf.KeyValue) error {
	am.Type, _ = data.GetString("type")
	am.Modifier, _ = data.GetString("modifier")
	am.Asset, _ = data.GetString("asset")
	return nil
}

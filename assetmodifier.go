package dota2

import (
	"github.com/baldurstod/vdf"
)

type AssetModifier struct {
	Type     string `json:"type,omitempty"`
	Modifier string `json:"modifier,omitempty"`
	Asset    string `json:"asset,omitempty"`
}

func (am *AssetModifier) initFromData(data *vdf.KeyValue) error {
	am.Type, _ = data.GetString("type")
	am.Modifier, _ = data.GetString("modifier")
	am.Asset, _ = data.GetString("asset")
	return nil
}

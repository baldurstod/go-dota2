package dota2

import (
	"github.com/baldurstod/vdf"
)

type AssetModifier struct {
	Type     string `json:"type,omitempty"`
	Modifier string `json:"modifier,omitempty"`
	Asset    string `json:"asset,omitempty"`
	Style    int    `json:"style,omitempty"`
	Persona  int    `json:"persona,omitempty"`
}

func (am *AssetModifier) initFromData(data *vdf.KeyValue) error {
	am.Type, _ = data.GetString("type")
	am.Modifier, _ = data.GetString("modifier")
	am.Asset, _ = data.GetString("asset")
	am.Style, _ = data.GetInt("style")
	am.Persona, _ = data.GetInt("persona")
	return nil
}

const (
	MODIFIER_UNKNOWN             = ""
	MODIFIER_ENTITY_MODEL        = "entity_model"
	MODIFIER_PARTICLE            = "particle"
	MODIFIER_PARTICLE_CREATE     = "particle_create"
	MODIFIER_PARTICLE_SNAPSHOT   = "particle_snapshot"
	MODIFIER_ADDITIONAl_WEARABLE = "additional_wearable"
	MODIFIER_PERSONA             = "persona"
)

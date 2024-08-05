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
	MODIFIER_UNKNOWN                   = ""
	MODIFIER_ENTITY_MODEL              = "entity_model"
	MODIFIER_PARTICLE                  = "particle"
	MODIFIER_PARTICLE_CREATE           = "particle_create"
	MODIFIER_PARTICLE_SNAPSHOT         = "particle_snapshot"
	MODIFIER_ADDITIONAL_WEARABLE       = "additional_wearable"
	MODIFIER_PERSONA                   = "persona"
	MODIFIER_ACTIVITY                  = "activity"
	MODIFIER_HERO_MODEL_CHANGE         = "hero_model_change"
	MODIFIER_MODEL                     = "model"
	MODIFIER_MODEL_SKIN                = "model_skin"
	MODIFIER_PET                       = "pet"
	MODIFIER_PORTRAIT_BACKGROUND_MODEL = "portrait_background_model"
	MODIFIER_COURIER                   = "courier"
	MODIFIER_COURIER_FLYING            = "courier_flying"
	MODIFIER_ENTITY_CLIENTSIDE_MODEL   = "entity_clientside_model"
)

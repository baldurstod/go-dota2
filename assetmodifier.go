package dota2

import (
	"github.com/baldurstod/vdf"
)

type AssetModifier struct {
	Type         string  `json:"type,omitempty"`
	Modifier     string  `json:"modifier,omitempty"`
	Asset        string  `json:"asset,omitempty"`
	Style        int     `json:"style,omitempty"`
	Skin         int     `json:"skin,omitempty"`
	Level        int     `json:"level,omitempty"`
	Persona      int     `json:"persona,omitempty"`
	ModelScale   float32 `json:"model_scale,omitempty"`
	VersusScale  float32 `json:"versus_scale,omitempty"`
	LoadoutScale float32 `json:"loadout_scale,omitempty"`
	ForceDisplay bool    `json:"force_display,omitempty"`
}

func (am *AssetModifier) initFromData(data *vdf.KeyValue) error {
	var err error
	am.Type, _ = data.GetString("type")
	am.Modifier, _ = data.GetString("modifier")
	am.Asset, _ = data.GetString("asset")
	am.Style, _ = data.GetInt("style")
	am.Skin, _ = data.GetInt("skin")
	am.Level, _ = data.GetInt("level")
	am.Persona, _ = data.GetInt("persona")
	if am.ModelScale, err = data.GetFloat32("ModelScale"); err != nil {
		am.ModelScale = 1
	}
	if am.VersusScale, err = data.GetFloat32("VersusScale"); err != nil {
		am.VersusScale = 1
	}
	if am.LoadoutScale, err = data.GetFloat32("LoadoutScale"); err != nil {
		am.LoadoutScale = 1
	}
	am.ForceDisplay, _ = data.GetBool("force_display")
	return nil
}

const (
	MODIFIER_UNKNOWN                       = ""
	MODIFIER_ENTITY_MODEL                  = "entity_model"
	MODIFIER_PARTICLE                      = "particle"
	MODIFIER_PARTICLE_CREATE               = "particle_create"
	MODIFIER_PARTICLE_SNAPSHOT             = "particle_snapshot"
	MODIFIER_ADDITIONAL_WEARABLE           = "additional_wearable"
	MODIFIER_PERSONA                       = "persona"
	MODIFIER_ACTIVITY                      = "activity"
	MODIFIER_HERO_MODEL_CHANGE             = "hero_model_change"
	MODIFIER_MODEL                         = "model"
	MODIFIER_MODEL_SKIN                    = "model_skin"
	MODIFIER_PET                           = "pet"
	MODIFIER_PORTRAIT_BACKGROUND_MODEL     = "portrait_background_model"
	MODIFIER_COURIER                       = "courier"
	MODIFIER_COURIER_FLYING                = "courier_flying"
	MODIFIER_ENTITY_CLIENTSIDE_MODEL       = "entity_clientside_model"
	MODIFIER_ABILITY_ICON                  = "ability_icon"
	MODIFIER_ANNOUCER_PREVIEW              = "announcer_preview"
	MODIFIER_SOUND                         = "sound"
	MODIFIER_HERO_SCALE                    = "hero_scale"
	MODIFIER_RESPONSE_CRITERIA             = "response_criteria"
	MODIFIER_ICON_REPLACEMENT_HERO         = "icon_replacement_hero"
	MODIFIER_ICON_REPLACEMENT_HERO_MINIMAP = "icon_replacement_hero_minimap"
	MODIFIER_BUFF_MODIFIER                 = "buff_modifier"
	MODIFIER_CUSTOM_KILL_EFFECT            = "custom_kill_effect"
	MODIFIER_ARCANA_LEVEL                  = "arcana_level"
)

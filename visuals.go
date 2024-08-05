package dota2

import (
	"errors"
	"strings"

	"github.com/baldurstod/vdf"
)

type Visuals struct {
	AssetModifiers               []*AssetModifier `json:"asset_modifiers,omitempty"`
	Styles                       *Styles          `json:"styles,omitempty"`
	SkipModelCombine             bool             `json:"skip_model_combine,omitempty"`
	HideStylesFromUI             bool             `json:"hide_styles_from_ui,omitempty"`
	HideOnPortrait               bool             `json:"hide_on_portrait,omitempty"`
	OnlyDisplayOnHeroModelChange bool             `json:"only_display_on_hero_model_change,omitempty"`
	Skin                         int              `json:"skin,omitempty"`
}

func NewVisuals() *Visuals {
	return &Visuals{
		AssetModifiers: make([]*AssetModifier, 0),
	}
}

func (v *Visuals) initFromData(data *vdf.KeyValue) error {
	for _, child := range data.GetChilds() {
		switch child.Key {
		case "skip_model_combine":
			v.SkipModelCombine, _ = child.ToBool()
		case "hide_on_portrait":
			v.HideOnPortrait, _ = child.ToBool()
		case "hide_styles_from_ui":
			v.HideStylesFromUI, _ = child.ToBool()
		case "only_display_on_hero_model_change":
			v.OnlyDisplayOnHeroModelChange, _ = child.ToBool()
		case "skin":
			v.Skin, _ = child.ToInt()
		case "styles":
			v.Styles = NewStyles()
			err := v.Styles.initFromData(child)
			if err != nil {
				return err
			}
		case "alternate_icons":
			//TODO
		case "animation_modifiers":
			//TODO
		case "player_card":
			//TODO
		default:
			if strings.HasPrefix(child.Key, "asset_modifier") {
				am := AssetModifier{}
				am.initFromData(child)
				v.AssetModifiers = append(v.AssetModifiers, &am)
			} else {
				return errors.New("unknown key " + child.Key)
			}
		}
	}
	return nil
}

func (v *Visuals) getSkin(style int) int {
	if v.Styles != nil {
		if skin, err := v.Styles.getSkin(style); err == nil {
			return skin
		}
	}
	return 0
}

/*
func (v *Visuals) MarshalJSON() ([]byte, error) {
	ret := make(map[string]interface{})

	ret["skip_model_combine"] = v.SkipModelCombine
	ret["hide_styles_from_ui"] = v.HideStylesFromUI
	ret["hide_on_portrait"] = v.HideOnPortrait
	ret["only_display_on_hero_model_change"] = v.OnlyDisplayOnHeroModelChange
	ret["skin"] = v.Skin

	if len(v.AssetModifiers) > 0 {
		ret["asset_modifiers"] = v.AssetModifiers
	}

	if len(v.Styles.Styles) > 0 {
		ret["styles"] = v.Styles
	}

	return json.Marshal(ret)
}
*/

package dota2

import (
	"errors"
	"strings"

	"github.com/baldurstod/vdf"
)

type Visuals struct {
	AssetModifiers               []*AssetModifier
	Styles                       *Styles
	SkipModelCombine             bool
	HideStylesFromUI             bool
	HideOnPortrait               bool
	OnlyDisplayOnHeroModelChange bool
	Skin                         int
}

func NewVisuals() *Visuals {
	return &Visuals{
		AssetModifiers: make([]*AssetModifier, 0),
		Styles:         NewStyles(),
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

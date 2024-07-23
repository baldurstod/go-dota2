package dota2

import (
	"errors"
	"strings"

	"github.com/baldurstod/vdf"
)

type Visuals struct {
	Visuals                      []*Visual
	Styles                       *Styles
	SkipModelCombine             bool
	HideStylesFromUI             bool
	HideOnPortrait               bool
	OnlyDisplayOnHeroModelChange bool
	Skin                         int
}

func NewVisuals() *Visuals {
	return &Visuals{
		Visuals: make([]*Visual, 0),
		Styles:  NewStyles(),
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
			v.Styles.initFromData(child)
		case "alternate_icons":
			//TODO
		case "animation_modifiers":
			//TODO
		case "player_card":
			//TODO
		default:
			if strings.HasPrefix(child.Key, "asset_modifier") {
				vis := Visual{}
				vis.initFromData(child)
				v.Visuals = append(v.Visuals, &vis)
			} else {
				return errors.New("unknown key " + child.Key)
			}
		}
	}
	return nil
}

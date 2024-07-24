package dota2

import (
	"github.com/baldurstod/vdf"
)

type Style struct {
	Name              string
	ModelPlayer       string
	Skin              int
	AutoStyleRule     string
	AutoStyleReason   string
	Asset             string
	AlternateIcon     bool
	EntityScaleFlying float32
	BodyGroups        map[string]bool
}

func NewStyle() *Style {
	return &Style{
		BodyGroups: make(map[string]bool),
	}
}

func (s *Style) initFromData(data *vdf.KeyValue) error {
	s.Name, _ = data.GetString("name")
	s.ModelPlayer, _ = data.GetString("model_player")
	s.Skin, _ = data.GetInt("skin")
	s.AutoStyleRule, _ = data.GetString("auto_style_rule")
	s.AutoStyleReason, _ = data.GetString("auto_style_reason")
	s.Asset, _ = data.GetString("asset")
	s.AlternateIcon, _ = data.GetBool("alternate_icon")
	s.EntityScaleFlying, _ = data.GetFloat32("alternate_icon")

	if bodyGroups, err := data.Get("body_groups"); err == nil {
		for _, child := range bodyGroups.GetChilds() {
			b, err := child.ToBool()
			if err != nil {
				return err
			}
			s.BodyGroups[child.Key] = b
		}
	}

	return nil
}

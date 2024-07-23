package dota2

import (
	"errors"

	"github.com/baldurstod/vdf"
)

type Style struct {
	Name            string
	ModelPlayer     string
	Skin            int
	AutoStyleRule   string
	AutoStyleReason string
	AlternateIcon   bool
}

func (s *Style) initFromData(data *vdf.KeyValue) error {
	for _, child := range data.GetChilds() {
		switch child.Key {
		case "auto_style_rule":
		case "name":
		case "model_player":
		case "alternate_icon":
		case "skin":

		case "unlock":
		case "auto_style_reason":
		case "body_groups":
		case "asset":
		case "entity_scale_flying":
		default:
			return errors.New("unknown key " + child.Key)
		}
	}

	s.Name, _ = data.GetString("name")
	s.ModelPlayer, _ = data.GetString("model_player")
	s.Skin, _ = data.GetInt("skin")
	s.AutoStyleRule, _ = data.GetString("auto_style_rule")
	s.AutoStyleReason, _ = data.GetString("auto_style_reason")
	s.AlternateIcon, _ = data.GetBool("alternate_icon")
	return nil
}

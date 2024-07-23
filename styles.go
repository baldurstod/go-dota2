package dota2

import (
	"github.com/baldurstod/vdf"
)

type Styles struct {
	Styles []*Style
}

func NewStyles() *Styles {
	return &Styles{
		Styles: make([]*Style, 0),
	}
}

func (s *Styles) initFromData(data *vdf.KeyValue) error {
	/*
		for _, child := range data.GetChilds() {
			switch child.Key {
			case "skip_model_combine":
				v.SkipModelCombine, _ = child.ToBool()
			case "skin":
				v.Skin, _ = child.ToInt()
			default:
				if strings.HasPrefix(child.Key, "asset_modifier") {
					vis := Visual{}
					vis.initFromData(child)
					v.Visuals = append(v.Visuals, &vis)
				} else {
					return errors.New("unknown key " + child.Key)
				}
			}
		}*/
	return nil
}

package dota2

import (
	"errors"

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
	for _, child := range data.GetChilds() {
		style := NewStyle()
		err := style.initFromData(child)
		if err != nil {
			return err
		}
		s.Styles = append(s.Styles, style)
	}
	return nil
}

func (s *Styles) getSkin(style int) (int, error) {
	if style < 0 || style > len(s.Styles) {
		return 0, errors.New("out of bounds")
	}

	return s.Styles[style].Skin, nil
}

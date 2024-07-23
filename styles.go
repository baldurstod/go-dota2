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
	for _, child := range data.GetChilds() {
		style := Style{}
		err := style.initFromData(child)
		if err != nil {
			return err
		}
		s.Styles = append(s.Styles, &style)
	}
	return nil
}

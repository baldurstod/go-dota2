package dota2

import (
	"github.com/baldurstod/vdf"
)

type Style struct {
	Name        string
	ModelPlayer string
}

func (s *Style) initFromData(data *vdf.KeyValue) error {
	s.Name, _ = data.GetString("name")
	s.ModelPlayer, _ = data.GetString("model_player")
	return nil
}

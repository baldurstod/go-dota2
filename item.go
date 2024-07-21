package dota2

import (
	"github.com/baldurstod/vdf"
)

type Item struct {
	Index       int
	Name        string
	ModelPlayer string
}

func NewItem(index int) *Item {
	return &Item{
		Index: index,
	}
}

func (i *Item) initFromData(data *vdf.KeyValue) error {
	var err error
	if i.Name, err = data.GetString("name"); err != nil {
		return err
	}

	i.ModelPlayer, _ = data.GetString("model_player")

	return nil
}

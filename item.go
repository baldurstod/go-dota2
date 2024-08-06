package dota2

type Item struct {
	template *ItemTemplate
	Style    int
}

func newItem(template *ItemTemplate) *Item {
	return &Item{
		template: template,
	}
}

func (item *Item) GetIndex() string {
	return item.template.Index
}

func (item *Item) GetName() string {
	return item.template.Name
}

func (item *Item) GetItemSlot() string {
	return item.template.ItemSlot
}

func (item *Item) GetModelPlayer() string {
	return item.template.ModelPlayer
}

func (item *Item) IsUsedByHero(hero string) bool {
	return item.template.IsUsedByHero(hero)
}

func (item *Item) GetAssetModifiers() []*AssetModifier {
	return item.template.GetAssetModifiers(item.Style)
}

func (item *Item) GetPersonaId() int {
	return item.template.GetPersonaId()
}

func (item *Item) GetSkin() int {
	return item.template.GetSkin(item.Style)
}

package dota2

type item struct {
	template *ItemTemplate
	Style    int
}

func newItem(template *ItemTemplate) *item {
	return &item{
		template: template,
	}
}

func (item *item) GetIndex() string {
	return item.template.Index
}

func (item *item) GetName() string {
	return item.template.Name
}

func (item *item) GetItemSlot() string {
	return item.template.ItemSlot
}

func (item *item) GetModelPlayer() string {
	return item.template.ModelPlayer
}

func (item *item) IsUsedByHero(hero string) bool {
	return item.template.IsUsedByHero(hero)
}

func (item *item) GetAssetModifiers() []*AssetModifier {
	return item.template.GetAssetModifiers(item.Style)
}

func (item *item) GetPersonaId() int {
	return item.template.GetPersonaId()
}

func (item *item) GetSkin() int {
	return item.template.GetSkin(item.Style)
}

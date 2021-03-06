package mock

import (
	ifc "pkg.deepin.io/dde/daemon/launcher/interfaces"
)

type CategoryInfo struct {
	id    ifc.CategoryID
	name  string
	items map[ifc.ItemID]bool
}

func NewCategoryInfo(id ifc.CategoryID, name string) *CategoryInfo {
	return &CategoryInfo{
		id:    id,
		name:  name,
		items: map[ifc.ItemID]bool{},
	}
}

func (c *CategoryInfo) ID() ifc.CategoryID {
	return c.id
}

func (c *CategoryInfo) Name() string {
	return c.name
}

func (c *CategoryInfo) LocaleName() string {
	// TODO: locale name
	return c.name
}

func (c *CategoryInfo) AddItem(itemID ifc.ItemID) {
	c.items[itemID] = true
}
func (c *CategoryInfo) RemoveItem(itemID ifc.ItemID) {
	delete(c.items, itemID)
}

func (c *CategoryInfo) Items() []ifc.ItemID {
	items := []ifc.ItemID{}
	for itemID := range c.items {
		items = append(items, itemID)
	}
	return items
}

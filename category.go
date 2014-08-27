package wu

import (
	"github.com/mattn/go-ole"
	"github.com/mattn/go-ole/oleutil"
)

type Categories struct {
	categories *ole.IDispatch
	Categories []*Category
}

type Category struct {
	category    *ole.IDispatch
	CategoryID  string
	Name        string
	Description string
}

func newCategory(category *ole.IDispatch) *Category {
	cat := new(Category)
	cat.category = category
	cat.CategoryID = oleutil.MustGetProperty(category, "CategoryID").ToString()
	cat.Name = oleutil.MustGetProperty(category, "Name").ToString()
	cat.Description = oleutil.MustGetProperty(category, "Description").ToString()
	return cat
}

func (cat *Category) GetString(attr string) string {
	return oleutil.MustGetProperty(cat.category, attr).ToString()
}

func (cat *Category) GetBool(attr string) bool {
	return oleutil.MustGetProperty(cat.category, attr).Value().(bool)
}

func (cat *Category) GetInt(attr string) int {
	return int(oleutil.MustGetProperty(cat.category, attr).Val)
}

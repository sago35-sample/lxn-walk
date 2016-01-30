package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"sort"
)

type Person struct {
	Index   int
	Name    string
	Age     int
	checked bool
}

type PersonModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Person
}

// TableViewを実装する際、RowCount()とValue()が必要
func (m *PersonModel) RowCount() int {
	return len(m.items)
}

// TableViewを実装する際、RowCount()とValue()が必要
func (m *PersonModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Index
	case 1:
		return item.Name
	case 2:
		return item.Age
	}
	panic("unexpected col")
}

func (m *PersonModel) Checked(row int) bool {
	return m.items[row].checked
}

func (m *PersonModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked
	return nil
}

// TableViewをsort可能にするには、walk.SorterBase、Sort()、Len()、Less()、Swap()の実装が必要
func (m *PersonModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

func (m *PersonModel) Len() int {
	return len(m.items)
}

func (m *PersonModel) Less(i, j int) bool {
	a, b := m.items[i], m.items[j]

	c := func(ls bool) bool {
		if m.sortOrder == walk.SortAscending {
			return ls
		}

		return !ls
	}

	switch m.sortColumn {
	case 0:
		return c(a.Index < b.Index)
	case 1:
		return c(a.Name < b.Name)
	case 2:
		return c(a.Age < b.Age)
	}

	panic("unreachable")
}

func (m *PersonModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

func NewPersonModel() *PersonModel {
	m := new(PersonModel)
	m.items = make([]*Person, 3)

	m.items[0] = &Person{
		Index: 0,
		Name:  "山田",
		Age:   20,
	}

	m.items[1] = &Person{
		Index: 1,
		Name:  "鈴木",
		Age:   21,
	}

	m.items[2] = &Person{
		Index: 2,
		Name:  "田中",
		Age:   22,
	}

	return m
}

type PersonMainWindow struct {
	*walk.MainWindow
	model *PersonModel
}

func main() {
	mw := &PersonMainWindow{model: NewPersonModel()}

	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "TableViewサンプル",
		Size:     Size{800, 600},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					HSpacer{},
					PushButton{
						Text: "Add",
						OnClicked: func() {
							mw.model.items = append(mw.model.items, &Person{
								Index: mw.model.Len() + 1,
								Name:  "xxx",
								Age:   mw.model.Len() * 5,
							})
							mw.model.PublishRowsReset()
						},
					},
				},
			},
			TableView{
				CheckBoxes:       true,
				ColumnsOrderable: true,
				MultiSelection:   true,
				Columns: []TableViewColumn{
					{Title: "#"},
					{Title: "名前"},
					{Title: "年齢"},
				},
				Model: mw.model,
				OnItemActivated: mw.tv_ItemAtivated,
			},
		},
	}.Run()
}

func (mw *PersonMainWindow) tv_ItemAtivated() {
	walk.MsgBox(mw, "title", "msg", walk.MsgBoxIconInformation)
}

package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Person struct {
	Index   int
	Name    string
	Age     int
	checked bool
}

type PersonModel struct {
	walk.TableModelBase
	items []*Person
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

func main() {
	model := NewPersonModel()

	MainWindow{
		Title:  "TableViewサンプル",
		Size:   Size{800, 600},
		Layout: VBox{},
		Children: []Widget{
			TableView{
				CheckBoxes: true,
				MultiSelection: true,
				Columns: []TableViewColumn{
					{Title: "#"},
					{Title: "名前"},
					{Title: "年齢"},
				},
				Model: model,
			},
		},
	}.Run()
}

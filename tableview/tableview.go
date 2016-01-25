package main

import (
	. "github.com/lxn/walk/declarative"
)

func main() {
	MainWindow{
		Title: "TableViewサンプル",
		Size:  Size{800, 600},
		Layout: VBox{},
		Children: []Widget{
			TableView{
				Columns: []TableViewColumn{
					{Title: "#"},
					{Title: "名前"},
					{Title: "年齢"},
				},
			},
		},
	}.Run()
}

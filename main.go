package main

import (
	"log"
	"os"
	"strings"
	"test-walk-gui/mylogger"

	"github.com/lxn/walk"
	ui "github.com/lxn/walk/declarative"
)

type MyWindow struct {
	*walk.MainWindow
	textEdit  *walk.LineEdit
	chk1      *walk.CheckBox
	chk2      *walk.CheckBox
	edit      *walk.TextEdit
	hostEdit  *walk.LineEdit
	destEdit  *walk.LineEdit
	portEdit  *walk.LineEdit
	statusbar *walk.StatusBarItem
}

var version string = "dev1.0"
var logger *mylogger.MyLogger

func main() {

	// Logger
	f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger = mylogger.New(log.New(f, "", log.Ldate|log.Ltime|log.Llongfile))

	logger.Info("Version", version)

	// Init Walk
	window := &MyWindow{}

	if err := (ui.MainWindow{
		AssignTo: &window.MainWindow,
		Title:    "App",
		MinSize:  ui.Size{Width: 300, Height: 200},
		Size:     ui.Size{Width: 500, Height: 500},
		Layout:   ui.VBox{},
		OnDropFiles: func(files []string) {
			window.textEdit.SetText(strings.Join(files, "\r\n"))
		},
		Children: []ui.Widget{
			ui.LineEdit{
				AssignTo:    &window.textEdit,
				ToolTipText: "Drop files here, from windows explorer...",
			},
			ui.GroupBox{
				Title:     "CheckBox",
				Layout:    ui.HBox{},
				Alignment: ui.AlignHNearVNear,
				Children: []ui.Widget{
					ui.CheckBox{
						AssignTo:         &window.chk1,
						Text:             "Check 1",
						Checked:          true,
						OnCheckedChanged: window.check1Changed,
					},
					ui.CheckBox{
						AssignTo:         &window.chk2,
						Text:             "Check 2",
						OnCheckedChanged: window.check2Changed,
					},
				},
			},
			ui.Composite{
				Layout: ui.Grid{},
				Children: []ui.Widget{
					ui.Label{
						Row:    0,
						Column: 0,
						Text:   "JmsHost:",
					},
					ui.LineEdit{
						AssignTo: &window.hostEdit,
						Row:      0,
						Column:   1,
						Text:     "192.168.1.105",
					},
					ui.Label{
						Row:    1,
						Column: 0,
						Text:   "JmsDest:",
					},
					ui.LineEdit{
						AssignTo: &window.destEdit,
						Row:      1,
						Column:   1,
						Text:     "/topic/IDC.QuoteExchange.bond2app",
					},
					ui.Label{
						Row:    2,
						Column: 0,
						Text:   "JmsPort:",
					},
					ui.LineEdit{
						AssignTo: &window.portEdit,
						Row:      2,
						Column:   1,
						Text:     "61612",
					},
				},
			},
			ui.TextEdit{
				AssignTo: &window.edit,
				VScroll:  true,
			},
			ui.PushButton{
				Text:      "Push",
				OnClicked: window.pbClicked,
			},
		},

		StatusBarItems: []ui.StatusBarItem{
			{
				AssignTo: &window.statusbar,
				Text:     "click",
				Width:    80,

				OnClicked: window.statusBarClicked,
			},
			{
				Text:        "left",
				ToolTipText: "no tooltip for me",
			},
			{
				Text: "\tcenter",
			},
			{
				Text: "\t\tright",
			},
		},
	}.Create()); err != nil {
		logger.Fatal(err)
	}

	// Log View
	lv, err := NewLogView(window)
	if err != nil {
		logger.Fatal(err)
	}

	logger.AddLogger(log.New(lv, "", 0))

	// Run Walk
	window.Run()

}

func (mw *MyWindow) check1Changed() {
	mw.edit.AppendText("Event: CHK1 changed\r\n")
	logger.Println("Event: CHK1 changed")
}

func (mw *MyWindow) check2Changed() {
	mw.edit.AppendText("Event: CHK2 changed\r\n")
	logger.Println("Event: CHK2 changed")
}

func (mw *MyWindow) pbClicked() {
	if mw.chk1.CheckState() == walk.CheckChecked {
		logger.Println("CHK1 changed")
		mw.edit.AppendText("CHK1 checked\r\n")
	}
	if mw.chk2.CheckState() == walk.CheckChecked {
		logger.Println("CHK2 changed")
		mw.edit.AppendText("CHK2 checked\r\n")
	}
}
func (mw *MyWindow) statusBarClicked() {
	if mw.statusbar.Text() == "click" {
		mw.statusbar.SetText("again")
	} else {
		mw.statusbar.SetText("click")
	}
}

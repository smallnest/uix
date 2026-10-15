package view

import (
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/command"
	"github.com/smallnest/uix/components/fileupload"
	"github.com/smallnest/uix/components/inputotp"
	"github.com/smallnest/uix/components/messagescroller"
	"github.com/smallnest/uix/components/questionnaire"
	"github.com/smallnest/uix/components/settings"
	"github.com/smallnest/uix/components/typewriter"
)

// commandState is the state of the command palette demo, which the
// palette keeps between frames.
var commandState command.State

// State is the state of the example; the controls edit it in place.
var State = GalleryState{}

// GalleryState holds the gallery: the page chosen in the sidebar and the
// search text filtering the sidebar items.
type GalleryState struct {
	// Selected is the ID of the page shown.
	Selected string
	// Query filters the sidebar items.
	Query string
}

// Demo holds the values the demo pages edit. Each page keeps its own
// state, so a control stays where the visitor left it across frames.
var Demo = DemoState{}

// DemoState is the state of every interactive demo of the gallery.
type DemoState struct {
	// The selection controls.
	BtnSel []string
	Seg    int
	Tabs   int
	Tg1    bool
	Tg2    bool
	Tg3    bool
	// The text input controls.
	Input     string
	Textarea  string
	Number    float64
	Search    string
	Combobox  string
	AutoComp  string
	Tokens    []string
	OTP       inputotp.State
	Amount    string
	Float     string
	AutoSize  string
	EditText  string
	SelStr    string
	Multi     []string
	NativeSel string
	StepperV  float64
	Buffer    *ui.TextBuffer
	Tw        typewriter.State
	// The value controls.
	Check   bool
	Switch  bool
	Radio   string
	Slider  float64
	RangeLo float64
	RangeHi float64
	Rating  int
	// The date, time and color controls.
	Date  time.Time
	Time  time.Time
	Both  time.Time
	Cal   time.Time
	Color ui.Color
	Well  ui.Color
	// The overlays; a true one shows its surface over the window.
	Dialog  bool
	Alert   bool
	Drawer  bool
	Sheet   bool
	Pop     bool
	Command bool
	Drop    bool
	NCenter int
	// The content controls.
	Composer string
	Chat     string
	Scroller messagescroller.State
	Question questionnaire.State
	FileUp   fileupload.State
	Settings settings.SettingsState
	// The layout controls.
	Split          float32
	Collapse, Acc1 bool
	Acc2           bool
	Tree1, Tree2   bool
	// The grid keeps its place and choice.
	Grid ui.GridState
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = GalleryState{Selected: "button"}
	Demo = DemoState{
		BtnSel:    []string{"left"},
		Tg1:       true,
		Tg3:       true,
		Input:     "Ada Lovelace",
		Textarea:  "A line of notes.",
		Number:    3,
		Combobox:  "Go",
		AutoComp:  "mygo",
		Tokens:    []string{"go", "ui"},
		EditText:  "Click to rename",
		SelStr:    "B",
		Multi:     []string{"go"},
		NativeSel: "Weekly",
		StepperV:  5,
		Buffer:    ui.NewTextBuffer("mygo"),
		Tw:        typewriter.State{Done: true},
		Check:     true,
		Switch:    true,
		Radio:     "Small",
		Slider:    40,
		RangeLo:   20,
		RangeHi:   80,
		Rating:    4,
		Date:      time.Now(),
		Time:      time.Now(),
		Both:      time.Now(),
		Cal:       time.Now(),
		Color:     ui.Color{R: 110, G: 90, B: 230, A: 255},
		Well:      ui.Color{R: 40, G: 160, B: 120, A: 255},
		Settings:  settings.SettingsState{Category: "general"},
		Split:     220,
		Collapse:  true,
		Acc1:      true,
		Tree1:     true,
	}
}

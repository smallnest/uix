package view

import (
	"fmt"
	"time"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/editor"

	components "github.com/oneo-me/onevra/source/components"
	ochat "github.com/oneo-me/onevra/source/components/chat"

	"github.com/smallnest/uix/components/carousel"
	"github.com/smallnest/uix/components/command"
	"github.com/smallnest/uix/components/fileupload"
	"github.com/smallnest/uix/components/infinitescroll"
	"github.com/smallnest/uix/components/inputotp"
	"github.com/smallnest/uix/components/messagescroller"
	"github.com/smallnest/uix/components/navigationmenu"
	"github.com/smallnest/uix/components/questionnaire"
	"github.com/smallnest/uix/components/settings"
	"github.com/smallnest/uix/components/typewriter"
)

// Web is the web view of the gallery window, made by main; nil in the
// headless tests, where the webview page draws a placeholder instead.
var Web *mygo.WebView

// commandState is the state of the command palette demo, which the
// palette keeps between frames.
var commandState command.State

// State is the state of the example; the controls edit it in place.
var State = GalleryState{}

// GalleryState holds the gallery: the page chosen in the sidebar, the
// search text filtering the sidebar items, and the library the sidebar
// is filtered to.
type GalleryState struct {
	// Selected is the ID of the page shown.
	Selected string
	// Query filters the sidebar items.
	Query string
	// Filter is the chosen segment of the library filter: 0 shows every
	// library mixed, 1 uix, 2 mygo, 3 onevra and 4 mujica alone.
	Filter int
}

// Demo holds the values the demo pages edit. Each page keeps its own
// state, so a control stays where the visitor left it across frames.
// The buffer and the history are born non-nil: a buffer input and the
// back buttons panic without them, and a visitor reaches the pages
// before any Reset of an embedding window.
var Demo = DemoState{Buffer: ui.NewTextBuffer(""), History: ui.NewRouter("/")}

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
	Split               float32
	Collapse, Acc1      bool
	Acc2                bool
	Tree1, Tree2, Tree3 bool
	// The grid keeps its place and choice.
	Grid ui.GridState
	// The navigation extras.
	Run   uint64
	Nav   navigationmenu.State
	Istep int
	// The pagination page shown, and the table's choice and sort.
	Page      int
	TableList ui.ListState
	TableSort ui.SortOrder
	// The carousel keeps where its track is and which slide shows.
	CarouselScroll ui.ScrollState
	Carousel       carousel.State
	// The find bar of the document demo.
	FindOpen    bool
	FindQuery   string
	FindCurrent int
	// The multi-select popup, closed unless the trigger opens it.
	MultiOpen bool
	// The infinite feed; OnMore appends from the counter.
	Feed     infinitescroll.State
	FeedRows []string
	FeedN    int
	// The responsive modal and the app shell.
	RModal   bool
	ShellSel string
	// The mygo pages: the history of the back and forward buttons, the
	// step of the stepped slider, the height of the vertical split, the
	// replay of the motion, and the diff and outline the pages draw.
	History   *ui.Router
	Step      float64
	SplitV    float32
	MotionRun uint64
	Diff      editor.DiffState
	OutlineS  ui.OutlineState[string]
	OutlineT  ui.OutlineState[string]
	// The onevra pages: the summary's expansion, the composer's draft
	// and measured glass, the clock of the time picker, the color of the
	// color view, the rows and branches of the virtualized collections,
	// and the transcript of the conversation, which the components keep.
	OSummary  ochat.SummaryState
	OComposer string
	OLayout   ochat.FloatingConversationState
	OTime     time.Time
	OColor    ui.Color
	OTable    components.TableVirtualizedState
	OTree     components.TreeVirtualizedState[string]
	OTurn     ochat.TurnState
	OConv     ochat.ConversationState
	OGroup    ochat.GroupState
	// The shader background of the signalfield page, and its tuning: the
	// intensity and the dot scale of the sliders, and the switches that
	// freeze the drift or ignore the pointer.
	ShaderIntensity float64
	ShaderScale     float64
	ShaderStill     bool
	ShaderInert     bool
	ShaderPrompt    string
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command. The gallery window's main sets Web after Reset, so
// the webview page draws a placeholder until then.
func Reset() {
	Web = nil
	State = GalleryState{Selected: "button"}
	Demo = DemoState{
		BtnSel:    []string{"left"},
		Tg1:       true,
		Tg3:       true,
		Input:     "Ada Lovelace",
		Textarea:  "一行备注。",
		Number:    3,
		Combobox:  "Go",
		AutoComp:  "mygo",
		Tokens:    []string{"go", "ui"},
		EditText:  "点击重命名",
		SelStr:    "B",
		Multi:     []string{"go"},
		NativeSel: "每周",
		StepperV:  5,
		Buffer:    ui.NewTextBuffer("mygo"),
		Tw:        typewriter.State{Done: true},
		Check:     true,
		Switch:    true,
		Radio:     "小",
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
		Tree2:     true,
		Tree3:     true,
		FindOpen:  true,
		FeedRows:  feedRows(1, 8),
		FeedN:     9,
		ShellSel:  "overview",
		History:   ui.NewRouter("/"),
		Step:      50,
		SplitV:    80,
		Diff: editor.DiffState{
			Old: editor.NewBuffer("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"hello\")\n\tfmt.Println(\"world\")\n}\n"),
			New: editor.NewBuffer("package main\n\nimport (\n\t\"fmt\"\n\t\"time\"\n)\n\nfunc main() {\n\tfmt.Println(\"你好, 画廊\")\n\tfmt.Println(time.Now())\n}\n"),
		},
		OTime:  time.Now(),
		OColor: ui.Color{R: 110, G: 90, B: 230, A: 255},
		Page:   3,
		// The shader starts at its defaults: the sliders sit mid-range and
		// the background follows the pointer.
		ShaderIntensity: 1,
		ShaderScale:     1,
	}
	// The table sorts by the column the visitor last chose; the header
	// clicks set it.
	Demo.TableList.Sort = &Demo.TableSort
	// The onevra pages begin with something worth looking at: the tree
	// with its folders open, the finished turn and its steps expanded
	// rather than folded behind the summary.
	Demo.OTree.Open.Add("src")
	Demo.OTree.Open.Add("docs")
	Demo.OTurn.SetExpanded(true)
	Demo.OGroup.SetExpanded(true)
}

// feedRows is a page of the infinite feed, "第 n 行" upward.
func feedRows(start, n int) []string {
	rows := make([]string, n)
	for i := range rows {
		rows[i] = fmt.Sprintf("第 %d 行", start+i)
	}
	return rows
}

package view

import (
	"fmt"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/autocomplete"
	"github.com/smallnest/uix/components/autosizetextarea"
	"github.com/smallnest/uix/components/breadcrumbs"
	"github.com/smallnest/uix/components/button"
	"github.com/smallnest/uix/components/buttongroup"
	"github.com/smallnest/uix/components/carousel"
	"github.com/smallnest/uix/components/combobox"
	"github.com/smallnest/uix/components/contextmenu"
	"github.com/smallnest/uix/components/dropdown"
	"github.com/smallnest/uix/components/editabletext"
	"github.com/smallnest/uix/components/floatinglabelinput"
	"github.com/smallnest/uix/components/input"
	"github.com/smallnest/uix/components/inputgroup"
	"github.com/smallnest/uix/components/inputotp"
	"github.com/smallnest/uix/components/link"
	"github.com/smallnest/uix/components/menu"
	"github.com/smallnest/uix/components/multiselect"
	"github.com/smallnest/uix/components/nativeselect"
	"github.com/smallnest/uix/components/numberinput"
	"github.com/smallnest/uix/components/pagination"
	"github.com/smallnest/uix/components/searchfield"
	"github.com/smallnest/uix/components/segmented"
	sel "github.com/smallnest/uix/components/select"
	"github.com/smallnest/uix/components/sidebar"
	"github.com/smallnest/uix/components/stepper"
	"github.com/smallnest/uix/components/tabs"
	"github.com/smallnest/uix/components/textarea"
	"github.com/smallnest/uix/components/toggle"
	"github.com/smallnest/uix/components/togglegroup"
	"github.com/smallnest/uix/components/tokenfield"
	"github.com/smallnest/uix/components/toolbar"
	"github.com/smallnest/uix/components/typewriter"
)

// demoButton shows the six faces of the button.
func demoButton(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Wrap().Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "Primary"})
		button.Button(c, button.Props{Label: "Secondary", Variant: button.Secondary})
		button.Button(c, button.Props{Label: "Outline", Variant: button.Outline})
		button.Button(c, button.Props{Label: "Ghost", Variant: button.Ghost})
		button.Button(c, button.Props{Label: "Destructive", Variant: button.Destructive})
		button.Button(c, button.Props{Label: "Disabled", Disabled: true})
	})
}

// demoButtonGroup shows a row of buttons joined as one control.
func demoButtonGroup(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		buttongroup.ButtonGroup(c, buttongroup.Props{
			Options: []buttongroup.Option{
				{Label: "Left", Value: "left"},
				{Label: "Center", Value: "center"},
				{Label: "Right", Value: "right"},
			},
			Selected: &Demo.BtnSel,
		})
		ui.Text(c, "One of the alignment controls is chosen.").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
}

// demoToggle shows two pressed toggles of a toolbar.
func demoToggle(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Gap(t.Space(2)).Children(func() {
		toggle.Toggle(c, toggle.Props{On: &Demo.Tg1, Label: "Bold"})
		toggle.Toggle(c, toggle.Props{On: &Demo.Tg2, Label: "Italic"})
		toggle.Toggle(c, toggle.Props{On: &Demo.Tg3, Label: "Underline"})
	})
}

// demoButtonToggleGroup shows a group of toggles, each with a tip.
func demoButtonToggleGroup(c *ui.Context, t *ui.Theme) {
	togglegroup.ToggleGroup(c, togglegroup.Props{
		Items: []togglegroup.Item{
			{On: &Demo.Tg1, Label: "B", Tip: "Bold"},
			{On: &Demo.Tg2, Label: "I", Tip: "Italic"},
			{On: &Demo.Tg3, Label: "U", Tip: "Underline"},
		},
	})
}

// demoLink shows a link that opens a URL.
func demoLink(c *ui.Context, t *ui.Theme) {
	link.Link(c, link.Props{Label: "The uix repository", URL: "https://github.com/smallnest/uix"})
}

// demoMenu shows a button that opens a menu of actions.
func demoMenu(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		menu.Menu(c, menu.Props{
			Label: "Open menu",
			Items: []menu.Item{
				{Label: "New", Action: func() {}},
				{Label: "Open", Action: func() {}},
				{Label: "Save", Action: func() {}},
			},
		})
		ui.Text(c, "The menu opens on a click.").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
}

// demoDropdown shows a button whose menu drops below it.
func demoDropdown(c *ui.Context, t *ui.Theme) {
	dropdown.Dropdown(c, dropdown.Props{
		Trigger: "Actions",
		Open:    &Demo.Drop,
		Items: []dropdown.Item{
			{Label: "Rename", OnClick: func() {}},
			{Label: "Duplicate", OnClick: func() {}},
			{Label: "Delete", OnClick: func() {}},
		},
	})
}

// demoContextMenu shows a panel whose right-click opens a menu.
func demoContextMenu(c *ui.Context, t *ui.Theme) {
	contextmenu.ContextMenu(c, contextmenu.Props{
		Items: []contextmenu.Item{
			{Label: "Copy", Action: func() {}},
			{Label: "Paste", Action: func() {}},
			{Label: "Cut", Action: func() {}},
		},
	}).FillWidth().Background(t.Surface).Border(1, t.Border).Radius(t.Radius).Padding(t.Space(3)).Children(func() {
		ui.Text(c, "Right-click this panel for the menu.").TextColor(t.TextMuted)
	})
}

// demoBreadcrumbs shows a path of items, the last where you are.
func demoBreadcrumbs(c *ui.Context, t *ui.Theme) {
	breadcrumbs.Breadcrumbs(c, breadcrumbs.Props{Items: []string{"Home", "Docs", "Install"}})
}

// demoTabs shows a tab list with an accent bar on the active tab.
func demoTabs(c *ui.Context, t *ui.Theme) {
	tabs.Tabs(c, tabs.Props{Selected: &Demo.Tabs, Labels: []string{"Overview", "Activity", "Settings"}})
}

// demoNavSidebar shows a navigation rail of sections and items.
func demoNavSidebar(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(200).Background(t.Background).Radius(t.Radius).Border(1, t.Border).Children(func() {
		sel := "desktop"
		sidebar.Sidebar(c, sidebar.Props{
			Selected: &sel,
			Sections: []sidebar.Section{
				{Title: "Favorites", Items: []sidebar.Item{{ID: "recents", Label: "Recents"}, {ID: "desktop", Label: "Desktop"}}},
				{Title: "Places", Items: []sidebar.Item{{ID: "documents", Label: "Documents"}, {ID: "downloads", Label: "Downloads"}}},
			},
		})
	})
}

// demoToolbar shows a row of actions along the top.
func demoToolbar(c *ui.Context, t *ui.Theme) {
	toolbar.Toolbar(c, toolbar.Props{
		Label: "Format",
		Children: func(c *ui.Context) {
			ui.Button(c, "New")
			ui.Button(c, "Open")
			ui.Spacer(c)
			ui.Button(c, "Share")
		},
	})
}

// demoPagination shows the page numbers with the previous and next.
func demoPagination(c *ui.Context, t *ui.Theme) {
	pagination.Pagination(c, pagination.Props{Page: 3, TotalPages: 9})
}

// demoButtonSegmented shows a switch between views, one chosen.
func demoButtonSegmented(c *ui.Context, t *ui.Theme) {
	segmented.Segmented(c, segmented.Props{Selected: &Demo.Seg, Labels: []string{"Hourly", "Daily", "Monthly"}})
}

// demoStepper shows a number with arrows that step it.
func demoStepper(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(160).Gap(t.Space(2)).Children(func() {
		stepper.Stepper(c, stepper.Props{Value: &Demo.StepperV, Min: 0, Max: 10, Step: 1, Label: "Seats"})
	})
}

// demoCarousel shows slides you flip with buttons or dots.
func demoCarousel(c *ui.Context, t *ui.Theme) {
	carousel.Carousel(c, carousel.Props{
		Count: 3,
		Slide: func(c *ui.Context, i int) {
			ui.Box(c).FillWidth().Height(140).Background(t.Accent.Alpha(1 - float32(i)*0.16)).
				Radius(t.Radius * 2).Center().Children(func() {
				ui.Text(c, fmt.Sprintf("Slide %d", i+1)).FontSize(t.FontSize * 1.2).FontWeight(600).TextColor(t.AccentText)
			})
		},
		Height: 140,
	})
}

// demoInput shows a single-line text input.
func demoInput(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		ui.Text(c, "Name").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		input.Input(c, input.Props{Value: &Demo.Input, Placeholder: "Your name"})
	})
}

// demoTextarea shows a multi-line text area.
func demoTextarea(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		ui.Text(c, "Notes").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		textarea.Textarea(c, textarea.Props{Value: &Demo.Textarea, Placeholder: "Write a note", Rows: 4})
	})
}

// demoNumberInput shows a number input with step buttons.
func demoNumberInput(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		ui.Text(c, "Quantity").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		numberinput.NumberInput(c, numberinput.Props{Value: &Demo.Number, Lo: 0, Hi: 10, Step: 1})
	})
}

// demoSearchField shows a search box with a clear button.
func demoSearchField(c *ui.Context, t *ui.Theme) {
	searchfield.SearchField(c, searchfield.Props{Value: &Demo.Search})
}

// demoCombobox shows a searchable drop-down.
func demoCombobox(c *ui.Context, t *ui.Theme) {
	combobox.Combobox(c, combobox.Props{
		Value:   &Demo.Combobox,
		Options: []string{"Go", "Rust", "Zig", "Swift", "Kotlin"},
	})
}

// demoAutocomplete shows a field that completes the text typed.
func demoAutocomplete(c *ui.Context, t *ui.Theme) {
	autocomplete.Autocomplete(c, autocomplete.Props{
		Value:       &Demo.AutoComp,
		Suggestions: []string{"mygo", "mygo/ui", "uix", "onevra"},
		Label:       "Package",
	})
}

// demoTokenField shows a field of tokens as chips.
func demoTokenField(c *ui.Context, t *ui.Theme) {
	tokenfield.TokenField(c, tokenfield.Props{
		Tokens:      &Demo.Tokens,
		Suggestions: []string{"go", "rust", "ui", "native"},
		Label:       "Tags",
	})
}

// demoOTP shows cells that take a code one character each.
func demoOTP(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		inputotp.OTP(c, inputotp.Props{State: &Demo.OTP, Length: 6, OnComplete: func(code string) {}})
		ui.Text(c, "The verification code, cell by cell.").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
}

// demoInputGroup shows a field with a mark at either end.
func demoInputGroup(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		ui.Text(c, "Amount").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		inputgroup.InputGroup(c, inputgroup.Props{
			Value:        &Demo.Amount,
			Placeholder:  "0.00",
			LeadingText:  "$",
			TrailingText: "USD",
		})
	})
}

// demoFloatingLabel shows a field whose label floats to the border.
func demoFloatingLabel(c *ui.Context, t *ui.Theme) {
	floatinglabelinput.FloatingLabelInput(c, floatinglabelinput.Props{Label: "Email", Value: &Demo.Float})
}

// demoAutoSize shows a text area that grows with its text.
func demoAutoSize(c *ui.Context, t *ui.Theme) {
	autosizetextarea.AutosizeTextarea(c, autosizetextarea.Props{
		Value:       &Demo.AutoSize,
		Placeholder: "A note that grows as you type",
		MinRows:     2,
		MaxRows:     6,
	})
}

// demoEditableText shows text that renames in place on a double click.
func demoEditableText(c *ui.Context, t *ui.Theme) {
	editabletext.EditableText(c, editabletext.Props{Value: &Demo.EditText, OnChange: func(string) {}})
}

// demoSelect shows a drop-down with a placeholder.
func demoSelect(c *ui.Context, t *ui.Theme) {
	sel.Select(c, sel.Props{
		Selected:    &Demo.SelStr,
		Options:     []string{"A", "B", "C", "D"},
		Placeholder: "Choose a plan",
	})
}

// demoMultiSelect shows a field choosing several options.
func demoMultiSelect(c *ui.Context, t *ui.Theme) {
	multiselect.Multiselect(c, multiselect.Props{
		Values:  &Demo.Multi,
		Options: []string{"go", "rust", "zig", "swift"},
	})
}

// demoNativeSelect shows a drop-down on the platform's own menu.
func demoNativeSelect(c *ui.Context, t *ui.Theme) {
	nativeselect.NativeSelect(c, nativeselect.Props{
		Selected: &Demo.NativeSel,
		Options: []nativeselect.Option{
			{Label: "Daily", Value: "Daily"},
			{Label: "Weekly", Value: "Weekly"},
			{Label: "Monthly", Value: "Monthly"},
		},
	})
}

// demoTypewriter shows text that types itself out, with a caret.
func demoTypewriter(c *ui.Context, t *ui.Theme) {
	typewriter.Typewriter(c, typewriter.Props{
		Text:  "Typing itself out, a character at a time.",
		State: &Demo.Tw,
	})
}

// demoTextInputBuffer shows a mygo buffer field over indexed storage.
func demoTextInputBuffer(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		ui.Text(c, "Buffer").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		ui.TextInputBuffer(c, Demo.Buffer)
	})
}

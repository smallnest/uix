package view

import (
	"fmt"
	"strings"

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
	"github.com/smallnest/uix/components/findbar"
	"github.com/smallnest/uix/components/floatinglabelinput"
	"github.com/smallnest/uix/components/input"
	"github.com/smallnest/uix/components/inputgroup"
	"github.com/smallnest/uix/components/inputotp"
	"github.com/smallnest/uix/components/interactivestepper"
	"github.com/smallnest/uix/components/link"
	"github.com/smallnest/uix/components/magicback"
	"github.com/smallnest/uix/components/menu"
	"github.com/smallnest/uix/components/menubar"
	"github.com/smallnest/uix/components/multiselect"
	"github.com/smallnest/uix/components/nativeselect"
	"github.com/smallnest/uix/components/navigationmenu"
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
		button.Button(c, button.Props{Label: "主要"})
		button.Button(c, button.Props{Label: "次要", Variant: button.Secondary})
		button.Button(c, button.Props{Label: "描边", Variant: button.Outline})
		button.Button(c, button.Props{Label: "幽灵", Variant: button.Ghost})
		button.Button(c, button.Props{Label: "危险", Variant: button.Destructive})
		button.Button(c, button.Props{Label: "禁用", Disabled: true})
	})
}

// demoButtonGroup shows a row of buttons joined as one control.
func demoButtonGroup(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		buttongroup.ButtonGroup(c, buttongroup.Props{
			Options: []buttongroup.Option{
				{Label: "左", Value: "left"},
				{Label: "中", Value: "center"},
				{Label: "右", Value: "right"},
			},
			Selected: &Demo.BtnSel,
		})
		ui.Text(c, "三个对齐方式中选择一个。").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
}

// demoToggle shows two pressed toggles of a toolbar.
func demoToggle(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Gap(t.Space(2)).Children(func() {
		toggle.Toggle(c, toggle.Props{On: &Demo.Tg1, Label: "加粗"})
		toggle.Toggle(c, toggle.Props{On: &Demo.Tg2, Label: "斜体"})
		toggle.Toggle(c, toggle.Props{On: &Demo.Tg3, Label: "下划线"})
	})
}

// demoButtonToggleGroup shows a group of toggles, each with a tip.
func demoButtonToggleGroup(c *ui.Context, t *ui.Theme) {
	togglegroup.ToggleGroup(c, togglegroup.Props{
		Items: []togglegroup.Item{
			{On: &Demo.Tg1, Label: "B", Tip: "加粗"},
			{On: &Demo.Tg2, Label: "I", Tip: "斜体"},
			{On: &Demo.Tg3, Label: "U", Tip: "下划线"},
		},
	})
}

// demoLink shows a link that opens a URL.
func demoLink(c *ui.Context, t *ui.Theme) {
	link.Link(c, link.Props{Label: "uix 仓库", URL: "https://github.com/smallnest/uix"})
}

// demoMenu shows a button that opens a menu of actions.
func demoMenu(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		menu.Menu(c, menu.Props{
			Label: "打开菜单",
			Items: []menu.Item{
				{Label: "新建", Action: func() {}},
				{Label: "打开", Action: func() {}},
				{Label: "保存", Action: func() {}},
			},
		})
		ui.Text(c, "点击按钮打开菜单。").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
}

// demoDropdown shows a button whose menu drops below it.
func demoDropdown(c *ui.Context, t *ui.Theme) {
	dropdown.Dropdown(c, dropdown.Props{
		Trigger: "操作",
		Open:    &Demo.Drop,
		Items: []dropdown.Item{
			{Label: "重命名", OnClick: func() {}},
			{Label: "复制", OnClick: func() {}},
			{Label: "删除", OnClick: func() {}},
		},
	})
}

// demoContextMenu shows a panel whose right-click opens a menu.
func demoContextMenu(c *ui.Context, t *ui.Theme) {
	contextmenu.ContextMenu(c, contextmenu.Props{
		Items: []contextmenu.Item{
			{Label: "复制", Action: func() {}},
			{Label: "粘贴", Action: func() {}},
			{Label: "剪切", Action: func() {}},
		},
	}).FillWidth().Background(t.Surface).Border(1, t.Border).Radius(t.Radius).Padding(t.Space(3)).Children(func() {
		ui.Text(c, "在面板上点击右键打开菜单。").TextColor(t.TextMuted)
	})
}

// demoBreadcrumbs shows a path of items, the last where you are.
func demoBreadcrumbs(c *ui.Context, t *ui.Theme) {
	breadcrumbs.Breadcrumbs(c, breadcrumbs.Props{Items: []string{"首页", "文档", "安装"}})
}

// demoTabs shows a tab list with an accent bar on the active tab.
func demoTabs(c *ui.Context, t *ui.Theme) {
	tabs.Tabs(c, tabs.Props{Selected: &Demo.Tabs, Labels: []string{"概览", "动态", "设置"}})
}

// demoNavSidebar shows a navigation rail of sections and items.
func demoNavSidebar(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(200).Background(t.Background).Radius(t.Radius).Border(1, t.Border).Children(func() {
		sel := "desktop"
		sidebar.Sidebar(c, sidebar.Props{
			Selected: &sel,
			Sections: []sidebar.Section{
				{Title: "收藏", Items: []sidebar.Item{{ID: "recents", Label: "最近"}, {ID: "desktop", Label: "桌面"}}},
				{Title: "位置", Items: []sidebar.Item{{ID: "documents", Label: "文稿"}, {ID: "downloads", Label: "下载"}}},
			},
		})
	})
}

// demoToolbar shows a row of actions along the top.
func demoToolbar(c *ui.Context, t *ui.Theme) {
	toolbar.Toolbar(c, toolbar.Props{
		Label: "格式",
		Children: func(c *ui.Context) {
			ui.Button(c, "新建")
			ui.Button(c, "打开")
			ui.Spacer(c)
			ui.Button(c, "分享")
		},
	})
}

// demoPagination shows the page numbers with the previous and next.
func demoPagination(c *ui.Context, t *ui.Theme) {
	pagination.Pagination(c, pagination.Props{Page: 3, TotalPages: 9})
}

// demoButtonSegmented shows a switch between views, one chosen.
func demoButtonSegmented(c *ui.Context, t *ui.Theme) {
	segmented.Segmented(c, segmented.Props{Selected: &Demo.Seg, Labels: []string{"按小时", "按天", "按月"}})
}

// demoStepper shows a number with arrows that step it.
func demoStepper(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(160).Gap(t.Space(2)).Children(func() {
		stepper.Stepper(c, stepper.Props{Value: &Demo.StepperV, Min: 0, Max: 10, Step: 1, Label: "座位"})
	})
}

// demoCarousel shows slides you flip with buttons or dots.
func demoCarousel(c *ui.Context, t *ui.Theme) {
	carousel.Carousel(c, carousel.Props{
		Count:  3,
		Scroll: &Demo.CarouselScroll,
		State:  &Demo.Carousel,
		Slide: func(c *ui.Context, i int) {
			ui.Box(c).FillWidth().Height(140).Background(t.Accent.Alpha(1 - float32(i)*0.16)).
				Radius(t.Radius * 2).Center().Children(func() {
				ui.Text(c, fmt.Sprintf("第 %d 张", i+1)).FontSize(t.FontSize * 1.2).FontWeight(600).TextColor(t.AccentText)
			})
		},
		Height: 140,
	})
}

// demoInput shows a single-line text input.
func demoInput(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		ui.Text(c, "姓名").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		input.Input(c, input.Props{Value: &Demo.Input, Placeholder: "你的名字"})
	})
}

// demoTextarea shows a multi-line text area.
func demoTextarea(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		ui.Text(c, "备注").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		textarea.Textarea(c, textarea.Props{Value: &Demo.Textarea, Placeholder: "写一条备注", Rows: 4})
	})
}

// demoNumberInput shows a number input with step buttons.
func demoNumberInput(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		ui.Text(c, "数量").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		numberinput.NumberInput(c, numberinput.Props{Value: &Demo.Number, Lo: 0, Hi: 10, Step: 1})
	})
}

// demoSearchField shows a search box with a clear button and a
// placeholder in Chinese.
func demoSearchField(c *ui.Context, t *ui.Theme) {
	searchfield.SearchField(c, searchfield.Props{Value: &Demo.Search, Placeholder: "搜索组件…"})
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
		Label:       "包名",
	})
}

// demoTokenField shows a field of tokens as chips.
func demoTokenField(c *ui.Context, t *ui.Theme) {
	tokenfield.TokenField(c, tokenfield.Props{
		Tokens:      &Demo.Tokens,
		Suggestions: []string{"go", "rust", "ui", "native"},
		Label:       "标签",
	})
}

// demoOTP shows cells that take a code one character each.
func demoOTP(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		inputotp.OTP(c, inputotp.Props{State: &Demo.OTP, Length: 6, OnComplete: func(code string) {}})
		ui.Text(c, "验证码，一格一位。").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
}

// demoInputGroup shows a field with a mark at either end.
func demoInputGroup(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		ui.Text(c, "金额").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
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
	floatinglabelinput.FloatingLabelInput(c, floatinglabelinput.Props{Label: "邮箱", Value: &Demo.Float})
}

// demoAutoSize shows a text area that grows with its text.
func demoAutoSize(c *ui.Context, t *ui.Theme) {
	autosizetextarea.AutosizeTextarea(c, autosizetextarea.Props{
		Value:       &Demo.AutoSize,
		Placeholder: "输入时自动增高的备注",
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
		Placeholder: "选择一个方案",
	})
}

// demoMultiSelect shows a field choosing several options.
func demoMultiSelect(c *ui.Context, t *ui.Theme) {
	multiselect.Multiselect(c, multiselect.Props{
		Values:  &Demo.Multi,
		Options: []string{"go", "rust", "zig", "swift"},
		Open:    &Demo.MultiOpen,
	})
}

// demoNativeSelect shows a drop-down on the platform's own menu.
func demoNativeSelect(c *ui.Context, t *ui.Theme) {
	nativeselect.NativeSelect(c, nativeselect.Props{
		Selected: &Demo.NativeSel,
		Options: []nativeselect.Option{
			{Label: "每天", Value: "每天"},
			{Label: "每周", Value: "每周"},
			{Label: "每月", Value: "每月"},
		},
	})
}

// demoTypewriter shows text that types itself out, with a caret.
func demoTypewriter(c *ui.Context, t *ui.Theme) {
	typewriter.Typewriter(c, typewriter.Props{
		Text:  "一个字符一个字符地打出来。",
		State: &Demo.Tw,
	})
}

// demoTextInputBuffer shows a mygo buffer field over indexed storage.
func demoTextInputBuffer(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		ui.Text(c, "缓冲区").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		ui.TextInputBuffer(c, Demo.Buffer)
	})
}

// demoMenubar shows the flat bar of menus of a desktop app.
func demoMenubar(c *ui.Context, t *ui.Theme) {
	menubar.Menubar(c, menubar.Props{
		Menus: []menubar.Menu{
			{Label: "文件", Items: []menubar.Item{
				{Label: "新建", ShortcutMods: ui.Cmd, ShortcutKey: ui.KeyN, Action: func() {}},
				{Separator: true},
				{Label: "保存", ShortcutMods: ui.Cmd, ShortcutKey: ui.KeyS, Action: func() {}},
				{Label: "退出", Action: func() {}},
			}},
			{Label: "编辑", Items: []menubar.Item{
				{Label: "撤销", ShortcutMods: ui.Cmd, ShortcutKey: ui.KeyZ, Action: func() {}},
				{Label: "加粗", Checked: true, Action: func() {}},
			}},
		},
	})
}

// demoNavigationMenu shows a bar of links, of which one opens a panel.
func demoNavigationMenu(c *ui.Context, t *ui.Theme) {
	navigationmenu.NavigationMenu(c, navigationmenu.Props{
		State: &Demo.Nav,
		Items: []navigationmenu.Item{
			{Label: "文档", Children: []navigationmenu.Item{
				{Label: "入门", OnClick: func() {}},
				{Label: "指南", OnClick: func() {}},
				{Label: "参考", OnClick: func() {}},
			}},
			{Label: "关于", OnClick: func() {}},
		},
	})
}

// demoMagicBack shows the back button of a browser, home when there is
// nothing to go back to.
func demoMagicBack(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(200).Gap(t.Space(2)).Children(func() {
		magicback.MagicBack(c, magicback.Props{
			Back:   true,
			Label:  "返回",
			OnBack: func() {},
			OnHome: func() {},
		})
		ui.Text(c, "有去处的返回按钮。").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
}

// demoInteractiveStepper shows a form of steps, each with its content.
func demoInteractiveStepper(c *ui.Context, t *ui.Theme) {
	interactivestepper.InteractiveStepper(c, interactivestepper.Props{
		Step: &Demo.Istep,
		Steps: []interactivestepper.Step{
			{Title: "信息", Body: func(c *ui.Context) { ui.Text(c, "填写账户信息。") }},
			{Title: "确认", Body: func(c *ui.Context) { ui.Text(c, "确认信息无误。") }},
			{Title: "完成", Body: func(c *ui.Context) { ui.Text(c, "提交完成。") }},
		},
		OnComplete: func() {},
	})
}

// demoBackButton shows mygo's back and forward buttons over a history:
// the buttons move through it, disabled at its ends, and the path shows
// where the history is.
func demoBackButton(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(320).Gap(t.Space(2)).Children(func() {
		ui.Row(c).Gap(t.Space(2)).Children(func() {
			ui.BackButton(c, Demo.History)
			ui.ForwardButton(c, Demo.History)
		})
		ui.Text(c, "当前路径: "+Demo.History.Path()).FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		ui.Row(c).Gap(t.Space(2)).Children(func() {
			button.Button(c, button.Props{Label: "访问 文档", OnClick: func() { Demo.History.Push("/docs") }})
			button.Button(c, button.Props{Label: "访问 设置", OnClick: func() { Demo.History.Push("/settings") }})
		})
	})
}

// demoFindBar shows a bar that finds a word in a short document; the
// gallery counts the matches of the query in the sample text.
func demoFindBar(c *ui.Context, t *ui.Theme) {
	const sample = "uix 是一个用纯 Go 写的组件库。组件库提供按钮、输入框和对话框等组件。"
	matches := 0
	if q := Demo.FindQuery; q != "" {
		matches = strings.Count(strings.ToLower(sample), strings.ToLower(q))
	}
	ui.Column(c).Width(440).Gap(t.Space(2)).Children(func() {
		ui.Text(c, sample)
		findbar.FindBar(c, findbar.Props{
			Open: &Demo.FindOpen, Query: &Demo.FindQuery,
			Matches: matches, Current: &Demo.FindCurrent,
		})
	})
}

package view

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/lucide"
)

// Library names the library a component comes from. The gallery keeps
// one page per role, by priority: uix's own component when there is one,
// else mygo's primitive, else onevra's. The MujicaUI pages are theirs
// alone, ported whole, so they never dedupe.
type Library int

const (
	Uix Library = iota
	Mygo
	Onevra
	Mujica
)

// Entry is one component of the gallery: its name, the library it comes
// from, and how its page draws.
type Entry struct {
	// ID is the sidebar item's ID, which the state holds.
	ID string
	// Name is the component's name, shown in the sidebar and heading.
	Name string
	// Lib is the library the page shows the component of.
	Lib Library
	// Icon is the lucide icon of the item, shown beside its name in the
	// sidebar; nil draws none.
	Icon func() *ui.SVG
	// Build draws the page's demo.
	Build func(c *ui.Context, t *ui.Theme)
}

// Section is a category of the sidebar, holding its entries.
type Section struct {
	// Title is the category, such as "文本输入".
	Title string
	// Entries are the components of the category, in order.
	Entries []Entry
}

// catalog is the gallery: every component of the three libraries, one
// page each, the duplicates resolved by priority, then MujicaUI's whole
// catalog as a fourth library.
var catalog = []Section{
	{
		Title: "操作与导航",
		Entries: []Entry{
			{"button", "Button", Uix, lucide.MousePointerClick, demoButton},
			{"buttongroup", "ButtonGroup", Uix, lucide.Grip, demoButtonGroup},
			{"toggle", "Toggle", Uix, lucide.ToggleRight, demoToggle},
			{"togglegroup", "ToggleGroup", Uix, lucide.ToggleLeft, demoButtonToggleGroup},
			{"link", "Link", Uix, lucide.Link, demoLink},
			{"menu", "Menu", Uix, lucide.Menu, demoMenu},
			{"dropdown", "Dropdown", Uix, lucide.ChevronDown, demoDropdown},
			{"contextmenu", "ContextMenu", Uix, lucide.MousePointer2, demoContextMenu},
			{"breadcrumbs", "Breadcrumbs", Uix, lucide.ChevronsRight, demoBreadcrumbs},
			{"tabs", "Tabs", Uix, lucide.PanelsTopLeft, demoTabs},
			{"sidebar", "Sidebar", Uix, lucide.PanelLeft, demoNavSidebar},
			{"toolbar", "Toolbar", Uix, lucide.PanelTop, demoToolbar},
			{"pagination", "Pagination", Uix, lucide.ListOrdered, demoPagination},
			{"segmented", "Segmented", Uix, lucide.SlidersHorizontal, demoButtonSegmented},
			{"stepper", "Stepper", Uix, lucide.ArrowUpDown, demoStepper},
			{"carousel", "Carousel", Uix, lucide.Rows3, demoCarousel},
			{"menubar", "Menubar", Uix, lucide.AppWindow, demoMenubar},
			{"navigationmenu", "NavigationMenu", Uix, lucide.Navigation, demoNavigationMenu},
			{"magicback", "MagicBack", Uix, lucide.Undo2, demoMagicBack},
			{"interactivestepper", "InteractiveStepper", Uix, lucide.Footprints, demoInteractiveStepper},
			{"backbutton", "BackButton", Mygo, lucide.CornerUpLeft, demoBackButton},
			{"holdtoconfirm", "HoldToConfirm", Onevra, lucide.Hand, demoHoldToConfirm},
			{"floatingbuttongroup", "FloatingButtonGroup", Onevra, lucide.Boxes, demoFloatingButtonGroup},
		},
	},
	{
		Title: "文本输入",
		Entries: []Entry{
			{"input", "Input", Uix, lucide.SquarePen, demoInput},
			{"textarea", "Textarea", Uix, lucide.Type, demoTextarea},
			{"numberinput", "NumberInput", Uix, lucide.Hash, demoNumberInput},
			{"searchfield", "SearchField", Uix, lucide.Search, demoSearchField},
			{"combobox", "Combobox", Uix, lucide.ChevronsUpDown, demoCombobox},
			{"autocomplete", "Autocomplete", Uix, lucide.Sparkles, demoAutocomplete},
			{"tokenfield", "TokenField", Uix, lucide.Tags, demoTokenField},
			{"inputotp", "OTP Input", Uix, lucide.KeyRound, demoOTP},
			{"inputgroup", "InputGroup", Uix, lucide.Brackets, demoInputGroup},
			{"floatinglabelinput", "FloatingLabelInput", Uix, lucide.ArrowUp, demoFloatingLabel},
			{"autosizetextarea", "AutoSizeTextarea", Uix, lucide.MoveVertical, demoAutoSize},
			{"editabletext", "EditableText", Uix, lucide.Pencil, demoEditableText},
			{"select", "Select", Uix, lucide.ListFilter, demoSelect},
			{"multiselect", "MultiSelect", Uix, lucide.ListChecks, demoMultiSelect},
			{"nativeselect", "NativeSelect", Uix, lucide.Monitor, demoNativeSelect},
			{"typewriter", "Typewriter", Uix, lucide.Keyboard, demoTypewriter},
			{"textinputbuffer", "TextInputBuffer", Mygo, lucide.TextCursor, demoTextInputBuffer},
			{"findbar", "FindBar", Uix, lucide.ScanSearch, demoFindBar},
		},
	},
	{
		Title: "选择与取值",
		Entries: []Entry{
			{"checkbox", "Checkbox", Uix, lucide.CircleCheck, demoCheckbox},
			{"checkboxgroup", "CheckboxGroup", Uix, lucide.CheckCheck, demoCheckboxGroup},
			{"switch", "Switch", Uix, lucide.ToggleLeft, demoSwitch},
			{"radio", "Radio", Uix, lucide.Radio, demoRadio},
			{"slider", "Slider", Uix, lucide.Volume2, demoSlider},
			{"rangeslider", "RangeSlider", Uix, lucide.MoveHorizontal, demoRangeSlider},
			{"rating", "Rating", Uix, lucide.Star, demoRating},
			{"meter", "Meter", Uix, lucide.Gauge, demoMeter},
			{"stepslider", "StepSlider", Mygo, lucide.ListMusic, demoStepSlider},
			{"cardstack", "CardStack", Onevra, lucide.Layers2, demoCardStack},
		},
	},
	{
		Title: "日期与时间",
		Entries: []Entry{
			{"datepicker", "DatePicker", Uix, lucide.CalendarDays, demoDatePicker},
			{"timeinput", "TimeInput", Uix, lucide.Clock, demoTimeInput},
			{"timepicker", "TimePicker", Onevra, lucide.AlarmClock, demoTimePicker},
			{"datetimepicker", "DateTimePicker", Uix, lucide.CalendarClock, demoDateTimePicker},
			{"calendar", "Calendar", Uix, lucide.Calendar, demoCalendar},
		},
	},
	{
		Title: "颜色",
		Entries: []Entry{
			{"colorpicker", "ColorPicker", Uix, lucide.Palette, demoColorPicker},
			{"colorwell", "ColorWell", Uix, lucide.SwatchBook, demoColorWell},
			{"colorview", "ColorView", Onevra, lucide.Eye, demoColorView},
		},
	},
	{
		Title: "内容展示",
		Entries: []Entry{
			{"text", "Text", Mygo, lucide.CaseSensitive, demoText},
			{"box", "Box", Mygo, lucide.Box, demoBox},
			{"row", "Row", Mygo, lucide.Rows2, demoRow},
			{"column", "Column", Mygo, lucide.Columns3, demoColumn},
			{"spacer", "Spacer", Mygo, lucide.Expand, demoSpacer},
			{"icon", "Icon", Uix, lucide.Sparkle, demoIcon},
			{"lucide", "Lucide", Uix, lucide.Shapes, demoLucide},
			{"image", "Image", Uix, lucide.Image, demoImage},
			{"avatar", "Avatar", Uix, lucide.User, demoAvatar},
			{"badge", "Badge", Uix, lucide.BadgeCheck, demoBadge},
			{"kbd", "Kbd", Uix, lucide.Command, demoKbd},
			{"label", "Label", Uix, lucide.Tag, demoLabel},
			{"marker", "Marker", Uix, lucide.MapPin, demoMarker},
			{"inlinecode", "InlineCode", Uix, lucide.Braces, demoInlineCode},
			{"blockquote", "Blockquote", Uix, lucide.Quote, demoBlockquote},
			{"headingwithanchor", "HeadingWithAnchor", Uix, lucide.Heading1, demoHeadingWithAnchor},
			{"richtext", "RichText", Uix, lucide.TextQuote, demoRichText},
			{"card", "Card", Uix, lucide.CreditCard, demoCard},
			{"statcard", "StatCard", Uix, lucide.Activity, demoStatCard},
			{"skeleton", "Skeleton", Uix, lucide.RectangleHorizontal, demoSkeleton},
			{"empty", "Empty", Uix, lucide.Inbox, demoEmpty},
			{"list", "List", Uix, lucide.List, demoList},
			{"gridview", "GridView", Uix, lucide.LayoutGrid, demoGridView},
			{"table", "Table", Uix, lucide.Table, demoTable},
			{"tablevirtualized", "TableVirtualized", Onevra, lucide.Table2, demoTableVirtualized},
			{"tree", "Tree", Uix, lucide.ListTree, demoTree},
			{"treevirtualized", "TreeVirtualized", Onevra, lucide.FolderTree, demoTreeVirtualized},
			{"outline", "Outline", Mygo, lucide.GitBranch, demoOutline},
			{"gauge", "Gauge", Uix, lucide.Gauge, demoGauge},
			{"linechart", "LineChart", Uix, lucide.ChartLine, demoLineChart},
			{"barchart", "BarChart", Uix, lucide.ChartColumn, demoBarChart},
			{"donutchart", "DonutChart", Onevra, lucide.ChartPie, demoDonutChart},
			{"area", "Area", Onevra, lucide.ChartArea, demoAreaChart},
			{"stackedbar", "StackedBar", Onevra, lucide.ChartColumnBig, demoStackedBar},
			{"graph", "Graph", Onevra, lucide.Network, demoGraph},
			{"pixelgrid", "PixelGrid", Onevra, lucide.Grid3x3, demoPixelGrid},
			{"signalfield", "SignalField", Onevra, lucide.Sparkles, demoSignalField},
			{"animate", "Animate", Uix, lucide.Clapperboard, demoAnimate},
			{"infinitescroll", "InfiniteScroll", Uix, lucide.Infinity, demoInfiniteScroll},
			{"item", "Item", Uix, lucide.Layers, demoItem},
			{"webview", "WebView", Uix, lucide.Globe, demoWebView},
			{"motion", "Motion", Mygo, lucide.Film, demoMotion},
			{"chart", "Chart", Mygo, lucide.ChartScatter, demoChart},
			{"outlinetable", "OutlineTable", Mygo, lucide.TableProperties, demoOutlineTable},
			{"avatargroup", "AvatarGroup", Onevra, lucide.Users, demoAvatarGroup},
			{"imagecompare", "ImageCompare", Onevra, lucide.Images, demoImageCompare},
			{"treemap", "Treemap", Onevra, lucide.Blocks, demoTreemap},
			{"flow", "Flow", Onevra, lucide.Workflow, demoFlow},
		},
	},
	{
		Title: "反馈与浮层",
		Entries: []Entry{
			{"dialog", "Dialog", Uix, lucide.Square, demoDialog},
			{"alertdialog", "AlertDialog", Uix, lucide.TriangleAlert, demoAlertDialog},
			{"drawer", "Drawer", Uix, lucide.ArrowUpFromLine, demoDrawer},
			{"sheet", "Sheet", Uix, lucide.PanelBottom, demoSheet},
			{"popover", "Popover", Uix, lucide.MessageSquare, demoPopover},
			{"tooltip", "Tooltip", Uix, lucide.Info, demoTooltip},
			{"toast", "Toast", Uix, lucide.Bell, demoToast},
			{"spinner", "Spinner", Uix, lucide.LoaderCircle, demoSpinner},
			{"progress", "Progress", Uix, lucide.CircleDot, demoProgress},
			{"progressring", "ProgressRing", Onevra, lucide.RotateCw, demoProgressRing},
			{"command", "Command", Uix, lucide.Terminal, demoCommand},
			{"alert", "Alert", Uix, lucide.CircleAlert, demoAlert},
			{"loadingbutton", "LoadingButton", Uix, lucide.Loader, demoLoadingButton},
			{"notificationcenter", "NotificationCenter", Uix, lucide.BellRing, demoNotificationCenter},
			{"responsivemodal", "ResponsiveModal", Uix, lucide.Maximize2, demoResponsiveModal},
		},
	},
	{
		Title: "内容与代码",
		Entries: []Entry{
			{"code", "Code", Uix, lucide.CodeXml, demoCode},
			{"markdown", "Markdown", Uix, lucide.FileText, demoMarkdown},
			{"codeeditor", "CodeEditor", Mygo, lucide.SquareCode, demoCodeEditor},
			{"chat", "Chat", Uix, lucide.MessagesSquare, demoChat},
			{"composer", "Composer", Uix, lucide.Send, demoComposer},
			{"message", "Message", Uix, lucide.MessageSquareText, demoMessage},
			{"attachment", "Attachment", Uix, lucide.Paperclip, demoAttachment},
			{"thinking", "Thinking", Uix, lucide.Brain, demoThinking},
			{"log", "Log", Uix, lucide.ListStart, demoLog},
			{"questionnaire", "Questionnaire", Uix, lucide.ClipboardList, demoQuestionnaire},
			{"settings", "Settings", Uix, lucide.Settings, demoSettings},
			{"bubble", "Bubble", Uix, lucide.MessageCircle, demoBubble},
			{"messagescroller", "MessageScroller", Uix, lucide.ScrollText, demoMessageScroller},
			{"diffview", "DiffView", Mygo, lucide.FileDiff, demoDiffView},
			{"resultsummary", "ResultSummary", Onevra, lucide.Summary, demoResultSummary},
			{"glasscomposer", "GlassComposer", Onevra, lucide.GlassWater, demoGlassComposer},
			{"conversation", "Conversation", Onevra, lucide.Speech, demoConversation},
		},
	},
	{
		Title: "布局与容器",
		Entries: []Entry{
			{"separator", "Separator", Uix, lucide.SeparatorHorizontal, demoSeparator},
			{"field", "Field", Uix, lucide.RectangleHorizontal, demoField},
			{"fieldset", "Fieldset", Uix, lucide.Group, demoFieldset},
			{"form", "Form", Uix, lucide.FileStack, demoForm},
			{"grid", "Grid", Uix, lucide.Grid2x2, demoGrid},
			{"scroll", "Scroll", Uix, lucide.Scroll, demoScroll},
			{"scrollhorizontal", "ScrollHorizontal", Uix, lucide.MoveRight, demoScrollHorizontal},
			{"scrollboth", "ScrollBoth", Uix, lucide.Move, demoScrollBoth},
			{"split", "Split", Uix, lucide.Spline, demoSplit},
			{"aspectratio", "AspectRatio", Uix, lucide.Ratio, demoAspectRatio},
			{"collapsible", "Collapsible", Uix, lucide.Minus, demoCollapsible},
			{"accordion", "Accordion", Uix, lucide.Plus, demoAccordion},
			{"fileupload", "FileUpload", Uix, lucide.FileUp, demoFileUpload},
			{"canvas", "Canvas", Mygo, lucide.PenTool, demoCanvas},
			{"splitvertical", "SplitVertical", Mygo, lucide.AlignVerticalJustifyStart, demoSplitVertical},
			{"appshell", "AppShell", Uix, lucide.LayoutPanelTop, demoAppShell},
			{"scrolledgebar", "ScrollEdgeBar", Onevra, lucide.PanelTopOpen, demoScrollEdgeBar},
		},
	},
}

// The MujicaUI pages follow the three libraries as a fourth: the ported
// catalog becomes one gallery section per MujicaUI category.
func init() {
	catalog = append(catalog, mujicaSections()...)
}

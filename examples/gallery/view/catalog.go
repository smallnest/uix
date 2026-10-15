package view

import "github.com/egoist/mygo/ui"

// Library names the library a component comes from. The gallery keeps
// one page per role, by priority: uix's own component when there is one,
// else mygo's primitive, else onevra's.
type Library int

const (
	Uix Library = iota
	Mygo
	Onevra
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
	// Build draws the page's demo.
	Build func(c *ui.Context, t *ui.Theme)
}

// Section is a category of the sidebar, holding its entries.
type Section struct {
	// Title is the category, such as "Text input".
	Title string
	// Entries are the components of the category, in order.
	Entries []Entry
}

// catalog is the gallery: every component of the three libraries, one
// page each, the duplicates resolved by priority.
var catalog = []Section{
	{
		Title: "Actions & navigation",
		Entries: []Entry{
			{"button", "Button", Uix, demoButton},
			{"buttongroup", "ButtonGroup", Uix, demoButtonGroup},
			{"toggle", "Toggle", Uix, demoToggle},
			{"togglegroup", "ToggleGroup", Uix, demoButtonToggleGroup},
			{"link", "Link", Uix, demoLink},
			{"menu", "Menu", Uix, demoMenu},
			{"dropdown", "Dropdown", Uix, demoDropdown},
			{"contextmenu", "ContextMenu", Uix, demoContextMenu},
			{"breadcrumbs", "Breadcrumbs", Uix, demoBreadcrumbs},
			{"tabs", "Tabs", Uix, demoTabs},
			{"sidebar", "Sidebar", Uix, demoNavSidebar},
			{"toolbar", "Toolbar", Uix, demoToolbar},
			{"pagination", "Pagination", Uix, demoPagination},
			{"segmented", "Segmented", Uix, demoButtonSegmented},
			{"stepper", "Stepper", Uix, demoStepper},
			{"carousel", "Carousel", Uix, demoCarousel},
		},
	},
	{
		Title: "Text input",
		Entries: []Entry{
			{"input", "Input", Uix, demoInput},
			{"textarea", "Textarea", Uix, demoTextarea},
			{"numberinput", "NumberInput", Uix, demoNumberInput},
			{"searchfield", "SearchField", Uix, demoSearchField},
			{"combobox", "Combobox", Uix, demoCombobox},
			{"autocomplete", "Autocomplete", Uix, demoAutocomplete},
			{"tokenfield", "TokenField", Uix, demoTokenField},
			{"inputotp", "OTP Input", Uix, demoOTP},
			{"inputgroup", "InputGroup", Uix, demoInputGroup},
			{"floatinglabelinput", "FloatingLabelInput", Uix, demoFloatingLabel},
			{"autosizetextarea", "AutoSizeTextarea", Uix, demoAutoSize},
			{"editabletext", "EditableText", Uix, demoEditableText},
			{"select", "Select", Uix, demoSelect},
			{"multiselect", "MultiSelect", Uix, demoMultiSelect},
			{"nativeselect", "NativeSelect", Uix, demoNativeSelect},
			{"typewriter", "Typewriter", Uix, demoTypewriter},
			{"textinputbuffer", "TextInputBuffer", Mygo, demoTextInputBuffer},
		},
	},
	{
		Title: "Selection & values",
		Entries: []Entry{
			{"checkbox", "Checkbox", Uix, demoCheckbox},
			{"checkboxgroup", "CheckboxGroup", Uix, demoCheckboxGroup},
			{"switch", "Switch", Uix, demoSwitch},
			{"radio", "Radio", Uix, demoRadio},
			{"slider", "Slider", Uix, demoSlider},
			{"rangeslider", "RangeSlider", Uix, demoRangeSlider},
			{"rating", "Rating", Uix, demoRating},
			{"meter", "Meter", Uix, demoMeter},
		},
	},
	{
		Title: "Date & time",
		Entries: []Entry{
			{"datepicker", "DatePicker", Uix, demoDatePicker},
			{"timeinput", "TimeInput", Uix, demoTimeInput},
			{"datetimepicker", "DateTimePicker", Uix, demoDateTimePicker},
			{"calendar", "Calendar", Uix, demoCalendar},
		},
	},
	{
		Title: "Color",
		Entries: []Entry{
			{"colorpicker", "ColorPicker", Uix, demoColorPicker},
			{"colorwell", "ColorWell", Uix, demoColorWell},
		},
	},
	{
		Title: "Content display",
		Entries: []Entry{
			{"text", "Text", Mygo, demoText},
			{"box", "Box", Mygo, demoBox},
			{"row", "Row", Mygo, demoRow},
			{"column", "Column", Mygo, demoColumn},
			{"spacer", "Spacer", Mygo, demoSpacer},
			{"icon", "Icon", Uix, demoIcon},
			{"lucide", "Lucide", Uix, demoLucide},
			{"image", "Image", Uix, demoImage},
			{"avatar", "Avatar", Uix, demoAvatar},
			{"badge", "Badge", Uix, demoBadge},
			{"kbd", "Kbd", Uix, demoKbd},
			{"label", "Label", Uix, demoLabel},
			{"marker", "Marker", Uix, demoMarker},
			{"inlinecode", "InlineCode", Uix, demoInlineCode},
			{"blockquote", "Blockquote", Uix, demoBlockquote},
			{"headingwithanchor", "HeadingWithAnchor", Uix, demoHeadingWithAnchor},
			{"richtext", "RichText", Uix, demoRichText},
			{"card", "Card", Uix, demoCard},
			{"statcard", "StatCard", Uix, demoStatCard},
			{"skeleton", "Skeleton", Uix, demoSkeleton},
			{"empty", "Empty", Uix, demoEmpty},
			{"list", "List", Uix, demoList},
			{"gridview", "GridView", Uix, demoGridView},
			{"table", "Table", Uix, demoTable},
			{"tree", "Tree", Uix, demoTree},
			{"outline", "Outline", Mygo, demoOutline},
			{"linechart", "LineChart", Uix, demoLineChart},
			{"barchart", "BarChart", Uix, demoBarChart},
			{"donutchart", "DonutChart", Onevra, demoDonutChart},
			{"graph", "Graph", Onevra, demoGraph},
			{"pixelgrid", "PixelGrid", Onevra, demoPixelGrid},
		},
	},
	{
		Title: "Feedback & overlays",
		Entries: []Entry{
			{"dialog", "Dialog", Uix, demoDialog},
			{"alertdialog", "AlertDialog", Uix, demoAlertDialog},
			{"drawer", "Drawer", Uix, demoDrawer},
			{"sheet", "Sheet", Uix, demoSheet},
			{"popover", "Popover", Uix, demoPopover},
			{"tooltip", "Tooltip", Uix, demoTooltip},
			{"toast", "Toast", Uix, demoToast},
			{"spinner", "Spinner", Uix, demoSpinner},
			{"progress", "Progress", Uix, demoProgress},
			{"command", "Command", Uix, demoCommand},
			{"alert", "Alert", Uix, demoAlert},
			{"loadingbutton", "LoadingButton", Uix, demoLoadingButton},
			{"notificationcenter", "NotificationCenter", Uix, demoNotificationCenter},
		},
	},
	{
		Title: "Content & code",
		Entries: []Entry{
			{"code", "Code", Uix, demoCode},
			{"markdown", "Markdown", Uix, demoMarkdown},
			{"codeeditor", "CodeEditor", Mygo, demoCodeEditor},
			{"chat", "Chat", Uix, demoChat},
			{"composer", "Composer", Uix, demoComposer},
			{"message", "Message", Uix, demoMessage},
			{"attachment", "Attachment", Uix, demoAttachment},
			{"thinking", "Thinking", Uix, demoThinking},
			{"log", "Log", Uix, demoLog},
			{"questionnaire", "Questionnaire", Uix, demoQuestionnaire},
			{"settings", "Settings", Uix, demoSettings},
		},
	},
	{
		Title: "Layout & containers",
		Entries: []Entry{
			{"separator", "Separator", Uix, demoSeparator},
			{"field", "Field", Uix, demoField},
			{"fieldset", "Fieldset", Uix, demoFieldset},
			{"form", "Form", Uix, demoForm},
			{"grid", "Grid", Uix, demoGrid},
			{"scroll", "Scroll", Uix, demoScroll},
			{"scrollhorizontal", "ScrollHorizontal", Uix, demoScrollHorizontal},
			{"scrollboth", "ScrollBoth", Uix, demoScrollBoth},
			{"split", "Split", Uix, demoSplit},
			{"aspectratio", "AspectRatio", Uix, demoAspectRatio},
			{"collapsible", "Collapsible", Uix, demoCollapsible},
			{"accordion", "Accordion", Uix, demoAccordion},
			{"fileupload", "FileUpload", Uix, demoFileUpload},
			{"canvas", "Canvas", Mygo, demoCanvas},
		},
	},
}

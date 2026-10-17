package view

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/lucide"
	"github.com/smallnest/uix/examples/gallery/mujica"
)

// mujicaIcons picks a lucide icon for a MujicaUI component, where the
// category's icon does not fit; the rest fall back to their category.
var mujicaIcons = map[string]func() *ui.SVG{
	"Text":        lucide.CaseSensitive,
	"Heading":     lucide.Heading1,
	"Image":       lucide.Image,
	"Avatar":      lucide.User,
	"Badge":       lucide.BadgeCheck,
	"Kbd":         lucide.Command,
	"Button":      lucide.MousePointerClick,
	"IconButton":  lucide.Sparkle,
	"TextArea":    lucide.Type,
	"NumberInput": lucide.Hash,
	"SearchInput": lucide.Search,
	"Checkbox":    lucide.CircleCheck,
	"RadioGroup":  lucide.Radio,
	"Switch":      lucide.ToggleRight,
	"Toggle":      lucide.ToggleLeft,
	"Slider":      lucide.Volume2,
	"RangeSlider": lucide.MoveHorizontal,
	"Select":      lucide.ListFilter,
	"MultiSelect": lucide.ListChecks,
	"Combobox":    lucide.ChevronsUpDown,
	"Tabs":        lucide.PanelsTopLeft,
	"Sidebar":     lucide.PanelLeft,
	"Toolbar":     lucide.PanelTop,
	"Breadcrumb":  lucide.ChevronsRight,
	"Pagination":  lucide.ListOrdered,
	"Dialog":      lucide.Square,
	"Tooltip":     lucide.Info,
	"Toast":       lucide.Bell,
	"Progress":    lucide.CircleDot,
	"Spinner":     lucide.LoaderCircle,
	"Skeleton":    lucide.RectangleHorizontal,
	"EmptyState":  lucide.Inbox,
	"Terminal":    lucide.SquareTerminal,
	"Calendar":    lucide.CalendarDays,
	"LineChart":   lucide.ChartLine,
	"BarChart":    lucide.ChartColumn,
	"PieChart":    lucide.ChartPie,
}

// mujicaCategoryIcons maps each MujicaUI gallery category to a lucide
// icon, so a page without a listed icon still draws a meaningful one in
// the sidebar.
var mujicaCategoryIcons = map[string]func() *ui.SVG{
	"Layout":                 lucide.LayoutTemplate,
	"Display":                lucide.Monitor,
	"Inputs":                 lucide.TextCursorInput,
	"Choices":                lucide.SquareCheck,
	"Forms":                  lucide.FileStack,
	"Date & Time":            lucide.CalendarDays,
	"Navigation":             lucide.Compass,
	"Data":                   lucide.Database,
	"Overlays":               lucide.Layers,
	"Feedback":               lucide.Bell,
	"Charts I":               lucide.ChartLine,
	"Charts II":              lucide.ChartColumn,
	"Market Charts":          lucide.ChartCandlestick,
	"Trading":                lucide.TrendingUp,
	"Markets":                lucide.ChartColumnBig,
	"Chat Messages":          lucide.MessageCircle,
	"Chat Input":             lucide.SquarePen,
	"Chat Sessions":          lucide.MessagesSquare,
	"Agent":                  lucide.Bot,
	"Agent Ops & Motion":     lucide.Zap,
	"Git":                    lucide.GitBranch,
	"Terminal & Code":        lucide.Terminal,
	"Media":                  lucide.Clapperboard,
	"Projects":               lucide.Kanban,
	"Files & Collaboration":  lucide.FolderOpen,
	"Messaging & Mail":       lucide.Mail,
	"Calendar":               lucide.Calendar,
	"Dashboards & Dev Tools": lucide.LayoutDashboard,
	"Settings & Account":     lucide.Settings,
	"Essentials":             lucide.Package,
}

// buildMujica wraps a MujicaUI page as a gallery demo: it applies the
// MujicaUI theme, then runs the page's demo with a no-op event log.
// The log only records events for the source gallery's panel; the uix
// page has none, so a no-op is faithful.
func buildMujica(p mujica.Page) func(c *ui.Context, t *ui.Theme) {
	return func(c *ui.Context, t *ui.Theme) {
		mujica.Use(c)
		p.Demo(c, func(string) {})
	}
}

// mujicaSections turns MujicaUI's catalog into gallery sections, one
// per category, in the order its gallery lists them. Every page keeps
// its source ID, "001".."500", which never collides with the uix IDs.
func mujicaSections() []Section {
	var out []Section
	index := map[string]int{}
	for _, p := range mujica.Pages() {
		i, ok := index[p.Category]
		if !ok {
			index[p.Category] = len(out)
			out = append(out, Section{Title: p.Category})
			i = len(out) - 1
		}
		icon := mujicaIcons[p.Name]
		if icon == nil {
			icon = mujicaCategoryIcons[p.Category]
		}
		out[i].Entries = append(out[i].Entries, Entry{
			ID:    p.ID,
			Name:  p.Name,
			Lib:   Mujica,
			Icon:  icon,
			Build: buildMujica(p),
		})
	}
	return out
}

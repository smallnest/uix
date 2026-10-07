// Package view builds the dashboard example: a window of MyGo native UI
// that shows the appshell, statcard, linechart and barchart components
// together as a store overview. The charts follow the pointer: resting
// on a month swaps the headline for that month and shows what it was a
// year earlier. The main package shows it in a window; the tests and the
// snapshot command draw it headless.
package view

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/appshell"
	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/barchart"
	"github.com/smallnest/uix/components/linechart"
	"github.com/smallnest/uix/components/sidebar"
	"github.com/smallnest/uix/components/statcard"
)

// Width and Height are the size of the example window.
const Width, Height = 1080, 820

// State is the state of the example; the controls edit it in place.
var State = DashboardState{}

// DashboardState holds the dashboard: the page shown and the months the
// charts rest on.
type DashboardState struct {
	// Place is the sidebar's choice: "overview", "reports",
	// "customers" or "settings".
	Place string
	// Revenue is the month the revenue chart rests on, -1 for none.
	Revenue int
	// Orders is the month the orders chart rests on, -1 for none.
	Orders int
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() { State = DashboardState{Place: "overview", Revenue: -1, Orders: -1} }

// sections are the sidebar's groups.
var sections = []sidebar.Section{
	{Title: "Dashboard", Items: []sidebar.Item{
		{ID: "overview", Label: "Overview"},
		{ID: "reports", Label: "Reports"},
	}},
	{Title: "Management", Items: []sidebar.Item{
		{ID: "customers", Label: "Customers"},
		{ID: "settings", Label: "Settings"},
	}},
}

// revenue and orders are the BoardUI series: a year of monthly figures
// against the year before.
var revenue = []linechart.Point{
	{Label: "Jan", Current: 9840, Previous: 8210},
	{Label: "Feb", Current: 10120, Previous: 8460},
	{Label: "Mar", Current: 11380, Previous: 9950},
	{Label: "Apr", Current: 10960, Previous: 10240},
	{Label: "May", Current: 12210, Previous: 10880},
	{Label: "Jun", Current: 12740, Previous: 11020},
	{Label: "Jul", Current: 13980, Previous: 11760},
	{Label: "Aug", Current: 13120, Previous: 12030},
	{Label: "Sep", Current: 14210, Previous: 12190},
	{Label: "Oct", Current: 14690, Previous: 12480},
	{Label: "Nov", Current: 14360, Previous: 12160},
	{Label: "Dec", Current: 14703.92, Previous: 11924},
}

var orders = []barchart.Point{
	{Label: "Jan", Current: 1680, Previous: 1510},
	{Label: "Feb", Current: 1740, Previous: 1480},
	{Label: "Mar", Current: 1920, Previous: 1650},
	{Label: "Apr", Current: 1850, Previous: 1620},
	{Label: "May", Current: 2040, Previous: 1710},
	{Label: "Jun", Current: 2110, Previous: 1760},
	{Label: "Jul", Current: 2290, Previous: 1840},
	{Label: "Aug", Current: 2180, Previous: 1890},
	{Label: "Sep", Current: 2320, Previous: 1930},
	{Label: "Oct", Current: 2410, Previous: 1970},
	{Label: "Nov", Current: 2280, Previous: 1810},
	{Label: "Dec", Current: 2342, Previous: 1798},
}

// DashboardView draws the example: the app shell around the page the
// sidebar chose.
func DashboardView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	appshell.AppShell(c, appshell.Props{
		Title:    pageTitle(),
		Team:     "Board team",
		Member:   "Mertcan",
		Selected: &State.Place,
		Sections: sections,
		Content: func(c *ui.Context) {
			if State.Place != "overview" {
				placeholder(c, t)
				return
			}
			dashboard(c, t)
		},
	})
}

// pageTitle is the name of the page shown, for the trail and heading.
func pageTitle() string {
	switch State.Place {
	case "reports":
		return "Reports"
	case "customers":
		return "Customers"
	case "settings":
		return "Settings"
	default:
		return "Overview"
	}
}

// placeholder is the page of the sidebar choices the example does not
// build, so the shell still works past the overview.
func placeholder(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Fill().Center().Gap(t.Space(2)).Children(func() {
		ui.Text(c, pageTitle()).FontSize(t.FontSize * 1.3).FontWeight(600).TextColor(t.Text)
		ui.Text(c, "This page is not part of the example.").TextColor(t.TextMuted).SingleLine()
	})
}

// dashboard is the overview page: the heading, the metric cards, the two
// charts and the display cards, as a BoardUI dashboard.
func dashboard(c *ui.Context, t *ui.Theme) {
	ui.Scroll(c).Fill().Children(func() {
		ui.Column(c).FillWidth().Gap(t.Space(5)).Padding(0, 0, t.Space(6), 0).Children(func() {
			ui.Column(c).Gap(2).Children(func() {
				ui.Text(c, "Overview").FontSize(t.FontSize * 1.6).FontWeight(600).TextColor(t.Text)
				ui.Text(c, "Your store's performance").TextColor(t.TextMuted).SingleLine()
			})
			statcard.StatCards(c, statcard.Props{Stats: []statcard.Stat{
				{Icon: statcard.IconUsers, Label: "Customers", Value: "14,592",
					Delta: "+5.3%", DeltaColor: statcard.Up},
				{Icon: statcard.IconBox, Label: "Unit sold", Value: "385",
					Delta: "-2.1%", DeltaColor: statcard.Down},
				{Icon: statcard.IconBasket, Label: "Orders", Value: "1,394",
					Delta: "0.00%", DeltaColor: statcard.Flat},
				{Icon: statcard.IconChat, Label: "Support tickets", Value: "708",
					Delta: "+12.8%", DeltaColor: statcard.Up},
			}})
			// The charts side by side, each filling half the row.
			ui.Row(c).FillWidth().Height(300).Gap(t.Space(4)).Children(func() {
				linechart.LineChart(c, linechart.Props{
					Data: revenue, Selected: &State.Revenue, Title: "Revenue",
				})
				barchart.BarChart(c, barchart.Props{
					Data: orders, Selected: &State.Orders, Title: "Orders",
				})
			})
			statcard.StatCards(c, statcard.Props{Variant: statcard.Footer, Stats: []statcard.Stat{
				{Icon: statcard.IconCoins, Label: "Total revenue", Value: "$152,313.92",
					Delta: "16%", DeltaColor: statcard.Up, Tone: statcard.Blue,
					Caption: "Against last month", Hint: "Gross revenue, before refunds."},
				{Icon: statcard.IconRefund, Label: "Refunds", Value: "$2,318.42",
					Delta: "3%", DeltaColor: statcard.Down, Tone: statcard.Pink,
					Caption: "Against last month"},
			}})
		})
	})
}

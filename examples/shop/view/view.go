// Package view builds the shop example: a window of MyGo native UI that
// shows the combobox, numberinput, searchfield and rating components
// together, as a product catalog the filters narrow. The main package
// shows it in a window; the tests and the snapshot command draw it
// headless.
package view

import (
	"fmt"
	"strings"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/combobox"
	"github.com/smallnest/uix/components/numberinput"
	"github.com/smallnest/uix/components/rating"
	"github.com/smallnest/uix/components/searchfield"
	"github.com/smallnest/uix/components/separator"
)

// Width and Height are the size of the example window.
const Width, Height = 460, 560

// Product is a row of the catalog.
type Product struct {
	Name, Category string
	Rate, Price    int
}

// Products are the rows of the catalog, in their original order.
var Products = []Product{
	{"Wrench", "Tools", 4, 12},
	{"Hammer", "Tools", 3, 9},
	{"Drill", "Tools", 5, 45},
	{"Notepad", "Office", 4, 3},
	{"Desk lamp", "Office", 2, 22},
	{"Coffee mug", "Kitchen", 5, 8},
	{"Kettle", "Kitchen", 4, 35},
}

// State is the state of the example; the filters edit it in place.
var State = ShopState{}

// ShopState holds the filters, as the controls edit them.
type ShopState struct {
	Query    string
	Category string
	Limit    float64
	MinRate  int
}

// Shown returns the products the filters leave, no more than the limit.
func Shown() []Product {
	var out []Product
	for _, p := range Products {
		if s := strings.ToLower(State.Query); s != "" && !strings.Contains(strings.ToLower(p.Name), s) {
			continue
		}
		if State.Category != "" && p.Category != State.Category {
			continue
		}
		if p.Rate < State.MinRate {
			continue
		}
		out = append(out, p)
		if State.Limit > 0 && len(out) >= int(State.Limit) {
			break
		}
	}
	return out
}

// ShopView draws the example: a search field, a category box, a limit
// and a minimum rating above the products, whose rows the filters leave.
// Every filter narrows the catalog at once.
func ShopView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	shown := Shown()
	ui.Column(c).Fill().Padding(24).Gap(14).Children(func() {
		ui.Text(c, "Shop").FontSize(24).Bold()
		// The search box names itself with its placeholder.
		searchfield.SearchField(c, searchfield.Props{Value: &State.Query}).Label("Search")
		ui.Row(c).Gap(t.Space(3)).Children(func() {
			// The combobox stretches to its parent by default; in a Row that
			// would claim the whole line and push the number input out, so
			// the Width here overrides it with a bounded trigger.
			ui.Field(c, "Category", func() {
				combobox.Combobox(c, combobox.Props{
					Value:   &State.Category,
					Options: []string{"Tools", "Office", "Kitchen"},
				}).Label("Category").Width(200)
			})
			ui.Field(c, "Limit", func() {
				numberinput.NumberInput(c, numberinput.Props{
					Value: &State.Limit,
					Lo:    0, Hi: 7, Step: 1,
				}).Label("Limit")
			})
		})
		// The field names the rating row, which names itself otherwise;
		// then the label is the only element the tests find as "Min rate".
		ui.Field(c, "Min rate", func() {
			rating.Rating(c, rating.Props{Value: &State.MinRate, Max: 5})
		})
		separator.Separator(c, separator.Props{})
		for _, p := range shown {
			ui.Row(c).Gap(t.Space(3)).Children(func() {
				ui.Column(c).Grow(1).Children(func() {
					ui.Text(c, p.Name).Bold()
				})
				rating.Rating(c, rating.Props{Value: &p.Rate, Max: 5, ReadOnly: true})
				ui.Text(c, fmt.Sprintf("$%d", p.Price)).TextColor(t.TextMuted)
			})
		}
		switch {
		case len(shown) == 0:
			ui.Text(c, "No products match.").TextColor(t.TextMuted)
		default:
			ui.Text(c, fmt.Sprintf("%d products", len(shown))).TextColor(t.TextMuted)
		}
	})
}

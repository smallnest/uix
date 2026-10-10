// Package view builds the placeholders example: a window of MyGo native
// UI that shows the skeleton, empty and aspect-ratio components of uix
// together in one page. The skeleton stands in for an article still
// loading, the empty state is what a search finds when it finds
// nothing, and the aspect-ratio box keeps the video's shape as the
// window changes. The main package shows it in a window; the tests and
// the snapshot command draw it headless.
package view

import (
	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/aspectratio"
	"github.com/smallnest/uix/components/empty"
	"github.com/smallnest/uix/components/input"
	"github.com/smallnest/uix/components/skeleton"
)

// Width and Height are the size of the example window.
const Width, Height = 460, 620

// State is the state of the example; the controls edit it in place.
var State = PlaceholdersState{Query: "video"}

// PlaceholdersState holds the example: the search query and whether it
// found something.
type PlaceholdersState struct {
	// Query is the text of the search above the empty state.
	Query string
	// Searched says the search has run, so the empty state shows.
	Searched bool
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() { State = PlaceholdersState{Query: ""} }

// video is the play button of the aspect-ratio box.
var video = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="currentColor"><path d="m6 4 14 8-14 8V4z"/></svg>`))

// PlaceholdersView draws the example: the three components, each under
// its label.
func PlaceholdersView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	ui.Column(c).Fill().Padding(24).Gap(t.Space(4)).Children(func() {
		ui.Column(c).FillWidth().Gap(1).Children(func() {
			ui.Text(c, "Placeholders").FontSize(t.FontSize * 1.5).FontWeight(600).TextColor(t.Text)
			ui.Text(c, "Skeleton, empty state and aspect ratio, from uix.").TextColor(t.TextMuted)
		})
		section(c, t, "Skeleton")
		articleCard(c, t)
		section(c, t, "Empty state")
		searchCard(c, t)
		section(c, t, "Aspect ratio")
		videoCard(c, t)
	})
}

// section is the label above a component.
func section(c *ui.Context, t *ui.Theme, name string) {
	ui.Text(c, name).FontSize(t.FontSize * 0.9).FontWeight(600).TextColor(t.TextMuted)
}

// articleCard is the skeleton demo: the bars of an article that is
// still loading.
func articleCard(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(2)).Children(func() {
		skeleton.Skeleton(c, skeleton.Props{Width: 240, Height: 22})
		skeleton.Skeleton(c, skeleton.Props{Height: 13, Radius: 8})
		skeleton.Skeleton(c, skeleton.Props{Height: 13, Radius: 8})
		skeleton.Skeleton(c, skeleton.Props{Width: 300, Height: 13, Radius: 8})
	})
}

// searchCard is the empty state demo: a search that may find nothing,
// and the empty state with its Clear action when it did.
func searchCard(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(2)).Children(func() {
		ui.Row(c).FillWidth().Gap(t.Space(2)).Children(func() {
			ui.Column(c).Grow(1).Children(func() {
				in := input.Input(c, input.Props{
					Value:       &State.Query,
					Placeholder: "Search…",
				})
				if in.Submitted() {
					State.Searched = true
				}
			})
		})
		if State.Searched {
			ui.Box(c).FillWidth().Height(220).Background(t.Surface).Radius(t.Radius).
				Children(func() {
					empty.Empty(c, empty.Props{
						Title:       "No results",
						Description: "Clear the search to see everything again.",
						ActionLabel: "Clear search",
						OnAction: func() {
							State.Query = ""
							State.Searched = false
						},
					})
				})
		} else {
			ui.Text(c, "Type a query, then press Enter to search.").
				FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		}
	})
}

// videoCard is the aspect-ratio demo: a dark box of 16:9 with a play
// button in the middle.
func videoCard(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(2)).Children(func() {
		aspectratio.AspectRatio(c, aspectratio.Props{
			Ratio: 16.0 / 9,
			Children: func() {
				ui.Box(c).Fill().Background(ui.RGBA(0, 0, 0, 0.8)).Center().
					Radius(t.Radius).Clip().Children(func() {
					ui.Box(c).Size(56, 56).Radius(28).Background(ui.RGBA(255, 255, 255, 0.9)).
						Center().Children(func() {
						ui.Icon(c, video).Size(20, 20).TextColor(ui.RGBA(0, 0, 0, 0.8))
					})
				})
			},
		})
		ui.Text(c, "The box keeps 16:9 as the window changes.").
			FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
}

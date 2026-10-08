// Package view builds the navigation example: a window of MyGo native
// UI that shows the carousel, pagination and file upload components of
// BoardUI together in one page. The carousel flips through slides with
// its buttons or dots, the pagination steps through pages, and the file
// zone uploads a picked file with its progress ring. The main package
// shows it in a window; the tests and the snapshot command draw it
// headless.
package view

import (
	"fmt"
	"time"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/button"
	"github.com/smallnest/uix/components/carousel"
	"github.com/smallnest/uix/components/fileupload"
	"github.com/smallnest/uix/components/pagination"
)

// Width and Height are the size of the example window.
const Width, Height = 760, 780

// State is the state of the example; the controls edit it in place.
var State = NavigationState{}

// NavigationState holds the example: the carousel, the page, the upload
// and the name the last finished upload reported.
type NavigationState struct {
	// Scroll is the scroll state of the carousel track.
	Scroll ui.ScrollState
	// Carousel is where the carousel is.
	Carousel carousel.State
	// Page is the page the pagination shows, from 1.
	Page int
	// Upload is where the file upload is.
	Upload fileupload.State
	// Uploaded is the name the last finished upload reported.
	Uploaded string
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() { State = NavigationState{Page: 1} }

// NavigationView draws the example: the three BoardUI components, each
// under its label.
func NavigationView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	ui.Column(c).Fill().Padding(24).Gap(t.Space(4)).Children(func() {
		ui.Column(c).FillWidth().Gap(1).Children(func() {
			ui.Text(c, "Navigation").FontSize(t.FontSize * 1.5).FontWeight(600).TextColor(t.Text)
			ui.Text(c, "Carousel, pagination and file upload, from BoardUI.").TextColor(t.TextMuted)
		})
		section(c, t, "Carousel")
		carousel.Carousel(c, carousel.Props{
			Count:  3,
			Slide:  slide,
			Scroll: &State.Scroll,
			State:  &State.Carousel,
		})
		section(c, t, "Pagination")
		pageCard(c, t)
		section(c, t, "File upload")
		uploadCard(c, t)
	})
}

// section is the label above a component.
func section(c *ui.Context, t *ui.Theme, name string) {
	ui.Text(c, name).FontSize(t.FontSize * 0.9).FontWeight(600).TextColor(t.TextMuted)
}

// slide is one carousel slide: a tinted banner with its number, lighter
// the further it is from the first.
func slide(c *ui.Context, i int) {
	t := c.Theme()
	ui.Box(c).FillWidth().Height(160).Background(t.Accent.Alpha(1 - float32(i)*0.18)).
		Radius(t.Radius * 2).Center().Children(func() {
		ui.Column(c).Center().Gap(t.Space(2)).Children(func() {
			ui.Text(c, fmt.Sprintf("Slide %d", i+1)).FontSize(t.FontSize * 1.4).
				FontWeight(600).TextColor(t.AccentText)
			ui.Text(c, "A MyGo component of BoardUI").FontSize(t.FontSize * 0.85).
				TextColor(t.AccentText.Alpha(0.85))
		})
	})
}

// pageCard is the pagination demo: a line naming the page, and the
// pagination under it that changes it.
func pageCard(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(3)).Children(func() {
		ui.Box(c).FillWidth().Background(t.Surface).Radius(t.Radius).
			Padding(t.Space(3), t.Space(4)).Children(func() {
			ui.Text(c, fmt.Sprintf("Page %d of 12", State.Page)).FontWeight(600).TextColor(t.Text).SingleLine()
		})
		ui.Row(c).FillWidth().Justify(ui.Center).Children(func() {
			pagination.Pagination(c, pagination.Props{
				Page:       State.Page,
				TotalPages: 12,
				OnChange:   func(p int) { State.Page = p },
			})
		})
	})
}

// uploadCard is the file upload demo: the zone, the controls to start a
// sample upload or clear it, and the name the completion reported.
func uploadCard(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(3)).Children(func() {
		fileupload.FileUpload(c, fileupload.Props{
			State: &State.Upload,
			OnComplete: func(name string, _ int64) {
				State.Uploaded = name
			},
		})
		ui.Row(c).FillWidth().Justify(ui.End).Gap(t.Space(3)).Children(func() {
			// The controls start a sample upload without the dialog, so
			// the example runs anywhere.
			button.Button(c, button.Props{
				Label: "Upload sample",
				Size:  button.Sm,
				OnClick: func() {
					State.Upload.Start("sales-report.xlsx", 1258291, time.Now())
				},
			})
			button.Button(c, button.Props{
				Label:   "Clear",
				Size:    button.Sm,
				Variant: button.Secondary,
				OnClick: func() { State.Upload = fileupload.State{} },
			})
		})
		if State.Uploaded != "" {
			ui.Text(c, fmt.Sprintf("%s uploaded", State.Uploaded)).TextColor(t.Success)
		}
	})
}

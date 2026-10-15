package view

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/accordion"
	"github.com/smallnest/uix/components/aspectratio"
	"github.com/smallnest/uix/components/collapsible"
	"github.com/smallnest/uix/components/field"
	"github.com/smallnest/uix/components/fieldset"
	"github.com/smallnest/uix/components/fileupload"
	"github.com/smallnest/uix/components/form"
	"github.com/smallnest/uix/components/grid"
	"github.com/smallnest/uix/components/input"
	"github.com/smallnest/uix/components/scroll"
	"github.com/smallnest/uix/components/scrollboth"
	"github.com/smallnest/uix/components/scrollhorizontal"
	"github.com/smallnest/uix/components/separator"
	"github.com/smallnest/uix/components/split"
)

// demoSeparator shows a horizontal divider.
func demoSeparator(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(3)).Children(func() {
		ui.Text(c, "Above the line")
		separator.Separator(c, separator.Props{})
		ui.Text(c, "Below the line")
	})
}

// demoField shows a field with a label, a hint and a control.
func demoField(c *ui.Context, t *ui.Theme) {
	field.Field(c, field.Props{Label: "Name", Hint: "As shown to others"}, func() {
		input.Input(c, input.Props{Value: &Demo.Input})
	})
}

// demoFieldset shows a group of fields under a bold legend.
func demoFieldset(c *ui.Context, t *ui.Theme) {
	fieldset.Fieldset(c, fieldset.Props{
		Legend: "Account",
		Children: func(c *ui.Context) {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				field.Field(c, field.Props{Label: "Name"}, func() {
					input.Input(c, input.Props{Value: &Demo.Input})
				})
				field.Field(c, field.Props{Label: "Email"}, func() {
					input.Input(c, input.Props{Value: &Demo.Float, Placeholder: "you@example.com"})
				})
			})
		},
	})
}

// demoForm shows a form whose labels line up to the widest.
func demoForm(c *ui.Context, t *ui.Theme) {
	form.Form(c, form.Props{
		Children: func(c *ui.Context) {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				field.Field(c, field.Props{Label: "Name"}, func() {
					input.Input(c, input.Props{Value: &Demo.Input})
				})
				field.Field(c, field.Props{Label: "Email"}, func() {
					input.Input(c, input.Props{Value: &Demo.Float, Placeholder: "you@example.com"})
				})
			})
		},
	})
}

// demoGrid shows a grid layout of columns the children fill in order.
func demoGrid(c *ui.Context, t *ui.Theme) {
	grid.Grid(c, grid.Props{
		Columns: 3,
		Gap:     t.Space(2),
		Children: func(c *ui.Context) {
			for i := 0; i < 6; i++ {
				ui.Box(c).Height(56).Radius(t.Radius).Background(t.Accent.Alpha(1 - float32(i)*0.12))
			}
		},
	})
}

// demoScroll shows a container that scrolls its content vertically.
func demoScroll(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Height(160).Children(func() {
		scroll.Scroll(c, scroll.Props{Children: func(c *ui.Context) {
			ui.Column(c).Gap(t.Space(1.5)).Children(func() {
				for i := 1; i <= 20; i++ {
					ui.Text(c, "Row "+string(rune('0'+i/10))+string(rune('0'+i%10))).TextColor(t.TextMuted)
				}
			})
		}}).Fill()
	})
}

// demoScrollHorizontal shows a row that scrolls its content sideways.
func demoScrollHorizontal(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Height(72).Children(func() {
		scrollhorizontal.ScrollHorizontal(c, scrollhorizontal.Props{Children: func(c *ui.Context) {
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				for i := 1; i <= 12; i++ {
					ui.Box(c).Size(96, 64).Radius(t.Radius).Background(t.Surface).Border(1, t.Border).
						Children(func() {
							ui.Text(c, string(rune('A'+i-1))).TextColor(t.TextMuted).Center()
						})
				}
			})
		}}).Fill()
	})
}

// demoScrollBoth shows a container that scrolls up and sideways.
func demoScrollBoth(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Height(160).Children(func() {
		scrollboth.ScrollBoth(c, scrollboth.Props{Children: func(c *ui.Context) {
			ui.Column(c).Gap(t.Space(1.5)).Children(func() {
				for row := 1; row <= 20; row++ {
					ui.Text(c, "Line "+string(rune('0'+row/10))+string(rune('0'+row%10))).TextColor(t.TextMuted)
				}
			})
		}}).Fill()
	})
}

// demoSplit shows resizable panes with a divider.
func demoSplit(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Height(180).Children(func() {
		split.Split(c, split.Props{
			Size: &Demo.Split,
			First: func(c *ui.Context) {
				ui.Column(c).Fill().Padding(t.Space(2)).Children(func() {
					ui.Text(c, "First pane").TextColor(t.TextMuted)
				})
			},
			Second: func(c *ui.Context) {
				ui.Column(c).Fill().Padding(t.Space(2)).Children(func() {
					ui.Text(c, "Second pane").TextColor(t.TextMuted)
				})
			},
		})
	})
}

// demoAspectRatio shows a box that keeps an aspect ratio.
func demoAspectRatio(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(280).Children(func() {
		aspectratio.AspectRatio(c, aspectratio.Props{
			Ratio: 16.0 / 9.0,
			Children: func() {
				ui.Box(c).Fill().Background(t.Accent.Alpha(0.25)).Children(func() {
					ui.Text(c, "16:9").TextColor(t.Accent).Center()
				})
			},
		})
	})
}

// demoCollapsible shows a disclosure whose label shows the content.
func demoCollapsible(c *ui.Context, t *ui.Theme) {
	collapsible.Collapsible(c, collapsible.Props{
		Label: "More details",
		Open:  &Demo.Collapse,
		Children: func(c *ui.Context) {
			ui.Text(c, "The hidden content of the disclosure.").TextColor(t.TextMuted)
		},
	})
}

// demoAccordion shows a column of sections a click opens and closes.
func demoAccordion(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(360).Children(func() {
		accordion.Accordion(c, accordion.Props{
			Items: []accordion.Item{
				{Title: "Section one", Open: &Demo.Acc1, Build: func(c *ui.Context) {
					ui.Text(c, "The first section.").TextColor(t.TextMuted)
				}},
				{Title: "Section two", Open: &Demo.Acc2, Build: func(c *ui.Context) {
					ui.Text(c, "The second section.").TextColor(t.TextMuted)
				}},
			},
		})
	})
}

// demoFileUpload shows a file zone with a progress ring.
func demoFileUpload(c *ui.Context, t *ui.Theme) {
	fileupload.FileUpload(c, fileupload.Props{
		State:             &Demo.FileUp,
		AllowedExtensions: []string{".png", ".jpg", ".pdf"},
		Height:            160,
	})
}

// demoCanvas shows a mygo canvas drawn with the painter.
func demoCanvas(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Size(96, 96).Radius(t.Radius).Background(t.Surface).Border(1, t.Border).Draw(func(p *ui.Painter, r ui.Rect) {
		p.Fill(ui.Rect{X: r.X + 8, Y: r.Y + 8, W: r.W - 16, H: r.H - 16}, t.Accent, 8)
		p.Fill(ui.Rect{X: r.X + 32, Y: r.Y + 32, W: 32, H: 32}, t.Success, 6)
	})
}

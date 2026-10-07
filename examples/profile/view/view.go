// Package view builds the profile example: a window of MyGo native UI
// that shows the card, avatar, separator and alert components together,
// as a profile page. The main package shows it in a window; the tests
// and the snapshot command draw it headless.
package view

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/alert"
	"github.com/smallnest/uix/components/avatar"
	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/button"
	"github.com/smallnest/uix/components/card"
	"github.com/smallnest/uix/components/separator"
)

// Width and Height are the size of the example window.
const Width, Height = 480, 560

// State is the state of the example; the controls edit it in place.
var State = ProfileState{}

// ProfileState holds whether the follow button is on.
type ProfileState struct {
	Following bool
}

// ProfileView draws the example: a profile card with an avatar, a
// divider and a follow button, and the alerts under it. The follow
// button toggles the following state, which swaps its label and shows a
// success alert.
func ProfileView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	ui.Column(c).Fill().Padding(24).Gap(14).Children(func() {
		ui.Text(c, "Profile").FontSize(24).Bold()
		card.Card(c, card.Props{
			Description: "Mathematician, writer of the first algorithm.",
			Footer: func() {
				button.Button(c, button.Props{Label: "Edit", Variant: button.Secondary})
				follow := "Follow"
				if State.Following {
					follow = "Following"
				}
				button.Button(c, button.Props{
					Label:   follow,
					Variant: button.Primary,
					OnClick: func() { State.Following = !State.Following },
				})
			},
		}, func() {
			ui.Row(c).Gap(t.Space(3)).Children(func() {
				avatar.Avatar(c, avatar.Props{Name: "Ada Lovelace"})
				ui.Column(c).Gap(t.Space(0.5)).Children(func() {
					ui.Text(c, "Ada Lovelace").Bold()
					ui.Text(c, "@ada").TextColor(t.TextMuted)
				})
			})
			separator.Separator(c, separator.Props{})
			ui.Text(c, "Devised the first algorithm meant for a machine, for the Analytical Engine.").
				FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
		})
		alert.Alert(c, alert.Props{
			Title:       "Heads up!",
			Description: "You can add components to your app.",
			Variant:     alert.Info,
		})
		alert.Alert(c, alert.Props{
			Title:       "Beta",
			Description: "Some features are not final.",
			Variant:     alert.Warning,
		})
		alert.Alert(c, alert.Props{
			Title:       "Danger",
			Description: "This action cannot be undone.",
			Variant:     alert.Danger,
		})
		if State.Following {
			alert.Alert(c, alert.Props{
				Title:       "Followed!",
				Description: "You now follow Ada Lovelace.",
				Variant:     alert.Success,
			})
		}
	})
}

// Package view builds the forms example: a window of MyGo native UI that
// shows the form controls of uix together. The main package shows it in a
// window; the tests and the snapshot command draw it headless.
package view

import (
	"fmt"
	"strings"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/button"
	"github.com/smallnest/uix/components/checkbox"
	"github.com/smallnest/uix/components/field"
	"github.com/smallnest/uix/components/input"
	selector "github.com/smallnest/uix/components/select"
	"github.com/smallnest/uix/components/slider"
	switcher "github.com/smallnest/uix/components/switch"
	"github.com/smallnest/uix/components/textarea"
)

// Width and Height are the size of the example window.
const Width, Height = 440, 760

// Form is the state of the example; the controls edit it in place.
var Form = FormState{Volume: 60}

// FormState holds the values the controls edit, and the errors that the
// submit button sets.
type FormState struct {
	Name      string
	Email     string
	Country   string
	Message   string
	Volume    float64
	Dark      bool
	Agree     bool
	Submitted bool
	NameErr   string
	EmailErr  string
	TermsErr  string
}

// FormsView draws the example: a sign-up form with a name, an email, a
// country select, a message text area, a volume slider, a dark-mode
// switch, a terms check box and a submit button. The dark-mode switch
// restyles the whole form at once, and the submit button validates the
// fields it must before it welcomes the sign-up.
func FormsView(c *ui.Context) {
	// The dark-mode switch forces the dark appearance; off, the window
	// follows the system.
	c.SetTheme(tokens.For(c.Theme().Dark || Form.Dark))
	t := c.Theme()
	ui.Column(c).Fill().Padding(24).Gap(14).Children(func() {
		ui.Text(c, "Sign up").FontSize(24).Bold()
		field.Field(c, field.Props{
			Label:    "Name",
			Required: true,
			Hint:     "How people find you.",
			Error:    Form.NameErr,
		}, func() {
			input.Input(c, input.Props{Value: &Form.Name, Placeholder: "Ada Lovelace"})
		})
		field.Field(c, field.Props{Label: "Email", Error: Form.EmailErr}, func() {
			input.Input(c, input.Props{Value: &Form.Email, Placeholder: "you@example.com"})
		})
		field.Field(c, field.Props{Label: "Country"}, func() {
			selector.Select(c, selector.Props{
				Selected:    &Form.Country,
				Options:     []string{"China", "Japan", "Singapore", "Germany", "United States"},
				Placeholder: "Choose a country",
			})
		})
		field.Field(c, field.Props{Label: "Message", Hint: "What is your project about?"}, func() {
			textarea.Textarea(c, textarea.Props{
				Value:       &Form.Message,
				Placeholder: "Tell us a little…",
				Rows:        5,
			})
		})
		slider.Slider(c, slider.Props{
			Value:  &Form.Volume,
			Min:    0,
			Max:    100,
			Label:  "Volume",
			Format: func(v float64) string { return fmt.Sprintf("%.0f%%", v) },
		})
		switcher.Switch(c, switcher.Props{On: &Form.Dark, Label: "Dark mode"})
		checkbox.Checkbox(c, checkbox.Props{Checked: &Form.Agree, Label: "I agree to the terms"})
		if Form.TermsErr != "" {
			ui.Text(c, Form.TermsErr).FontSize(t.FontSize * 0.875).TextColor(t.Danger)
		}
		button.Button(c, button.Props{Label: "Create account", Variant: button.Primary, OnClick: submit})
		if Form.Submitted {
			ui.Text(c, "Welcome!").TextColor(t.Success).Bold()
		}
	})
}

// submit validates the required fields, stores the errors it finds and
// shows the welcome message only when the form is clear.
func submit() {
	Form.NameErr, Form.EmailErr, Form.TermsErr = "", "", ""
	Form.Submitted = true
	if Form.Name == "" {
		Form.NameErr = "Name is required."
		Form.Submitted = false
	}
	if !strings.Contains(Form.Email, "@") {
		Form.EmailErr = "Enter a valid email."
		Form.Submitted = false
	}
	if !Form.Agree {
		Form.TermsErr = "You must agree to the terms."
		Form.Submitted = false
	}
}

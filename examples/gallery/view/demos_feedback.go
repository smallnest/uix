package view

import (
	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/editor"

	"github.com/smallnest/uix/components/alert"
	"github.com/smallnest/uix/components/alertdialog"
	"github.com/smallnest/uix/components/attachment"
	"github.com/smallnest/uix/components/chat"
	"github.com/smallnest/uix/components/code"
	"github.com/smallnest/uix/components/command"
	"github.com/smallnest/uix/components/composer"
	"github.com/smallnest/uix/components/dialog"
	"github.com/smallnest/uix/components/drawer"
	"github.com/smallnest/uix/components/loadingbutton"
	"github.com/smallnest/uix/components/log"
	"github.com/smallnest/uix/components/markdown"
	"github.com/smallnest/uix/components/message"
	"github.com/smallnest/uix/components/notificationcenter"
	"github.com/smallnest/uix/components/popover"
	"github.com/smallnest/uix/components/progress"
	"github.com/smallnest/uix/components/questionnaire"
	sel "github.com/smallnest/uix/components/select"
	"github.com/smallnest/uix/components/settings"
	"github.com/smallnest/uix/components/sheet"
	"github.com/smallnest/uix/components/spinner"
	sw "github.com/smallnest/uix/components/switch"
	"github.com/smallnest/uix/components/thinking"
	"github.com/smallnest/uix/components/toast"
	"github.com/smallnest/uix/components/tooltip"

	"github.com/smallnest/uix/components/button"
)

// demoDialog shows a modal dialog with a title and actions; the button
// opens it.
func demoDialog(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "Open dialog", OnClick: func() { Demo.Dialog = true }})
		ui.Text(c, "The dialog opens over the window.").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
	dialog.Dialog(c, dialog.Props{
		Open:        &Demo.Dialog,
		Title:       "Delete your account?",
		Description: "This removes your data and cannot be undone.",
		Confirm:     "Delete",
		Cancel:      "Keep account",
		OnConfirm:   func() { Demo.Dialog = false },
		OnCancel:    func() { Demo.Dialog = false },
	})
}

// demoAlertDialog shows a modal that interrupts for a confirmation.
func demoAlertDialog(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "Open alert dialog", Variant: button.Secondary, OnClick: func() { Demo.Alert = true }})
	})
	alertdialog.AlertDialog(c, alertdialog.Props{
		Open:        &Demo.Alert,
		Title:       "Are you sure?",
		Description: "You cannot undo this.",
		Confirm:     "Continue",
		Cancel:      "Cancel",
		OnConfirm:   func() { Demo.Alert = false },
		OnCancel:    func() { Demo.Alert = false },
	})
}

// demoDrawer shows a panel that rises over the bottom.
func demoDrawer(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "Open drawer", OnClick: func() { Demo.Drawer = true }})
	})
	drawer.Drawer(c, drawer.Props{
		Open:  &Demo.Drawer,
		Title: "Drawer",
		Children: func() {
			ui.Text(c, "A panel that rises over the bottom, dimming the rest.")
		},
	})
}

// demoSheet shows a panel that slides over from an edge.
func demoSheet(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "Open sheet", OnClick: func() { Demo.Sheet = true }})
	})
	sheet.Sheet(c, sheet.Props{
		Open:  &Demo.Sheet,
		Title: "Sheet",
		Children: func() {
			ui.Text(c, "A panel that slides over from the edge.")
		},
	})
}

// demoPopover shows a panel below a trigger.
func demoPopover(c *ui.Context, t *ui.Theme) {
	popover.Popover(c, popover.Props{
		Open: &Demo.Pop,
		Trigger: func(c *ui.Context) ui.Element {
			return button.Button(c, button.Props{Label: "Open popover", Variant: button.Outline})
		},
		Content: func(c *ui.Context) {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				ui.Text(c, "A panel below the trigger.").Bold()
				ui.Text(c, "It closes on a click outside or Escape.").TextColor(t.TextMuted)
			})
		},
	})
}

// demoTooltip shows a label that appears over a control.
func demoTooltip(c *ui.Context, t *ui.Theme) {
	tooltip.Tooltip(c, tooltip.Props{
		Anchor: button.Button(c, button.Props{Label: "Hover me"}),
		Text:   "A tip appears over the button.",
	})
}

// demoToast shows a transient notification at a window corner.
func demoToast(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "Show toast", Variant: button.Secondary, OnClick: func() {
			toast.Push(c, toast.Props{Title: "Settings saved", Description: "They take effect at once.", Type: toast.Success})
		}})
		ui.Text(c, "The toast appears at the corner and fades.").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
	toast.Viewport(c)
}

// demoSpinner shows a ring of spokes that turns while work goes on.
func demoSpinner(c *ui.Context, t *ui.Theme) {
	spinner.Spinner(c, spinner.Props{})
}

// demoProgress shows a progress bar with a label and a readout.
func demoProgress(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(300).Gap(t.Space(2)).Children(func() {
		progress.Progress(c, progress.Props{Value: 0.64, Label: "Disk usage"})
		progress.Progress(c, progress.Props{Value: 0, Label: "Indexing", Indeterminate: true})
	})
}

// demoCommand shows a palette that opens over the window.
func demoCommand(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "Open command palette", OnClick: func() { Demo.Command = true }})
	})
	command.Command(c, command.Props{
		Open:  &Demo.Command,
		State: &commandState,
		Items: []command.Item{
			{Label: "New file", Keywords: []string{"create", "document"}},
			{Label: "Open settings", Keywords: []string{"preferences"}},
			{Label: "Run build", Keywords: []string{"compile", "make"}},
		},
	})
}

// demoAlert shows colored notices in four variants.
func demoAlert(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(360).Gap(t.Space(2)).Children(func() {
		alert.Alert(c, alert.Props{Title: "Info", Description: "A notice in the accent color.", Variant: alert.Info})
		alert.Alert(c, alert.Props{Title: "Success", Description: "Everything worked.", Variant: alert.Success})
		alert.Alert(c, alert.Props{Title: "Warning", Description: "Mind the limit.", Variant: alert.Warning})
		alert.Alert(c, alert.Props{Title: "Danger", Description: "Something failed.", Variant: alert.Danger})
	})
}

// demoLoadingButton shows a button that turns into a spinner.
func demoLoadingButton(c *ui.Context, t *ui.Theme) {
	loadingbutton.LoadingButton(c, loadingbutton.Props{Label: "Saving", Loading: true})
}

// demoNotificationCenter shows a panel of notifications in tabs.
func demoNotificationCenter(c *ui.Context, t *ui.Theme) {
	notificationcenter.NotificationCenter(c, notificationcenter.Props{
		Tab: &Demo.NCenter,
		Items: []notificationcenter.Item{
			{ID: "1", Category: "system", Title: "Backup complete", Description: "Your data is safe.", Timestamp: "2m", Unread: true},
			{ID: "2", Category: "mentions", Title: "Ada mentioned you", Description: "In the design review.", Timestamp: "1h", Unread: true},
			{ID: "3", Category: "system", Title: "Update available", Description: "v0.2.2 is ready.", Timestamp: "1d"},
		},
	})
}

// demoCode shows retained syntax highlighting.
func demoCode(c *ui.Context, t *ui.Theme) {
	code.Parse("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"hello\")\n}\n", "go").Render(c)
}

// demoMarkdown shows markdown with tables and highlighted code.
func demoMarkdown(c *ui.Context, t *ui.Theme) {
	markdown.Markdown(c, "# Gallery\n\nA **markdown** view with `code`.\n\n| Col | Col |\n| --- | --- |\n| one | two |\n")
}

// demoCodeEditor shows mygo's code editor, read-only here.
func demoCodeEditor(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Height(220).Children(func() {
		editor.View(c, &editor.State{
			Buffer:   editor.NewBuffer("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"hello\")\n}\n"),
			ReadOnly: true,
		})
	})
}

// demoChat shows a transcript and a composer.
func demoChat(c *ui.Context, t *ui.Theme) {
	chat.Chat(c, chat.Props{
		Messages: []chat.Message{
			{Role: chat.User, Text: "What is uix?"},
			{Role: chat.Assistant, Text: "A component library for MyGo native UI, in pure Go."},
		},
		Value:       &Demo.Chat,
		Model:       "claude-fable-5",
		Provider:    "anthropic",
		OnSubmit:    func(string) {},
		OnStop:      func() {},
		Suggestions: []string{"Show a component", "List the libraries"},
	})
}

// demoComposer shows a message box with a prompt and a send button.
func demoComposer(c *ui.Context, t *ui.Theme) {
	composer.Composer(c, composer.Props{
		Value:       &Demo.Composer,
		Placeholder: "Ask anything",
		OnSubmit:    func(string) {},
		OnStop:      func() {},
		Model:       "claude-fable-5",
		Provider:    "anthropic",
	})
}

// demoMessage shows a turn of a conversation.
func demoMessage(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Gap(t.Space(2)).Children(func() {
		message.Message(c, message.Props{
			Name:   "Ada",
			Header: "Ada · 2m",
			Text:   "A message bubble of a conversation.",
		})
	})
}

// demoAttachment shows a card of a file of a conversation.
func demoAttachment(c *ui.Context, t *ui.Theme) {
	attachment.Attachment(c, attachment.Props{
		Title:       "design.pdf",
		Description: "2.4 MB · uploaded",
		Progress:    -1,
	})
}

// demoThinking shows the thinking indicator of an agent.
func demoThinking(c *ui.Context, t *ui.Theme) {
	thinking.Thinking(c, thinking.Props{Label: "Thinking"})
}

// demoLog shows a scrolling list of steps that reveal one by one.
func demoLog(c *ui.Context, t *ui.Theme) {
	log.Log(c, log.Props{
		Items:    []string{"Reading the docs", "Finding the pattern", "Writing the code"},
		Revealed: 3,
		Running:  false,
	})
}

// demoQuestionnaire shows a form of questions, one at a time.
func demoQuestionnaire(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(420).Children(func() {
		questionnaire.Questionnaire(c, questionnaire.Props{
			State: &Demo.Question,
			Steps: []questionnaire.Step{
				{Title: "How large is your team?", Description: "The size of your group.", Choices: []questionnaire.Choice{
					{Label: "1–10", Value: "small"},
					{Label: "11–50", Value: "mid"},
					{Label: "50+", Value: "large"},
				}},
				{Title: "What is your role?", Description: "What you do.", Choices: []questionnaire.Choice{
					{Label: "Engineer", Value: "engineer"},
					{Label: "Designer", Value: "designer"},
				}},
			},
			OnDone: func() {},
		})
	})
}

// demoSettings shows the settings composer of a category.
func demoSettings(c *ui.Context, t *ui.Theme) {
	settings.Settings(c, &Demo.Settings, []settings.SettingsCategory{
		{
			ID: "general", Title: "General", Description: "The app's preferences.",
			Groups: []settings.SettingsGroup{
				{
					ID: "appearance", Title: "Appearance",
					Items: []settings.SettingsItem{
						{ID: "dark", Title: "Dark mode", Description: "Use the dark appearance.", Build: func(c *ui.Context) {
							sw.Switch(c, sw.Props{On: &Demo.Switch, Label: ""})
						}},
						{ID: "language", Title: "Language", Description: "The language of the interface.", Build: func(c *ui.Context) {
							sel.Select(c, sel.Props{Selected: &Demo.SelStr, Options: []string{"English", "简体中文", "日本語"}})
						}},
					},
				},
			},
		},
	}, settings.SettingsOptions{})
}

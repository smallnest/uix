// Package view builds the agent example: a window of MyGo native UI that
// shows the appshell, chat, composer, log, notificationcenter and
// thinking components together in a starter for an agentic app. A demo
// model streams a reply word by word in the chat, the log reveals its
// steps one by one, and the notifications filter by tab. The main
// package shows it in a window; the tests and the snapshot command draw
// it headless.
package view

import (
	"strings"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/appshell"
	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/chat"
	"github.com/smallnest/uix/components/log"
	"github.com/smallnest/uix/components/notificationcenter"
	"github.com/smallnest/uix/components/sidebar"
)

// Width and Height are the size of the example window.
const Width, Height = 980, 680

// State is the state of the example; the controls edit it in place.
var State = AgentState{}

// AgentState holds the agent app: the page shown, the chat, the log and
// the notifications.
type AgentState struct {
	// Place is the sidebar's choice: "chat", "log" or "notifications".
	Place string
	// Messages are the conversation, edited by the chat and the demo
	// model.
	Messages []chat.Message
	// Input is the composer's draft.
	Input string
	// Busy is true while the demo model writes an answer.
	Busy bool
	// Thinking is true before the model's first word.
	Thinking bool
	// Scroll follows the chat transcript.
	Scroll ui.ScrollState
	// Revealed is how many log steps have finished.
	Revealed int
	// Notifs are the notifications, whose unread flags Mark all read
	// clears.
	Notifs []notificationcenter.Item
	// Tab is the notifications tab shown: 0 All, 1 Mentions, 2 System.
	Tab int
	// The reply the demo model streams.
	replyWords []string
	nextWord   time.Time
	// When the next log step lands.
	nextTick time.Time
}

// Paces of the demo model, as package vars so the tests and the snapshot
// command can shorten them.
var (
	// StreamPace is the delay between the words of a reply.
	StreamPace = 45 * time.Millisecond
	// ThinkPace is the delay before the model's first word.
	ThinkPace = 600 * time.Millisecond
	// LogPace is the delay between the log's steps.
	LogPace = 700 * time.Millisecond
)

// logSteps are the steps the agent takes, revealed one by one.
var logSteps = []string{
	"Planning the task",
	"Searching the docs",
	"Reading the source",
	"Writing the summary",
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = AgentState{
		Place:   "chat",
		Notifs:  sampleNotifs(),
		nextTick: time.Time{},
	}
}

// sampleNotifs is the notification center's starting set.
func sampleNotifs() []notificationcenter.Item {
	return []notificationcenter.Item{
		{ID: "n1", Category: "mentions", Title: "Ada mentioned you",
			Description: "in a comment on the revenue chart", Timestamp: "2m",
			Unread: true, Status: notificationcenter.Info},
		{ID: "n2", Category: "system", Title: "Backup complete",
			Description: "The nightly backup finished at 3:00", Timestamp: "1h",
			Unread: true, Status: notificationcenter.Success},
		{ID: "n3", Category: "activity", Title: "Weekly report is ready",
			Description: "Open it from the dashboard", Timestamp: "3h",
			Status: notificationcenter.Neutral},
	}
}

// sections are the sidebar's groups.
var sections = []sidebar.Section{
	{Title: "Agent", Items: []sidebar.Item{
		{ID: "chat", Label: "Chat"},
		{ID: "log", Label: "Log"},
		{ID: "notifications", Label: "Notifications"},
	}},
	{Title: "Account", Items: []sidebar.Item{
		{ID: "settings", Label: "Settings"},
	}},
}

// AgentView draws the example: the app shell around the page the sidebar
// chose.
func AgentView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	tick(c)
	appshell.AppShell(c, appshell.Props{
		Title:    pageTitle(),
		Team:     "Board team",
		Member:   "Mertcan",
		Selected: &State.Place,
		Sections: sections,
		Content: func(c *ui.Context) {
			switch State.Place {
			case "log":
				log.Log(c, log.Props{
					Items:    logSteps,
					Revealed: State.Revealed,
					Running:  State.Revealed < len(logSteps),
					Working:  workingLabel(),
				})
			case "notifications":
				notificationcenter.NotificationCenter(c, notificationcenter.Props{
					Items: State.Notifs,
					Tab:   &State.Tab,
					OnTabChange: func(i int) {
						State.Tab = i
					},
					OnMarkAllRead: func() {
						for i := range State.Notifs {
							State.Notifs[i].Unread = false
						}
					},
					Height: 400,
				})
			default:
				chat.Chat(c, chat.Props{
					Messages:  State.Messages,
					Value:     &State.Input,
					Busy:      State.Busy,
					Thinking:  State.Thinking,
					OnSubmit:  submit,
					OnStop:    stop,
					Model:     "openai/gpt-5-nano",
					Provider:  "Demo mode",
					Scroll:    &State.Scroll,
					Suggestions: []string{
						"Explain this starter",
						"Suggest a name for the app",
						"Write the product update",
					},
				})
			}
		},
	})
}

// pageTitle is the name of the page shown, for the trail and heading.
func pageTitle() string {
	switch State.Place {
	case "log":
		return "Log"
	case "notifications":
		return "Notifications"
	default:
		return "Chat"
	}
}

// workingLabel is the step the agent works on now: the next unrevealed
// one, none when it finished.
func workingLabel() string {
	if State.Revealed >= len(logSteps) {
		return ""
	}
	return logSteps[State.Revealed]
}

// submit sends the user's words to the demo model, which starts an
// empty assistant message.
func submit(text string) {
	State.Messages = append(State.Messages,
		chat.Message{Role: chat.User, Text: text},
		chat.Message{Role: chat.Assistant, Streaming: true},
	)
	State.Busy = true
	State.Thinking = true
	State.replyWords = replyFor(text)
	State.nextWord = time.Now().Add(ThinkPace)
}

// stop cuts the demo model's answer off, keeping the words it wrote.
func stop() {
	State.Busy = false
	State.Thinking = false
	State.replyWords = nil
	if n := len(State.Messages); n > 0 {
		State.Messages[n-1].Streaming = false
	}
}

// replyFor is the demo model: a canned answer, one word at a time.
func replyFor(text string) []string {
	switch {
	case strings.Contains(text, "update"):
		return strings.Fields("Here is the update: the team shipped three fixes this week, and the new dashboard went live on Monday.")
	case strings.Contains(text, "name"):
		return strings.Fields("How about Pulse? It is short, it is fast, and it says what the app does: it keeps the team in the loop.")
	default:
		return strings.Fields("Here is what I found: these components are native MyGo UI, they copy into your app with uix add, and the theme tokens keep them in your style.")
	}
}

// tick runs the demo model and the log's pace, asking for the frames
// they take. The window sleeps between them; the tests and the snapshot
// command shorten the paces and drive the frames themselves.
func tick(c *ui.Context) {
	now := c.Now()
	if State.Busy {
		// The model thinks first, then streams one word a frame.
		if State.Thinking {
			if now.Before(State.nextWord) {
				c.After(State.nextWord.Sub(now))
			} else {
				State.Thinking = false
				State.nextWord = now.Add(StreamPace)
			}
		} else {
			last := &State.Messages[len(State.Messages)-1]
			if len(State.replyWords) == 0 {
				// The answer is done.
				State.Busy = false
				last.Streaming = false
			} else if now.Before(State.nextWord) {
				c.After(State.nextWord.Sub(now))
			} else {
				if last.Text != "" {
					last.Text += " "
				}
				last.Text += State.replyWords[0]
				State.replyWords = State.replyWords[1:]
				last.Streaming = len(State.replyWords) > 0
				State.nextWord = now.Add(StreamPace)
				c.After(StreamPace)
			}
		}
	}
	if State.Revealed < len(logSteps) {
		// The first step waits one pace, so the log starts with a beat.
		if State.nextTick.IsZero() {
			State.nextTick = now.Add(LogPace)
		}
		if now.Before(State.nextTick) {
			c.After(State.nextTick.Sub(now))
		} else {
			State.Revealed++
			State.nextTick = now.Add(LogPace)
			c.After(LogPace)
		}
	}
}

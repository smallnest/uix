package mujica

import (
	"github.com/egoist/mygo/ui"
)

// Page is one component's gallery entry.
type Page struct {
	ID       string
	Name     string
	ZhName   string
	Category string
	// Demo builds the interactive example; log records events.
	Demo func(c *ui.Context, log func(string))
	Code string
	Keys string
}

func catalog() []Page {
	var all []Page
	for _, g := range [][]Page{layoutPages, displayPages, inputPages, choicePages, formPages, datetimePages, navigationPages, dataPages, overlayPages, feedbackPages,
		charts1Pages, charts2Pages, mktchartPages, tradingPages, marketsPages, chatmsgPages, chatinputPages, chatsessionPages, agentPages, motionPages, gitPages, codePages, mediaPages, projectsPages, filesPages, messagingPages, calendarPages, devtoolsPages, accountPages, essentialsPages} {
		all = append(all, g...)
	}
	return all
}

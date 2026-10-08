// Package link provides a text that opens a URL in the browser when it
// is clicked, as a shadcn/ui component for the Link of MyGo.
//
//	link.Link(c, link.Props{
//		Label: "Read the guide",
//		URL:   "https://example.com/guide",
//	})
//
// The label shows in the accent color, and gains an underline when the
// pointer rests on it. In a page of a Router, a URL without a scheme, as
// "/notes/42", goes there in the router instead of the browser.
package link

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the link to draw.
type Props struct {
	// Label is the text of the link.
	Label string
	// URL is what the link opens: an address for the browser, or a path
	// for the Router the page lives in.
	URL string
}

// Link draws the label as a link that opens URL when clicked or Entered,
// and returns it.
func Link(c *ui.Context, p Props) ui.Element {
	return ui.Link(c, p.Label, p.URL)
}

// Package tree provides a hierarchical list whose branches open and
// close, from the mygo Tree and TreeItem widgets. A node with Children
// is a branch, one without them a leaf. Choosing a node highlights it
// and runs its action.
package tree

import (
	"github.com/egoist/mygo/ui"
)

// Item is one node of the tree.
type Item struct {
	// Label is the text of the node. Labels must differ, as the chosen
	// node is found by its label.
	Label string
	// Open opens and closes a branch. A branch — an item with Children —
	// needs it; the tree flips it when the arrow is clicked.
	Open *bool
	// Children are the nodes inside the branch.
	Children []Item
	// Action runs when the node is chosen, by a click or Enter.
	Action func()
}

// Props configure the tree.
type Props struct {
	// Items are the nodes of the tree.
	Items []Item
	// Chosen is the label of the chosen node. The tree sets it when a
	// node is chosen and shows the node highlighted. nil keeps the tree
	// from highlighting.
	Chosen *string
	// Disabled keeps the tree from opening or choosing nodes.
	Disabled bool
}

// Tree draws the items as a tree. A click on the arrow of a branch
// opens or closes it; a click on a node, or Enter, chooses it. While a
// node has the keyboard focus, Up and Down move around it, Right opens
// it or moves to its first child, and Left closes it or moves to its
// parent.
func Tree(c *ui.Context, p Props) ui.Element {
	t := ui.Tree(c, func() {
		for _, it := range p.Items {
			build(c, it, p)
		}
	})
	if p.Disabled {
		t.Disabled(true)
	}
	return t
}

// build draws a node, the nodes inside it, and applies the choice.
func build(c *ui.Context, it Item, p Props) {
	item := ui.TreeItem(c, it.Label, it.Open, func() {
		for _, child := range it.Children {
			build(c, child, p)
		}
	})
	if p.Chosen != nil {
		item.Selected(*p.Chosen == it.Label)
	}
	if item.Clicked() {
		if p.Chosen != nil {
			*p.Chosen = it.Label
		}
		if it.Action != nil {
			it.Action()
		}
	}
}

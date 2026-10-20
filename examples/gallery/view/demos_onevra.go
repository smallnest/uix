package view

import (
	"fmt"
	"image"
	"image/color"

	"github.com/egoist/mygo/ui"

	"github.com/oneo-me/onevra/source/components"
	"github.com/oneo-me/onevra/source/components/flow"
	"github.com/oneo-me/onevra/source/components/treemap"

	"github.com/smallnest/uix/components/lucide"
)

// This file draws the pages of the onevra components its author added in
// v0.2026.13: AvatarGroup, CardStack, FloatingButtonGroup, HoldToConfirm,
// ImageCompare, ScrollEdgeBar, the flow diagram and the treemap. Their
// demos keep the uix gallery's compact, mostly static style.

// demoAvatarGroup shows onevra's overlapping avatars, the rest collapsing
// into a +N chip.
func demoAvatarGroup(c *ui.Context, t *ui.Theme) {
	components.AvatarGroup(c, []components.AvatarGroupMember{
		{Name: "Ada Lovelace", Status: components.AvatarOnline},
		{Name: "Alan Turing"},
		{Name: "Grace Hopper", Status: components.AvatarOffline},
		{Name: "Margaret Hamilton"},
		{Name: "Katherine Johnson"},
		{Name: "Dennis Ritchie"},
		{Name: "Linus Torvalds"},
	}, components.AvatarGroupOptions{Max: 5, Label: "团队成员"})
}

// demoCardStack shows onevra's swipeable deck of shortlist cards.
func demoCardStack(c *ui.Context, t *ui.Theme) {
	face := func(name, place, note string, icon func() *ui.SVG, tint ui.Color) func(*ui.Context, float32) {
		return func(c *ui.Context, radius float32) {
			ui.Column(c).Fill().Children(func() {
				ui.Box(c).FillWidth().Height(110).Center().Radius(radius, radius, 0, 0).
					LinearGradient(ui.LinearGradient{From: tint.Alpha(.32), To: tint.Alpha(.08), Angle: 135}).
					Children(func() {
						ui.Icon(c, icon()).Size(30, 30).TextColor(tint)
					})
				ui.Column(c).FillWidth().Padding(t.Space(4)).Gap(t.Space(1)).Children(func() {
					ui.Text(c, name).FontWeight(600)
					ui.Text(c, place).TextColor(t.TextMuted)
					ui.Text(c, note).TextColor(t.TextMuted).Wrap()
				})
			})
		}
	}
	components.CardStack(c, &Demo.OCardStack, components.CardStackOptions{
		Label:        "旅行候选",
		LeftLabel:    "跳过",
		RightLabel:   "想去",
		LeftOutcome:  "已跳过",
		RightOutcome: "已想去",
		Items: []components.CardStackItem{
			{ID: "lisbon", Label: "里斯本", Content: face("里斯本", "葡萄牙", "电车在七座山丘间穿行。", lucide.TramFront, ui.RGB(226, 132, 72))},
			{ID: "kyoto", Label: "京都", Content: face("京都", "日本", "枫叶季的寺庙与庭园。", lucide.Landmark, ui.RGB(206, 86, 110))},
			{ID: "reykjavik", Label: "雷克雅未克", Content: face("雷克雅未克", "冰岛", "午夜太阳与蓝湖。", lucide.Snowflake, ui.RGB(72, 150, 214))},
			{ID: "queenstown", Label: "皇后镇", Content: face("皇后镇", "新西兰", "湖畔跳伞与远足。", lucide.Mountain, ui.RGB(126, 112, 214))},
		},
	})
}

// demoFloatingButtonGroup shows onevra's floating tray of a board's
// actions; undo and redo stay disabled until a click enables the other.
func demoFloatingButtonGroup(c *ui.Context, t *ui.Theme) {
	items := []components.FloatingButtonGroupItem{
		{ID: "undo", Label: "撤销", Icon: lucide.Undo2, IconOnly: true, Shortcut: "⌘Z", Disabled: Demo.OFloating.edits == 0},
		{ID: "redo", Label: "重做", Icon: lucide.Redo2, IconOnly: true, Shortcut: "⇧⌘Z", Disabled: Demo.OFloating.undone == 0},
		{Separator: true},
		{ID: "comment", Label: "评论", Icon: lucide.MessageCircle, Shortcut: "C"},
		{ID: "present", Label: "演示", Icon: lucide.Play, Toggle: true, On: Demo.OFloating.presenting},
	}
	switch components.FloatingButtonGroup(c, "画布操作", components.FloatingButtonGroupMuted, ui.Horizontal, items).Clicked() {
	case "undo":
		Demo.OFloating.edits--
		Demo.OFloating.undone++
	case "redo":
		Demo.OFloating.edits++
		Demo.OFloating.undone--
	case "present":
		Demo.OFloating.presenting = !Demo.OFloating.presenting
	}
}

// demoHoldToConfirm shows onevra's long-press button, which completes a
// destructive action only after the hold fills; the reset button restores it.
func demoHoldToConfirm(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Wrap().Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
		h := components.HoldToConfirm(c, &Demo.OHold, "长按归档", components.HoldToConfirmOptions{ConfirmedLabel: "已归档", Icon: lucide.Archive()})
		switch {
		case h.Holding():
			ui.Text(c, "继续按住…").TextColor(t.TextMuted)
		case Demo.OHold:
			if components.Button(c, "重置", components.ButtonGhost).Clicked() {
				Demo.OHold = false
			}
		}
	})
}

// compareColor and compareGray are the picture the image-compare demo
// slides between: a small landscape painted once, color on one side and
// its luminance on the other, so the demo needs no asset.
var (
	compareColor = compareBitmap(false)
	compareGray  = compareBitmap(true)
)

// compareBitmap paints a landscape of sky, sun and hills of w×h pixels;
// gray draws the same picture in luminance.
func compareBitmap(gray bool) ui.ImageSource {
	const w, h = 480, 300
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	horizon := h * 2 / 3
	lum := func(r, g, b uint8) uint8 {
		return uint8(0.299*float32(r) + 0.587*float32(g) + 0.114*float32(b))
	}
	lerp := func(a, b uint8, f float32) uint8 { return uint8(float32(a) + (float32(b)-float32(a))*f) }
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			var r, g, b uint8
			if y < horizon {
				f := float32(y) / float32(horizon)
				r, g, b = lerp(48, 150, f), lerp(96, 205, f), lerp(190, 235, f)
			} else {
				f := float32(y-horizon) / float32(h-horizon)
				r, g, b = lerp(88, 52, f), lerp(148, 92, f), lerp(92, 56, f)
			}
			dx, dy := float32(x-380), float32(y-64)
			if dx*dx+dy*dy < 46*46 {
				r, g, b = 250, 222, 120
			}
			if gray {
				v := lum(r, g, b)
				r, g, b = v, v, v
			}
			img.SetRGBA(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}
	return ui.NewBitmap(img)
}

// demoImageCompare shows onevra's slider between the color picture and
// its grayscale, so both sides stay aligned without a second asset.
func demoImageCompare(c *ui.Context, t *ui.Theme) {
	side := func(src ui.ImageSource) func() {
		return func() {
			components.Image(c, src).Fill().Fit(ui.Cover)
		}
	}
	components.ImageCompare(c, &Demo.OCompare, side(compareGray), side(compareColor), components.ImageCompareOptions{
		BeforeLabel: "原图",
		AfterLabel:  "灰阶",
	}).FillWidth().MaxWidth(520)
}

// demoScrollEdgeBar docks a toolbar over a scrolling list, its surface
// fading over the content near the edge.
func demoScrollEdgeBar(c *ui.Context, t *ui.Theme) {
	inset := components.ScrollEdgeBarExtent(c, components.ScrollEdgeBarOptions{})
	ui.Box(c).FillWidth().MinWidth(0).Height(320).Radius(t.Radius*2).Clip().
		Background(t.Background).Border(1, t.Border).Children(func() {
		components.ScrollView(c, components.ScrollViewOptions{}).Fill().Padding(inset+8, 16).
			ScrollbarInsets(inset, 0).Gap(10).Children(func() {
			for i := range 12 {
				ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
					ui.Box(c).Size(48, 36).Radius(8).Shrink(0).Background(t.SurfaceHover)
					ui.Text(c, fmt.Sprintf("第 %d 行", i+1)).TextColor(t.TextMuted)
				})
			}
		})
		components.ScrollEdgeBar(c, components.ScrollEdgeBarOptions{}, func() {
			components.IconButton(c, lucide.ChevronLeft(), "返回", components.IconButtonPlain)
			ui.Text(c, "文件列表").Bold().Grow(1)
			components.IconButton(c, lucide.Share(), "分享", components.IconButtonPlain)
		}).Label("工具栏")
	})
}

// demoFlow shows onevra's flow diagram of a plan's steps, laid out left
// to right and linked by their edges.
func demoFlow(c *ui.Context, t *ui.Theme) {
	card := func(title, desc string) func(*ui.Context) {
		return func(c *ui.Context) {
			flow.Card(c, flow.CardOptions{Title: title, Description: desc})
		}
	}
	flow.View(c, &Demo.OFlow, flow.Data{
		Nodes: []flow.Node{
			{ID: "research", Label: "调研", Content: card("调研", "理解需求")},
			{ID: "design", Label: "设计", Content: card("设计", "原型与走查")},
			{ID: "build", Label: "开发", Content: card("开发", "实现与自测")},
			{ID: "release", Label: "发布", Content: card("发布", "灰度与上线")},
		},
		Edges: []flow.Edge{
			{Source: "research", Target: "design"},
			{Source: "design", Target: "build"},
			{Source: "build", Target: "release"},
		},
	}, flow.Options{Height: 240})
}

// demoTreemap shows onevra's treemap of a disk's folders, sized by
// occupied space; clicking a tile drills into it and breadcrumbs lead
// back.
func demoTreemap(c *ui.Context, t *ui.Theme) {
	treemap.View(c, &Demo.OTreemap, []treemap.Node{
		{ID: "photos", Label: "照片", Value: 3200, Children: []treemap.Node{
			{ID: "photos/2026", Label: "2026", Value: 1800},
			{ID: "photos/old", Label: "往年", Value: 1400},
		}},
		{ID: "docs", Label: "文档", Value: 1200},
		{ID: "videos", Label: "视频", Value: 2400},
		{ID: "music", Label: "音乐", Value: 600},
	}, treemap.Options{Label: "磁盘占用", Height: 260})
}

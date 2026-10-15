package view

import (
	"fmt"
	"time"

	"github.com/egoist/mygo/ui"
	mchart "github.com/egoist/mygo/ui/chart"

	"github.com/oneo-me/onevra/source/components"
	"github.com/oneo-me/onevra/source/components/chart"
	"github.com/oneo-me/onevra/source/components/graph"

	"github.com/smallnest/uix/components/animate"
	"github.com/smallnest/uix/components/avatar"
	"github.com/smallnest/uix/components/badge"
	"github.com/smallnest/uix/components/barchart"
	"github.com/smallnest/uix/components/blockquote"
	"github.com/smallnest/uix/components/button"
	"github.com/smallnest/uix/components/card"
	"github.com/smallnest/uix/components/empty"
	"github.com/smallnest/uix/components/gridview"
	"github.com/smallnest/uix/components/headingwithanchor"
	"github.com/smallnest/uix/components/icon"
	"github.com/smallnest/uix/components/image"
	"github.com/smallnest/uix/components/infinitescroll"
	"github.com/smallnest/uix/components/inlinecode"
	"github.com/smallnest/uix/components/item"
	"github.com/smallnest/uix/components/kbd"
	"github.com/smallnest/uix/components/label"
	"github.com/smallnest/uix/components/linechart"
	"github.com/smallnest/uix/components/list"
	"github.com/smallnest/uix/components/lucide"
	"github.com/smallnest/uix/components/marker"
	"github.com/smallnest/uix/components/richtext"
	"github.com/smallnest/uix/components/skeleton"
	"github.com/smallnest/uix/components/statcard"
	"github.com/smallnest/uix/components/table"
	"github.com/smallnest/uix/components/tree"
	"github.com/smallnest/uix/components/webview"
)

// demoText shows mygo's text primitive.
func demoText(c *ui.Context, t *ui.Theme) {
	ui.Text(c, "你好，这是 mygo 的 Text 原语。")
}

// demoBox shows mygo's box container.
func demoBox(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(220).Padding(t.Space(3)).Radius(t.Radius).Background(t.Surface).
		Border(1, t.Border).Children(func() {
		ui.Text(c, "Box 把子元素按列排布。")
	})
}

// demoRow shows mygo's row container: children side by side.
func demoRow(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Gap(t.Space(2)).Children(func() {
		ui.Box(c).Size(24, 24).Radius(6).Background(t.Accent)
		ui.Box(c).Size(24, 24).Radius(6).Background(t.Success)
		ui.Box(c).Size(24, 24).Radius(6).Background(t.Warning)
	})
}

// demoColumn shows mygo's column container.
func demoColumn(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(1.5)).Children(func() {
		ui.Text(c, "第一行")
		ui.Text(c, "第二行")
		ui.Text(c, "第三行")
	})
}

// demoSpacer shows mygo's spacer pushing a sibling to the far end.
func demoSpacer(c *ui.Context, t *ui.Theme) {
	ui.Row(c).FillWidth().Gap(t.Space(2)).Children(func() {
		ui.Text(c, "左侧")
		ui.Spacer(c)
		ui.Text(c, "右侧")
	})
}

// demoIcon shows an SVG shown as an icon in the text color.
func demoIcon(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
		icon.Icon(c, icon.Props{SVG: lucide.Search(), Label: "搜索", Width: 20, Height: 20})
		icon.Icon(c, icon.Props{SVG: lucide.Bell(), Label: "通知", Width: 20, Height: 20})
		icon.Icon(c, icon.Props{SVG: lucide.Settings(), Label: "设置", Width: 20, Height: 20})
	})
}

// demoLucide shows icons of the Lucide set at different sizes.
func demoLucide(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
		for _, f := range []func() *ui.SVG{lucide.House, lucide.Mail, lucide.Users, lucide.Palette} {
			icon.Icon(c, icon.Props{SVG: f(), Width: 20, Height: 20})
		}
	})
}

// demoImage shows a picture fitted into a box, here an SVG.
func demoImage(c *ui.Context, t *ui.Theme) {
	image.Image(c, image.Props{Src: lucide.Palette(), Width: 64, Height: 64})
}

// demoAvatar shows a circle with initials or a photo.
func demoAvatar(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Gap(t.Space(3)).Children(func() {
		avatar.Avatar(c, avatar.Props{Name: "Ada Lovelace", Size: 40})
		avatar.Avatar(c, avatar.Props{Name: "Alan Turing", Size: 40})
	})
}

// demoBadge shows the four looks of a badge.
func demoBadge(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Wrap().Gap(t.Space(2)).Children(func() {
		badge.Badge(c, badge.Props{Label: "v1.2.0"})
		badge.Badge(c, badge.Props{Label: "次要", Variant: badge.Secondary})
		badge.Badge(c, badge.Props{Label: "Beta", Variant: badge.Outline})
		badge.Badge(c, badge.Props{Label: "需要关注", Variant: badge.Destructive})
	})
}

// demoKbd shows a key of the keyboard as a chip.
func demoKbd(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Gap(t.Space(2)).Children(func() {
		kbd.Kbd(c, kbd.Props{Label: "⌘"})
		kbd.Kbd(c, kbd.Props{Label: "K"})
	})
}

// demoLabel shows a standalone label of a control.
func demoLabel(c *ui.Context, t *ui.Theme) {
	label.Label(c, label.Props{Text: "邮箱地址"})
}

// demoMarker shows short inline labels of a turn.
func demoMarker(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Gap(t.Space(2)).Children(func() {
		marker.Marker(c, marker.Props{Label: "系统"})
		marker.Marker(c, marker.Props{Label: "描边", Variant: marker.Border})
	})
}

// demoInlineCode shows a code chip in the prose.
func demoInlineCode(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Gap(t.Space(1.5)).AlignItems(ui.Center).Children(func() {
		ui.Text(c, "运行").TextColor(t.TextMuted)
		inlinecode.InlineCode(c, inlinecode.Props{Text: "go run ./examples/gallery"})
	})
}

// demoBlockquote shows a quote on a muted face with an accent bar.
func demoBlockquote(c *ui.Context, t *ui.Theme) {
	blockquote.Blockquote(c, blockquote.Props{
		Text:   "预测未来最好的方式是创造它。",
		Author: "— Alan Kay",
	})
}

// demoHeadingWithAnchor shows a heading with a link that copies the
// anchor.
func demoHeadingWithAnchor(c *ui.Context, t *ui.Theme) {
	headingwithanchor.Heading(c, headingwithanchor.Props{Text: "安装", Anchor: "install"})
}

// demoRichText shows styled runs of text.
func demoRichText(c *ui.Context, t *ui.Theme) {
	richtext.RichText(c, richtext.Props{Spans: []ui.Span{
		{Text: "一段"},
		{Text: "加粗", Weight: 600},
		{Text: "、一段"},
		{Text: "彩色", Color: t.Accent},
		{Text: "的文字。"},
	}})
}

// demoCard shows a rounded container with a header and a body.
func demoCard(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(280).Children(func() {
		card.Card(c, card.Props{Title: "卡片", Description: "带圆角的内容容器。"}, func() {
			ui.Text(c, "卡片的正文。")
		})
	})
}

// demoStatCard shows metric cards with a delta.
func demoStatCard(c *ui.Context, t *ui.Theme) {
	statcard.StatCards(c, statcard.Props{Stats: []statcard.Stat{
		{Icon: statcard.IconUsers, Label: "客户", Value: "14,592", Delta: "+5.3%", DeltaColor: statcard.Up},
		{Icon: statcard.IconBox, Label: "售出", Value: "385", Delta: "-2.1%", DeltaColor: statcard.Down},
		{Icon: statcard.IconBasket, Label: "订单", Value: "1,394", Delta: "0.00%", DeltaColor: statcard.Flat},
	}})
}

// demoSkeleton shows placeholders of content still loading.
func demoSkeleton(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		skeleton.Skeleton(c, skeleton.Props{Width: 200, Height: 16, Radius: 4})
		skeleton.Skeleton(c, skeleton.Props{Width: 160, Height: 16, Radius: 4})
	})
}

// demoEmpty shows an empty state: an icon over a title and a description.
func demoEmpty(c *ui.Context, t *ui.Theme) {
	empty.Empty(c, empty.Props{
		Icon:        lucide.Search(),
		Title:       "没有结果",
		Description: "换个关键词试试。",
	})
}

// demoList shows a scrolling list of rows, one chosen.
func demoList(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Height(180).Children(func() {
		list.List(c, list.Props{Items: []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}})
	})
}

// demoGridView shows a grid of items that a click chooses.
func demoGridView(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Height(200).Children(func() {
		gridview.GridView(c, gridview.Props{
			State:    &Demo.Grid,
			Count:    6,
			MinWidth: 100,
			Height:   64,
			Item: func(c *ui.Context, i int) {
				ui.Column(c).Center().Gap(4).Children(func() {
					ui.Box(c).Size(28, 28).Radius(8).Background(t.Accent.Alpha(1 - float32(i)*0.08))
					ui.Text(c, string(rune('A'+i))).FontSize(t.FontSize * 0.8).TextColor(t.TextMuted)
				})
			},
		})
	})
}

// demoTable shows a data table with a header band.
func demoTable(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Height(180).Children(func() {
		table.Table(c, table.Props{
			Columns: []ui.TableColumn{
				{Title: "姓名"},
				{Title: "角色"},
				{Title: "状态"},
			},
			Rows: 4,
			Cell: func(row, col int) {
				names := []string{"Ada Lovelace", "Alan Turing", "Grace Hopper", "Linus Torvalds"}
				roles := []string{"分析师", "数学家", "海军上将", "工程师"}
				statuses := []string{"在线", "在线", "离开", "在线"}
				cell := [][]string{names, roles, statuses}[col]
				ui.Text(c, cell[row]).SingleLine()
			},
		})
	})
}

// demoTree shows a hierarchical list whose branches open and close.
func demoTree(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Height(200).Children(func() {
		tree.Tree(c, tree.Props{
			Items: []tree.Item{
				{Label: "src", Open: &Demo.Tree1, Children: []tree.Item{{Label: "main.go"}, {Label: "view"}}},
				{Label: "README.md"},
				{Label: "go.mod"},
			},
		})
	})
}

// demoOutline shows mygo's collapsible outline of a hierarchy.
func demoOutline(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Height(180).Width(220).Children(func() {
		ui.Outline(c, &ui.OutlineState[string]{}, []string{"src", "go.mod"},
			func(k string) []string {
				if k == "src" {
					return []string{"main.go", "view.go"}
				}
				return nil
			},
			func(k string) { ui.Text(c, k).SingleLine() })
	})
}

// demoLineChart shows a line chart of a year against the year before.
func demoLineChart(c *ui.Context, t *ui.Theme) {
	linechart.LineChart(c, linechart.Props{
		Title: "营收",
		Data: []linechart.Point{
			{Label: "Jan", Current: 40, Previous: 30},
			{Label: "Feb", Current: 52, Previous: 38},
			{Label: "Mar", Current: 46, Previous: 44},
			{Label: "Apr", Current: 68, Previous: 50},
			{Label: "May", Current: 61, Previous: 55},
			{Label: "Jun", Current: 82, Previous: 62},
		},
	}).Height(240)
}

// demoBarChart shows a paired bar chart against a year earlier.
func demoBarChart(c *ui.Context, t *ui.Theme) {
	barchart.BarChart(c, barchart.Props{
		Title: "订单",
		Data: []barchart.Point{
			{Label: "Jan", Current: 22, Previous: 18},
			{Label: "Feb", Current: 28, Previous: 20},
			{Label: "Mar", Current: 26, Previous: 24},
			{Label: "Apr", Current: 34, Previous: 26},
			{Label: "May", Current: 31, Previous: 30},
			{Label: "Jun", Current: 41, Previous: 33},
		},
	}).Height(240)
}

// demoDonutChart shows a onevra donut of proportions.
func demoDonutChart(c *ui.Context, t *ui.Theme) {
	chart.Donut(c, &chart.State{Index: -1}, []chart.Slice{
		{ID: "design", Label: "设计", Value: 34},
		{ID: "development", Label: "开发", Value: 28},
		{ID: "research", Label: "研究", Value: 20},
		{ID: "planning", Label: "规划", Value: 18},
	}, chart.Options{PlotHeight: 200}).Height(240)
}

// demoAreaChart shows a onevra area chart of a series filling its
// baseline.
func demoAreaChart(c *ui.Context, t *ui.Theme) {
	chart.Area(c, &chart.State{Index: -1}, chart.Data{
		Labels: []string{"一月", "二月", "三月", "四月", "五月", "六月"},
		Series: []chart.Series{
			{ID: "revenue", Label: "营收", Values: []float64{64, 90, 48, 110, 76, 95}},
		},
	}, chart.Options{PlotHeight: 200}).Height(240)
}

// demoStackedBar shows a onevra bar chart of series stacked one on
// another.
func demoStackedBar(c *ui.Context, t *ui.Theme) {
	chart.StackedBar(c, &chart.State{Index: -1}, chart.Data{
		Labels: []string{"一月", "二月", "三月"},
		Series: []chart.Series{
			{ID: "design", Label: "设计", Values: []float64{34, 42, 28}},
			{ID: "development", Label: "开发", Values: []float64{48, 55, 60}},
		},
	}, chart.Options{PlotHeight: 200}).Height(240)
}

// demoGraph shows a onevra relationship graph of nodes and edges.
func demoGraph(c *ui.Context, t *ui.Theme) {
	graph.View(c, &graph.State{}, graph.Data{
		Nodes: []graph.Node{
			{ID: "ada", Label: "Ada"},
			{ID: "alan", Label: "Alan"},
			{ID: "grace", Label: "Grace"},
		},
		Edges: []graph.Edge{
			{Source: "ada", Target: "alan"},
			{Source: "alan", Target: "grace"},
			{Source: "grace", Target: "ada"},
		},
	}, graph.Options{Height: 220})
}

// demoPixelGrid shows a onevra squared backdrop under content.
func demoPixelGrid(c *ui.Context, t *ui.Theme) {
	components.PixelGrid(c, components.PixelGridOptions{Spacing: 18}).FillWidth().Height(120).
		Background(t.Background).Border(1, t.Border).Children(func() {
		ui.Column(c).Fill().Center().Children(func() {
			ui.Text(c, "方格背景").TextColor(t.TextMuted)
		})
	})
}

// demoAnimate shows elements that move in on their own: cards that
// enter one after another, and a pulse that breathes. The replay button
// plays the entrances again.
func demoAnimate(c *ui.Context, t *ui.Theme) {
	card := func(label string) ui.Element {
		return ui.Box(c).Size(96, 96).Radius(t.Radius * 2).Background(t.Accent.Alpha(0.82)).Center().
			Children(func() {
				ui.Text(c, label).TextColor(t.AccentText)
			})
	}
	ui.Column(c).Gap(t.Space(3)).Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(t.Space(2)).Children(func() {
			animate.Pulse(ui.Box(c).Size(10, 10).Radius(5).Background(t.Danger), animate.Options{Duration: 1400 * time.Millisecond})
			ui.Text(c, "进入时依次浮现。").TextColor(t.TextMuted)
			ui.Spacer(c)
			button.Button(c, button.Props{Label: "重播", Variant: button.Outline, OnClick: func() { Demo.Run++ }})
		})
		ui.Row(c).Gap(t.Space(3)).Children(func() {
			animate.FadeIn(card("淡入"), animate.Options{Run: Demo.Run})
			animate.RiseIn(card("升起"), animate.Options{Delay: 120 * time.Millisecond, Run: Demo.Run})
			animate.SlideIn(card("滑入"), animate.Options{Delay: 240 * time.Millisecond, Run: Demo.Run}, animate.Left)
			animate.ScaleIn(card("缩放"), animate.Options{Delay: 360 * time.Millisecond, Run: Demo.Run})
		})
	})
}

// demoInfiniteScroll shows a feed that appends a page as it nears the
// end; the counter names the rows appended.
func demoInfiniteScroll(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(360).Height(220).Children(func() {
		infinitescroll.InfiniteScroll(c, infinitescroll.Props{
			State: &Demo.Feed,
			OnMore: func() {
				Demo.FeedRows = append(Demo.FeedRows, feedRows(Demo.FeedN, 8)...)
				Demo.FeedN += 8
			},
			Children: func() {
				ui.Column(c).FillWidth().Children(func() {
					for _, name := range Demo.FeedRows {
						ui.Row(c).FillWidth().Height(32).AlignItems(ui.Center).Padding(0, 12).Children(func() {
							ui.Text(c, name).Grow(1)
						})
					}
				})
			},
		})
	})
}

// demoItem shows a row of a list: a tile, a title, a description and
// the actions at the far end.
func demoItem(c *ui.Context, t *ui.Theme) {
	item.ItemGroup(c, []item.Props{
		{Media: lucide.User(), Title: "账户", Description: "你的名字和照片", OnClick: func() {}},
		{Media: lucide.Bell(), Title: "通知", Description: "提醒你的内容", OnClick: func() {}},
		{Media: lucide.Shield(), Title: "隐私", Description: "应用可以看到什么", OnClick: func() {}},
	})
}

// demoWebView shows the page of the window's web view, with a reload
// button. Headless there is no web view, so the demo draws a
// placeholder that names the page.
func demoWebView(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(520).Gap(t.Space(2)).Children(func() {
		ui.Row(c).FillWidth().MinWidth(0).AlignItems(ui.Center).Gap(t.Space(2)).Children(func() {
			ui.Text(c, "WebView").Bold()
			ui.Text(c, "https://example.com").Grow(1).MinWidth(0).SingleLine().TextColor(t.TextMuted)
			if Web != nil {
				button.Button(c, button.Props{Label: "刷新", Variant: button.Ghost, OnClick: func() {
					Web.Page().Reload()
				}})
			}
		})
		ui.Box(c).FillWidth().Height(240).Radius(t.Radius).Border(1, t.Border).Children(func() {
			if Web != nil {
				webview.WebView(c, webview.Props{Web: Web}).Fill()
			} else {
				ui.Box(c).Fill().Center().Children(func() {
					ui.Text(c, "打开窗口时,这里显示 example.com 的网页。").TextColor(t.TextMuted)
				})
			}
		})
	})
}

// demoMotion shows mygo's own animation engine: a card that a Keyframes
// sequence fades and grows in on replay, and a dot that a Loop keeps
// breathing.
func demoMotion(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(3)).Children(func() {
		ui.Row(c).Gap(t.Space(4)).Children(func() {
			b := ui.Box(c).Size(96, 96).Radius(t.Radius * 2).Background(t.Accent.Alpha(0.82)).Center().
				Children(func() {
					ui.Text(c, "Keyframes").TextColor(t.AccentText)
				})
			p := b.Keyframes("demo.motion.in", []ui.Keyframe{
				{At: 0, Value: 0},
				{At: 400 * time.Millisecond, Value: 1, Ease: ui.EaseOut},
			}, ui.KeyframeOptions{Run: Demo.MotionRun})
			b.Opacity(p).Transform(ui.Scaling(0.5+0.5*p, 0.5+0.5*p)).TransformOrigin(0.5, 0.5)
			d := ui.Box(c).Size(24, 24).Radius(12).Background(t.Success)
			phase := d.Loop("demo.motion.pulse", 1200*time.Millisecond, ui.Bounce(ui.EaseInOut))
			d.Opacity(0.35 + 0.65*phase)
		})
		ui.Row(c).AlignItems(ui.Center).Gap(t.Space(2)).Children(func() {
			button.Button(c, button.Props{Label: "重播", Variant: button.Outline, OnClick: func() { Demo.MotionRun++ }})
			ui.Text(c, "mygo 的 Keyframes 与 Loop。").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		})
	})
}

// demoChart shows a bar plot built from mygo's chart scales: the band
// scale places the bars, the linear scale their heights.
func demoChart(c *ui.Context, t *ui.Theme) {
	x := mchart.NewBand("周一", "周二", "周三", "周四", "周五").Padding(0.3, 0.15)
	y := mchart.NewLinear(0, 120).Nice(4)
	values := []float64{64, 90, 48, 110, 76}
	mchart.Plot(c, mchart.Axes{Left: y.Ticks(4), Bottom: x.Ticks()}, func(area ui.Element) {
		for i, v := range values {
			mchart.Rect(c, x.Band(i), y.Span(0, v)).Background(t.Accent).Radius(4)
		}
	}).Height(200).Gap(6)
}

// demoTimePicker shows onevra's clock dial choosing a time of day.
func demoTimePicker(c *ui.Context, t *ui.Theme) {
	components.TimePicker(c, &Demo.OTime).Label("时间")
}

// demoColorView shows onevra's color as its value in the color formats,
// each editable.
func demoColorView(c *ui.Context, t *ui.Theme) {
	components.ColorView(c, &Demo.OColor).Label("颜色")
}

// demoProgressRing shows onevra's circular progress at a completion and
// at work whose end it does not know.
func demoProgressRing(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
		components.ProgressRing(c, 0.64)
		components.ProgressRing(c, -1)
	})
}

// demoTableVirtualized shows onevra's table that builds only the rows on
// screen: ten thousand files scroll without a row for every file.
func demoTableVirtualized(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(520).Height(240).Children(func() {
		components.TableVirtualized(c, &Demo.OTable, []ui.TableColumn{
			{Title: "名称"},
			{Title: "大小", Width: 90},
		}, 10000, func(row, col int) {
			if col == 0 {
				ui.Text(c, fmt.Sprintf("文件_%d.go", row+1)).SingleLine()
			} else {
				ui.Text(c, fmt.Sprintf("%d KB", row*7919%977+1)).SingleLine()
			}
		}).FillWidth()
	})
}

// demoTreeVirtualized shows onevra's tree that builds only the rows on
// screen: the folders hold many files, yet only those in view are rows.
func demoTreeVirtualized(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(320).Height(240).Children(func() {
		components.TreeVirtualized(c, &Demo.OTree, []string{"src", "docs"},
			func(k string) []string {
				if k == "src" {
					// Enough files that scrolling matters.
					names := make([]string, 2000)
					for i := range names {
						names[i] = fmt.Sprintf("file_%04d.go", i+1)
					}
					return names
				}
				if k == "docs" {
					return []string{"guide.md", "api.md"}
				}
				return nil
			},
			func(k string) { ui.Text(c, k).SingleLine() })
	})
}

// demoOutlineTable shows mygo's outline as a table of columns, folders
// and files opening and closing.
func demoOutlineTable(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(360).Height(200).Children(func() {
		ui.OutlineTable(c, &Demo.OutlineT, []ui.TableColumn{
			{Title: "名称", Sortable: true},
			{Title: "类型", Width: 100},
		}, []string{"src", "README.md"},
			func(k string) []string {
				if k == "src" {
					return []string{"main.go", "view.go"}
				}
				return nil
			},
			func(path string, col int) {
				if col == 0 {
					ui.Text(c, path).SingleLine()
				} else if path == "src" {
					ui.Text(c, "目录")
				} else {
					ui.Text(c, "文件")
				}
			})
	})
}

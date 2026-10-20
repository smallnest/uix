package view

import (
	"image"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
)

// gaugePrimitives is every primitive a gauge face can compose, named as
// its method is, so a test can look for the calls rather than the
// pixels. The gauge package exports exactly these; a new one added there
// and never drawn in the gallery shows up as a row no page satisfies.
var gaugePrimitives = []string{
	"Arc", "Dot", "Hub", "Inset", "Marks", "Needle", "Stack", "Text",
	"TickFill", "TickLabels", "Ticks", "Track", "Value", "Zones",
}

// TestGaugePagesCoverEveryPrimitive reads the scene pages and requires
// every gauge primitive to be called by at least one of them. This is
// the check that fails when a primitive is dropped from the gallery,
// which is the gap the scene pages were added to close: the one page
// the gallery used to have drew ten of the fourteen and nothing said so.
//
// It reads the source rather than the pixels on purpose. Every primitive
// is a drawing call inside a shared draw callback, so its ink is mixed
// with the ink of the primitives around it and cannot be counted from
// outside; the call itself is the only honest evidence of its use. The
// pixels are covered by TestGaugeSceneDialsDraw, which requires each
// dial to paint.
func TestGaugePagesCoverEveryPrimitive(t *testing.T) {
	files, err := filepath.Glob("demos_gauge*.go")
	if err != nil {
		t.Fatal(err)
	}
	var src strings.Builder
	for _, name := range files {
		if strings.HasSuffix(name, "_test.go") {
			continue
		}
		b, err := os.ReadFile(name)
		if err != nil {
			t.Fatal(err)
		}
		src.Write(b)
	}
	text := src.String()
	for _, p := range gaugePrimitives {
		if !strings.Contains(text, "."+p+"(gauge."+p) {
			t.Errorf("no scene page draws %s: no call to .%s(gauge.%s…)", p, p, p)
		}
	}
}

// ink counts the pixels inside r that are not the surface color, so a
// test can tell a drawn dial from an empty box. The gauge glyphs are
// painted rather than laid out as text, so they are not in the
// accessibility tree: the pixels are the only evidence they drew.
func ink(img *image.RGBA, r ui.Rect, surface ui.Color) int {
	b := img.Bounds()
	x0, y0 := max(int(r.X), b.Min.X), max(int(r.Y), b.Min.Y)
	x1, y1 := min(int(r.X+r.W), b.Max.X), min(int(r.Y+r.H), b.Max.Y)
	count := 0
	for y := y0; y < y1; y++ {
		for x := x0; x < x1; x++ {
			if !sameColor(img.RGBAAt(x, y), surface) {
				count++
			}
		}
	}
	return count
}

// TestGaugeSceneDialsDraw checks every dial of every scene page leaves
// ink on the canvas, in both themes. A page can keep its dial elements
// and still paint nothing inside them — a draw callback that returns
// early, a sweep collapsed to zero, a color set to the surface — and
// the accessible label alone would not notice, because the label comes
// from the element rather than from what it painted.
//
// What this proves is that each dial drew something, not that each
// primitive of it drew: the primitive a dial uses is checked by
// TestGaugePagesCoverEveryPrimitive, and the finer question of whether
// one primitive inside a busy dial still paints has no cheap answer,
// since the dials overlap in ink. Both tests together are what stands
// in for "no primitive was dropped from a page".
//
// The row gives each page the dials it must draw, by label; the test
// draws the page, finds each dial's box, and requires ink in it. Only
// dials on the first screen are listed: one below the fold lays out
// with no size, so it has no box to measure.
func TestGaugeSceneDialsDraw(t *testing.T) {
	for _, tc := range []struct {
		page   string
		labels []string
	}{
		{"gauge", []string{"读数", "刻度表", "车速"}},
		{"gauge-monitor", []string{"CPU 占用", "内存", "网络", "核心 0", "核心 7"}},
		{"gauge-health", []string{"今天", "星期1", "星期7", "心率", "睡眠"}},
		{"gauge-time", []string{"本地", "纽约", "雷克雅未克", "东京", "秒表", "闹钟", "计时"}},
		{"gauge-weather", []string{"气温", "风向风速", "罗盘", "气压", "空气质量"}},
		{"gauge-cockpit", []string{"空速", "姿态", "高度", "转弯侧滑", "航向", "升降速度"}},
	} {
		t.Run(tc.page, func(t *testing.T) {
			for _, dark := range []bool{false, true} {
				Reset()
				filterTo(tc.page)
				State.Selected = tc.page
				tt := ui.NewTester(GalleryView, Width, Height)
				tt.SetDark(dark)
				tt.Frame()
				img := tt.Image()
				surface := tokens.For(dark).Surface
				for _, label := range tc.labels {
					r, ok := tt.Find(label)
					if !ok {
						t.Fatalf("dark=%v: no dial labelled %q", dark, label)
					}
					if got := ink(img, r, surface); got == 0 {
						t.Errorf("dark=%v: dial %q drew nothing in %v", dark, label, r)
					}
				}
			}
		})
	}
}

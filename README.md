# uix

uix is a shadcn/ui-style component library for the [MyGo](https://github.com/egoist/mygo) native UI toolkit.

Import a component the Go way — it is a normal package of the module — and it draws with MyGo's native UI. No webview, no HTML.

## Quick start

Import the components you use:

```sh
go get github.com/smallnest/uix/components/base github.com/smallnest/uix/components/button
```

The `base` package holds the design tokens; `go mod tidy` fetches any sibling a component composes. Use the components in a view:

```go
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/button"
)

func view(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	ui.Column(c).Fill().Center().Gap(12).Children(func() {
		button.Button(c, button.Props{
			Label:   "Save",
			Variant: button.Primary,
			OnClick: save,
		})
	})
}

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Demo",
			Width:   420,
			Height:  280,
			Content: ui.View(view),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
```

## Components

| Component | Role |
| --- | --- |
| `base` | Design tokens, light and dark |
| `animate` | Entrance and attention animations for any element: fade, rise, slide, scale, pulse, spin |
| `button` | Button in six variants and three sizes |
| `buttongroup` | Row of buttons joined as one control, of which one is chosen, or several |
| `field` | Form field: label, hint, error |
| `input` | Text input, one line |
| `inputgroup` | Field with an icon or text at either end, as a search or a currency |
| `inputotp` | Row of cells that take a one-time password, one character each |
| `kbd` | A key of the keyboard shown as a chip |
| `label` | Standalone label of a control, in the medium weight |
| `checkbox` | Check box with a label |
| `switch` | Switch with a label |
| `slider` | Slider with a label and a readout |
| `select` | Drop-down with a placeholder |
| `badge` | Pill label in four variants |
| `progress` | Progress bar with a label and a readout |
| `dialog` | Modal dialog with a title and actions |
| `drawer` | Panel that rises over the bottom, with a handle, dimming the rest |
| `sheet` | Panel that slides over from an edge, dimming the rest |
| `command` | Palette that opens over the window to run a command by name |
| `toast` | Transient notification at a window corner |
| `dropdown` | Menu that opens from a button |
| `tabs` | Tab list with an accent bar for the active tab |
| `radio` | Radio group with one Tab stop |
| `textarea` | Text area, many lines, with a height in rows |
| `table` | Data table with sort, choice and a header band |
| `card` | Rounded container with a header, content and footer |
| `avatar` | Circle with a photo or initials |
| `aspectratio` | Box that keeps an aspect ratio, its children filling it |
| `separator` | Horizontal or vertical divider |
| `alert` | Colored notice in four variants |
| `combobox` | Searchable drop-down that filters the options as you type |
| `numberinput` | Number input with step buttons and arrow keys |
| `searchfield` | Search box with a clear button and an Enter submission |
| `rating` | Row of stars that set a score, or show one read-only |
| `datepicker` | Calendar button that picks a date, by click or keyboard |
| `timeinput` | Time field with stepping, typing and arrow keys |
| `colorpicker` | Swatch that opens a picker with channels, hex and swatches |
| `toggle` | Pressed button, like the icon toggles of a toolbar |
| `accordion` | Bordered column of sections that a click opens and closes |
| `collapsible` | Disclosure whose label shows and hides its content |
| `breadcrumbs` | Path of items separated by chevrons, the last where you are |
| `popover` | Panel below a trigger, closed by a click outside or Escape |
| `togglegroup` | Row of toggle buttons, each with a tip, for a toolbar |
| `tooltip` | Label that appears over a control after the pointer rests |
| `menu` | Button that opens a menu of actions, with separators and disabled items |
| `menubar` | Row of flat buttons, each opening a menu, as an app's File and Edit |
| `navigationmenu` | Bar of links where a link with children opens a panel of links |
| `contextmenu` | Menu the system shows for a right-click, from the items of the Props |
| `sidebar` | Navigation column of sections and items, one chosen |
| `tree` | Hierarchical list whose branches open and close, one node chosen |
| `split` | Resizable panes with a divider the user drags |
| `list` | Scrolling list of rows that keeps its place, one row chosen |
| `toolbar` | Row of actions along the top of a window, one stop of Tab |
| `meter` | Progress meter that colors by how near the top it is |
| `spinner` | Ring of spokes that turns while work goes on |
| `skeleton` | Placeholder of content still loading, a bar that pulses softly |
| `empty` | Empty state: an icon over a title and a description, with an action |
| `rangeslider` | Track with two knobs that set the low and the high of a range |
| `stepper` | Number with arrows that step it, which repeat as they are held |
| `richtext` | Text of styled runs, each with its own weight, color or underline |
| `editabletext` | Text that renames in place on a double click |
| `findbar` | Bar to find in a document, with the count and the step buttons |
| `segmented` | Row of segments of which one is chosen, a switch between views |
| `grid` | Grid layout of columns the children fill in order, or place themselves in |
| `gridview` | Scrolling grid of items, of which a click or the arrows choose one |
| `scroll` | Container that scrolls its content vertically |
| `fieldset` | Group of fields under a bold legend |
| `autocomplete` | Field that completes the text typed, of the suggestions below it |
| `tokenfield` | Field of tokens as chips, added on Enter and taken out by a button |
| `calendar` | Month's calendar of which a click or the keys choose a day |
| `checkboxgroup` | Group of check boxes under a check of all of them, mixed while some are |
| `form` | Form whose field labels line up to the widest, as macOS draws them |
| `questionnaire` | Form of questions, one at a time, with a progress bar and Back and Next |
| `scrollhorizontal` | Row that scrolls its content sideways |
| `scrollboth` | Container that scrolls its content up and sideways |
| `webview` | Web page inside the native UI, from a web view of the window |
| `link` | Text that opens a URL, or a path of the Router it lives in |
| `icon` | SVG shown as an icon, in the color of the text around it |
| `image` | Bitmap or SVG shown as a picture, at its size or fitted into a box |
| `colorwell` | Swatch that opens a color picker below it, as AppKit's color well |
| `thinking` | "Thinking" indicator of an agent: a pulsing dot and steps that light up |
| `log` | Scrolling list of an agent's steps that reveal one by one as they run |
| `chat` | Message list with bubbles for the user and the assistant |
| `marker` | Short inline label of a turn of a conversation, in three looks |
| `attachment` | Card of a file of a conversation, with a preview and the upload's progress |
| `composer` | Message box with a prompt and the send button for a chat |
| `notificationcenter` | Panel of notifications in tabs, opened from a bell button |
| `appshell` | App frame with a title, a sidebar and the page it switches |
| `statcard` | Metric card with the delta against the previous period |
| `linechart` | Line chart that follows the pointer, against a year earlier |
| `barchart` | Paired bar chart that follows the pointer, against a year earlier |
| `carousel` | Slides you flip with buttons, dots or a trackpad |
| `pagination` | Page numbers with dots and the previous and next buttons |
| `fileupload` | File zone that uploads a picked file with a progress ring |
| `alertdialog` | Modal that interrupts for a confirmation, with a media icon and a compact size |
| `bubble` | Framed surface of a message, in looks from the strong default to an unframed ghost |
| `item` | Row of a list with a media tile, a title, a description and actions |
| `message` | Turn of a conversation, with the avatar, the sender line, the bubble and a status line |
| `messagescroller` | Viewport of a conversation that follows the end, and offers a button back to it |
| `nativeselect` | Drop-down on the platform's own menu, with grouped options |

The last rows are ports of [BoardUI](https://github.com/BoardUI/boardui)
components to MyGo's native UI: the same roles and the same look, drawn
without HTML. Where BoardUI leans on the browser (an input element), the
port uses the closest native MyGo equivalent, such as the open dialog for
a file pick, and the drop zones take files from other apps directly.

The rows below are ports of the
[shadcn/ui expansions](https://shadcnui-expansions.typeart.cc) to MyGo's
native UI. Where an expansion leans on the browser, the port uses the
closest native MyGo equivalent: the responsive modal becomes a bottom
sheet on a narrow window, the heading anchor copies `uix#section` to the
clipboard instead of setting a URL, and the magic back button returns to
the app's home when it has no history.

| Component | Role |
| --- | --- |
| `autosizetextarea` | Text area that grows with its text, from a floor to a ceiling of rows |
| `blockquote` | Quote on a muted face with an accent bar, and the author at the far end |
| `datetimepicker` | Date field and time field editing one moment together |
| `floatinglabelinput` | Field whose label sits inside it, and floats to the top border on focus or text |
| `headingwithanchor` | Heading with a link button that copies the section's anchor to the clipboard |
| `infinitescroll` | Scrolling list that loads the next page as its end comes near |
| `inlinecode` | Code chip in the prose, with a muted face and the monospace font |
| `interactivestepper` | Steps of a process, each reachable from its indicator, with Back and Next |
| `loadingbutton` | Button that turns into a spinner while its work runs, and stops clicks |
| `magicback` | Round back button that goes back, or home when there is no history |
| `multiselect` | Field choosing several options, whose popup has a check box for each |
| `responsivemodal` | Dialog centered on a wide window, a bottom sheet on a narrow one |
| `typewriter` | Text that types itself out, a character at a time, with a blinking caret |

The rows below are ports of the
[onevra](https://github.com/oneo-me/onevra) components to MyGo's native
UI. They keep the same roles: parsing happens once on the application
side, and rendering draws the retained result, so the UI thread never
runs a parser.

| Component | Role |
| --- | --- |
| `code` | Retained syntax highlighting: a parsed document that renders its light and dark spans |
| `markdown` | Markdown with GFM tables, tasks and images, and highlighted code blocks |
| `lucide` | The Lucide icon set, each icon one lazy parse shared by every draw |
| `settings` | Settings surface of categories, groups and items, with search over their metadata |

## Examples

The `examples` directory shows the components at work. An example imports
the components of this repository directly, the way a library consumer
does, so it always shows the current source. See `examples/README.md`.

## Design tokens

The MyGo theme names its colors semantically. The `tokens` package curates their values into one palette. It maps to the shadcn tokens:

| shadcn token | MyGo field |
| --- | --- |
| background | Background |
| foreground | Text |
| primary | Accent |
| primary-foreground | AccentText |
| secondary | Surface |
| muted | SurfaceHover |
| muted-foreground | TextMuted |
| destructive | Danger |
| border | Border |
| ring | Focus |

Apply the palette at the top of a view:

```go
c.SetTheme(tokens.For(c.Theme().Dark))
```

The palette ships in the `base` package. To change it, start from it and
override the colors you want, on the theme you set:

```go
t := tokens.Light()
t.Accent = ui.Hex("#0f766e")
c.SetTheme(t)
```

`tokens.Light` and `tokens.Dark` start from the MyGo theme, so the
neutrals, fonts and metrics stay system-tuned; only the colors you name
change.

## Writing a component

Add a directory under `components`. Each directory is one package,
importable from the module. A component reads its colors from the theme
of the `Context`.

Write a headless test for the component. MyGo renders frames in memory
without a window:

```go
func TestClick(t *testing.T) {
	clicks := 0
	tt := ui.NewTester(view, 200, 100)
	_ = tt.Click("Go")
	if clicks != 1 {
		t.Fatal("want one click")
	}
}
```

## Related projects

- [onevra](https://github.com/oneo-me/onevra) — the sibling component library these four rows port: the settings composer, the Markdown and code views and the Lucide set come from it. MIT.
- [MujicaUI](https://github.com/ZacharyZhang-NY/MujicaUI) — a court-gothic component library for MyGo native UI. Ivory and charcoal base, wine-red accent, old-gold ornament lines, serif titles. 100 components, light and dark themes, three densities. MyGo draws the UI on the GPU; no webview. MIT.
- [GoRex](https://github.com/egoist/gorex) — a clone of Superlogical's Rex terminal in MyGo native UI. Tabs and split panes, persistent sessions, program activity. No webview, no cgo; one ~15 MB app.
- [Godiff](https://github.com/egoist/godiff) — a reimplementation of codiff in MyGo native UI. Review Git changes and commit them, all local. No webview, no JavaScript.
- [MyGo 复刻 DeepSeek harness 界面](https://x.com/hylarucoder/status/2107492125252464831) — 海拉鲁编程客 (@hylarucoder). Pure Go, no JS/TS. AI copied DeepSeek's harness interface, including the streaming chat page and the trace page.
- [MyGo Native UI docs](https://mygo.egoist.dev/docs/ui) — the toolkit these components draw with. Windows written in Go with package `ui`: no HTML, no JavaScript, no webview; MyGo draws them on the GPU.
- [MyGo Terminal plugin](https://mygo.egoist.dev/docs/plugins/terminal) — a terminal for native UI that runs the shell or any program with Ghostty's emulator (libghostty-vt). All Go, no cgo.

## Rules

- A component imports only `github.com/egoist/mygo`, the sibling components it composes, and the mature parser a port needs (goldmark for Markdown, chroma for highlighting).
- Motion comes from the `animate` component, which honors the desktop's reduce-motion preference.
- A component reads its colors from the theme of the `Context`. It does not hard-code colors.
- Every component ships with a headless test that proves its behavior.

## Known limits

- The custom-face variants draw their hover and press states with `Draw`, which tracks the pointer. Primary and Secondary keep the MyGo look, whose disabled state stays colored.
- A tooltip shows after the pointer rests, so its test waits the same moment of real time before it checks the label.

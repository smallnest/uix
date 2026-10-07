# uix

uix is a shadcn/ui-style component registry for the [MyGo](https://github.com/egoist/mygo) native UI toolkit.

The model follows shadcn/ui. A component is source code that you copy into your project and own. It is not a library you import. The `uix` command installs components from a registry into your Go module.

The components draw with MyGo's native UI. They do not use a webview.

## Quick start

Install the base set, which contains the design tokens:

```sh
uix init
```

Add a component:

```sh
uix add button
```

Use it in a view:

```go
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"demo/internal/ui/tokens"
	"demo/internal/ui/components/button"
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
| `button` | Button in six variants and three sizes |
| `field` | Form field: label, hint, error |
| `input` | Text input, one line |
| `checkbox` | Check box with a label |
| `switch` | Switch with a label |
| `slider` | Slider with a label and a readout |
| `select` | Drop-down with a placeholder |
| `badge` | Pill label in four variants |
| `progress` | Progress bar with a label and a readout |
| `dialog` | Modal dialog with a title and actions |
| `toast` | Transient notification at a window corner |
| `dropdown` | Menu that opens from a button |
| `tabs` | Tab list with an accent bar for the active tab |
| `radio` | Radio group with one Tab stop |
| `textarea` | Text area, many lines, with a height in rows |
| `table` | Data table with sort, choice and a header band |
| `card` | Rounded container with a header, content and footer |
| `avatar` | Circle with a photo or initials |
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
| `sidebar` | Navigation column of sections and items, one chosen |
| `tree` | Hierarchical list whose branches open and close, one node chosen |
| `split` | Resizable panes with a divider the user drags |
| `list` | Scrolling list of rows that keeps its place, one row chosen |
| `toolbar` | Row of actions along the top of a window, one stop of Tab |
| `meter` | Progress meter that colors by how near the top it is |
| `spinner` | Ring of spokes that turns while work goes on |
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
| `scrollhorizontal` | Row that scrolls its content sideways |
| `scrollboth` | Container that scrolls its content up and sideways |
| `link` | Text that opens a URL, or a path of the Router it lives in |
| `icon` | SVG shown as an icon, in the color of the text around it |
| `image` | Bitmap or SVG shown as a picture, at its size or fitted into a box |
| `colorwell` | Swatch that opens a color picker below it, as AppKit's color well |

## Examples

The `examples` directory shows the components at work. An example imports
the components of this repository directly, so it always shows the current
source. See `examples/README.md`.

## Commands

| Command | Action |
| --- | --- |
| `uix init` | Install the base set (design tokens) into the current Go module. |
| `uix add <name>` | Install a component and its dependencies into the current Go module. |
| `uix list` | List the components in the registry. |
| `uix theme list` | List the palettes a module can pick. |
| `uix theme set <name>` | Write a palette into the tokens of the current module. |
| `uix theme set custom <hex>` | Derive a palette from your own accent color. |
| `uix registry build` | Rebuild the embedded registry from the `components` directory. |
| `uix registry serve` | Serve the registry over HTTP, so other machines can `uix add` from it. |

Use `-dir` to point at another Go module, and `-registry <url>` to add from a remote registry.

## HTTP registry

The registry is embedded in the `uix` binary. To serve it over HTTP so
`uix add` works from another machine:

```sh
uix registry serve -addr :7373           # serve the embedded registry
curl http://localhost:7373/registry.json # the registry as one JSON file
```

On the consumer machine, point `uix add` at it:

```sh
uix add button -registry http://localhost:7373/registry.json
```

Or write the URL into `uix.config.json`, and every `uix add` in that module
uses it:

```json
{
  "registry": "http://localhost:7373/registry.json"
}
```

To publish a registry you build from the current components (for example, a
fork), build the JSON and host it anywhere:

```sh
go run ./cmd/uix registry build -out dist/registry.json
```

## How it works

The registry is one JSON file. Each item carries its files with their content. The `uix add` command reads the registry, writes the files into `internal/ui`, and fetches the Go modules they need.

| shadcn/ui | uix |
| --- | --- |
| `npx shadcn add button` | `uix add button` |
| CSS variables (design tokens) | `internal/ui/tokens` themes |
| Radix primitives | MyGo `ui` primitives |
| Tailwind classes | Theme colors in the `Context` |
| `components/ui/*.tsx` | `internal/ui/components/*` |
| `cn()` util | Variant system in each component |

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

### Choosing a palette

Pick another palette from the built-in set:

```sh
uix theme list
uix theme set emerald
```

Each palette is one accent scale for both appearances, with the danger
color and the text color on the accent. `uix theme set` regenerates
`internal/ui/tokens/tokens.go`, which you own: the palettes are a
starting point, and the file stays yours to edit.

Derive a palette from your own accent color, the way the shadcn/ui
theme builder turns one swatch into a theme:

```sh
uix theme set custom '#f97316'
```

The light appearance keeps your accent and darkens it for hover and
press; the dark appearance lightens it so it reads on the dark
background. The text color on the accent follows the luminance, white
or near-black. The danger color stays the default of the built-in
palettes, and a very light accent keeps its color in the dark
appearance instead of washing out.

## Writing a component

Add a directory under `components`. Each directory is one item. An optional `uix.json` in it says how the item installs:

```json
{
  "type": "components:ui",
  "install": "internal/ui/components/button",
  "dependencies": ["github.com/egoist/mygo"],
  "registryDependencies": ["base"]
}
```

The defaults are the base set at `internal/ui` and components at `internal/ui/components/<name>`.

Then rebuild the registry:

```sh
go run ./cmd/uix registry build
```

Write a headless test for the component. MyGo renders frames in memory without a window:

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

- [MujicaUI](https://github.com/ZacharyZhang-NY/MujicaUI) — a court-gothic component library for MyGo native UI. Ivory and charcoal base, wine-red accent, old-gold ornament lines, serif titles. 100 components, light and dark themes, three densities. MyGo draws the UI on the GPU; no webview. MIT.
- [GoRex](https://github.com/egoist/gorex) — a clone of Superlogical's Rex terminal in MyGo native UI. Tabs and split panes, persistent sessions, program activity. No webview, no cgo; one ~15 MB app.
- [Godiff](https://github.com/egoist/godiff) — a reimplementation of codiff in MyGo native UI. Review Git changes and commit them, all local. No webview, no JavaScript.
- [MyGo 复刻 DeepSeek harness 界面](https://x.com/hylarucoder/status/2107492125252464831) — 海拉鲁编程客 (@hylarucoder). Pure Go, no JS/TS. AI copied DeepSeek's harness interface, including the streaming chat page and the trace page.
- [MyGo Native UI docs](https://mygo.egoist.dev/docs/ui) — the toolkit these components draw with. Windows written in Go with package `ui`: no HTML, no JavaScript, no webview; MyGo draws them on the GPU.
- [MyGo Terminal plugin](https://mygo.egoist.dev/docs/plugins/terminal) — a terminal for native UI that runs the shell or any program with Ghostty's emulator (libghostty-vt). All Go, no cgo.

## Rules

- A component may import only `github.com/egoist/mygo`. The copy model has no path rewriting.
- A component reads its colors from the theme of the `Context`. It does not hard-code colors.
- Every component ships with a headless test that proves its behavior.

## Known limits

- The custom-face variants draw their hover and press states with `Draw`, which tracks the pointer. Primary and Secondary keep the MyGo look, whose disabled state stays colored.
- A tooltip shows after the pointer rests, so its test waits the same moment of real time before it checks the label.

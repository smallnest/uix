# Examples

The examples show the uix components at work. Each example is a folder in
this directory. An example imports the components of this repository
directly, so it always shows the current source. (An app you build
copies the components with `uix add`; an example has no such copy.)

## Forms

The forms example shows the form controls together in a sign-up form:
field, input, select, textarea, slider, switch, checkbox and button. The
dark-mode switch restyles the whole form at once, and the submit button
validates the fields it must before it welcomes the sign-up.

```
go run ./examples/forms                open the window
go test ./examples/forms/...           render and check it headless
go run ./examples/forms/cmd/snapshot   write light.png and dark.png
```

## Feedback

The feedback example shows the feedback controls together: badge,
progress, toast and dialog, plus the button they build on. It shows a
toast, and it opens a delete-account dialog; both run on the tokens
palette in the light and the dark appearance.

```
go run ./examples/feedback                open the window
go test ./examples/feedback/...           render and check it headless
go run ./examples/feedback/cmd/snapshot   write light.png, dark.png and dialog.png
```

## Settings

The settings example shows the navigation controls together in a settings
screen with three panes: tabs switch the pane, a dropdown picks a
language, and a radio group picks a theme preference, which redraws the
screen in the chosen appearance at once. The Appearance pane shows the
switch and the check box, and its dark-mode switch is a shortcut for the
theme preference; the check box shows the status bar at the bottom of the
window. The About pane shows the badges and a button that fires a toast.
Everything runs on the tokens palette in both appearances.

```
go run ./examples/settings                open the window
go test ./examples/settings/...           render and check it headless
go run ./examples/settings/cmd/snapshot   write light.png, dark.png, dropdown.png, appearance.png and theme-dark.png
```

## Files

The files example shows the table component as a small file browser: a
click on a header sorts the rows by that column, another reverses the
order, a click on a row chooses it, and Enter (or a double click) opens
it. A line under the table names the row chosen or opened.

```
go run ./examples/files                open the window
go test ./examples/files/...           render and check it headless
go run ./examples/files/cmd/snapshot   write files.png, files-sorted.png and files-dark.png
```

## Profile

The profile example shows the card, avatar, separator and alert components
together in a profile page: a card with an avatar, a name and a divider,
and the info, warning and danger alerts under it. The follow button toggles
the following state, which swaps its label and shows a success alert.

```
go run ./examples/profile                open the window
go test ./examples/profile/...           render and check it headless
go run ./examples/profile/cmd/snapshot   write profile.png, profile-following.png and profile-dark.png
```

## Shop

The shop example shows the combobox, numberinput, searchfield and rating
components together in a product catalog. A search box, a category box,
a limit and a minimum rating all narrow the catalog at once, and the
products show their own read-only ratings.

```
go run ./examples/shop                open the window
go test ./examples/shop/...           render and check it headless
go run ./examples/shop/cmd/snapshot   write shop.png, shop-filtered.png and shop-dark.png
```

## Events

The events example shows the datepicker, timeinput, colorpicker and toggle
components together in an event panel. The date and time fields set the
event, the color field tags it, and a toggle arms its reminder; a summary
line shows every change at once.

```
go run ./examples/events                open the window
go test ./examples/events/...           render and check it headless
go run ./examples/events/cmd/snapshot   write events.png, events-changed.png and events-dark.png
```

## Help center

The help center example shows the breadcrumbs, accordion, collapsible and
popover components together in a help page. The breadcrumbs say where the
user is, the accordion opens the answers to the questions, a collapsible
hides more options, and a popover jumps to a section.

```
go run ./examples/helpcenter                open the window
go test ./examples/helpcenter/...           render and check it headless
go run ./examples/helpcenter/cmd/snapshot   write helpcenter.png, helpcenter-open.png and helpcenter-dark.png
```

## Editor

The editor example shows the togglegroup, tooltip, menu and sidebar
components together in a note editor. The sidebar chooses the file being
edited, the toggles restyle its text, tips name the controls, and the file
menu opens a file or saves.

```
go run ./examples/editor                open the window
go test ./examples/editor/...           render and check it headless
go run ./examples/editor/cmd/snapshot   write editor.png, editor-style.png and editor-dark.png
```

## Explorer

The explorer example shows the toolbar, tree, split and list components
together in a small file browser. The toolbar creates files and opens the
chosen one, the tree chooses a folder, the list shows its files, and the
divider between them is draggable.

```
go run ./examples/explorer                open the window
go test ./examples/explorer/...           render and check it headless
go run ./examples/explorer/cmd/snapshot   write explorer.png, explorer-changed.png and explorer-dark.png
```

## Monitor

The monitor example shows the meter, spinner, rangeslider and stepper
components together in a system monitor. The meters show the load of the
CPU, the memory and the disk, Refresh cycles to another reading, Scan
runs a fake scan with a spinner while it is on, the range slider sets
the load target, and the stepper sets the port.

```
go run ./examples/monitor                open the window
go test ./examples/monitor/...           render and check it headless
go run ./examples/monitor/cmd/snapshot   write monitor.png, monitor-scanning.png and monitor-dark.png
```

## Docreader

The docreader example shows the richtext, editabletext, findbar and
segmented components together in a document reader. The segmented
control switches between the reader and the outline, the title renames
in place, and the find bar searches the text, whose matches highlight
and which it steps through.

```
go run ./examples/docreader                open the window
go test ./examples/docreader/...           render and check it headless
go run ./examples/docreader/cmd/snapshot   write docreader.png, docreader-find.png and docreader-dark.png
```

## Gallery

The gallery example shows the grid, gridview, scroll and fieldset
components together in a file gallery. Two field sets filter the files
by type and sort them by name or date, and a grid view shows the files
that remain, of which a click chooses one.

```
go run ./examples/gallery                open the window
go test ./examples/gallery/...           render and check it headless
go run ./examples/gallery/cmd/snapshot   write gallery.png, gallery-docs.png and gallery-dark.png
```

## Composer

The composer example shows the autocomplete, tokenfield, calendar and
checkboxgroup components together in a form creating an appointment.
The task field completes what is typed from the tasks known, the tag
field keeps a set of tags, the calendar picks the due day, and the
reminder group chooses how to be reminded.

```
go run ./examples/composer                open the window
go test ./examples/composer/...           render and check it headless
go run ./examples/composer/cmd/snapshot   write composer.png, composer-filled.png and composer-dark.png
```

## Workspace

The workspace example shows the form, scrollhorizontal, scrollboth and
link components together in a project panel. The macOS-style form lines
up its fields, the days scroll sideways, the board scrolls both ways,
and the header links open the guide.

```
go run ./examples/workspace                open the window
go test ./examples/workspace/...           render and check it headless
go run ./examples/workspace/cmd/snapshot   write workspace.png, workspace-filled.png and workspace-dark.png
```

## Studio

The studio example shows the icon, image and colorwell components
together in a brand workbench. The mark of the studio and a like toggle
are icons, the preview shows the photo and the mark in its own colors,
and a color well picks the brand color, which the preview reflects.

```
go run ./examples/studio                open the window
go test ./examples/studio/...           render and check it headless
go run ./examples/studio/cmd/snapshot   write studio.png, studio-liked.png and studio-dark.png
```

## Agent

The agent example shows the appshell, sidebar, thinking, log, chat,
composer and notificationcenter components together in a copilot page.
The composer asks a question, the thinking indicator spins while the
model works, and the reply streams into the chat word by word; the log
reveals its steps one at a time, and the bell opens the notifications
in their tabs. The stream is simulated so the example runs anywhere.

```
go run ./examples/agent                open the window
go test ./examples/agent/...           render and check it headless
go run ./examples/agent/cmd/snapshot   write agent.png, agent-reply.png, agent-log.png, agent-notifications.png and agent-dark.png
```

## Dashboard

The dashboard example shows the appshell, statcard, linechart and
barchart components together in a store overview: the metric cards up
top, the two charts side by side, and the display cards under them. The
charts follow the pointer, swapping their headline for the month under
it and what it was a year earlier.

```
go run ./examples/dashboard                open the window
go test ./examples/dashboard/...           render and check it headless
go run ./examples/dashboard/cmd/snapshot   write dashboard.png, dashboard-hover.png and dashboard-dark.png
```

## Navigation

The navigation example shows the carousel, pagination and fileupload
components together in one page. The carousel flips through slides with
its buttons, dots or a trackpad; the pagination steps through pages; and
the file zone uploads a picked file with its progress ring. The upload
is simulated, so the zone needs no storage; the Upload sample button
starts one without the dialog.

```
go run ./examples/navigation                open the window
go test ./examples/navigation/...           render and check it headless
go run ./examples/navigation/cmd/snapshot   write navigation.png and navigation-upload.png
```

module github.com/smallnest/uix

go 1.27.1

require (
	git.zacharyzhang.com/ZacharyZhang-NY/MujicaUI v0.0.0-20261011191201-bc716d27becc
	github.com/alecthomas/chroma/v2 v2.27.0
	github.com/egoist/mygo v0.5.5
	github.com/oneo-me/onevra/source v0.2026.11
	github.com/yuin/goldmark v1.8.6
)

require (
	github.com/dlclark/regexp2/v2 v2.8.4 // indirect
	github.com/ebitengine/purego v0.11.1 // indirect
	github.com/go-text/typesetting v0.3.5 // indirect
	golang.org/x/image v0.47.0 // indirect
	golang.org/x/text v0.43.0 // indirect
)

// MujicaUI (git.zacharyzhang.com) is forked into third_party/mujicaui
// because its internal package uses mygo v0.2.10 API that mygo v0.5.5
// (required by onevra) no longer provides; the fork patches those sites.
replace git.zacharyzhang.com/ZacharyZhang-NY/MujicaUI => ./third_party/mujicaui

package view

import (
	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/editor"

	ochat "github.com/oneo-me/onevra/source/components/chat"

	"github.com/smallnest/uix/components/alert"
	"github.com/smallnest/uix/components/alertdialog"
	"github.com/smallnest/uix/components/attachment"
	"github.com/smallnest/uix/components/bubble"
	"github.com/smallnest/uix/components/chat"
	"github.com/smallnest/uix/components/code"
	"github.com/smallnest/uix/components/command"
	"github.com/smallnest/uix/components/composer"
	"github.com/smallnest/uix/components/dialog"
	"github.com/smallnest/uix/components/drawer"
	"github.com/smallnest/uix/components/loadingbutton"
	"github.com/smallnest/uix/components/log"
	"github.com/smallnest/uix/components/markdown"
	"github.com/smallnest/uix/components/message"
	"github.com/smallnest/uix/components/messagescroller"
	"github.com/smallnest/uix/components/notificationcenter"
	"github.com/smallnest/uix/components/popover"
	"github.com/smallnest/uix/components/progress"
	"github.com/smallnest/uix/components/questionnaire"
	"github.com/smallnest/uix/components/responsivemodal"
	sel "github.com/smallnest/uix/components/select"
	"github.com/smallnest/uix/components/settings"
	"github.com/smallnest/uix/components/sheet"
	"github.com/smallnest/uix/components/spinner"
	sw "github.com/smallnest/uix/components/switch"
	"github.com/smallnest/uix/components/thinking"
	"github.com/smallnest/uix/components/toast"
	"github.com/smallnest/uix/components/tooltip"

	"github.com/smallnest/uix/components/button"
)

// demoDialog shows a modal dialog with a title and actions; the button
// opens it.
func demoDialog(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "打开对话框", OnClick: func() { Demo.Dialog = true }})
		ui.Text(c, "对话框在窗口上方打开。").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
	dialog.Dialog(c, dialog.Props{
		Open:        &Demo.Dialog,
		Title:       "要删除你的账户吗？",
		Description: "这将删除你的数据，且无法撤销。",
		Confirm:     "删除",
		Cancel:      "保留账户",
		OnConfirm:   func() { Demo.Dialog = false },
		OnCancel:    func() { Demo.Dialog = false },
	})
}

// demoAlertDialog shows a modal that interrupts for a confirmation.
func demoAlertDialog(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "打开警告对话框", Variant: button.Secondary, OnClick: func() { Demo.Alert = true }})
	})
	alertdialog.AlertDialog(c, alertdialog.Props{
		Open:        &Demo.Alert,
		Title:       "你确定吗？",
		Description: "此操作无法撤销。",
		Confirm:     "继续",
		Cancel:      "取消",
		OnConfirm:   func() { Demo.Alert = false },
		OnCancel:    func() { Demo.Alert = false },
	})
}

// demoDrawer shows a panel that rises over the bottom.
func demoDrawer(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "打开抽屉", OnClick: func() { Demo.Drawer = true }})
	})
	drawer.Drawer(c, drawer.Props{
		Open:  &Demo.Drawer,
		Title: "抽屉",
		Children: func() {
			ui.Text(c, "从底部升起的面板，其余区域变暗。")
		},
	})
}

// demoSheet shows a panel that slides over from an edge.
func demoSheet(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "打开侧边面板", OnClick: func() { Demo.Sheet = true }})
	})
	sheet.Sheet(c, sheet.Props{
		Open:  &Demo.Sheet,
		Title: "侧边面板",
		Children: func() {
			ui.Text(c, "从边缘滑入的面板。")
		},
	})
}

// demoPopover shows a panel below a trigger.
func demoPopover(c *ui.Context, t *ui.Theme) {
	popover.Popover(c, popover.Props{
		Open: &Demo.Pop,
		Trigger: func(c *ui.Context) ui.Element {
			return button.Button(c, button.Props{Label: "打开气泡", Variant: button.Outline})
		},
		Content: func(c *ui.Context) {
			ui.Column(c).Gap(t.Space(2)).Children(func() {
				ui.Text(c, "触发器下方的面板。").Bold()
				ui.Text(c, "点击外部或按 Esc 关闭。").TextColor(t.TextMuted)
			})
		},
	})
}

// demoTooltip shows a label that appears over a control.
func demoTooltip(c *ui.Context, t *ui.Theme) {
	tooltip.Tooltip(c, tooltip.Props{
		Anchor: button.Button(c, button.Props{Label: "悬停我"}),
		Text:   "按钮上方出现提示。",
	})
}

// demoToast shows a transient notification at a window corner.
func demoToast(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "显示提示", Variant: button.Secondary, OnClick: func() {
			toast.Push(c, toast.Props{Title: "设置已保存", Description: "立即生效。", Type: toast.Success})
		}})
		ui.Text(c, "提示出现在角落并淡出。").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
	toast.Viewport(c)
}

// demoSpinner shows a ring of spokes that turns while work goes on.
func demoSpinner(c *ui.Context, t *ui.Theme) {
	spinner.Spinner(c, spinner.Props{})
}

// demoProgress shows a progress bar with a label and a readout.
func demoProgress(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(300).Gap(t.Space(2)).Children(func() {
		progress.Progress(c, progress.Props{Value: 0.64, Label: "磁盘占用"})
		progress.Progress(c, progress.Props{Value: 0, Label: "索引中", Indeterminate: true})
	})
}

// demoCommand shows a palette that opens over the window.
func demoCommand(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "打开命令面板", OnClick: func() { Demo.Command = true }})
	})
	command.Command(c, command.Props{
		Open:  &Demo.Command,
		State: &commandState,
		Items: []command.Item{
			{Label: "新建文件", Keywords: []string{"create", "document"}},
			{Label: "打开设置", Keywords: []string{"preferences"}},
			{Label: "运行构建", Keywords: []string{"compile", "make"}},
		},
	})
}

// demoAlert shows colored notices in four variants.
func demoAlert(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(360).Gap(t.Space(2)).Children(func() {
		alert.Alert(c, alert.Props{Title: "信息", Description: "强调色的通知。", Variant: alert.Info})
		alert.Alert(c, alert.Props{Title: "成功", Description: "一切正常。", Variant: alert.Success})
		alert.Alert(c, alert.Props{Title: "警告", Description: "注意上限。", Variant: alert.Warning})
		alert.Alert(c, alert.Props{Title: "危险", Description: "出了点问题。", Variant: alert.Danger})
	})
}

// demoLoadingButton shows a button that turns into a spinner.
func demoLoadingButton(c *ui.Context, t *ui.Theme) {
	loadingbutton.LoadingButton(c, loadingbutton.Props{Label: "保存中", Loading: true})
}

// demoNotificationCenter shows a panel of notifications in tabs.
func demoNotificationCenter(c *ui.Context, t *ui.Theme) {
	notificationcenter.NotificationCenter(c, notificationcenter.Props{
		Tab: &Demo.NCenter,
		Items: []notificationcenter.Item{
			{ID: "1", Category: "system", Title: "备份完成", Description: "你的数据已安全。", Timestamp: "2分钟前", Unread: true},
			{ID: "2", Category: "mentions", Title: "Ada 提到了你", Description: "在设计评审中。", Timestamp: "1小时前", Unread: true},
			{ID: "3", Category: "system", Title: "有新版本", Description: "v0.2.2 已就绪。", Timestamp: "1天前"},
		},
	})
}

// demoCode shows retained syntax highlighting.
func demoCode(c *ui.Context, t *ui.Theme) {
	code.Parse("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"hello\")\n}\n", "go").Render(c)
}

// demoMarkdown shows markdown with tables and highlighted code.
func demoMarkdown(c *ui.Context, t *ui.Theme) {
	markdown.Markdown(c, "# 组件画廊\n\n一个带 **markdown** 和 `代码` 的视图。\n\n| 列 | 列 |\n| --- | --- |\n| 一 | 二 |\n")
}

// demoCodeEditor shows mygo's code editor, read-only here.
func demoCodeEditor(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Height(220).Children(func() {
		editor.View(c, &editor.State{
			Buffer:   editor.NewBuffer("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"hello\")\n}\n"),
			ReadOnly: true,
		})
	})
}

// demoChat shows a transcript and a composer.
func demoChat(c *ui.Context, t *ui.Theme) {
	chat.Chat(c, chat.Props{
		Messages: []chat.Message{
			{Role: chat.User, Text: "uix 是什么？"},
			{Role: chat.Assistant, Text: "一个用纯 Go 写的 MyGo 原生 UI 组件库。"},
		},
		Value:       &Demo.Chat,
		Model:       "claude-fable-5",
		Provider:    "anthropic",
		OnSubmit:    func(string) {},
		OnStop:      func() {},
		Suggestions: []string{"展示一个组件", "列出所有库"},
	})
}

// demoComposer shows a message box with a prompt and a send button.
func demoComposer(c *ui.Context, t *ui.Theme) {
	composer.Composer(c, composer.Props{
		Value:       &Demo.Composer,
		Placeholder: "随便问点什么",
		OnSubmit:    func(string) {},
		OnStop:      func() {},
		Model:       "claude-fable-5",
		Provider:    "anthropic",
	})
}

// demoMessage shows a turn of a conversation.
func demoMessage(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Gap(t.Space(2)).Children(func() {
		message.Message(c, message.Props{
			Name:   "Ada",
			Header: "Ada · 2分钟前",
			Text:   "对话中的一条消息气泡。",
		})
	})
}

// demoAttachment shows a card of a file of a conversation.
func demoAttachment(c *ui.Context, t *ui.Theme) {
	attachment.Attachment(c, attachment.Props{
		Title:       "design.pdf",
		Description: "2.4 MB · 已上传",
		Progress:    -1,
	})
}

// demoThinking shows the thinking indicator of an agent.
func demoThinking(c *ui.Context, t *ui.Theme) {
	thinking.Thinking(c, thinking.Props{Label: "思考中"})
}

// demoLog shows a scrolling list of steps that reveal one by one.
func demoLog(c *ui.Context, t *ui.Theme) {
	log.Log(c, log.Props{
		Items:    []string{"阅读文档", "寻找模式", "编写代码"},
		Revealed: 3,
		Running:  false,
	})
}

// demoQuestionnaire shows a form of questions, one at a time.
func demoQuestionnaire(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(420).Children(func() {
		questionnaire.Questionnaire(c, questionnaire.Props{
			State: &Demo.Question,
			Steps: []questionnaire.Step{
				{Title: "你的团队有多大？", Description: "团队的人数。", Choices: []questionnaire.Choice{
					{Label: "1–10", Value: "small"},
					{Label: "11–50", Value: "mid"},
					{Label: "50+", Value: "large"},
				}},
				{Title: "你的角色是什么？", Description: "你所做的工作。", Choices: []questionnaire.Choice{
					{Label: "工程师", Value: "engineer"},
					{Label: "设计师", Value: "designer"},
				}},
			},
			OnDone: func() {},
		})
	})
}

// demoSettings shows the settings composer of a category.
func demoSettings(c *ui.Context, t *ui.Theme) {
	settings.Settings(c, &Demo.Settings, []settings.SettingsCategory{
		{
			ID: "general", Title: "通用", Description: "应用的偏好设置。",
			Groups: []settings.SettingsGroup{
				{
					ID: "appearance", Title: "外观",
					Items: []settings.SettingsItem{
						{ID: "dark", Title: "深色模式", Description: "使用深色外观。", Build: func(c *ui.Context) {
							sw.Switch(c, sw.Props{On: &Demo.Switch, Label: ""})
						}},
						{ID: "language", Title: "语言", Description: "界面的语言。", Build: func(c *ui.Context) {
							sel.Select(c, sel.Props{Selected: &Demo.SelStr, Options: []string{"English", "简体中文", "日本語"}})
						}},
					},
				},
			},
		},
	}, settings.SettingsOptions{})
}

// demoBubble shows the framed bubble of a conversation, on either side.
func demoBubble(c *ui.Context, t *ui.Theme) {
	ui.Row(c).Gap(t.Space(2)).Children(func() {
		bubble.Bubble(c, bubble.Props{Text: "早上好。", Align: bubble.End})
		bubble.Bubble(c, bubble.Props{Text: "收到,还有别的吗?", Variant: bubble.Secondary})
	})
}

// demoMessageScroller shows a conversation that stays at the newest
// message, with a button to jump back down when the reader scrolled up.
func demoMessageScroller(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(440).Height(240).Children(func() {
		messagescroller.MessageScroller(c, messagescroller.Props{
			State:    &Demo.Scroller,
			NewLabel: "新消息",
			Children: func() {
				ui.Column(c).FillWidth().Gap(t.Space(2)).Padding(t.Space(3)).Children(func() {
					for _, m := range []message.Props{
						{Name: "你", Header: "现在", Text: "早上好。", Align: bubble.End},
						{Name: "Ada", Header: "现在", Text: "收到,这就回复。", Variant: bubble.Secondary},
						{Name: "你", Header: "现在", Text: "谢谢。", Align: bubble.End},
					} {
						message.Message(c, m)
					}
				})
			},
		})
	})
}

// demoResponsiveModal shows a modal that becomes a bottom sheet in a
// narrow window; the button opens it.
func demoResponsiveModal(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(2)).Children(func() {
		button.Button(c, button.Props{Label: "打开响应式对话框", OnClick: func() { Demo.RModal = true }})
		ui.Text(c, "窗口窄时它变成底部面板。").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
	responsivemodal.ResponsiveModal(c, responsivemodal.Props{
		Open:  &Demo.RModal,
		Title: "响应式对话框",
		Body: func(c *ui.Context) {
			ui.Text(c, "窗口窄时它变成底部面板。")
		},
		Actions: func(c *ui.Context) {
			ui.Row(c).FillWidth().Justify(ui.End).Children(func() {
				button.Button(c, button.Props{Label: "完成", OnClick: func() { Demo.RModal = false }})
			})
		},
	})
}

// demoDiffView shows mygo's editor reviewing a change: the old and new
// texts side by side, the changed lines marked.
func demoDiffView(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(560).Height(240).Children(func() {
		editor.DiffView(c, &Demo.Diff)
	})
}

// demoResultSummary shows onevra's summary of a finished turn: a title,
// a description and the changed items, the extra ones behind a toggle.
func demoResultSummary(c *ui.Context, t *ui.Theme) {
	ochat.ResultSummary(c, &Demo.OSummary, ochat.SummaryOptions{
		Title:       "已更改 3 个文件",
		Description: "重构了消息的发送路径。",
		ActionLabel: "查看变更",
		MoreLabel:   "更多文件",
		LessLabel:   "更少文件",
		Items: []ochat.SummaryItem{
			{ID: "main", Label: "main.go", Detail: "重命名函数"},
			{ID: "view", Label: "view.go", Detail: "新增状态"},
			{ID: "readme", Label: "README.md", Detail: "更新用法"},
		},
	})
}

// demoConversation shows onevra's transcript of a finished turn: the
// prompt, the steps that ran, and the answer, the steps folding behind a
// summary once done. The turn is expanded in the demo's state, so the
// transcript shows it whole.
func demoConversation(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(560).Height(320).Children(func() {
		ochat.Conversation(c, &Demo.OConv, "回到底部", func(c *ui.Context) {
			ochat.Turn(c, &Demo.OTurn, ochat.TurnOptions{
				ID:      "sample",
				Status:  ochat.Completed,
				Summary: "已运行 2 个步骤",
				Prompt: func(c *ui.Context) {
					ui.Text(c, "帮我审查这段代码。")
				},
				Steps: []ochat.Step{
					{ID: "reason", Message: func(c *ui.Context) {
						ui.Text(c, "先读源码,再对照示例。").TextColor(t.TextMuted)
					}},
					{ID: "tools", Title: "并行检查", Group: &Demo.OGroup, Activities: []ochat.Activity{
						{ID: "read", Action: "读取", Detail: "main.go", Status: ochat.Completed},
						{ID: "scan", Action: "检查", Detail: "示例目录", Status: ochat.Completed},
					}},
				},
				Final: func(c *ui.Context) {
					ui.Text(c, "审查完成,没有发现明显问题。")
				},
			})
		})
	})
}

// demoGlassComposer shows onevra's message box on a glass panel over the
// backdrop; the draft and the layout the composer keeps come from the
// demo's state.
func demoGlassComposer(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Width(440).Padding(t.Space(2)).Children(func() {
		ochat.GlassComposer(c, &Demo.OLayout, &Demo.OComposer, ochat.ComposerOptions{
			Label:       "消息",
			Placeholder: "随便问点什么",
			SendLabel:   "发送",
			StopLabel:   "停止",
		})
	})
}

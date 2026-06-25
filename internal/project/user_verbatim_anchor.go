package project

import (
	"fmt"
	"strings"

	"cyberstrike-ai/internal/database"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

const (
	// UserVerbatimSectionHeading 用户原文锚点可读标题（块内保留，供 Agent 阅读）。
	UserVerbatimSectionHeading = "## 用户历史输入（原文保留，勿省略或改写）"

	// UserVerbatimSectionStartMarker / EndMarker：HTML 注释边界，供程序化替换；对模型无指令语义。
	UserVerbatimSectionStartMarker = "<!-- user-verbatim-start -->"
	UserVerbatimSectionEndMarker   = "<!-- user-verbatim-end -->"
)

// ExtractUserContentsFromMessages 按时间顺序提取 user 角色消息的原文（跳过空白）。
func ExtractUserContentsFromMessages(msgs []database.Message) []string {
	out := make([]string, 0, len(msgs))
	for i := range msgs {
		if !strings.EqualFold(strings.TrimSpace(msgs[i].Role), "user") {
			continue
		}
		content := strings.TrimSpace(msgs[i].Content)
		if content == "" {
			continue
		}
		out = append(out, content)
	}
	return out
}

// BuildUserVerbatimAnchorBlockFromMessages 从 messages 表行构建用户原文锚点块。
// maxRunes: 0 = 不截断；>0 = 总 rune 上限（仍保留每一轮，仅对超长单条做尾部截断提示）。
func BuildUserVerbatimAnchorBlockFromMessages(msgs []database.Message, maxRunes int) string {
	return BuildUserVerbatimAnchorBlock(ExtractUserContentsFromMessages(msgs), maxRunes)
}

// BuildUserVerbatimAnchorBlock 将各轮用户原文格式化为 system prompt 锚点块。
func BuildUserVerbatimAnchorBlock(userContents []string, maxRunes int) string {
	if len(userContents) == 0 {
		return ""
	}
	lines := make([]string, 0, len(userContents))
	for _, content := range userContents {
		content = strings.TrimSpace(content)
		if content == "" {
			continue
		}
		lines = append(lines, fmt.Sprintf("[第%d轮] %s", len(lines)+1, content))
	}
	if len(lines) == 0 {
		return ""
	}
	body := strings.Join(lines, "\n")
	if maxRunes > 0 {
		body = capUserVerbatimBody(body, maxRunes)
	}
	return wrapUserVerbatimBlock(UserVerbatimSectionHeading + "\n\n" + body)
}

func capUserVerbatimBody(body string, maxRunes int) string {
	rs := []rune(body)
	if len(rs) <= maxRunes {
		return body
	}
	suffix := "\n\n...(用户原文锚点已达配置上限，更早轮次可能被截断；完整原文见 messages 表)..."
	suffixRunes := []rune(suffix)
	keep := maxRunes - len(suffixRunes)
	if keep <= 0 {
		return string(rs[:maxRunes])
	}
	return string(rs[:keep]) + suffix
}

func wrapUserVerbatimBlock(content string) string {
	content = strings.TrimSpace(content)
	if content == "" {
		return ""
	}
	return UserVerbatimSectionStartMarker + "\n" + content + "\n" + UserVerbatimSectionEndMarker + "\n"
}

// ReplaceUserVerbatimAnchorSection 用 freshBlock 替换 content 中已有的用户原文锚点段。
func ReplaceUserVerbatimAnchorSection(content, freshBlock string) (string, bool) {
	content = strings.TrimSpace(content)
	freshBlock = strings.TrimSpace(freshBlock)
	if freshBlock == "" {
		return content, false
	}
	start, ok := userVerbatimSectionStart(content)
	if !ok {
		return content, false
	}
	end, ok := userVerbatimSectionEnd(content, start)
	if !ok {
		return content, false
	}
	return strings.TrimSpace(content[:start] + freshBlock + content[end:]), true
}

func userVerbatimSectionStart(content string) (int, bool) {
	idx := strings.Index(content, UserVerbatimSectionStartMarker)
	if idx < 0 {
		return 0, false
	}
	return idx, true
}

func userVerbatimSectionEnd(content string, start int) (int, bool) {
	if start < 0 || start >= len(content) {
		return 0, false
	}
	tail := content[start:]
	idx := strings.LastIndex(tail, UserVerbatimSectionEndMarker)
	if idx < 0 {
		return 0, false
	}
	return start + idx + len(UserVerbatimSectionEndMarker), true
}

// RefreshUserVerbatimAnchorInMessages 在 summarization 等压缩后，用 freshBlock 刷新 system 中的用户原文锚点。
// 若尚无锚点段，则追加到首条 system 消息；若无 system 消息则在开头插入一条。
func RefreshUserVerbatimAnchorInMessages(msgs []adk.Message, freshBlock string) []adk.Message {
	freshBlock = strings.TrimSpace(freshBlock)
	if freshBlock == "" || len(msgs) == 0 {
		return msgs
	}

	out := make([]adk.Message, len(msgs))
	changed := false
	for i, msg := range msgs {
		if msg == nil || msg.Role != schema.System {
			out[i] = msg
			continue
		}
		newContent, ok := ReplaceUserVerbatimAnchorSection(msg.Content, freshBlock)
		if !ok {
			out[i] = msg
			continue
		}
		cloned := *msg
		cloned.Content = newContent
		out[i] = &cloned
		changed = true
	}

	if changed {
		return out
	}

	for i, msg := range msgs {
		if msg == nil || msg.Role != schema.System {
			continue
		}
		cloned := *msg
		cloned.Content = AppendSystemPromptBlock(cloned.Content, freshBlock)
		out[i] = &cloned
		return out
	}

	prefix := make([]adk.Message, 0, len(msgs)+1)
	prefix = append(prefix, schema.SystemMessage(freshBlock))
	return append(prefix, msgs...)
}

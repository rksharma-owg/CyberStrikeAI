package project

import (
	"strings"
	"testing"

	"cyberstrike-ai/internal/database"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
)

func TestBuildUserVerbatimAnchorBlock_MultiTurn(t *testing.T) {
	msgs := []database.Message{
		{Role: "user", Content: "目标 https://a.com 仅测 /api"},
		{Role: "assistant", Content: "好的"},
		{Role: "user", Content: "用 admin:test 登录"},
	}
	block := BuildUserVerbatimAnchorBlockFromMessages(msgs, 0)
	if block == "" {
		t.Fatal("expected non-empty block")
	}
	if !strings.Contains(block, UserVerbatimSectionStartMarker) {
		t.Error("missing start marker")
	}
	if !strings.Contains(block, "[第1轮]") || !strings.Contains(block, "https://a.com") {
		t.Error("missing first user turn")
	}
	if !strings.Contains(block, "[第2轮]") || !strings.Contains(block, "admin:test") {
		t.Error("missing second user turn")
	}
	if strings.Contains(block, "好的") {
		t.Error("assistant content should not appear")
	}
}

func TestReplaceUserVerbatimAnchorSection(t *testing.T) {
	old := "prefix\n\n" + wrapUserVerbatimBlock("## old\n\n[第1轮] a") + "\nsuffix"
	newBlock := wrapUserVerbatimBlock(UserVerbatimSectionHeading + "\n\n[第1轮] b\n[第2轮] c")
	out, ok := ReplaceUserVerbatimAnchorSection(old, newBlock)
	if !ok {
		t.Fatal("expected replace ok")
	}
	if !strings.Contains(out, "[第2轮] c") {
		t.Errorf("expected new block, got %q", out)
	}
	if !strings.HasPrefix(strings.TrimSpace(out), "prefix") {
		t.Error("prefix should remain")
	}
	if !strings.Contains(out, "suffix") {
		t.Error("suffix should remain")
	}
}

func TestRefreshUserVerbatimAnchorInMessages_ReplaceExisting(t *testing.T) {
	oldBlock := wrapUserVerbatimBlock(UserVerbatimSectionHeading + "\n\n[第1轮] old")
	msgs := []adk.Message{
		schema.SystemMessage("instr\n\n" + oldBlock),
		schema.UserMessage("hi"),
	}
	newBlock := wrapUserVerbatimBlock(UserVerbatimSectionHeading + "\n\n[第1轮] new")
	out := RefreshUserVerbatimAnchorInMessages(msgs, newBlock)
	if len(out) != 2 {
		t.Fatalf("message count: got %d", len(out))
	}
	if !strings.Contains(out[0].Content, "[第1轮] new") {
		t.Errorf("system content: %q", out[0].Content)
	}
	if strings.Contains(out[0].Content, "[第1轮] old") {
		t.Error("old anchor should be replaced")
	}
}

func TestRefreshUserVerbatimAnchorInMessages_InsertWhenMissing(t *testing.T) {
	msgs := []adk.Message{
		schema.SystemMessage("base instruction"),
		schema.UserMessage("hi"),
	}
	block := wrapUserVerbatimBlock(UserVerbatimSectionHeading + "\n\n[第1轮] anchor")
	out := RefreshUserVerbatimAnchorInMessages(msgs, block)
	if !strings.Contains(out[0].Content, "[第1轮] anchor") {
		t.Errorf("expected appended anchor, got %q", out[0].Content)
	}
}

func TestBuildUserVerbatimAnchorBlock_MaxRunes(t *testing.T) {
	long := strings.Repeat("字", 200)
	block := BuildUserVerbatimAnchorBlock([]string{long}, 50)
	body := block
	if idx := strings.Index(body, UserVerbatimSectionStartMarker); idx >= 0 {
		body = strings.TrimPrefix(body[idx+len(UserVerbatimSectionStartMarker):], "\n")
	}
	if len([]rune(body)) > 120 {
		t.Errorf("expected capped body, got %d runes", len([]rune(body)))
	}
}

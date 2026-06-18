package runtime

import "testing"

func TestBufferAddCommand(t *testing.T) {
	b := NewBuffer()

	b.Add("ls")

	commands := b.List()

	if len(commands) != 1 {
		t.Fatalf("expected 1 command, got %d", len(commands))
	}

	if commands[0].Text != "ls" {
		t.Fatalf("expected ls, got %s", commands[0].Text)
	}

	if commands[0].Status != Pending {
		t.Fatalf("expected Pending status")
	}
}

func TestBufferListCommands(t *testing.T) {
	b := NewBuffer()

	b.Add("ls")
	b.Add("pwd")

	commands := b.List()

	if len(commands) != 2 {
		t.Fatalf("expected 2 commands, got %d", len(commands))
	}
}

func TestRejectEmptyCommand(t *testing.T) {
	b := NewBuffer()

	b.Add("")

	if len(b.List()) != 0 {
		t.Fatalf("empty command should be rejected")
	}
}

func TestRejectCommentCommand(t *testing.T) {
	b := NewBuffer()

	b.Add("# comment")

	if len(b.List()) != 0 {
		t.Fatalf("comment command should be rejected")
	}
}

func TestAcceptNormalCommand(t *testing.T) {
	cmd := Command{
		Text: "go test ./...",
	}

	if !Validate(cmd) {
		t.Fatalf("normal command should pass validation")
	}
}

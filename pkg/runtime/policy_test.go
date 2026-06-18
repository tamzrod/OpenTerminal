package runtime

import "testing"

func TestBlockInteractiveCommands(t *testing.T) {
	blocked := []string{
		"nano file.go",
		"vim main.go",
		"vi test.txt",
		"less output.log",
		"top",
	}

	for _, text := range blocked {
		cmd := Command{
			Text: text,
		}

		if Allow(cmd) {
			t.Fatalf("expected blocked command: %s", text)
		}
	}
}

func TestAllowNormalCommand(t *testing.T) {
	cmd := Command{
		Text: "go test ./...",
	}

	if !Allow(cmd) {
		t.Fatalf("expected normal command allowed")
	}
}

package runtime

import "testing"

func TestRunnerExecutesCommand(t *testing.T) {
	runner := NewRunner()

	result := runner.Run("echo hello")

	if result.Status != Success {
		t.Fatalf("expected success")
	}

	if result.Output == "" {
		t.Fatalf("expected output")
	}
}

func TestRunnerCreatesAudit(t *testing.T) {
	runner := NewRunner()

	runner.Run("echo hello")

	history := runner.History()

	if len(history) != 1 {
		t.Fatalf("expected audit record")
	}
}

func TestRunnerRejectsInvalidCommand(t *testing.T) {
	runner := NewRunner()

	result := runner.Run("# comment")

	if result.Status != Skipped {
		t.Fatalf("expected skipped")
	}
}

func TestRunnerPolicyBlock(t *testing.T) {
	runner := NewRunner()

	result := runner.Run("nano test.go")

	if result.Status != Skipped {
		t.Fatalf("expected policy block")
	}
}

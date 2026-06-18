package runtime

import "testing"

func TestRunnerExecutesCommand(t *testing.T) {
	runner := NewRunner()

	result := runner.Run("echo hello")

	if result.Status != Success {
		t.Fatalf("expected success")
	}
}

func TestRunnerCreatesAudit(t *testing.T) {
	runner := NewRunner()

	runner.Run("echo hello")

	history := runner.History()

	if len(history) != 1 {
		t.Fatalf("expected audit history")
	}
}

func TestRunnerRejectsInvalidCommand(t *testing.T) {
	runner := NewRunner()

	result := runner.Run("# bad")

	if result.Status != Skipped {
		t.Fatalf("expected skipped")
	}
}

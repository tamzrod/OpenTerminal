package runtime

import "testing"

func TestExecutorSuccess(t *testing.T) {
	executor := NewExecutor()

	cmd := Command{
		Text: "echo hello",
	}

	result := executor.Run(cmd)

	if result.Status != Success {
		t.Fatalf("expected success")
	}

	if result.Output == "" {
		t.Fatalf("expected output")
	}
}

func TestExecutorFailure(t *testing.T) {
	executor := NewExecutor()

	cmd := Command{
		Text: "command_that_does_not_exist",
	}

	result := executor.Run(cmd)

	if result.Status != Failed {
		t.Fatalf("expected failure")
	}

	if result.Error == "" {
		t.Fatalf("expected error")
	}
}

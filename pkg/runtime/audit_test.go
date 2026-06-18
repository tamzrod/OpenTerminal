package runtime

import "testing"

func TestAuditRecordsCommand(t *testing.T) {
	audit := NewAudit()

	cmd := Command{
		Text: "go test ./...",
		Status: Success,
		Output: "PASS",
	}

	audit.Record(cmd)

	entries := audit.Entries()

	if len(entries) != 1 {
		t.Fatalf("expected audit entry")
	}

	if entries[0].Command.Text != cmd.Text {
		t.Fatalf("wrong command stored")
	}
}

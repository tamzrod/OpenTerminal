package runtime

import "time"

type AuditEntry struct {
	Time    time.Time
	Command Command
}

type Audit struct {
	entries []AuditEntry
}

func NewAudit() *Audit {
	return &Audit{}
}

func (a *Audit) Record(cmd Command) {
	a.entries = append(a.entries, AuditEntry{
		Time:    time.Now(),
		Command: cmd,
	})
}

func (a *Audit) Entries() []AuditEntry {
	result := make([]AuditEntry, len(a.entries))
	copy(result, a.entries)

	return result
}

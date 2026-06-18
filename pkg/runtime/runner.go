package runtime

type Runner struct {
	executor *Executor
	audit    *Audit
}

func NewRunner() *Runner {
	return &Runner{
		executor: NewExecutor(),
		audit:    NewAudit(),
	}
}

func (r *Runner) Run(text string) Command {
	cmd := Command{
		Text: text,
		Status: Pending,
	}

	if !Validate(cmd) {
		cmd.Status = Skipped
		cmd.Error = "invalid command"
		r.audit.Record(cmd)
		return cmd
	}

	result := r.executor.Run(cmd)

	r.audit.Record(result)

	return result
}

func (r *Runner) History() []AuditEntry {
	return r.audit.Entries()
}

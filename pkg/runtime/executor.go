package runtime

import (
	"os/exec"
)

type Executor struct {
	context Context
}

func NewExecutor() *Executor {
	return &Executor{
		context: NewContext(),
	}
}

func (e *Executor) Run(cmd Command) Command {
	if !Allow(cmd) {
		cmd.Status = Skipped
		cmd.Error = "blocked by policy"
		return cmd
	}

	process := exec.Command("sh", "-c", cmd.Text)
	process.Dir = e.context.WorkDir

	out, err := process.CombinedOutput()

	cmd.Output = string(out)

	if err != nil {
		cmd.Status = Failed
		cmd.Error = err.Error()
		return cmd
	}

	cmd.Status = Success
	return cmd
}

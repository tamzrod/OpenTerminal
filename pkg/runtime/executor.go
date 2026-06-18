package runtime

import (
	"os/exec"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) Run(cmd Command) Command {
	out, err := exec.Command("sh", "-c", cmd.Text).CombinedOutput()

	cmd.Output = string(out)

	if err != nil {
		cmd.Status = Failed
		cmd.Error = err.Error()
		return cmd
	}

	cmd.Status = Success
	return cmd
}

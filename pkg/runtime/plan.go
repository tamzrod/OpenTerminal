package runtime

type CommandStatus string

const (
	Pending CommandStatus = "Pending"
	Success CommandStatus = "Success"
	Failed  CommandStatus = "Failed"
	Skipped CommandStatus = "Skipped"
)

type Command struct {
	ID     int
	Text   string
	Status CommandStatus
	Output string
	Error  string
}

type Plan struct {
	Commands []Command
}

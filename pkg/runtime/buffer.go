package runtime

import "sync"

type Buffer struct {
	mu       sync.Mutex
	commands []Command
}

func NewBuffer() *Buffer {
	return &Buffer{}
}

func (b *Buffer) Add(command string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	cmd := Command{
		ID:     len(b.commands) + 1,
		Text:   command,
		Status: Pending,
	}

	if !Validate(cmd) {
		return
	}

	b.commands = append(b.commands, cmd)
}

func (b *Buffer) List() []Command {
	b.mu.Lock()
	defer b.mu.Unlock()

	result := make([]Command, len(b.commands))
	copy(result, b.commands)

	return result
}

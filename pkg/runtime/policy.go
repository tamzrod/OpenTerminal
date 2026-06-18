package runtime

import "strings"

var blockedCommands = []string{
	"nano",
	"vim",
	"vi",
	"less",
	"top",
}

func Allow(cmd Command) bool {
	text := strings.TrimSpace(cmd.Text)

	for _, blocked := range blockedCommands {
		if text == blocked || strings.HasPrefix(text, blocked+" ") {
			return false
		}
	}

	return true
}

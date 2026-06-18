package runtime

import "strings"

func Validate(cmd Command) bool {
	text := strings.TrimSpace(cmd.Text)

	if text == "" {
		return false
	}

	if strings.HasPrefix(text, "#") {
		return false
	}

	return true
}

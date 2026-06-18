package runtime

import "testing"

func TestContextHasWorkDir(t *testing.T) {
	ctx := NewContext()

	if ctx.WorkDir == "" {
		t.Fatalf("expected working directory")
	}
}

func TestContextHasUser(t *testing.T) {
	ctx := NewContext()

	if ctx.User == "" {
		t.Fatalf("expected user")
	}
}

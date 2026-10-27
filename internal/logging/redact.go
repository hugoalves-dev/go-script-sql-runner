package logging

import "github.com/hugoalves-dev/go-script-sql-runner/internal/redact"

// Redact is kept for callers that use the logging package directly.
func Redact(text string, secrets ...string) string {
	return redact.Secrets(text, secrets...)
}

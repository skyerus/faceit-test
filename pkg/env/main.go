package env

import (
	"os"
	"strings"
)

// SetEnv - set app env
func SetEnv() {
	os.Setenv("ENV", "dev")
	os.Setenv("DB_HOST", "mysql.faceit-test.local:3306")
	os.Setenv("DB_USER", "faceit")
	os.Setenv("DB_PASSWORD", "faceit")
	os.Setenv("DB_NAME", "faceit")
	os.Setenv("TZ", "Europe/London")
	var sb strings.Builder
	sb.WriteString(os.Getenv("DB_USER"))
	sb.WriteString(":")
	sb.WriteString(os.Getenv("DB_PASSWORD"))
	sb.WriteString("@tcp(")
	sb.WriteString(os.Getenv("DB_HOST"))
	sb.WriteString(")/")
	sb.WriteString(os.Getenv("DB_NAME"))
	os.Setenv("DB_URL", sb.String())
}

package ahhhhh

import "os"

// password check function. Need to use secure methods like bcrypt
func CheckPass(password string) bool {
	cli_pass := os.Getenv("CLI_PASS")
	return password == cli_pass
}

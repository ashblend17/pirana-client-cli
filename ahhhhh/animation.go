package ahhhhh

import (
	"fmt"
	"os"
	"os/exec"
)

func startAnimation() *exec.Cmd {
	path := "./ahhhhh/external/animation.sh"
	cmd := exec.Command("bash", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting animation:", err)
		return nil
	}
	return cmd
}

func stopAnimation(cmd *exec.Cmd) {
	if cmd != nil {
		err := cmd.Process.Kill()
		if err != nil {
			fmt.Println("Error stopping animation:", err)
		}
	}
	fmt.Println("\n[✔] API request completed!")
}

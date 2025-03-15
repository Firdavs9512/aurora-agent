package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

// ActiveCmd stores the currently active command (for CTRL+C handling)
var ActiveCmd *exec.Cmd

// RunCommandWithPTY runs a command with PTY to preserve colors and returns the output
func RunCommandWithPTY(cmd *exec.Cmd) string {
	ptmx, err := pty.Start(cmd)
	if err != nil {
		fmt.Println("Error:", err)
		return fmt.Sprintf("Error executing command: %v", err)
	}
	defer ptmx.Close()

	// Store the active process
	ActiveCmd = cmd

	// Buffer to collect command output
	var outputBuffer bytes.Buffer

	// Redirect PTY output to terminal and collect it
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := ptmx.Read(buf)
			if err != nil {
				break
			}
			// Write to stdout
			os.Stdout.Write(buf[:n])
			// Also collect in buffer
			outputBuffer.Write(buf[:n])
		}
	}()

	// Wait for command to complete
	cmd.Wait()

	// Clear activeCmd after process completes
	ActiveCmd = nil

	// Return the collected output
	return outputBuffer.String()
}

package utils

import (
	"aurora-agent/config"
	"fmt"
)

func SteamPrint(text string, ansiBuffer *string) {
	// If buffer contains ANSI code
	if config.AnsiPattern.MatchString(*ansiBuffer) {
		// Buffer has ANSI code, process it
		processedBuffer := ProcessANSICodes(*ansiBuffer)
		fmt.Print(processedBuffer)
		*ansiBuffer = ""
	} else if config.AnsiStartPattern.MatchString(*ansiBuffer) && len(*ansiBuffer) > 30 {
		// If buffer contains the start of an ANSI code, but not the end
		// and buffer length is more than 30, process it
		// This can happen when ANSI code is in incorrect format
		processedBuffer := ProcessANSICodes(*ansiBuffer)
		fmt.Print(processedBuffer)
		*ansiBuffer = ""
	} else if len(*ansiBuffer) > 20 && !config.AnsiStartPattern.MatchString(*ansiBuffer) {
		// If buffer length is more than 20 and no ANSI code start is found,
		// process it
		processedBuffer := ProcessANSICodes(*ansiBuffer)
		fmt.Print(processedBuffer)
		*ansiBuffer = ""
	}
}

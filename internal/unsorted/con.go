// Copyright © 2020 Hedzr Yeh.

package tool

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PressEnterToContinue prompts and waiting user's keystroke in terminal
func PressEnterToContinue(msg ...string) (input string) {
	if len(msg) > 0 && len(msg[0]) > 0 {
		fmt.Print(msg[0])
	} else {
		fmt.Print("Press 'Enter' to continue...")
	}
	b, _ := bufio.NewReader(os.Stdin).ReadBytes('\n')
	return strings.TrimRight(string(b), "\n")
}

// PressAnyKeyToContinue prompts and waiting user's keystroke in terminal
func PressAnyKeyToContinue(msg ...string) (input string) {
	if len(msg) > 0 && len(msg[0]) > 0 {
		fmt.Print(msg[0])
	} else {
		fmt.Print("Press any key to continue...")
	}
	_, _ = fmt.Scanf("%s", &input)
	return
}

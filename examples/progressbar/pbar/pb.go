// Copyright © 2020 Hedzr Yeh.

package pbar

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	total = 50
	count = 0
	wscol = 20
)

func init() {
	err := updateWSCol()
	if err != nil {
		panic(err)
	}
}

func updateWSCol() error {
	ws, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
	if err != nil {
		return err
	}
	wscol = int(ws.Col)
	return nil
}

func renderbar() {
	fmt.Print("\x1b7")       // save the cursor position
	fmt.Print("\x1b[2k")     // erase the current line
	defer fmt.Print("\x1b8") // restore the cursor position

	barwidth := wscol - len("Progress: 100% []")
	done := int(float64(barwidth) * float64(count) / float64(total))

	fmt.Printf("Progress: \x1b[33m%3d%%\x1b[0m ", count*100/total)
	fmt.Printf("[%s%s]",
		strings.Repeat("=", done),
		strings.Repeat("-", barwidth-done))
}

func main1() {
	done := make(chan struct{})
	Run(done)
	fmt.Println()
}

func Run(done chan struct{}) {
	// set signal handler
	sigwinch := make(chan os.Signal, 1)
	defer close(sigwinch)
	signal.Notify(sigwinch, syscall.SIGWINCH)
	go func() {
		for {
			select {
			case _, ok := <-sigwinch:
				if !ok {
					close(done)
					return
				}
			}
			_ = updateWSCol()
			renderbar()
		}
	}()
	for count = 1; count <= 50; count++ {
		renderbar()
		time.Sleep(330 * time.Millisecond)
	}
}

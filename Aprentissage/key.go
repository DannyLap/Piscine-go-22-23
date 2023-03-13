package main

import (
	"fmt"

	"devt.de/krotik/common/termutil/getch"
)

func main() {
	var err error
	var e *getch.KeyEvent

	if err = getch.Start(); err != nil {
		fmt.Println(err)
		return
	}
	defer getch.Stop()

	for e == nil || e.Code != getch.KeyTab {
		e, err = getch.Getch()

		fmt.Println(e)
	}
}

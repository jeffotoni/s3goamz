/***********
*
*
* project s3 Upload
*
* @package     main
* @author      jeffotoni
* @copyright   Copyright (c) 2017
* @license     --
* @link        --
* @since       Version 0.1
*
 */
package runer

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/fatih/color"
)

func RunerTimer() <-chan time.Time {

	//
	//
	//
	white := color.New(color.FgWhite)
	boldWhite := white.Add(color.Bold)

	//
	//
	//
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	timer := time.Tick(time.Duration(50) * time.Millisecond)

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)

		<-sc

		boldRed.Println("\ncanceled!")
		fmt.Print("\033[?25h")
		os.Exit(0)
	}()

	fmt.Print("\033[?25l")

	s := []rune(`|/~\`)
	//s := []rune(`-=*=`)
	//s := []rune(`◐◓◑◒`)
	i := 0

	go func() {
		for {

			<-timer

			fmt.Print("\r")
			boldWhite.Print(string(s[i]))

			i++

			if i == len(s) {
				i = 0
			}
		}
	}()

	return timer
}

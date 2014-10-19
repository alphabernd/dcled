/*
 * ----------------------------------------------------------------------------
 * "THE PIZZA-WARE LICENSE" (Revision 42):
 * <whoami@dev-urandom.eu> wrote this file.  As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a pizza in return.
 * ----------------------------------------------------------------------------
 */

package main

import (
	"flag"
	"github.com/imaohw/libdcled"
	"os"
	"sync"
)

func main() {
	var message string
	var scroll bool
	var debug bool

	flag.StringVar(&message, "text", "no text was given", "Text to display")
	flag.BoolVar(&scroll, "scroll", false, "Scroll text")
	flag.BoolVar(&debug, "debug", false, "Debug mode")
	flag.Parse()

	libdcled.Debug = debug

	led, err := libdcled.NewDcLed()
	if err != nil {
		os.Exit(2)
	}

	if scroll {
		led.ScrollText(libdcled.NewText(message, libdcled.STANDARD_FONT))
	} else {
		led.PrintText(libdcled.NewText(message, libdcled.STANDARD_FONT))
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 181.

// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"

	"gopl.io/ch7/tempconv"
)

// !+
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature(必须以C或F结尾)")
var fTemp = tempconv.FahrenheitFlag("fTemp", 100.0, "the f temperature(必须以C或F结尾)")

func main() {
	flag.Parse()
	fmt.Println(*temp)
	fmt.Println("华氏温度:", *fTemp)
}

//!-

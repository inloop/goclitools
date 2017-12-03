package goclitools

import (
	"fmt"
	"strings"
)

// Log ...
func Log(m ...interface{}) {
	fmt.Println(m...)
}

// LogSection ...
func LogSection(title string, m ...interface{}) {
	header := fmt.Sprintf("==== %s ====", title)
	fmt.Println("\n" + header)
	fmt.Println(strings.Repeat("=", len(header)))
	if len(m) > 0 {
		fmt.Print("\n")
		fmt.Println(m...)
		fmt.Println(strings.Repeat("=", len(header)))
	}
}

// PrintChecking ...
func PrintChecking(message string) {
	fmt.Print("checking " + message + ": ")
}

// PrintNotOK ...
func PrintNotOK() {
	fmt.Println("✘")
}

// PrintOK ...
func PrintOK() {
	fmt.Println("✓")
}

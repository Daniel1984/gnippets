/*
* The issue: you have 3rd party dependency (in this case logentries logger)
* and you want it to output to logentries api when in production/staging or
* development but output to standard output while developing locally.
*	One way to achieve this is to create a custom struct and interface
*	where you specify most commonly used logentries methods or all of them if
*	you'd like.
 */

package main

import (
	"fmt"
	"github.com/bsphere/le_go"
	"os"
)

// we only use Println and Printf logentries methods in our app
// but if needed all methods could be replicated
type logger interface {
	Println(...interface{})
	Printf(string, ...interface{})
}

// fmtLogger has now 2 methods to satisfy interface
type fmtLogger struct{}

func (l *fmtLogger) Println(args ...interface{}) {
	fmt.Println(args...)
}

func (l *fmtLogger) Printf(s string, args ...interface{}) {
	fmt.Printf(s, args...)
}

func main() {
	logger := newLogger()
	logger.Printf("Testing Printf %v", 123)
}

// both, our local fmtLogger and logentries instance, satisfy our logger interface
// so we can conditionally return either of them from newLogger function
func newLogger() logger {
	leKey := os.Getenv("LOG_ENTRIES_KEY")
	fmtFallback := &fmtLogger{}

	if leKey == "" {
		return fmtFallback
	}

	le, err := le_go.Connect(leKey)

	if err != nil {
		return fmtFallback
	}

	defer le.Close()

	return le
}

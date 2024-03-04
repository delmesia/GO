package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// The interface that main and test will use
type Sleeper interface {
	Sleep()
}

// A spy struct that will spy on how many times the Sleep is called.
// This will be used by the test to skip the waiting time on Sleep().
// Each call on the Spysleeper will increment Calls by 1.
type SpySleeper struct {
	Calls int
}

// This struct that implements the sleeper interface
// to be used in the main function.
type DefaultSleeper struct{}

// A method for DefaultSleeper to be used in the main function.
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// A method for Spysleeper to be used by the tests.
func (s *SpySleeper) Sleep() {
	s.Calls++
}

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer, sleeper Sleeper) {
	// io.Writer has been used so it can be used by
	// any code that implements the io.Writer interface
	// i.e. os.Stdout, bytes.Buffer, etc.
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		// All structs that implements the Sleeper interface
		// can be used in this function as a parameter.
		// i.e DefaultSleeper for main, Spysleeper for tests
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

package main

import (
	"os"
	"time"

	"github.com/mjishu/intro_learn_go_with_tests/mocking"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{1 * time.Second, time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}

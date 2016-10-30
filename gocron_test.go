// Tests for gocron
package gocron

import (
	"fmt"
	"github.com/chacken/gocron"
	"testing"
	"time"
)

var err = 1

func task() {
	fmt.Println("I am a running job.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func TestSecond(*testing.T) {
	s := gocron.NewScheduler()
	s.Job("123").Every(1).Second().Do(task)
	s.Job("234").Every(1).Second().Do(taskWithParams, 1, "hello")
	s.Start()
	time.Sleep(10 * time.Second)
}

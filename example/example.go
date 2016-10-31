package main

import (
	"fmt"

	"github.com/chacken/gocron"
)

func task() {
	fmt.Println("I am runnning task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {

	s := gocron.NewScheduler()

	// Do jobs with params
	s.Job("123").Every(1).Second().Do(taskWithParams, 1, "hello")

	// Do jobs without params
	s.Job("abc").Every(1).Second().Do(task)
	s.Job("def").Every(2).Seconds().Do(task)
	s.Job("ghi").Every(1).Minute().Do(task)
	s.Job("jkl").Every(2).Minutes().Do(task)
	s.Job("mno").Every(1).Hour().Do(task)
	s.Job("pqr").Every(2).Hours().Do(task)
	s.Job("stu").Every(1).Day().Do(task)
	s.Job("vwx").Every(2).Days().Do(task)

	// Do jobs on specific weekday
	s.Job("yz").Every(1).Monday().Do(task)
	s.Job("123").Every(1).Thursday().Do(task)

	// function At() take a string like 'hour:min'
	s.Job("456").Every(1).Day().At("10:30").Do(task)
	s.Job("789").Every(1).Monday().At("18:30").Do(task)

	// remove, clear and next_run
	_, time := s.NextRun()
	fmt.Println(time)

	// gocron.Remove(task)
	// gocron.Clear()

	// function Start start all the pending jobs
	<-s.Start()
}

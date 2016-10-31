## goCron: A Golang Job Scheduling Package.
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/chacken/gocron)
[![Stories in Ready](https://badge.waffle.io/chacken/gocron.png?label=ready&title=Ready)](https://waffle.io/chacken/gocron)

goCron is a Golang job scheduling package which lets you run Go functions periodically at pre-determined interval using a simple, human-friendly syntax.

goCron is a Golang implementation of Ruby module [clockwork](<https://github.com/tomykaira/clockwork>) and Python job scheduling package [schedule](<https://github.com/dbader/schedule>), and personally, this package is my first Golang program, just for fun and practice.

See also this two great articles:
* [Rethinking Cron](http://adam.heroku.com/past/2010/4/13/rethinking_cron/)
* [Replace Cron with Clockwork](http://adam.heroku.com/past/2010/6/30/replace_cron_with_clockwork/)

Back to this package, you could just use this simple API as below, to run a cron scheduler.

``` go
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
	s.Job("8c1f99f3-2b6e-4fdb-9656-b50c91bfa740").Every(1).Second().Do(taskWithParams, 1, "hello")

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
	_, time := gocron.NextRun()
	fmt.Println(time)

	s.Remove("abc")
	s.Clear()

	// function Start start all the pending jobs
	<- s.Start()

}
```
and full test cases and [document](http://godoc.org/github.com/chacken/gocron) will be coming soon.

Once again, thanks to the great works of Ruby clockwork and Python schedule package. BSD license is used, see the file License for detail.

Hava fun!

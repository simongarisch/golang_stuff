package main

import (
	"fmt"

	"github.com/whiteshtef/clockwork"
)

func main() {
	sched := clockwork.NewScheduler()

	sched.Schedule().Every(10).Seconds().Do(every10Sec)
	sched.Schedule().Every(3).Minutes().Do(something)
	sched.Schedule().Every(4).Hours().Do(something)
	sched.Schedule().Every(2).Days().At("12:32").Do(something)
	sched.Schedule().Every(12).Weeks().Do(something)

	sched.Schedule().Every().Second().Do(every1Sec) // Every() is "shorthand" for Every(1)
	sched.Schedule().Every().Monday().Do(something)
	sched.Schedule().Every().Saturday().At("8:00").Do(something)

	sched.Run()
}

func something() {
	fmt.Println("foo")
}

func every10Sec() {
	fmt.Println("Every 10 seconds")
}

func every1Sec() {
	fmt.Println("Every 1 second")
}

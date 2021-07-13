package tasks

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	fmt.Println(123)
	f := func(args ...interface{}) {
		fmt.Println(args)
	}

	DefaultTasks.AddTask(
		"example",
		"描述",
		"* * * * *",
		true,
		f,
		[]interface{}{
			1,
		},
		true,
	)
	DefaultTasks.StartAsync()
	select {}
}

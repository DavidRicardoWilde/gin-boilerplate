package examples

import "fmt"

type ExampleTask struct {
}

func (t *ExampleTask) Name() string {
	return "example"
}

func (t *ExampleTask) Exec() {
	fmt.Println("example task")
}

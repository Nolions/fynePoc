package internal

import "fmt"

type DemoError struct {
	Title   string
	Message string
}

func (e DemoError) Error() string {
	return fmt.Sprintf("%v, %v", e.Title, e.Message)
}

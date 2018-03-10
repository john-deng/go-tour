package counters

import "fmt"

// AlertCounter type
type AlertCounter int

func init()  {
	fmt.Println("counter.init()")
}


// New creates and returns values of the unexported type alertCounter.
func New(value int) AlertCounter {
	return AlertCounter(value)
}

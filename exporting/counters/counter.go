package counters

// AlertCounter type
type AlertCounter int

// New creates and returns values of the unexported type alertCounter.
func New(value int) AlertCounter {
	return AlertCounter(value)
}

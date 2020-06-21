package server

import (
	"context"
	"fmt"
)

// RootResolver is the root resolver
type RootResolver struct{}

// Greet greets
func (*RootResolver) Greet() string {
	return "Hello, world!"
}

// GreetPerson greets prerson
func (*RootResolver) GreetPerson(args struct{ Person string }) string {
	return fmt.Sprintf("Hello, %s!", args.Person)
}

// GreetPersonTimeOfDay greets person time of day
func (*RootResolver) GreetPersonTimeOfDay(ctx context.Context, args PersonTimeOfDayArgs) string {
	timeOfDay, ok := TimesOfDay[args.TimeOfDay]
	if !ok {
		timeOfDay = "Go to bed"
	}
	return fmt.Sprintf("%s, %s!", timeOfDay, args.Person)
}

// PersonTimeOfDayArgs is a structure
type PersonTimeOfDayArgs struct {
	Person    string // Note that fields need to be exported.
	TimeOfDay string
}

// TimesOfDay is a struct
var TimesOfDay = map[string]string{
	"MORNING":   "Good morning",
	"AFTERNOON": "Good afternoon",
	"EVENING":   "Good evening",
}

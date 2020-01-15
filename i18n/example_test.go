package i18n_test

import (
	"fmt"

	"github.com/icza/gox/i18n"
)

// Define your locales
const (
	EN = iota
	HU
	DE
	GR
)

var Monday = i18n.Dict{
	EN: "Monday",
	DE: "Montag",
	HU: i18n.Empty, // We want this to be empty
}.Get

var Introduce = i18n.Dict{
	EN: "My name is %s, and I'm %d years old.",
	DE: "Mein Name ist %s und ich bin %d Jahre alt.",
}.Get

func Example() {
	fmt.Printf("Monday in EN: %s\n", Monday(EN))
	fmt.Printf("Monday in DE: %s\n", Monday(DE))
	fmt.Printf("Monday in HU: %s (empty)\n", Monday(HU))
	fmt.Printf("Monday in GR: %s (missing, defaults to EN)\n", Monday(GR))

	fmt.Println(Introduce(EN, "Bob", 22))
	fmt.Println(Introduce(DE, "Alice", 12))

	// Output:
	// Monday in EN: Monday
	// Monday in DE: Montag
	// Monday in HU:  (empty)
	// Monday in GR: Monday (missing, defaults to EN)
	// My name is Bob, and I'm 22 years old.
	// Mein Name ist Alice und ich bin 12 Jahre alt.
}

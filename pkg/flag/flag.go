package flag

import "strings"

// Flag represents all of the necessary functions to accurately describe and use
// a flag.
type Flag interface {
	Seed() string
	Input() []byte
	InputAsString() string
	Data() []byte
	DataAsString() string
	String() string
	Raw() string
}

type flag struct {
	input []byte
	value string
	seed  string
	data  []byte
}

// Seed returns the seed used to generate the flag.
func (fl *flag) Seed() string {
	return fl.seed
}

// Input returns the input used to generate the flag.
func (fl *flag) Input() []byte {
	return fl.input
}

// InputAsString returns the input used to generate the flag, but as a string
func (fl *flag) InputAsString() string {
	return string(fl.input)
}

// Data returns the concatenation of the input and the seed which is fed into the
// hashing function as a byte slice.
func (fl *flag) Data() []byte {
	return fl.data
}

// DataAsString returns the concatenation of the input and the seed which is fed into the
// hashing function as a a string.
func (fl *flag) DataAsString() string {
	return string(fl.data)
}

// String returns the flag's value as a string. This includes the yctf prefix and the hash.
func (fl *flag) String() string {
	return strings.Join([]string{prefix, fl.value}, "-")
}

// Raw returns the flag's hash as a string without the prefix.
func (fl *flag) Raw() string {
	return fl.value
}

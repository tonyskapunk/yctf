package flag

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

const (
	prefix = "yctf"
)

// Factory represents the flag factory interface. It's what is needed
// to generate a flag and validate flags have been generated through it.
type Factory interface {
	Seed() string
	NewFlag(in []byte) *flag
	Validate(f *Flag) bool
}

type factory struct {
	seed      string
	generated []string
}

// NewFactory returns a simple factory that helps generate flags.
func NewFactory(seed string) *factory {
	return &factory{
		seed: seed,
	}
}

// Seed returns the seed provided as input to the factory.
func (f *factory) Seed() string {
	return f.seed
}

// Validate checks a flag to see if the seed stored in the flag
// matches the seed in the factory. It also regenerates the hash based on the
// inputs stored in the flag to ensure the value matches based on the input
func (f *factory) Validate(fl Flag) bool {
	return fl.Seed() == f.Seed() && fl.Raw() == fmt.Sprintf("%x", sha256.Sum256(fl.Data()))
}

// NewFlag returns a flag containing a the input value, the seed stored in the factory,
// and the calculated value of that flag. The calculated value of that flag is a
// sha256 sum of the input concatenated with the seed.
func (f *factory) NewFlag(in []byte) *flag {
	data := bytes.Join([][]byte{in, []byte(f.seed)}, []byte{})
	sum := sha256.Sum256(data)
	return &flag{
		input: in,
		value: fmt.Sprintf("%x", sum),
		seed:  f.seed,
		data:  data,
	}
}

package generators

import (
	"time"
)

// Generator is the default generator contract
type Generator interface {
	Generate(date time.Time, params map[string]interface{}) ([]string, error)
}

// Generators contains all supported generator instances
var Generators map[string]Generator

// Register a generator to a ransomware
func Register(name string, generator Generator) {
	if _, ok := Generators[name]; !ok {
		Generators[name] = generator
	}
}

func init() {
	Generators = map[string]Generator{}
}

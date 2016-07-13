package locky

import (
	"time"

	"zenithar.org/go/dgabloom/generators"
)

type lockyGenerator struct {
}

func init() {
	generators.Register("locky", &lockyGenerator{})
}

// -----------------------------------------------------------------------------

func (g *lockyGenerator) Generate(date time.Time, params map[string]interface{}) ([]string, error) {
	res := []string{}

	for config := 0; config < len(configs); config++ {
		for i := uint32(0); i < 12; i++ {
			res = append(res, dga(date, uint32(config), i))
		}
	}

	return res, nil
}

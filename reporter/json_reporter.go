package reporter

import (
	"encoding/json"

	"github.com/mfbmina/enxame/swarm"
)

type JSONReporter struct{}

func (r JSONReporter) Report(responses []swarm.HTTPResponse) string {
	j, err := json.Marshal(responses)
	if err != nil {
		// TODO: handle error better
		return err.Error()
	}

	return string(j)
}

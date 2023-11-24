package reporter

import (
	"encoding/json"

	"github.com/mfbmina/enxame/swarm"
)

type JSONReporter struct {
	Responses []swarm.HTTPResponse `json:"results"`
}

func (r JSONReporter) Report() string {
	j, err := json.Marshal(r.Responses)
	if err != nil {
		// TODO: handle error
		return "malformed json"
	}

	return string(j)
}

package reporter

import (
	"encoding/json"
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

type JSONReporter struct{}

func (r JSONReporter) Report(responses []swarm.HTTPResponse) (string, error) {
	j, err := json.Marshal(responses)
	if err != nil {
		msg := fmt.Sprintf("JSONReporter.Report(): Error marshaling responses: %s\n", err.Error())
		return "", fmt.Errorf(msg)
	}

	return string(j), nil
}

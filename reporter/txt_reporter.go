package reporter

import (
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

type TXTReporter struct {
	Responses []swarm.HTTPResponse
}

func (r TXTReporter) Report(responses []swarm.HTTPResponse) string {
	report := ""
	for _, r := range responses {
		report += fmt.Sprintf("%d %s %s\n", r.StatusCode, r.Time, r.Path)
	}
	return report
}

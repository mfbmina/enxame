package reporter

import (
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

type TXTReporter struct {
	Responses []swarm.HTTPResponse
}

func (r TXTReporter) Report() string {
	report := ""
	for _, r := range r.Responses {
		report += fmt.Sprintf("%d %s %s\n", r.StatusCode, r.Time, r.Path)
	}
	return report
}

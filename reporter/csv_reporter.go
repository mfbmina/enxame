package reporter

import (
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

const CSV_HEADERS = "status_code,time,path"

type CSVReporter struct {
	Responses []swarm.HTTPResponse
}

func (r CSVReporter) Report() string {
	report := CSV_HEADERS
	for _, r := range r.Responses {
		report += fmt.Sprintf("\n%d,%s,%s", r.StatusCode, r.Time, r.Path)
	}
	return report
}

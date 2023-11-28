package reporter

import (
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

const CSV_HEADERS = "status_code,time,path"

type CSVReporter struct{}

func (r CSVReporter) Report(responses []swarm.HTTPResponse) string {
	report := CSV_HEADERS
	for _, r := range responses {
		report += fmt.Sprintf("\n%d,%s,%s", r.StatusCode, r.Time, r.Path)
	}
	return report
}

package reporter

import (
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

const CSV_HEADERS = "status_code,time,path"

type CSVReport struct {
	Responses []swarm.HTTPResponse
}

func (r CSVReport) WriteToFile() {
	panic("CSVReport.WriteToFile() not implemented")
}

func (r CSVReport) WriteToStdout() {
	fmt.Println("Reporting results as csv...")
	fmt.Println("-----------------------------")
	fmt.Println(r.createReport())
}

func (r CSVReport) createReport() string {
	report := CSV_HEADERS
	for _, r := range r.Responses {
		report += fmt.Sprintf("\n%d,%s,%s", r.StatusCode, r.Time, r.Path)
	}
	return report
}

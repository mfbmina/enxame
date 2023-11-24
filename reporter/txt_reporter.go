package reporter

import (
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

const TYPE = "csv"

type TXTReport struct {
	Responses []swarm.HTTPResponse
}

func (r TXTReport) WriteToFile() {
	panic("TXTReport.WriteToFile() not implemented")
}

func (r TXTReport) WriteToStdout() {
	fmt.Println("Reporting results to stdout...")
	fmt.Println("-----------------------------")
	fmt.Printf(r.createReport())
	fmt.Println("-----------------------------")
}

func (r TXTReport) createReport() string {
	report := ""
	for _, r := range r.Responses {
		report += fmt.Sprintf("%d %s %s\n", r.StatusCode, r.Time, r.Path)
	}
	return report
}

package reporter

import (
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

type CsvReport struct {
	Responses []swarm.HTTPResponse
}

func (cR CsvReport) Report() {
	fmt.Println("Reporting results as csv...")
	fmt.Println("-----------------------------")
	fmt.Println("status_code,time,path")
	for _, r := range cR.Responses {
		fmt.Printf("%d,%s,%s\n", r.StatusCode, r.Time, r.Path)
	}
}

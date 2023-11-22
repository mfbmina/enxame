package reporter

import (
	"fmt"

	"github.com/mfbmina/enxame/lib/swarm"
)

type CsvReport struct {
	Responses []swarm.Response
}

func (cR CsvReport) Report() {
	fmt.Println("Reporting results as csv...")
	fmt.Println("-----------------------------")
	fmt.Println("status_code,time,path")
	for _, r := range cR.Responses {
		fmt.Printf("%d,%s,%s\n", r.StatusCode, r.Time, r.Path)
	}
}

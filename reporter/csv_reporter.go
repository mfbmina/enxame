package reporter

import (
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

type CSVReport struct {
	Responses []swarm.HTTPResponse
}

func (r CSVReport) Report() {
	fmt.Println("Reporting results as csv...")
	fmt.Println("-----------------------------")
	fmt.Println("status_code,time,path")
	for _, r := range r.Responses {
		fmt.Printf("%d,%s,%s\n", r.StatusCode, r.Time, r.Path)
	}
}

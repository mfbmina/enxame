package reporter

import (
	"fmt"

	"github.com/mfbmina/enxame/lib/swarm"
)

type StdoutReport struct {
	Responses []swarm.Response
}

func (sR StdoutReport) Report() {
	fmt.Println("Reporting results to stdout...")
	fmt.Println("-----------------------------")
	for _, r := range sR.Responses {
		fmt.Println(r.StatusCode, r.Time, r.Path)
	}
	fmt.Println("-----------------------------")
}

package reporter

import (
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

type TXTReport struct {
	Responses []swarm.HTTPResponse
}

func (r TXTReport) Report() {
	fmt.Println("Reporting results to stdout...")
	fmt.Println("-----------------------------")
	for _, r := range r.Responses {
		fmt.Println(r.StatusCode, r.Time, r.Path)
	}
	fmt.Println("-----------------------------")
}

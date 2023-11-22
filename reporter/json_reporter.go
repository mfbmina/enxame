package reporter

import (
	"encoding/json"
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

type JSONReport struct {
	Responses []swarm.HTTPResponse `json:"results"`
}

func (r JSONReport) Report() {
	fmt.Println("Reporting results as JSON...")
	fmt.Println("-----------------------------")
	j, err := json.Marshal(r.Responses)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(j))
}

package reporter

import (
	"encoding/json"
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

type JSONReport struct {
	Results []swarm.Response `json:"results"`
}

func (cR JSONReport) Report() {
	fmt.Println("Reporting results as JSON...")
	fmt.Println("-----------------------------")
	j, err := json.Marshal(cR.Results)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(j))
}

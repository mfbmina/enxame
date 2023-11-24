package reporter

import (
	"encoding/json"
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

type JSONReport struct {
	Responses []swarm.HTTPResponse `json:"results"`
}

func (r JSONReport) WriteToStdout() {
	fmt.Println("Reporting results as JSON...")
	fmt.Println("-----------------------------")

	fmt.Println(r.createReport())
}

func (r JSONReport) createReport() string {
	j, err := json.Marshal(r.Responses)
	if err != nil {
		fmt.Println(err)
		return "malformed json"
	}

	return string(j)
}

package reporter

import (
	"fmt"

	"github.com/mfbmina/enxame/swarm"
)

type Report struct {
	Output    string
	Reporter  Reporter
	Responses []swarm.HTTPResponse
	Type      string
}

type Reporter interface {
	Report() string
}

func NewReporter(reportType, output string, responses []swarm.HTTPResponse) Report {
	r := Report{Output: output, Responses: responses, Type: reportType}
	switch reportType {
	case "json":
		r.Reporter = JSONReporter{Responses: responses}
	case "csv":
		r.Reporter = CSVReporter{Responses: responses}
	default:
		r.Reporter = TXTReporter{Responses: responses}
	}

	return r
}

func (r Report) Report() {
	if r.Output == "" {
		r.writeToStdout()
	} else {
		r.writeToFile()
	}
}

func (r Report) writeToStdout() {
	fmt.Printf("Reporting results as %s...\n", r.Type)
	fmt.Println("------------------------------")
	fmt.Println(r.Reporter.Report())
}

func (r Report) writeToFile() {
}

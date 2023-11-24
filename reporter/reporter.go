package reporter

import (
	"fmt"
	"os"

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
	fmt.Printf("Reporting results as %s...\n", r.Type)
	if r.Output == "" {
		r.writeToStdout()
	} else {
		r.writeToFile()
	}
}

func (r Report) writeToStdout() {
	fmt.Println("------------------------------")
	fmt.Println(r.Reporter.Report())
}

func (r Report) writeToFile() {
	name := fmt.Sprintf("%s.%s", r.Output, r.Type)
	f, err := os.Create(name)
	if err != nil {
		// TODO: handle error better
		fmt.Printf("Error creating file %s: %s\n", name, err.Error())
		return
	}

	defer f.Close()

	_, err = f.WriteString(r.Reporter.Report())
	if err != nil {
		// TODO: handle error
		fmt.Printf("Error writing to file %s: %s\n", name, err.Error())
		return
	}

	fmt.Println("Report saved to", name, "successfully!")
}

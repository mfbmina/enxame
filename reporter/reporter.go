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
	Report([]swarm.HTTPResponse) (string, error)
}

func NewReporter(reportType, output string, responses []swarm.HTTPResponse) Report {
	r := Report{Output: output, Responses: responses, Type: reportType}
	switch reportType {
	case "json":
		r.Reporter = JSONReporter{}
	case "csv":
		r.Reporter = CSVReporter{}
	default:
		r.Reporter = TXTReporter{}
	}

	return r
}

func (r Report) Report() error {
	fmt.Printf("Reporting results as %s...\n", r.Type)
	raw, err := r.Reporter.Report(r.Responses)
	if err != nil {
		return fmt.Errorf("Report.Report(): Error reporting results: %s\n", err.Error())
	}

	if r.Output == "" {
		r.writeToStdout(raw)
	} else {
		r.writeToFile(raw)
	}

	return nil
}

func (r Report) writeToStdout(raw string) {
	fmt.Println("------------------------------")
	fmt.Println(raw)
}

func (r Report) writeToFile(raw string) {
	name := fmt.Sprintf("%s.%s", r.Output, r.Type)
	f, err := os.Create(name)
	if err != nil {
		// TODO: handle error better
		fmt.Printf("Error creating file %s: %s\n", name, err.Error())
		return
	}

	defer f.Close()

	_, err = f.WriteString(raw)
	if err != nil {
		// TODO: handle error
		fmt.Printf("Error writing to file %s: %s\n", name, err.Error())
		return
	}

	fmt.Println("Report saved to", name, "successfully!")
}

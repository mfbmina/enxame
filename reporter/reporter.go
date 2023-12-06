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

func (r Report) Report() (string, error) {
	raw, err := r.Reporter.Report(r.Responses)
	if err != nil {
		return "", fmt.Errorf("Report.Report(): Error formating results: %s\n", err.Error())
	}

	if r.Output == "" {
		return raw, nil
	}

	return r.writeToFile(raw)
}

func (r Report) writeToFile(raw string) (string, error) {
	name := fmt.Sprintf("%s.%s", r.Output, r.Type)
	f, err := os.Create(name)
	if err != nil {
		return "", fmt.Errorf("Report.writeToFile(): Error creating file: %s\n", err.Error())
	}

	defer f.Close()

	_, err = f.WriteString(raw)
	if err != nil {
		return "", fmt.Errorf("Report.writeToFile(): Error writing to file: %s\n", err.Error())
	}

	return fmt.Sprintf("Report created at %s", name), nil
}

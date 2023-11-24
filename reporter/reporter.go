package reporter

import (
	"github.com/mfbmina/enxame/swarm"
)

type Reporter interface {
	createReport() string
	WriteToStdout()
}

func NewReporter(reportType string, responses []swarm.HTTPResponse) Reporter {
	switch reportType {
	case "csv":
		return CSVReport{Responses: responses}
	case "json":
		return JSONReport{Responses: responses}
	default:
		return TXTReport{Responses: responses}
	}
}

package reporter

import (
	"testing"
	"time"

	"github.com/mfbmina/enxame/swarm"
	"github.com/stretchr/testify/assert"
)

func TestTXTReportWhenResponsesAreNil(t *testing.T) {
	r := TXTReport{}
	expected := "Reporting results to stdout...\n-----------------------------\n-----------------------------\n"

	output := captureOutput(func() { r.Report() })
	assert.Equal(t, expected, output)
}

func TestTXTReportWhenResponsesAreEmpty(t *testing.T) {
	r := TXTReport{Responses: []swarm.HTTPResponse{}}
	expected := "Reporting results to stdout...\n-----------------------------\n-----------------------------\n"

	output := captureOutput(func() { r.Report() })
	assert.Equal(t, expected, output)
}

func TestTXTReportWhenResponsesExists(t *testing.T) {
	duration, _ := time.ParseDuration("1ms")
	r := TXTReport{Responses: []swarm.HTTPResponse{{StatusCode: 200, Time: duration, Path: "/users"}}}
	expected := "Reporting results to stdout...\n-----------------------------\n200 1ms /users\n-----------------------------\n"

	output := captureOutput(func() { r.Report() })
	assert.Equal(t, expected, output)
}

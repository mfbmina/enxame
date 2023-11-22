package reporter

import (
	"testing"
	"time"

	"github.com/mfbmina/enxame/swarm"
	"github.com/stretchr/testify/assert"
)

func TestStdoutReportWhenResponsesAreNil(t *testing.T) {
	r := StdoutReport{}
	expected := "Reporting results to stdout...\n-----------------------------\n-----------------------------\n"

	output := captureOutput(func() { r.Report() })
	assert.Equal(t, expected, output)
}

func TestStdoutReportWhenResponsesAreEmpty(t *testing.T) {
	r := StdoutReport{Responses: []swarm.Response{}}
	expected := "Reporting results to stdout...\n-----------------------------\n-----------------------------\n"

	output := captureOutput(func() { r.Report() })
	assert.Equal(t, expected, output)
}

func TestStdoutReportWhenResponsesExists(t *testing.T) {
	duration, _ := time.ParseDuration("1ms")
	r := StdoutReport{Responses: []swarm.Response{{StatusCode: 200, Time: duration, Path: "/users"}}}
	expected := "Reporting results to stdout...\n-----------------------------\n200 1ms /users\n-----------------------------\n"

	output := captureOutput(func() { r.Report() })
	assert.Equal(t, expected, output)
}

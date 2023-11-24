package reporter

import (
	"testing"
	"time"

	"github.com/mfbmina/enxame/swarm"
	"github.com/stretchr/testify/assert"
)

func TestJSONReportWhenResponsesAreNil(t *testing.T) {
	r := JSONReport{}
	expected := "Reporting results as JSON...\n-----------------------------\nnull\n"

	output := captureOutput(func() { r.WriteToStdout() })
	assert.Equal(t, expected, output)
}

func TestJSONReportWhenResponsesAreEmpty(t *testing.T) {
	r := JSONReport{Responses: []swarm.HTTPResponse{}}
	expected := "Reporting results as JSON...\n-----------------------------\n[]\n"

	output := captureOutput(func() { r.WriteToStdout() })
	assert.Equal(t, expected, output)
}

func TestJSONReportWhenResponsesExists(t *testing.T) {
	duration, _ := time.ParseDuration("1ms")
	r := JSONReport{Responses: []swarm.HTTPResponse{{StatusCode: 200, Time: duration, Path: "/users"}}}
	expected := "Reporting results as JSON...\n-----------------------------\n[{\"status_code\":200,\"time\":1000000,\"path\":\"/users\"}]\n"

	output := captureOutput(func() { r.WriteToStdout() })
	assert.Equal(t, expected, output)
}

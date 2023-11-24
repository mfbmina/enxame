package reporter

import (
	"testing"
	"time"

	"github.com/mfbmina/enxame/swarm"
	"github.com/stretchr/testify/assert"
)

func TestCSVReportWhenResponsesAreNil(t *testing.T) {
	r := CSVReport{}
	expected := "Reporting results as csv...\n-----------------------------\nstatus_code,time,path\n"

	output := captureOutput(func() { r.WriteToStdout() })
	assert.Equal(t, expected, output)
}

func TestCSVReportWhenResponsesAreEmpty(t *testing.T) {
	r := CSVReport{Responses: []swarm.HTTPResponse{}}
	expected := "Reporting results as csv...\n-----------------------------\nstatus_code,time,path\n"

	output := captureOutput(func() { r.WriteToStdout() })
	assert.Equal(t, expected, output)
}

func TestCSVReportWhenResponsesExists(t *testing.T) {
	duration, _ := time.ParseDuration("1ms")
	r := CSVReport{Responses: []swarm.HTTPResponse{{StatusCode: 200, Time: duration, Path: "/users"}}}
	expected := "Reporting results as csv...\n-----------------------------\nstatus_code,time,path\n200,1ms,/users\n"

	output := captureOutput(func() { r.WriteToStdout() })
	assert.Equal(t, expected, output)
}

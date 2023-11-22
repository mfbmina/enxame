package reporter

import (
	"testing"
	"time"

	"github.com/mfbmina/enxame/swarm"
	"github.com/stretchr/testify/assert"
)

func TestCsvReportWhenResponsesAreNil(t *testing.T) {
	sR := CsvReport{}
	expected := "Reporting results as csv...\n-----------------------------\nstatus_code,time,path\n"

	output := captureOutput(func() { sR.Report() })
	assert.Equal(t, expected, output)
}

func TestCsvReportWhenResponsesAreEmpty(t *testing.T) {
	sR := CsvReport{Responses: []swarm.Response{}}
	expected := "Reporting results as csv...\n-----------------------------\nstatus_code,time,path\n"

	output := captureOutput(func() { sR.Report() })
	assert.Equal(t, expected, output)
}

func TestCsvReportWhenResponsesExists(t *testing.T) {
	duration, _ := time.ParseDuration("1ms")
	sR := CsvReport{Responses: []swarm.Response{{StatusCode: 200, Time: duration, Path: "/users"}}}
	expected := "Reporting results as csv...\n-----------------------------\nstatus_code,time,path\n200,1ms,/users\n"

	output := captureOutput(func() { sR.Report() })
	assert.Equal(t, expected, output)
}

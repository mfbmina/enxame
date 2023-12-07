package reporter

import (
	"testing"
	"time"

	"github.com/mfbmina/enxame/swarm"
	"github.com/stretchr/testify/assert"
)

func Test_TXTReporter_Report_WhenResponsesIsNil(t *testing.T) {
	r := TXTReporter{}
	expected := ""

	resp, _ := r.Report(nil)
	assert.Equal(t, expected, resp)
}

func Test_TXTReporter_Report_WhenResponsesIsEmpty(t *testing.T) {
	r := TXTReporter{}
	expected := ""

	resp, _ := r.Report([]swarm.HTTPResponse{})
	assert.Equal(t, expected, resp)
}

func Test_TXTReporter_Report_WhenResponsesExists(t *testing.T) {
	duration, _ := time.ParseDuration("1ms")
	r := TXTReporter{}
	expected := "200 1ms /users\n"

	resp, _ := r.Report([]swarm.HTTPResponse{{StatusCode: 200, Time: duration, Path: "/users"}})
	assert.Equal(t, expected, resp)
}

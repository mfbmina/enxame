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

	assert.Equal(t, expected, r.Report(nil))
}

func Test_TXTReporter_Report_WhenResponsesIsEmpty(t *testing.T) {
	r := TXTReporter{}
	expected := ""

	assert.Equal(t, expected, r.Report([]swarm.HTTPResponse{}))
}

func Test_TXTReporter_Report_WhenResponsesExists(t *testing.T) {
	duration, _ := time.ParseDuration("1ms")
	r := TXTReporter{}
	expected := "200 1ms /users\n"

	assert.Equal(t, expected, r.Report([]swarm.HTTPResponse{{StatusCode: 200, Time: duration, Path: "/users"}}))
}

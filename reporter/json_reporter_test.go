package reporter

import (
	"testing"
	"time"

	"github.com/mfbmina/enxame/swarm"
	"github.com/stretchr/testify/assert"
)

func Test_JSONReporter_WhenResponsesAreNil(t *testing.T) {
	r := JSONReporter{}
	expected := "null"

	resp, _ := r.Report(nil)
	assert.Equal(t, expected, resp)
}

func Test_JSONReporter_WhenResponsesAreEmpty(t *testing.T) {
	r := JSONReporter{}
	expected := "[]"

	resp, _ := r.Report([]swarm.HTTPResponse{})
	assert.Equal(t, expected, resp)
}

func Test_JSONReporter_WhenResponsesExists(t *testing.T) {
	duration, _ := time.ParseDuration("1ms")
	r := JSONReporter{}
	expected := "[{\"status_code\":200,\"time\":1000000,\"path\":\"/users\"}]"

	resp, _ := r.Report([]swarm.HTTPResponse{{StatusCode: 200, Time: duration, Path: "/users"}})
	assert.Equal(t, expected, resp)
}

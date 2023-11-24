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

	assert.Equal(t, expected, r.Report())
}

func Test_JSONReporter_WhenResponsesAreEmpty(t *testing.T) {
	r := JSONReporter{Responses: []swarm.HTTPResponse{}}
	expected := "[]"

	assert.Equal(t, expected, r.Report())
}

func Test_JSONReporter_WhenResponsesExists(t *testing.T) {
	duration, _ := time.ParseDuration("1ms")
	r := JSONReporter{Responses: []swarm.HTTPResponse{{StatusCode: 200, Time: duration, Path: "/users"}}}
	expected := "[{\"status_code\":200,\"time\":1000000,\"path\":\"/users\"}]"

	assert.Equal(t, expected, r.Report())
}

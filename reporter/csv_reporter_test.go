package reporter

import (
	"testing"
	"time"

	"github.com/mfbmina/enxame/swarm"
	"github.com/stretchr/testify/assert"
)

func Test_CSVReporter_WhenResponsesAreNil(t *testing.T) {
	r := CSVReporter{}
	expected := "status_code,time,path"

	assert.Equal(t, expected, r.Report(nil))
}

func Test_CSVReporter_WhenResponsesAreEmpty(t *testing.T) {
	r := CSVReporter{}
	expected := "status_code,time,path"

	assert.Equal(t, expected, r.Report([]swarm.HTTPResponse{}))
}

func Test_CSVReporter_WhenResponsesExists(t *testing.T) {
	duration, _ := time.ParseDuration("1ms")
	r := CSVReporter{}
	expected := "status_code,time,path\n200,1ms,/users"

	assert.Equal(t, expected, r.Report([]swarm.HTTPResponse{{StatusCode: 200, Time: duration, Path: "/users"}}))
}

package reporter

import (
	"testing"

	"github.com/mfbmina/enxame/swarm"
	"github.com/stretchr/testify/assert"
)

func TestNewReportWhenTypeIsJSON(t *testing.T) {
	r := NewReporter("json", []swarm.HTTPResponse{})

	assert.Equal(t, JSONReport{[]swarm.HTTPResponse{}}, r)
}

func TestNewReportWhenTypeIsCSV(t *testing.T) {
	r := NewReporter("csv", []swarm.HTTPResponse{})

	assert.Equal(t, CSVReport{[]swarm.HTTPResponse{}}, r)
}

func TestNewReportWhenTypeIsTXT(t *testing.T) {
	r := NewReporter("txt", []swarm.HTTPResponse{})

	assert.Equal(t, TXTReport{[]swarm.HTTPResponse{}}, r)
}

func TestNewReportWhenTypeIsRandom(t *testing.T) {
	r := NewReporter("random", []swarm.HTTPResponse{})

	assert.Equal(t, TXTReport{[]swarm.HTTPResponse{}}, r)
}

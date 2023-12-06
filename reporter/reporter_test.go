package reporter

import (
	"os"
	"testing"

	"github.com/mfbmina/enxame/swarm"
	"github.com/stretchr/testify/assert"
)

func Test_NewReport_WhenTypeIsJSON(t *testing.T) {
	r := NewReporter("json", "", []swarm.HTTPResponse{})

	assert.Equal(t, JSONReporter{}, r.Reporter)
}

func Test_NewReport_WhenTypeIsCSV(t *testing.T) {
	r := NewReporter("csv", "", []swarm.HTTPResponse{})

	assert.Equal(t, CSVReporter{}, r.Reporter)
}

func Test_NewReport_WhenTypeIsTXT(t *testing.T) {
	r := NewReporter("txt", "", []swarm.HTTPResponse{})

	assert.Equal(t, TXTReporter{}, r.Reporter)
}

func Test_NewReport_WhenTypeIsRandom(t *testing.T) {
	r := NewReporter("random", "", []swarm.HTTPResponse{})

	assert.Equal(t, TXTReporter{}, r.Reporter)
}

func Test_Report_WhenOutputIsEmpty(t *testing.T) {
	r := NewReporter("txt", "", []swarm.HTTPResponse{})
	f, e := r.Report()

	assert.Equal(t, "", f)
	assert.Nil(t, e)
}

func Test_Report_WhenOutputIsNotEmpty(t *testing.T) {
	defer os.Remove("output.csv")

	r := NewReporter("csv", "output", []swarm.HTTPResponse{})
	f, e := r.Report()
	assert.Equal(t, "Report created at output.csv", f)
	assert.Nil(t, e)

	content, err := os.ReadFile("output.csv")
	assert.Nil(t, err)
	assert.Equal(t, "status_code,time,path", string(content))

	defer os.Remove("output.csv")
}

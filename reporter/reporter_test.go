package reporter

import (
	"io"
	"os"
	"testing"

	"github.com/mfbmina/enxame/swarm"
	"github.com/stretchr/testify/assert"
)

func Test_NewReport_WhenTypeIsJSON(t *testing.T) {
	r := NewReporter("json", "", []swarm.HTTPResponse{})

	assert.Equal(t, JSONReporter{[]swarm.HTTPResponse{}}, r.Reporter)
}

func Test_NewReport_WhenTypeIsCSV(t *testing.T) {
	r := NewReporter("csv", "", []swarm.HTTPResponse{})

	assert.Equal(t, CSVReporter{[]swarm.HTTPResponse{}}, r.Reporter)
}

func Test_NewReport_WhenTypeIsTXT(t *testing.T) {
	r := NewReporter("txt", "", []swarm.HTTPResponse{})

	assert.Equal(t, TXTReporter{[]swarm.HTTPResponse{}}, r.Reporter)
}

func Test_NewReport_WhenTypeIsRandom(t *testing.T) {
	r := NewReporter("random", "", []swarm.HTTPResponse{})

	assert.Equal(t, TXTReporter{[]swarm.HTTPResponse{}}, r.Reporter)
}

func Test_Report_WhenOutputIsEmpty(t *testing.T) {
	r := NewReporter("txt", "", []swarm.HTTPResponse{})

	o := captureOutput(func() { r.Report() })

	assert.Equal(t, "Reporting results as txt...\n------------------------------\n\n", o)
}

func captureOutput(f func()) string {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out)
}

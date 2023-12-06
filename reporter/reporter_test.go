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

	o := captureOutput(func() { r.Report() })

	assert.Equal(t, "\n", o)
}

func Test_Report_WhenOutputIsNotEmpty(t *testing.T) {
	defer os.Remove("output.csv")

	r := NewReporter("csv", "output", []swarm.HTTPResponse{})
	r.Report()

	content, err := os.ReadFile("output.csv")
	assert.Nil(t, err)
	assert.Equal(t, "status_code,time,path", string(content))

	defer os.Remove("output.csv")
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

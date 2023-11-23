package swarm

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
)

func TestSwarmWhenApiCallisSuccessful(t *testing.T) {
	url := "http://example.com/users"

	mock := apitest.NewMock().
		Get(url).
		RespondWith().
		Status(http.StatusOK)

	defer mock.EndStandalone()()

	r := Swarm(url, 1, 1)
	assert.Equal(t, 1, len(r))
	assert.Equal(t, 200, r[0].StatusCode)
	assert.NotNil(t, r[0].Time)
}

func TestSwarmWhenApiCallReturnsError(t *testing.T) {
	url := "strange_url"

	r := Swarm(url, 1, 1)
	assert.Equal(t, 1, len(r))
	assert.Equal(t, 0, r[0].StatusCode)
	assert.NotNil(t, r[0].Time)
}

func TestSwarmReturnsTotalRequests(t *testing.T) {
	url := "strange_url"

	r := Swarm(url, 2, 5)
	assert.Equal(t, 10, len(r))
}

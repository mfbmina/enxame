package swarm

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
)

func TestSwarmWhenApiCallWithGet(t *testing.T) {
	url := "http://example.com/users"

	mock := apitest.NewMock().
		Get(url).
		RespondWith().
		Status(http.StatusOK)

	defer mock.EndStandalone()()

	r, err := Swarm(url, "GET", 1, 1)
	assert.Equal(t, 1, len(r))
	assert.Equal(t, 200, r[0].StatusCode)
	assert.NotNil(t, r[0].Time)
	assert.Nil(t, err)
}

func TestSwarmWhenApiCallWithPost(t *testing.T) {
	url := "http://example.com/users"

	mock := apitest.NewMock().
		Post(url).
		RespondWith().
		Status(http.StatusCreated)

	defer mock.EndStandalone()()

	r, err := Swarm(url, "POST", 1, 1)
	assert.Equal(t, 1, len(r))
	assert.Equal(t, 201, r[0].StatusCode)
	assert.NotNil(t, r[0].Time)
	assert.Nil(t, err)
}

func TestSwarmWhenApiCallReturnsError(t *testing.T) {
	url := "strange_url"

	r, err := Swarm(url, "GET", 1, 1)
	assert.Equal(t, 1, len(r))
	assert.Equal(t, 0, r[0].StatusCode)
	assert.NotNil(t, r[0].Time)
	assert.Nil(t, err)
}

func TestSwarmReturnsTotalRequests(t *testing.T) {
	url := "strange_url"

	r, err := Swarm(url, "GET", 2, 5)
	assert.Equal(t, 10, len(r))
	assert.Nil(t, err)
}

package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	var srv Server
	var err error

	err = srv.Init()
	assert.Nil(t, err, "Could not init")
}

func TestQuit(t *testing.T) {
	var srv Server
	srv.Quit()
}

func TestInitQuit(t *testing.T) {
	var srv Server
	srv.Init()
	srv.Quit()
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOne(t *testing.T) {
	var srv Server
	var err error

	err = srv.Init()
	assert.Nil(t, err, "Could not init")

	srv.Quit()
}

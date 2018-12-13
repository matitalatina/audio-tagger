package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAudioExtension(t *testing.T) {
	assert.False(t, HasAudioExtension("ciao.txt"))
	assert.True(t, HasAudioExtension("ciao.m4a"))
	assert.True(t, HasAudioExtension("ciao.mp3"))
}

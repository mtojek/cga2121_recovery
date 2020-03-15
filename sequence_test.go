package main

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestNext_LowerCase(t *testing.T) {
	seq, err := newSequence(3, "a")
	assert.NoError(t, err)
	first := seq.next()
	second := seq.next()

	var next string
	for i := 0; i < 123; i++ {
		next = seq.next()
	}

	assert.Equal(t, "aaa", first)
	assert.Equal(t, "baa", second)
	assert.Equal(t, "uea", next)
}

func TestNext_LowerUpperCase(t *testing.T) {
	seq, err := newSequence(3, "aA")
	assert.NoError(t, err)
	first := seq.next()
	second := seq.next()

	var next string
	for i := 0; i < 1234; i++ {
		next = seq.next()
	}

	assert.Equal(t, "aaa", first)
	assert.Equal(t, "baa", second)
	assert.Equal(t, "Nxa", next)
}

func TestNext_LowerUpperDigitsCase(t *testing.T) {
	seq, err := newSequence(3, "aA1")
	assert.NoError(t, err)
	first := seq.next()
	second := seq.next()

	var next string
	for i := 0; i < 1234; i++ {
		next = seq.next()
	}

	assert.Equal(t, "aaa", first)
	assert.Equal(t, "baa", second)
	assert.Equal(t, "5ta", next)
}

func TestNext_LowerUpperDigitsSpecialCase(t *testing.T) {
	seq, err := newSequence(3, "aA1_")
	assert.NoError(t, err)
	first := seq.next()
	second := seq.next()

	var next string
	for i := 0; i < 25340; i++ {
		next = seq.next()
	}

	assert.Equal(t, "aaa", first)
	assert.Equal(t, "baa", second)
	assert.Equal(t, "){c", next)
}

func TestNext_LowerCaseSingleCharacter(t *testing.T) {
	seq, err := newSequence(1, "a")
	assert.NoError(t, err)
	first := seq.next()
	second := seq.next()

	var next string
	for i := 0; i < 25340; i++ {
		next = seq.next()
	}

	assert.Equal(t, "a", first)
	assert.Equal(t, "b", second)
	assert.Equal(t, "rmlb", next)
}

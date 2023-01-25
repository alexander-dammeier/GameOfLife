package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_initField(t *testing.T) {
	actual := initField(0, 0)
	expected := make([][]placement, 0)
	assert.Equalf(t, actual, expected, "fields need to be equal")

	actual = initField(2, 1)
	expected = make([][]placement, 1)
	expected[0] = make([]placement, 2)
	expected[0][0] = dead
	expected[0][1] = dead
	assert.Equalf(t, actual, expected, "fields need to be equal")
}

func Test_gameOfLife_getSize(t *testing.T) {
	game := gameOfLife{field: initField(0, 0)}
	width, height := game.getSize()
	assert.Equal(t, width, 0)
	assert.Equal(t, height, 0)

	game = gameOfLife{field: initField(2, 4)}
	width, height = game.getSize()
	assert.Equal(t, width, 2)
	assert.Equal(t, height, 4)
}

func Test_gameOfLife_getOrDead(t *testing.T) {
	game := gameOfLife{field: initField(0, 0)}
	assert.Equal(t, game.getOrDead(0, 0), dead)
	assert.Equal(t, game.getOrDead(1, 0), dead)
	assert.Equal(t, game.getOrDead(0, 1), dead)

	game = gameOfLife{field: initField(1, 1)}
	game.field[0][0] = alive
	assert.Equal(t, game.getOrDead(0, 0), alive)
	assert.Equal(t, game.getOrDead(1, 0), dead)
	assert.Equal(t, game.getOrDead(0, 1), dead)
}

func Test_gameOfLife_advance(t *testing.T) {
	game := gameOfLife{field: [][]placement{
		{dead, dead, dead},
		{dead, dead, dead},
		{dead, dead, dead},
	}}
	expected := [][]placement{
		{dead, dead, dead},
		{dead, dead, dead},
		{dead, dead, dead},
	}
	result := game.advance()
	assert.Equal(t, expected, result.field, "dead cells need to stay dead")
	assert.Equal(t, 1, result.iteration, "should be in iteration 1")

	game = gameOfLife{field: [][]placement{
		{alive, alive, dead},
		{alive, dead, dead},
		{dead, dead, dead},
	}}
	expected = [][]placement{
		{alive, alive, dead},
		{alive, alive, dead},
		{dead, dead, dead},
	}
	assert.Equal(t, expected, game.advance().field, "cell should be born")

	game = gameOfLife{field: [][]placement{
		{dead, alive, dead},
		{alive, dead, dead},
		{dead, dead, dead},
	}}
	expected = [][]placement{
		{dead, dead, dead},
		{dead, dead, dead},
		{dead, dead, dead},
	}
	assert.Equal(t, expected, game.advance().field, "all cells should die because of too few neighbours")

	game = gameOfLife{field: [][]placement{
		{alive, alive, alive},
		{alive, alive, alive},
		{alive, alive, alive},
	}}
	expected = [][]placement{
		{alive, dead, alive},
		{dead, dead, dead},
		{alive, dead, alive},
	}
	assert.Equal(t, expected, game.advance().field, "cells should die because of overpopulation")
}

func Test_generate(t *testing.T) {
	game := generate(10, 15)
	assert.Equal(t, 0, game.iteration)
	assert.Equal(t, 10, len(game.field[0]))
	assert.Equal(t, 15, len(game.field))
}

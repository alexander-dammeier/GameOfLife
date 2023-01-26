package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_initField(t *testing.T) {
	actual := initField(0, 0, deadCellGenerator)
	expected := make([][]cell, 0)
	assert.Equalf(t, actual, expected, "fields need to be equal")

	actual = initField(2, 1, deadCellGenerator)
	expected = [][]cell{
		{dead, dead},
	}
	assert.Equalf(t, actual, expected, "fields need to be equal")

	actual = initField(2, 1, func(x, y int) cell {
		return alive
	})
	expected = [][]cell{
		{alive, alive},
	}
	assert.Equalf(t, actual, expected, "fields need to be equal")
}

func Test_gameOfLife_getSize(t *testing.T) {
	game := gameOfLife{field: initField(0, 0, deadCellGenerator)}
	width, height := game.getSize()
	assert.Equal(t, width, 0)
	assert.Equal(t, height, 0)

	game = gameOfLife{field: initField(2, 4, deadCellGenerator)}
	width, height = game.getSize()
	assert.Equal(t, width, 2)
	assert.Equal(t, height, 4)
}

func Test_gameOfLife_getOrDead(t *testing.T) {
	game := gameOfLife{field: initField(0, 0, deadCellGenerator)}
	assert.Equal(t, game.getOrDead(0, 0), dead)
	assert.Equal(t, game.getOrDead(1, 0), dead)
	assert.Equal(t, game.getOrDead(0, 1), dead)

	game = gameOfLife{field: initField(1, 1, deadCellGenerator)}
	game.field[0][0] = alive
	assert.Equal(t, game.getOrDead(0, 0), alive)
	assert.Equal(t, game.getOrDead(1, 0), dead)
	assert.Equal(t, game.getOrDead(0, 1), dead)
}

func Test_gameOfLife_advance(t *testing.T) {
	game := gameOfLife{field: [][]cell{
		{dead, dead, dead},
		{dead, dead, dead},
		{dead, dead, dead},
	}}
	expected := [][]cell{
		{dead, dead, dead},
		{dead, dead, dead},
		{dead, dead, dead},
	}
	result := game.advance()
	assert.Equal(t, expected, result.field, "dead cells need to stay dead")
	assert.Equal(t, 1, result.iteration, "should be in iteration 1")

	game = gameOfLife{field: [][]cell{
		{alive, alive, dead},
		{alive, dead, dead},
		{dead, dead, dead},
	}}
	expected = [][]cell{
		{alive, alive, dead},
		{alive, alive, dead},
		{dead, dead, dead},
	}
	assert.Equal(t, expected, game.advance().field, "cell should be born")

	game = gameOfLife{field: [][]cell{
		{dead, alive, dead},
		{alive, dead, dead},
		{dead, dead, dead},
	}}
	expected = [][]cell{
		{dead, dead, dead},
		{dead, dead, dead},
		{dead, dead, dead},
	}
	assert.Equal(t, expected, game.advance().field, "all cells should die because of too few neighbours")

	game = gameOfLife{field: [][]cell{
		{alive, alive, alive},
		{alive, alive, alive},
		{alive, alive, alive},
	}}
	expected = [][]cell{
		{alive, dead, alive},
		{dead, dead, dead},
		{alive, dead, alive},
	}
	assert.Equal(t, expected, game.advance().field, "cells should die because of overpopulation")
}

func Test_generate(t *testing.T) {
	game := generateNewGame(10, 15, randomCellGenerator)
	assert.Equal(t, 0, game.iteration)
	assert.Equal(t, 10, len(game.field[0]))
	assert.Equal(t, 15, len(game.field))
}

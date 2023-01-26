package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generateNewGame(t *testing.T) {
	size := fieldSize{width: 0, height: 0}
	actual := generateNewGame(size, deadCellGenerator)
	expected := gameOfLife{
		field:     make([][]cell, 0),
		fieldSize: size,
		iteration: 0,
	}
	assert.Equalf(t, expected, *actual, "fields need to be equal")

	size = fieldSize{width: 2, height: 1}
	actual = generateNewGame(size, deadCellGenerator)
	expected = gameOfLife{
		field: [][]cell{
			{dead, dead},
		},
		fieldSize: size,
		iteration: 0,
	}
	assert.Equalf(t, expected, *actual, "fields need to be equal")

	size = fieldSize{width: 2, height: 1}
	actual = generateNewGame(size, func(pos position) cell {
		return alive
	})
	expected = gameOfLife{
		field: [][]cell{
			{alive, alive},
		},
		fieldSize: size,
		iteration: 0,
	}
	assert.Equalf(t, expected, *actual, "fields need to be equal")
}

func Test_gameOfLife_getOrDead(t *testing.T) {
	game := generateNewGame(fieldSize{width: 0, height: 0}, deadCellGenerator)
	assert.Equal(t, dead, game.getOrDead(position{0, 0}))
	assert.Equal(t, dead, game.getOrDead(position{1, 0}))
	assert.Equal(t, dead, game.getOrDead(position{0, 1}))

	game = generateNewGame(fieldSize{width: 1, height: 1}, deadCellGenerator)
	game.field[0][0] = alive
	assert.Equal(t, alive, game.getOrDead(position{0, 0}))
	assert.Equal(t, dead, game.getOrDead(position{1, 0}))
	assert.Equal(t, dead, game.getOrDead(position{0, 1}))
}

func Test_gameOfLife_advance(t *testing.T) {
	game := gameOfLife{field: [][]cell{
		{dead, dead, dead},
		{dead, dead, dead},
		{dead, dead, dead},
	}, fieldSize: fieldSize{3, 3}}
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
	}, fieldSize: fieldSize{3, 3}}
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
	}, fieldSize: fieldSize{3, 3}}
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
	}, fieldSize: fieldSize{3, 3}}
	expected = [][]cell{
		{alive, dead, alive},
		{dead, dead, dead},
		{alive, dead, alive},
	}
	assert.Equal(t, expected, game.advance().field, "cells should die because of overpopulation")
}

func Test_generate(t *testing.T) {
	game := generateNewGame(fieldSize{width: 10, height: 15}, randomCellGenerator)
	assert.Equal(t, 0, game.iteration)
	assert.Equal(t, 10, len(game.field[0]))
	assert.Equal(t, 15, len(game.field))
}

func Test_position_isOutOfBounds(t *testing.T) {
	size := fieldSize{1, 1}
	t.Run("pos is in field", func(t *testing.T) {
		assert.False(t, position{0, 0}.isOutOfBounds(size))
	})
	t.Run("pos too wide left", func(t *testing.T) {
		assert.True(t, position{-1, 0}.isOutOfBounds(size))
	})
	t.Run("pos too wide right", func(t *testing.T) {
		assert.True(t, position{1, 0}.isOutOfBounds(size))
	})
	t.Run("pos too wide up", func(t *testing.T) {
		assert.True(t, position{0, -1}.isOutOfBounds(size))
	})
	t.Run("pos too wide down", func(t *testing.T) {
		assert.True(t, position{0, 1}.isOutOfBounds(size))
	})
}

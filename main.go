package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type gameOfLife struct {
	field     [][]cell
	iteration int
}

type cell string

const (
	alive cell = "x"
	dead  cell = "_"
)

type cellGenerator = func(x, y int) cell

var randomCellGenerator = func(x, y int) cell {
	if rand.Intn(2) == 1 {
		return alive
	}
	return dead
}

var deadCellGenerator = func(x, y int) cell {
	return dead
}

func main() {
	game := generateNewGame(120, 17, randomCellGenerator)
	untilIteration := 10
	for game.iteration <= untilIteration {
		printGame(game)
		game = game.advance()
	}
}

func printGame(game *gameOfLife) {
	fmt.Printf("Iteration: %d \n", game.iteration)
	_, height := game.getSize()
	for x := 0; x < height; x++ {
		fmt.Println(game.field[x])
		/*for y := 0; y < height; y++ {
		}*/
	}
}

func generateNewGame(width, height int, generator cellGenerator) *gameOfLife {
	field := initField(width, height, generator)
	return &gameOfLife{field: field, iteration: 0}
}

func initField(width, height int, generator cellGenerator) [][]cell {
	field := make([][]cell, height)
	for x := 0; x < height; x++ {
		field[x] = make([]cell, width)
		for y := 0; y < width; y++ {
			field[x][y] = generator(x, y)
		}
	}
	return field
}

func (game *gameOfLife) advance() *gameOfLife {
	width, height := game.getSize()
	advancedField := initField(width, height, deadCellGenerator)
	wg := new(sync.WaitGroup)
	wg.Add(height)
	for x := 0; x < height; x++ {
		go func(x int) {
			defer wg.Done()
			for y := 0; y < cap(advancedField[x]); y++ {
				advancedField[x][y] = game.getAdvancedState(x, y)
			}
		}(x)
	}
	wg.Wait()
	return &gameOfLife{iteration: game.iteration + 1, field: advancedField}
}

func (game *gameOfLife) getAdvancedState(x, y int) cell {
	isAlive := game.field[x][y] == alive
	aliveNeighbours := game.countAliveNeighbors(x, y)
	if isAlive {
		if aliveNeighbours < 2 {
			return dead
		} else if aliveNeighbours > 3 {
			return dead
		}
		return alive
	} else {
		if aliveNeighbours == 3 {
			return alive
		}
		return dead
	}
}

func (game *gameOfLife) countAliveNeighbors(x int, y int) int {
	aliveCounter := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if game.getOrDead(x+i, y+j) == alive {
				aliveCounter++
			}
		}
	}
	return aliveCounter
}

func (game *gameOfLife) getSize() (width int, height int) {
	//TODO: len oder cap?
	height = len(game.field)
	if height > 0 {
		width = len(game.field[0])
	}
	return
}

func (game *gameOfLife) getOrDead(x int, y int) cell {
	width, height := game.getSize()
	if x >= 0 && x < height && y >= 0 && y < width {
		return game.field[x][y]
	}
	return dead
}

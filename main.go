package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	cellSize	= 20
	cols		= 32
	rows		= 32
)

var (
	Up		= Point{0, -1}
	Down	= Point{0, 1}
	Left	= Point{-1, 0}
	Right	= Point{1, 0}
)

type Point struct {
	X, Y int
}

type Game struct {
	snake []Point
	food Point
	dir Point
}

func (p Point) Add(o Point) Point {
	return Point{p.X + o.X, p.Y + o.Y}
}

func (g *Game) Update() error {
	newHead := g.snake[0].Add(g.dir)
	g.snake[0] = newHead
	return nil

}

func NewGame() *Game {
	return &Game{
		snake: []Point{{5,9}, {4,9}, {3,9}},
		food: Point{20, 9},
		dir: Right,
	}
}

func drawCell(screen *ebiten.Image, p Point, clr color.Color) {
	vector.FillRect(screen,
		float32(p.X*cellSize), float32(p.Y*cellSize),
		cellSize, cellSize, clr, false)
}

func drawGrid(screen *ebiten.Image) {
	lineClr := color.RGBA{40, 40, 40, 255}
	for x := range cols + 1 {
		px := float32(x * cellSize)
		vector.StrokeLine(screen, px, 0, px, 640, 1, lineClr, false)
	}
	for y := range cols + 1 {
		py := float32(y * cellSize)
		vector.StrokeLine(screen, 0, py, 640, py, 1, lineClr, false)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawGrid(screen)
	drawCell(screen, g.food, colornames.Red)
	for _, p := range g.snake {
		drawCell(screen, p, colornames.Limegreen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func main() {
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetTPS(10)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

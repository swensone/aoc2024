package pathfinder

import (
	"bufio"
	"fmt"
)

type direction int

const (
	North direction = iota
	East
	South
	West
)

type Pathfinder struct {
	Debug      bool
	Direction  direction
	PositionX  int
	PositionY  int
	Visited    map[string]bool
	Obstacle   map[string]bool
	VisitCount int
	MapWidth   int
	MapHeight  int
}

func New(scanner *bufio.Scanner, debug bool) *Pathfinder {
	p := &Pathfinder{
		Debug:     debug,
		Direction: North,
		Visited:   map[string]bool{},
		Obstacle:  map[string]bool{},
	}

	line := 0
	for scanner.Scan() {
		text := scanner.Text()
		p.MapWidth = len(text)
		for column, c := range text {
			switch c {
			case '^':
				p.PositionX = column
				p.PositionY = line
			case '#':
				p.Obstacle[fmt.Sprintf("%d,%d", column, line)] = true
			}
		}
		line++
	}
	p.MapHeight = line

	if p.Debug {
		fmt.Printf("%+v\n", p)
	}
	return p
}

func (s *Pathfinder) FindPath() int {
	step := 0
	s.Visit()

	continueLoop := true
	for continueLoop {
		if s.Debug {
			fmt.Printf("step: %d, x: %d, y: %d\n", step, s.PositionX, s.PositionY)
			s.PrintMap()
			fmt.Println()
		}
		step++
		continueLoop = s.NextStep()
	}

	return s.VisitCount
}

func (s *Pathfinder) NextStep() bool {
	for {
		x, y := s.Step()
		if s.Debug {
			fmt.Printf("checking for obstacle at %d,%d\n", x, y)
		}
		if s.Obstacle[fmt.Sprintf("%d,%d", x, y)] {
			s.TurnRight()
		} else {
			s.PositionX = x
			s.PositionY = y
			break
		}
	}

	if s.PositionX < 0 || s.PositionX >= s.MapWidth || s.PositionY < 0 || s.PositionY >= s.MapHeight {
		return false
	}

	s.Visit()

	return true
}

func (s *Pathfinder) Visit() {
	if !s.Visited[fmt.Sprintf("%d,%d", s.PositionX, s.PositionY)] {
		if s.Debug {
			fmt.Printf("VISIT: %d,%d\n", s.PositionX, s.PositionY)
		}
		s.Visited[fmt.Sprintf("%d,%d", s.PositionX, s.PositionY)] = true
		s.VisitCount++
	}
}

func (s *Pathfinder) PrintMap() {
	for y := range s.MapHeight {
		for x := range s.MapWidth {
			if s.PositionX == x && s.PositionY == y {
				switch s.Direction {
				case North:
					fmt.Print("^")
				case East:
					fmt.Print(">")
				case South:
					fmt.Print("v")
				case West:
					fmt.Print("<")
				}
			} else if s.Obstacle[fmt.Sprintf("%d,%d", x, y)] {
				fmt.Print("#")
			} else if s.Visited[fmt.Sprintf("%d,%d", x, y)] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (s *Pathfinder) Step() (int, int) {
	switch s.Direction {
	case North:
		return s.PositionX, s.PositionY - 1
	case East:
		return s.PositionX + 1, s.PositionY
	case South:
		return s.PositionX, s.PositionY + 1
	case West:
		return s.PositionX - 1, s.PositionY
	default:
		panic("invalid direction")
	}
}

func (s *Pathfinder) TurnRight() {
	s.Direction = (s.Direction + 1) % 4
}

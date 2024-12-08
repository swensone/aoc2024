package pathfinder

import (
	"bufio"
	"fmt"
	"slices"
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
	Block      string
	Visited    map[string]bool
	Obstacle   map[string]bool
	Path       []string
	VisitCount int
	MapWidth   int
	MapHeight  int
}

func New(scanner *bufio.Scanner, block string, debug bool) *Pathfinder {
	p := &Pathfinder{
		Debug:     debug,
		Direction: North,
		Block:     block,
		Visited:   map[string]bool{},
		Obstacle:  map[string]bool{},
		Path:      []string{},
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

func (s *Pathfinder) FindPath() ([]string, bool) {
	step := 0
	s.Visit()

	continueLoop := true
	looped := false
	for continueLoop {
		if s.Debug {
			fmt.Printf("step: %d, x: %d, y: %d\n", step, s.PositionX, s.PositionY)
			s.PrintMap()
			fmt.Println()
		}
		step++
		continueLoop, looped = s.NextStep()
		if looped {
			break
		}
	}

	visited := []string{}
	for k, _ := range s.Visited {
		visited = append(visited, k)
	}

	return visited, looped
}

func (s *Pathfinder) NextStep() (bool, bool) {
	for {
		x, y := s.Step()
		if s.Debug {
			fmt.Printf("checking for obstacle at %d,%d\n", x, y)
		}
		if s.Obstacle[fmt.Sprintf("%d,%d", x, y)] {
			s.TurnRight()
		} else {
			s.Path = append(s.Path, fmt.Sprintf("%d:%d", x, y))
			s.PositionX = x
			s.PositionY = y
			break
		}
	}

	if s.PositionX < 0 || s.PositionX >= s.MapWidth || s.PositionY < 0 || s.PositionY >= s.MapHeight {
		return false, false
	}

	looped := s.Visit()

	return true, looped
}

func (s *Pathfinder) Visit() bool {
	if !s.Visited[fmt.Sprintf("%d,%d", s.PositionX, s.PositionY)] {
		if s.Debug {
			fmt.Printf("VISIT: %d,%d\n", s.PositionX, s.PositionY)
		}
		s.Visited[fmt.Sprintf("%d,%d", s.PositionX, s.PositionY)] = true
		s.VisitCount++
	} else {
		return s.LoopCheck()
	}
	return false
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

func (s *Pathfinder) LoopCheck() bool {
	i := 4
	pathLen := len(s.Path)
	for pathLen >= i*2 {
		if slices.Equal(s.Path[pathLen-i:], s.Path[pathLen-2*i:pathLen-i]) {
			return true
		}
		i += 2
	}
	return false
}

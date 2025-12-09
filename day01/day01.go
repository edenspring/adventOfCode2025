package day01

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SafeDial interface {
	TurnLeft(clicks int)
	TurnRight(clicks int)
	SetPosition(position int)
	GetPosition() int
	SetDirections(directions string)
	ProcessDirections() int
	IngestInput(path string) (string, error)
	GetAnswer() int
}

type SecretSafe struct {
	position   int
	directions string
	answer     int
}

func (s *SecretSafe) IngestInput(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (s *SecretSafe) SetDirections(directions string) {
	s.directions = directions
}

func (s *SecretSafe) TurnLeft(clicks int) {
	newPos := s.GetPosition() - clicks
	newPos = s.checkPos(newPos)
	msg := fmt.Sprintf("The dial is rotated %d clicks left from position %d to %d", clicks, s.GetPosition(), newPos)
	fmt.Println(msg)
	s.SetPosition(newPos)
}

func (s *SecretSafe) TurnLeftPartTwo(clicks int) {
	origClicks := clicks
	origPos := s.GetPosition()
	newPos := s.GetPosition()
	for clicks > 0 {
		newPos--
		clicks--
		if newPos == 0 {
			fmt.Println("hit 0 going left")
			s.answer++
		}
		if newPos < 0 {
			newPos = newPos + 100
		}
	}
	msg := fmt.Sprintf("The dial is rotated %d clicks left from position %d to %d", origClicks, origPos, newPos)
	fmt.Println(msg)
	s.SetPosition(newPos)
}

func (s *SecretSafe) TurnRight(clicks int) {
	newPos := s.GetPosition() + clicks
	newPos = s.checkPos(newPos)
	msg := fmt.Sprintf("The dial is rotated %d clicks right from position %d to %d", clicks, s.GetPosition(), newPos)
	fmt.Println(msg)
	s.SetPosition(newPos)
}

func (s *SecretSafe) TurnRightPartTwo(clicks int) {
	origClicks := clicks
	origPos := s.GetPosition()
	newPos := s.GetPosition()
	for clicks > 0 {
		newPos++
		clicks--
		if newPos > 99 {
			fmt.Println("Passed 99, hit 0")
			newPos = 0
			s.answer++
		}
	}

	msg := fmt.Sprintf("The dial is rotated %d clicks right from position %d to %d", origClicks, origPos, newPos)
	fmt.Println(msg)
	s.SetPosition(newPos)
}

func (s *SecretSafe) checkPos(position int) int {
	for position > 99 {
		position = position - 100
	}
	for position < 0 {
		position = position + 100
	}
	if position == 0 {
		s.answer = s.answer + 1
	}
	return position
}

func (s *SecretSafe) GetPosition() int {
	return s.position
}

func (s *SecretSafe) SetPosition(position int) {
	s.position = position
}

func (s *SecretSafe) GetAnswer() int {
	steps := strings.Split(s.directions, "\n")
	for _, step := range steps {
		leftOrRight := strings.ToLower(step[:1])
		clicks, clicksErr := strconv.Atoi(step[1:])
		if clicksErr != nil {
			fmt.Println("ERROR CONVERTING CLICKS TO INT")
			continue
		}
		switch leftOrRight {
		case "l":
			s.TurnLeft(clicks)
		case "r":
			s.TurnRight(clicks)
		}
	}
	return s.answer
}

func (s *SecretSafe) GetAnswerPartTwo() int {
	steps := strings.Split(s.directions, "\n")
	for _, step := range steps {
		leftOrRight := strings.ToLower(step[:1])
		clicks, clicksErr := strconv.Atoi(step[1:])
		if clicksErr != nil {
			fmt.Println("ERROR CONVERTING CLICKS TO INT")
			continue
		}
		switch leftOrRight {
		case "l":
			s.TurnLeftPartTwo(clicks)
		case "r":
			s.TurnRightPartTwo(clicks)
		}
	}
	return s.answer
}

func NewSecretSafe(directions string) SecretSafe {
	safe := SecretSafe{
		position:   50,
		directions: "",
		answer:     0,
	}
	if directions != "" {
		safe.SetDirections(directions)
	} else {
		directions, err := safe.IngestInput("./day01/input.txt")
		if err != nil {
			fmt.Println("unexpected err: " + err.Error())
		}
		safe.SetDirections(directions)
	}
	return safe
}

func SolveDay01() int {
	safe := NewSecretSafe("")
	return safe.GetAnswer()
}

func SolveDay01PartTwo() int {
	safe := NewSecretSafe("")
	return safe.GetAnswerPartTwo()
}

package main

import (
	"bytes"
	"fmt"
)

var (
	azLowerCasePatterns = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q',
		'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	azUpperCasePatterns = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q',
		'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	digitPatterns   = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	specialPatterns = []byte{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '-', '+', '=', '{', '[', '}', ']',
		':', ';', '"', '\'', '|', '\\', '~', '`', '<', ',', '>', ',', '?', '/', '.'}
)

type sequence struct {
	charPool []byte

	places []int

	firstIteration bool
}

func newSequence(minLength int, patterns string) (*sequence, error) {
	var charPool []byte
	for _, p := range patterns {
		switch p {
		case 'a':
			charPool = append(charPool, azLowerCasePatterns...)
		case 'A':
			charPool = append(charPool, azUpperCasePatterns...)
		case '1':
			charPool = append(charPool, digitPatterns...)
		case '_':
			charPool = append(charPool, specialPatterns...)
		default:
			return nil, fmt.Errorf("unknown pattern: %s", patterns)

		}
	}

	places := make([]int, minLength)
	places[0] = -1

	return &sequence{
		places:         places,
		charPool:       charPool,
		firstIteration: true,
	}, nil
}

func (s *sequence) next() string {
	var overflow bool

	i := 0

	for {
		s.places[i] = (s.places[i] + 1) % len(s.charPool)
		if s.places[i] == 0 {
			overflow = true
		}

		i++

		if s.firstIteration {
			s.firstIteration = false
			break
		}
		if overflow {
			overflow = false
			if i == len(s.places) {
				s.places = append(s.places, 0)
			}
			continue
		}
		break
	}

	var buf bytes.Buffer
	for _, place := range s.places {
		buf.WriteByte(s.charPool[place])
	}
	return buf.String()
}

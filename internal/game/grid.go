package game

import (
	"fmt"
	"image/color"
	"math/rand/v2"
	"strconv"
)

const (
	DEFAULT_NUM_COLUMNS int = 3
	DEFAULT_NUM_ROWS    int = 3
)

type Square struct {
	id     int
	colour color.RGBA
}

func NewSquare(id int, colourHex string) Square {
	var newSquare Square = Square{
		id: id,
	}
	newSquare.SetColourFromHex(colourHex)
	return newSquare
}

// DESCRIPTION
// Setting Square colour from a hex string
func (s *Square) SetColourFromHex(newHex string) error {
	// check that hex meets string requirements
	if len(newHex) != 7 || newHex[0] != '#' {
		return fmt.Errorf("Invalid hex format")
	}

	// convert string hex to uint
	hexVal, hexErr := strconv.ParseUint(newHex, 16, 32) // save hex value in 32-bit UINT
	if hexErr != nil {
		return hexErr
	}

	// replacing current colour
	s.colour = color.RGBA{
		R: uint8(hexVal>>16) & 0xff,
		G: uint8(hexVal>>8) & 0xff,
		B: uint8(hexVal>>0) & 0xff,
		A: 0xff, // alpha 100%
	}

	return nil
}

// DESCRIPTION
// Selects random colour and applies it to the current square
func (s *Square) SetColourToRandom() {
	var randomInt = rand.Int()
	s.colour = color.RGBA{
		R: uint8(randomInt>>16) & 0xff,
		G: uint8(randomInt>>8) & 0xff,
		B: uint8(randomInt>>0) & 0xff,
		A: 0xff,
	}
}

package game

import (
	"fmt"
	"image/color"
	"math/rand/v2"
	"strconv"
)

type Square struct {
	ID     int
	Colour color.RGBA
}

// DESCRIPTION
// Sets a new square colour
func (s *Square) SetColour(newColour color.RGBA) {
	s.Colour = newColour
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
	s.Colour = color.RGBA{
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
	s.Colour = color.RGBA{
		R: uint8(randomInt>>16) & 0xff,
		G: uint8(randomInt>>8) & 0xff,
		B: uint8(randomInt>>0) & 0xff,
		A: 0xff,
	}
}

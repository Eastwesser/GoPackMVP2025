package main

import (
	"errors"
	"fmt"
)

type Mouth struct{}

type Food struct{}

type Energy struct{}

func Digest(f Food) interface{} {
	return f
}

/*
	Receiver (m Mouth) → Your mouth is the "object" that receives something (like a method in Go).
		Receiver → Like a body part doing an action (Mouth.Eat()).

	Parameters (f Food) → The food you put in your mouth (input arguments).
		Parameters → Stuff you consume (input).

	Function Body → All the digestion happens here (processing logic).
		Function Logic → Your organs process it.

	Return Values (Energy, error) → What comes out the other end (useful output or errors).
		Return → Results (or errors if digestion fails).
*/

func (m Mouth) Eat(f Food) (Energy, error) {
	// Throat → Stomach → Intestines (function body does processing)
	energy := Digest(f)

	// "Return" is like... well, you know
	if energy <= 0 {
		return energy, errors.New("indigestion")
	} else if err != nil {
		fmt.Println("Emergency bathroom trip!")
	} else {
		fmt.Println("Good energy! 💪")
	}

	return energy, nil
}

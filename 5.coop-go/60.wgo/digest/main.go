package main

import (
	"errors"
	"fmt"
)

// Chew and swallow=====================================================================================================
func (m Mouth) Chew(f Food) (Pulp, error) {
	if f.IsHard {
		return Pulp{}, errors.New("too crunchy, jaw stuck")
	}

	return Pulp{Texture: "smooth"}, nil
}

/*
	(m Mouth) JAW is a receiver, <- with what you do the action as an instrument
	Chew is a method. It can be Crunch, Eat, Slurp, any.
	(f Food) BURGER is a parameter, <- what we actually put in function to transform (give life - receive gold)
	(Pulp, error) PULP is a return value along with errors, <- the result what we expect, done or not.
	Pulp - the mush you swallow (or error if you bite a rock)
*/

// Process =============================================================================================================
func (g Gut) Process(f Food) (Energy float64, Gas bool, err error) {
	if f.IsSpicy {
		return 0, true, errors.New("🔥🌶️ RIP digestion")
	}
	return 100, false, nil
}

/*
	Now your gut is the receiver (g Gut).
	Returns three things: energy (useful), gas (side effect), error (if disaster strikes).
	Type this while thinking: "What happens after I eat this burrito?"
*/

// EatSushi ============================================================================================================
func (m Mouth) EatSushi(f Food) error {
	defer m.BrushTeeth() // Runs even if you puke later
	if f.IsRaw {
		return errors.New("parasites incoming")
	}
	return nil
}

// SneakSnack ==========================================================================================================
func SneakSnack() {
	eatQuickly := func(f Food) { // Anonymous func, like a sneaky bite
		fmt.Println("Nom nom nom...", f.Name)
	}
	eatQuickly(Food{Name: "donut"})
}

// Anonymous functions = doing something quick without naming it.

// PackLunch ===========================================================================================================
func PackLunc() func() Food {
	sandwich := Food{}
}

package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"runtime"
	"time"
)

/*
	"In the grim darkness of the 41st millennium, there is only concurrency..."

	Warhammer 40k-themed goroutine monitor, refactored as an Ultramarines Battle Report System.

	Because even the Adeptus Mechanicus needs to debug their cogitators.
*/

// BattleReport monitors goroutines like an Ultramarine Tactical Display
func BattleReport(previousForces int) {
	for {
		currentForces := runtime.NumGoroutine()

		// Initialize forces (deploying scouts)
		if previousForces == 0 {
			previousForces = currentForces
			continue
		}

		// Calculate reinforcements/losses (Codex-approved)
		forceRatio := float64(currentForces) / float64(previousForces) * 100

		switch {
		// Reinforcements detected (+++INCOMING+++)
		case forceRatio >= 120:
			reinforcements := int(forceRatio - 100)
			fmt.Printf(
				"📡 [TACTICAL ALERT] Reinforcements incoming! +%d%% Legionaries deployed!\n",
				reinforcements,
			)

		// Heavy losses (Xenos threat detected!)
		case forceRatio <= 80:
			casualties := 100 - int(forceRatio)
			fmt.Printf(
				"💀 [BATTLE LOSSES] %d%% of battle brothers fallen! Requesting backup!\n",
				casualties,
			)

		// All quiet on the Golang front
		default:
			fmt.Println("🛡️ [STATUS] No threats detected. The Codex Astartes supports this concurrency.")
		}

		// Current battlefield status
		fmt.Printf("⚔️ Active Legionaries: %d\n", currentForces)
		previousForces = currentForces

		// Standard Imperial scanning interval (300ms)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	// The Chapter Master (errgroup) leads the battle
	battleGroup, _ := errgroup.WithContext(context.Background())

	// Actiate the Vox-Caster (goroutine monitor)
	go func() {
		BattleReport(runtime.NumGoroutine())
	}()

	// Deploying 64 Battle Brothers (goroutines)
	for i := 0; i < 64; i++ {
		battleGroup.Go(func() error {
			// Simulate combat (5-second mission)
			time.Sleep(5 * time.Second)
			return nil // No Heresy detected
		})
		time.Sleep(80 * time.Millisecond) // Staggered deployment
	}

	// Await battle report
	if err := battleGroup.Wait(); err != nil {
		fmt.Println("☠️ HERESY DETECTED:", err) // Chaos corruption!
	} else {
		fmt.Println("✅ BATTLE CONCLUDED: Victory for the Emperor!")
	}

}

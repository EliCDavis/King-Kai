package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

type KingKai struct {
	controls Controller
	screen   Game
}

func NewKingKai(controls Controller, screen Game) *KingKai {
	return &KingKai{controls, screen}
}

func (king *KingKai) incrementStateCount(stateCount map[int]int, state int) {
	if _, ok := stateCount[state]; !ok {
		stateCount[state] = 0
	}
	stateCount[state] = stateCount[state] + 1
}

func (king *KingKai) stateActionHash(state int, action Input) string {
	return fmt.Sprintf("%d-%s", state, action.Hash())
}

func (king *KingKai) bestPolicy(state int, policy map[string]float64) Input {
	var best Input
	best_value := -1000000.0

	for _, action := range allPosibleActions {
		ha := king.stateActionHash(state, action)

		if val, ok := policy[ha]; ok && val > best_value {
			best = action
			best_value = val
		}

	}

	if best == nil {
		return allPosibleActions[rand.Intn(len(allPosibleActions))]
	}

	return best

}

func (king *KingKai) bestPolicyValue(state int, policy map[string]float64) float64 {
	hash := king.stateActionHash(state, king.bestPolicy(state, policy))

	if _, ok := policy[hash]; ok {
		return policy[hash]
	}
	return 0.0
}

func (king *KingKai) incrementStateActionCount(state int, action Input, stateActionCount map[string]int, policy map[string]float64) {
	hash := king.stateActionHash(state, action)
	if _, ok := stateActionCount[hash]; !ok {
		stateActionCount[hash] = 0
		policy[hash] = 0.0
	}

	stateActionCount[hash] = stateActionCount[hash] + 1
}

func (king *KingKai) sampleActions(state int, stateCount map[int]int, policy map[string]float64) Input {
	n0 := 100.0

	epsilon := 0.0
	if val, ok := stateCount[state]; ok {
		epsilon = n0 / (float64(val) + n0)
	}

	greedy_chance := (epsilon / float64(len(allPosibleActions))) + 1.0 - epsilon

	best_action := king.bestPolicy(state, policy)

	if rand.Float64() < greedy_chance {
		return best_action
	}

	return allPosibleActions[rand.Intn(len(allPosibleActions))]
}

func (king *KingKai) Train(numberOfEpisodes int) error {

	// lastDraw := time.Now()

	discountFactor := 0.9

	stateCount := make(map[int]int)
	stateActionCount := make(map[string]int)
	QPolicy := make(map[string]float64)

	for curEpisode := 0; curEpisode < numberOfEpisodes; curEpisode++ {

		err := king.controls.reset()
		if err != nil {
			return errors.New("Problem reseting episode: " + err.Error())
		}

		lastComboDamage := 0
		king.incrementStateCount(stateCount, lastComboDamage)

		action := king.sampleActions(lastComboDamage, stateCount, QPolicy)
		king.incrementStateActionCount(lastComboDamage, action, stateActionCount, QPolicy)

		saHash := king.stateActionHash(lastComboDamage, action)

		reward := 0

		for reward >= lastComboDamage {
			log.Println(action.Hash())
			err := action.Execute(king.controls)
			if err != nil {
				log.Printf("issues: %s", king.controls.everything())
				return errors.New("Problem executing input: " + err.Error())
			}

			reward, err = king.screen.GetDamage()
			if err != nil {
				return errors.New("Problem getting damage: " + err.Error())
			}

			// inc state count
			king.incrementStateCount(stateCount, reward)

			action_prime := king.sampleActions(reward, stateCount, QPolicy)

			// inc sa count with action prime
			king.incrementStateActionCount(lastComboDamage, action, stateActionCount, QPolicy)

			// get sa hash prime
			saHashPrime := king.stateActionHash(lastComboDamage, action_prime)

			// update Q
			QPolicy[saHash] = QPolicy[saHash] + ((1.0 / float64(stateActionCount[saHash])) * (float64(reward) + (king.bestPolicyValue(reward, QPolicy) * discountFactor) - QPolicy[saHash]))

			action = action_prime
			saHash = saHashPrime

			if reward > lastComboDamage {
				lastComboDamage = reward
			}
		}

		comboDamage, err := king.screen.GetDamage()
		if err != nil {
			log.Printf("Error parsing image: %s\n", err.Error())
		} else if comboDamage != lastComboDamage {
			log.Printf("Combo Damage: %d", comboDamage)
			lastComboDamage = comboDamage
		}
		// now := time.Now()
		// log.Printf("fps: %d", int(time.Second/now.Sub(lastDraw)))
		// lastDraw = now
	}

	return nil
}

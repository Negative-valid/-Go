package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Fighter struct {
	name string

	health int
}

type Drone struct {
	id int

	health int
}

func createFighters() []Fighter {

	fighters := make([]Fighter, 5)

	for i := 0; i < 5; i++ {

		fighters[i].name = "Fighter_" + fmt.Sprint(i+1)

		fighters[i].health = 100

	}

	return fighters

}

func createDrones() []Drone {

	drones := []Drone{}

	for i := 1; i <= 4; i++ {

		d := Drone{id: i, health: 70}

		drones = append(drones, d)

	}

	return drones

}

func dealDamage(base int) int {

	dmg := base + rand.Intn(5)

	return int(dmg)

}

func attackFighter(d Fighter, dmg int) Fighter {

	d.health = d.health - dmg

	if d.health < 0 {
		d.health = 0
	}

	return d

}

func attackDrone(dr Drone, dmg int) Drone {

	if dr.health > 0 {

		dr.health = dr.health - dmg

		if dr.health < 0 {
			dr.health = 0
		}

	}

	return dr

}

func allDronesDestroyed(drones []Drone) bool {

	for _, d := range drones {

		if d.health > 0 {

			return false

		}

	}

	return true

}

func allFightersDead(fighters []Fighter) bool {

	dead := 0

	for i := 0; i < len(fighters); i++ {

		if fighters[i].health <= 0 {

			dead++

		}

	}

	return dead == len(fighters)

}

func main() {

	rand.Seed(time.Now().UnixNano())

	fighters := createFighters()

	drones := createDrones()

	round := 1

	for {

		fmt.Println("=== ROUND", round, "===")

		for i := 0; i < len(fighters); i++ {

			if fighters[i].health <= 0 {

				continue

			}

			target := rand.Intn(len(drones))

			if drones[target].health <= 0 {
				for k := 0; k < len(drones); k++ {
					if drones[k].health > 0 {
						target = k
						break
					}
				}
			}

			dmg := dealDamage(10)

			drones[target] = attackDrone(drones[target], dmg)

			fmt.Println(fighters[i].name, "hits Drone", drones[target].id, "for", dmg)

		}

		for j := 0; j < len(drones); j++ {

			if drones[j].health <= 0 {
				continue
			}

			target := rand.Intn(len(fighters))

			if fighters[target].health <= 0 {
				for k := 0; k < len(fighters); k++ {
					if fighters[k].health > 0 {
						target = k
						break
					}
				}
			}

			dmg := dealDamage(5)

			fighters[target] = attackFighter(fighters[target], dmg)

			fmt.Println("Drone", drones[j].id, "hits", fighters[target].name, "for", dmg)

		}

		if allDronesDestroyed(drones) {

			fmt.Println("Fighters win!")

			break

		}

		if allFightersDead(fighters) {

			fmt.Println("Drones win!")

			break

		}

		round++

		time.Sleep(100 * time.Millisecond)

	}

	fmt.Println("Battle finished in", round, "rounds.")

}

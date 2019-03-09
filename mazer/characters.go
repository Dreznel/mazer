package mazer

import (
	"fmt"
)

type Fighter struct {
	HitPoints int
	Stamina int
	Treasure int
	Name string
}

func NewFighter(name string) (Character) {
	return &Fighter {
		HitPoints: Roll(2, 10),
		Stamina: Roll(10, 10),
		Name: name,
	}
}

func (character *Fighter) Attack() int {
	attackRoll := Roll(2, 10)
	// fmt.Println(fmt.Sprintf("%s attacks! (%d)", character.Name, attackRoll))
	return attackRoll
}

func (character *Fighter) UseSkill() int {
	skillRoll := Roll(1, 6)
	fmt.Println(fmt.Sprintf("%s tries to solve his problems with violence. (%d)"), character.Name, skillRoll)
	return skillRoll
}

func (character *Fighter) UseMagic() int {
	magicRoll := 0
	fmt.Println(fmt.Sprintf("%s scratches his head. What is 'Magic'? (%d)"), character.Name, magicRoll)
	return magicRoll
}

func (character *Fighter) GetTreasure(amount int) {
	character.Treasure += amount
}

func (character *Fighter) IsDead() bool {
	return character.HitPoints <= 0
}

func (character *Fighter) ChangeHp(amount int) {
	character.HitPoints += amount
}

func (character *Fighter) ChangeStamina(amount int) {
	character.Stamina += amount
}

func (character *Fighter) PrintStats() {
	deathIndicator := ""
	if(character.IsDead()) {
		deathIndicator = "(Deceased)"
	}
	fmt.Println(fmt.Sprintf("Name: %s %s\nHP Remaining: %d\nStamina Remaining: %d\nTotal Treasure: %d", deathIndicator, character.Name, character.HitPoints, character.Stamina, character.Treasure))
}

func (character *Fighter) GetName() string {
	return character.Name
}
//package main
//
//import "fmt"
//import "math/rand"
//import "time"
//
//func main() {
//	rand.Seed(time.Now().UnixNano())
//	min := 10
//	max := 30
//	fmt.Println(r.Intn(max - min) + min)
//}
//
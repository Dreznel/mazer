package mazer

import "fmt"

type MonsterRoom struct {
	MonsterHp int
	Strength int
	TreasureRemaining int
	Name string
}

func NewEasyMonsterRoom(name string) Room {
	return &MonsterRoom {
		MonsterHp: Roll(2, 4),
		Strength: Roll(1, 2),
		TreasureRemaining: Roll(3, 10),
		Name: name,
	}
}

func NewDifficultMonsterRoom(name string) Room  {
	return &MonsterRoom {
		MonsterHp: Roll(3, 6),
		Strength: Roll(3, 2), //Strength of 3 to 6
		TreasureRemaining: Roll(3, 10),
		Name: name,
	}
}

func (room *MonsterRoom) DoChallenge(character Character) bool {
	monsterAttack := Roll(room.Strength, 8)
	characterAttack := character.Attack()

	fmt.Println(fmt.Sprintf("Monster attacks! (%d)", monsterAttack))

	if(monsterAttack >= characterAttack) {
		fmt.Println("Character fails his attack!")
		character.ChangeHp(characterAttack - monsterAttack)
		if(character.IsDead()) {
			fmt.Println("Character has been slain!")
		}
		return false
	} else {
		fmt.Println("Character attacks successfully!")
		room.MonsterHp -= characterAttack - monsterAttack
		if room.MonsterHp > 0 {
			return false
		} else {
			fmt.Println("Monster is slain!")
			character.GetTreasure(room.TreasureRemaining)
			return true
		}
	}
}
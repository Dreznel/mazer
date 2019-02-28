package main

import (
	"fmt"
	"github.com/Dreznel/mazer/mazer"
)

func main() {
	var rooms []mazer.Room
	for i := 0; i < 10; i++ {
		rooms = append(rooms, mazer.NewEasyMonsterRoom(fmt.Sprintf("Monster Room %d", i)))
	}
	//for i := 10; i < 20; i++ {
	//	rooms = append(rooms, mazer.NewDifficultMonsterRoom(fmt.Sprintf("Monster Room Difficult %d", i)))
	//}

	c1 := mazer.NewFighter("The Bravest Noob")
	c1.PrintStats()


	for roomNumber, room := range(rooms) {
		roomCleared := false
		for roomUncleared := true; roomUncleared && !c1.IsDead(); roomUncleared = !roomCleared {
			fmt.Println(fmt.Sprintf("Character is attempting room %d", roomNumber))
			roomCleared = room.DoChallenge(c1)
			fmt.Println()
		}
	}

	c1.PrintStats()
}
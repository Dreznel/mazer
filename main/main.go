package main

import (
	"fmt"
	"github.com/Dreznel/mazer/mazer"
	"time"
)

func timeTrack(startTime time.Time, name string) int64 {
	elapsedTime := time.Since(startTime)
	fmt.Println(fmt.Sprintf("%s took %d nanoseconds.\n", name, elapsedTime.Nanoseconds()))
	return elapsedTime.Nanoseconds()
}

func main() {
	var rooms []mazer.Room
	for i := 0; i < 10; i++ {
		rooms = append(rooms, mazer.NewEasyMonsterRoom(fmt.Sprintf("Monster Room %d", i)))
	}
	//for i := 10; i < 20; i++ {
	//	rooms = append(rooms, mazer.NewDifficultMonsterRoom(fmt.Sprintf("Monster Room Difficult %d", i)))
	//}

	c1 := mazer.NewFighter("The Bravest Noob")
	c2 := mazer.NewFighter("Teddin McLanhaw")
	c3 := mazer.NewFighter("Sir Kero Kero, Frog Warrior")

	characterList := []mazer.Character{c1, c2, c3,}

	runGame(characterList, rooms)

	fmt.Println("\n#########\nFinal results:\n######### ")
	for _, character := range(characterList) {
		character.PrintStats()
		fmt.Println()
	}



}

func runGame(characters []mazer.Character, maze []mazer.Room) {
	defer timeTrack(time.Now(), "Overall Game")

	for _, character := range(characters) {
		sendThroughMaze(character, maze)
	}
}

func sendThroughMaze(character mazer.Character, maze []mazer.Room) {
	defer timeTrack(time.Now(), fmt.Sprintf("%s's journey", character.GetName()))

	fmt.Print(fmt.Sprintf("%s's journey: Entrance-", character.GetName()))
	for roomNumber, room := range(maze) {
		roomCleared := false
		for roomUncleared := true; roomUncleared && !character.IsDead(); roomUncleared = !roomCleared {
			fmt.Print(fmt.Sprintf("%d-", roomNumber))
			roomCleared = room.DoChallenge(character)
		}
	}
	if(character.IsDead()) {
		fmt.Print("☠")
	} else {
		fmt.Print("Exit")
	}
	fmt.Println()

	// character.PrintStats()
}
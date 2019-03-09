package main

import (
	"fmt"
	"github.com/Dreznel/mazer/mazer"
	"strings"
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

	/*
	c4 := mazer.NewFighter("The Overconfident Noob")
	c5 := mazer.NewFighter("Geoffram Richson")
	c6 := mazer.NewFighter("Lady Silverear, Werewolf Killer")
*/
	characterList1 := []mazer.Character{c1, c2, c3,}
	characterList2 := []mazer.Character{c1.Clone(), c2.Clone(), c3.Clone(),}
	//characterList2 := []mazer.Character{c4, c5, c6,}

	fmt.Println("Linear-Game--------------")
	linearGameDuration := runGame(characterList1, rooms)

	fmt.Println("\n#########\nFinal results:\n######### ")
	for _, character := range(characterList1) {
		character.PrintStats()
		fmt.Println()
	}

	fmt.Println("\n\n\n")

	fmt.Println("Parallel-Game-------------")
	parallelGameDuration := runGameParallel(characterList2, rooms)
	fmt.Println("\n#########\nFinal results:\n######### ")
	for _, character := range(characterList2) {
		character.PrintStats()
		fmt.Println()
	}

	fmt.Println("\n\n\n")
	fmt.Println("Time Comparisons")
	fmt.Println(fmt.Sprintf("Linear  : %d", linearGameDuration))
	fmt.Println(fmt.Sprintf("Parallel: %d", parallelGameDuration))
	fmt.Println(fmt.Sprintf("Parallel game finished %d nanoseconds faster than linear game.", linearGameDuration - parallelGameDuration))
}

func runGame(characters []mazer.Character, maze []mazer.Room) int64 {
	//defer timeTrack(time.Now(), "Overall Game")
	startTime := time.Now()

	for _, character := range(characters) {
		sendThroughMaze(character, maze)
	}
	return time.Since(startTime).Nanoseconds()
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

func runGameParallel(characters []mazer.Character, maze []mazer.Room) int64 {
	// defer timeTrack(time.Now(), "Overall Game Parallel")
	startTime := time.Now()

	progress := make(chan string)
	for _, character := range(characters) {
		go sendThroughMazeParallel(character, maze, progress)
	}
	heroesFinished := 0
	for {
		heroResult, _ := <- progress
		fmt.Println(heroResult)
		heroesFinished ++
		if heroesFinished == len(characters) {
			close(progress)
			break
		}
	}
	return time.Since(startTime).Nanoseconds()

}

func sendThroughMazeParallel(character mazer.Character, maze []mazer.Room, c chan string) {
	//defer timeTrack(time.Now(), fmt.Sprintf("%s's journey", character.GetName()))

	outputString := strings.Builder{}
	outputString.WriteString(fmt.Sprintf("%s's journey: Entrance-", character.GetName()))
	for roomNumber, room := range(maze) {
		roomCleared := false
		for roomUncleared := true; roomUncleared && !character.IsDead(); roomUncleared = !roomCleared {
			outputString.WriteString(fmt.Sprintf("%d-", roomNumber))
			roomCleared = room.DoChallenge(character)
		}
	}
	if(character.IsDead()) {
		outputString.WriteString("☠")
	} else {
		outputString.WriteString("Exit")
	}
	//outputString.WriteString("\n")
	c <- outputString.String()

}

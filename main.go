package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Stack []string

func (s Stack) Push(str string) Stack {
	return append(s, str)
}

func (s Stack) Pop() (Stack, string, int) {
	l := len(s)
	return s[:l-1], s[l-1], l
}

func main() {
	var S1, S2, S3, S4, S5, S6, S7, S8 Stack
	poles := []Stack{S1, S2, S3, S4, S5, S6, S7, S8}

	for i := 0; i < 16; i++ {
		fmt.Println("ROUND ", i+1, " :")
		poles = White(poles)
		isWin, winner := CheckWiningStatus(poles)
		if isWin {
			fmt.Println(poles)
			fmt.Println("Kazanan: ", winner)
			break
		}
		poles = Green(poles)
		isWin, winner = CheckWiningStatus(poles)
		if isWin {
			fmt.Println(poles)
			fmt.Println("Kazanan: ", winner)
			break
		}
		fmt.Println(poles)

	}

}

func RandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(8)
}

func White(poles []Stack) []Stack {
	poleNumber := ChoosePole(poles)
	poles[poleNumber] = poles[poleNumber].Push("W")
	return poles
}

func Green(poles []Stack) []Stack {
	poleNumber := ChoosePole(poles)
	poles[poleNumber] = poles[poleNumber].Push("G")
	return poles
}

func ChoosePole(poles []Stack) int {
	index := RandomNumber()
	pole := poles[index]
	if len(pole) == 4 {
		index = ChoosePole(poles)
	}
	return index
}

func CheckWiningStatus(poles []Stack) (bool, string) {
	var isWin bool
	var winner string
	isWin, winner = CheckVertially(poles)
	if isWin {
		return isWin, winner
	}
	isWin, winner = CheckHorizontally(poles)
	if isWin {
		return isWin, winner
	}
	return false, ""
}

func CheckVertially(poles []Stack) (bool, string) {
	for i := 0; i < len(poles); i++ {
		pole := poles[i]
		if len(pole) > 0 {
			firstPop, lastBall, _ := pole.Pop()

			if len(firstPop) > 0 {
				secondPop, secondLastBall, _ := firstPop.Pop()

				if lastBall == secondLastBall {
					if len(secondPop) > 0 {
						_, thirdLastBall, _ := secondPop.Pop()
						if secondLastBall == thirdLastBall {
							return true, thirdLastBall
						}
					}

				}
			}

		}

	}
	return false, ""
}

func CheckHorizontally(poles []Stack) (bool, string) {
	for i := 0; i < len(poles)-2; i++ {
		pole := poles[i]
		if len(pole) > 0 {
			_, lastBall, firstIndex := pole.Pop()

			secondPole := poles[i+1]

			if len(secondPole) > 0 {
				_, secondLastBall, secondIndex := secondPole.Pop()

				if firstIndex == secondIndex && lastBall == secondLastBall {
					thirdPole := poles[i+2]

					if len(thirdPole) > 0 {
						_, thirdLastBall, thirdIndex := secondPole.Pop()

						if firstIndex == thirdIndex && lastBall == thirdLastBall {
							return true, lastBall
						}
					}

				}

			}

		}

	}
	return false, ""
}

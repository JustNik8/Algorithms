package main

import (
	"bufio"
	"fmt"
	"os"
)

type empty struct{}

// https://contest.yandex.ru/contest/45468/problems/2/
func main() {
	reader := bufio.NewReader(os.Stdin)

	var k int
	fmt.Fscan(reader, &k)

	var str string
	fmt.Fscan(reader, &str)

	alphabet := make(map[uint8]empty)

	// Собираем алфавит из всех букв
	for i := 0; i < len(str); i++ {
		alphabet[str[i]] = empty{}
	}

	// Идея следующая: решаем через 2 указателя. Правым указателем идем и проверяем равна ли текущая буква
	// букве из алфавита. Если не равна, то вычитаем -1 из changes, Если равна, то просто двигаем правый указатель
	// правый указатель двигаем до упора строки. Когда буквы для замены закончились, то проверяем длинну красивой строки
	// Далее левым указателем выкидываем самую левую букву. Если эта буква равна букве из алфавита, то changes не меняется,
	// иначе прибавляем +1 к changes. Это означает, что если мы выкинули букву не из алфавита, то можем сделать
	// еще 1 замену
	maxBeauty := 1
	for letter := range alphabet {
		changes := k
		left := 0
		right := 0

		for left < len(str) && right < len(str) {
			for right < len(str) {
				if changes == 0 && str[right] != letter {
					break
				}
				if str[right] != letter {
					changes--
				}
				right++
			}

			if right-left > maxBeauty {
				maxBeauty = right - left
			}
			//fmt.Println(string(letter), left, right, right-left)

			if str[left] != letter {
				changes++
			}
			left++
		}
	}

	fmt.Println(maxBeauty)
}

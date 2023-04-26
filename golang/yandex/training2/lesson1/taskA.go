package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://contest.yandex.ru/contest/28730/problems/

/*
Вердикт чекера - 0..7
Вердикт интерактора - 0..7
Код завершения задачи -128..127

1) Интерактор - 0 и Код завершения != 0 => итоговый вердикт = 3 иначе вердикт чекера
2) Интерактор - 1 => итогововый вердикт = вердикт чекера
3) Интерактор = 4 => итоговый вердикт= 3, если код завершений != 0 и иначе итоговый вердикт = 4
4) Интерактор = 6 => Итоговый вердикт = 0
5) Интерактор = 7 => Итоговый вердикт = 1
6) В остальных случаях вердикт = вердикт интерактора
*/

func main() {

	reader := bufio.NewReader(os.Stdin)
	var code, interactor, checker int

	fmt.Fscan(reader, &code, &interactor, &checker)

	var result int
	if interactor == 0 {
		// 1
		if code != 0 {
			result = 3
		} else {
			result = checker
		}
	} else if interactor == 1 {
		// 2
		result = checker
	} else if interactor == 4 {
		// 3
		if code != 0 {
			result = 3
		} else {
			result = 4
		}
	} else if interactor == 6 {
		// 4
		result = 0
	} else if interactor == 7 {
		// 5
		result = 1
	} else {
		// 6
		result = interactor
	}

	fmt.Println(result)
}

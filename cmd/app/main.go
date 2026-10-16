// Лабораторная работа 4. Собственные пакеты, документация и подключение пакетов из GitHub.
// Вариант 2: wallpapercalc (обои)
// 1. Perimeter(length, width float64) (float64, error)
// 2. RollsNeeded(perimeter, height, rollWidth, rollLength float64) (int, error)
// 3. AddExtraRolls(rolls *int, extra int) error
// 4. FormatWallpaperReport(room string, rolls int, rollPrice float64) (string, error)
// В main вызвать функции пакета, обработать ошибки if err != nil,
// вывести результат через fmt.Printf с точностью (%.2f, %.3f).
// Подключить внешний пакет из GitHub (github.com/fatih/color)
// и чужой готовый пакет (github.com/google/uuid), вызвать его функцию.
package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/pilot443/lab4-variant02/pkg/wallpapercalc"
)

func main() {
	var room string
	var length, width, height float64
	var rollWidth, rollLength, rollPrice float64
	var extra int

	fmt.Print("комната: ")
	fmt.Scan(&room)
	fmt.Print("длина и ширина комнаты (м): ")
	fmt.Scan(&length, &width)
	fmt.Print("высота стен (м): ")
	fmt.Scan(&height)
	fmt.Print("ширина и длина рулона (м): ")
	fmt.Scan(&rollWidth, &rollLength)
	fmt.Print("цена рулона: ")
	fmt.Scan(&rollPrice)
	fmt.Print("запасных рулонов: ")
	fmt.Scan(&extra)

	id := uuid.New()
	color.Cyan("операция %s", id.String())

	p, err := wallpapercalc.Perimeter(length, width)
	if err != nil {
		color.Red("ошибка: %v", err)
		os.Exit(1)
	}
	fmt.Printf("периметр: %.2f м\n", p)
	fmt.Printf("площадь стен: %.3f м2\n", p*height)

	rolls, err := wallpapercalc.RollsNeeded(p, height, rollWidth, rollLength)
	if err != nil {
		color.Red("ошибка: %v", err)
		os.Exit(1)
	}
	fmt.Printf("нужно рулонов: %d\n", rolls)

	err = wallpapercalc.AddExtraRolls(&rolls, extra)
	if err != nil {
		color.Red("ошибка: %v", err)
		os.Exit(1)
	}
	fmt.Printf("с запасом: %d\n", rolls)

	report, err := wallpapercalc.FormatWallpaperReport(room, rolls, rollPrice)
	if err != nil {
		color.Red("ошибка: %v", err)
		os.Exit(1)
	}
	color.Green("%s", report)
}

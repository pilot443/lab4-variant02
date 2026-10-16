// Лабораторная работа 4. Вариант 2: wallpapercalc (обои)
// В пакете реализовать функции:
// 1. Perimeter(length, width float64) (float64, error)
// 2. RollsNeeded(perimeter, height, rollWidth, rollLength float64) (int, error)
// 3. AddExtraRolls(rolls *int, extra int) error
// 4. FormatWallpaperReport(room string, rolls int, rollPrice float64) (string, error)
// Вычислительные функции возвращают (значение, error), функция с указателем меняет значение,
// строка отчёта собирается через fmt.Sprintf.
// Имена экспортируемых функций с заглавной буквы, doc-комментарии перед package и функциями,
// входные данные проверять и ошибки возвращать через fmt.Errorf.

// Package wallpapercalc считает, сколько рулонов обоев нужно на комнату и сколько это стоит.
package wallpapercalc

import (
	"fmt"
	"math"
)

// Perimeter возвращает периметр комнаты по длине и ширине.
func Perimeter(length, width float64) (float64, error) {
	if length <= 0 || width <= 0 {
		return 0, fmt.Errorf("длина и ширина должны быть больше нуля")
	}
	return 2 * (length + width), nil
}

// RollsNeeded считает число рулонов: сколько полос нужно по периметру
// и сколько полос получается из одного рулона.
func RollsNeeded(perimeter, height, rollWidth, rollLength float64) (int, error) {
	if perimeter <= 0 || height <= 0 || rollWidth <= 0 || rollLength <= 0 {
		return 0, fmt.Errorf("все размеры должны быть больше нуля")
	}

	strips := math.Ceil(perimeter / rollWidth)
	stripsInRoll := math.Floor(rollLength / height)
	if stripsInRoll == 0 {
		return 0, fmt.Errorf("рулон %.2f м короче высоты стены %.2f м", rollLength, height)
	}
	return int(math.Ceil(strips / stripsInRoll)), nil
}

// AddExtraRolls добавляет запасные рулоны к уже посчитанным.
func AddExtraRolls(rolls *int, extra int) error {
	if rolls == nil {
		return fmt.Errorf("не передано число рулонов")
	}
	if extra < 0 {
		return fmt.Errorf("запас не может быть отрицательным")
	}
	*rolls += extra
	return nil
}

// FormatWallpaperReport собирает строку отчёта: комната, рулоны, цена рулона и итог.
func FormatWallpaperReport(room string, rolls int, rollPrice float64) (string, error) {
	if room == "" {
		return "", fmt.Errorf("не указана комната")
	}
	if rolls <= 0 {
		return "", fmt.Errorf("рулонов должно быть больше нуля")
	}
	if rollPrice < 0 {
		return "", fmt.Errorf("цена рулона не может быть отрицательной")
	}

	total := float64(rolls) * rollPrice
	line := fmt.Sprintf("%-10s | рулонов: %3d | цена рулона: %8.2f | итого: %10.2f", room, rolls, rollPrice, total)
	return line, nil
}

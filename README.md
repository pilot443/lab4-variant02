# lab4-variant02, wallpapercalc

Лабораторная работа 4, вариант 2. Пакет считает обои для комнаты.

## Функции пакета

- `Perimeter(length, width)` периметр комнаты
- `RollsNeeded(perimeter, height, rollWidth, rollLength)` сколько нужно рулонов
- `AddExtraRolls(rolls *int, extra)` добавляет запас через указатель
- `FormatWallpaperReport(room, rolls, rollPrice)` строка отчёта через Sprintf

Внешние пакеты: `github.com/fatih/color` (цветной вывод) и `github.com/google/uuid` (номер операции).

## Запуск

```
go run ./cmd/app
```

Пример ввода: `Спальня`, `4 3.5`, `2.7`, `0.53 10.05`, `4500`, `1`.

## Документация

```
go doc ./pkg/wallpapercalc
go doc ./pkg/wallpapercalc RollsNeeded
```

Через godoc:

```
go install golang.org/x/tools/cmd/godoc@latest
godoc -http=:6060
```

Потом открыть http://localhost:6060/pkg/github.com/pilot443/lab4-variant02/

## Подключить пакет в другом проекте

```
go get github.com/pilot443/lab4-variant02@latest
```

```go
import "github.com/pilot443/lab4-variant02/pkg/wallpapercalc"
```

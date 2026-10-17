package main

import "fmt"

func main() {
	var cost int
	fmt.Print("Введите сумму заказа: ")
	fmt.Scan(&cost)

	if cost < 0 {
		fmt.Println("Ошибка: сумма заказа не может быть отрицательной")
		return
	}

	var cat string
	switch {
	case cost < 10_000:
		cat = "small"
	case cost < 50_000:
		cat = "regular"
	default:
		cat = "large"
	}
	fmt.Printf(
		"Сумма заказа: %d коп.\nКатегория: %s\n",
		cost, cat,
	)
}

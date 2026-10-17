package main

import "fmt"

func OrderStats(orders []int) (int, int, int, float64) {
	minOrder := orders[0]
	maxOrder := orders[0]
	total := 0
	for _, order := range orders {
		minOrder = min(minOrder, order)
		maxOrder = max(maxOrder, order)
		total += order
	}
	avg := float64(total) / float64(len(orders))
	return minOrder, maxOrder, total, avg
}

func main() {
	var n int
	fmt.Print("Введите кол-во заказов:")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Нет заказов для анализа")
		return
	}

	orders := make([]int, n)

	fmt.Printf("Введите %d значений:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&orders[i])
	}

	minOrder, maxOrder, total, avg := OrderStats(orders)

	fmt.Printf("Количество заказов: %d\n", len(orders))
	fmt.Printf("Минимальный заказ: %d коп.\n", minOrder)
	fmt.Printf("Максимальный заказ: %d коп.\n", maxOrder)
	fmt.Printf("Общая сумма: %d коп.\n", total)
	fmt.Printf("Средний заказ: %g коп.\n", avg)
}

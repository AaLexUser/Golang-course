package safe_sum

import "math"

// SafeSum последовательно суммирует транзакции, пропуская те, добавление
// которых привело бы к переполнению int64. skipped — индексы пропущенных
// элементов исходного среза, в порядке появления.
func SafeSum(transactions []int64) (total int64, skipped []int) {
	// TODO: проверьте возможность переполнения ДО сложения.
	for idx, item := range transactions {
		if item > 0 && total > math.MaxInt64-item ||
			item < 0 && total < math.MinInt64-item {
			skipped = append(skipped, idx)
		} else {
			total += item
		}
	}
	return total, skipped
}

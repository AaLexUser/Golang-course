package unicode_name

import "unicode"

// AnalyzeName считает байты, Unicode-кодовые точки, латинские и кириллические
// буквы и возвращает байтовые позиции начала каждой Unicode-кодовой точки.
// Предполагается, что входная строка содержит корректный UTF-8.
func AnalyzeName(name string) (byteCount, runeCount, latinCount, cyrillicCount int, positions []int) {
	// TODO: используйте range по строке; для классификации букв можно
	// использовать стандартный пакет unicode.
	byteCount = len(name)
	for i, letter := range name {
		positions = append(positions, i)
		if unicode.IsLetter(letter) {
			if unicode.Is(unicode.Latin, letter) {
				latinCount += 1
			} else if unicode.Is(unicode.Cyrillic, letter) {
				cyrillicCount += 1
			}
		}
	}
	runeCount = len(positions)
	return byteCount, runeCount, latinCount, cyrillicCount, positions
}

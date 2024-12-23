package main

/*
1) Переменная justString объявлена глобально. Это потенциальный race condition. Если несколько горутин будут изменять эту строку без синхронизации, то состояние гонки обеспечено.

2) Потенциальная утечка памяти. Переменная justString ссылается на срез строки. Но сборщик мусора не очистит область "больше не используемой" строки.
Соответственно, createHugeString может создать огромную строку, которую GC не сможет очистить
*/

func main() {
	justString := someFunc()
	_ = justString
}
func someFunc() string {
	v := createHugeString(1 << 10)
	return v[:100]
}

func createHugeString(length int) string {
	return string(make([]rune, length)) // Создаем слайс рун вместо []byte, потому что подразумеваем под length - количество символов utf8
}

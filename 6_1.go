package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Заданный путь к каталогу (ваш стартовый путь)
	startPath := "/local/Roaming/minecraft"

	// Инициализируем переменные для хранения информации о файле максимального размера
	var maxFileSize int64
	var maxFilePath string

	// Рекурсивно обходим файловую систему
	err := filepath.Walk(startPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Если это файл, проверяем его размер
			if info.Size() > maxFileSize {
				maxFileSize = info.Size()
				maxFilePath = path
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Ошибка при обходе файловой системы:", err)
		return
	}

	if maxFilePath != "" {
		fmt.Printf("Файл максимального размера: %s (размер: %d байт)\n", maxFilePath, maxFileSize)
	} else {
		fmt.Println("Файлы не найдены.")
	}
}

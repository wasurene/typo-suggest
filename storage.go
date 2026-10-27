package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		}
	}()

	var words []string

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		newLine := strings.TrimSpace(scan.Text())

		if newLine != "" {
			words = append(words, newLine)
		}
	}

	if err := scan.Err(); err != nil {
		fmt.Println("Ошибка чтения файла")
		return []string{}, err
	}

	return words, nil
}

func writeResult(words []Result, filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла")
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Ошибка при закрытии файла")
			fmt.Printf("Текст ошибки: %v", err)
		}
	}()

	writer := bufio.NewWriter(file)

	for i := range words {
		fmt.Fprintln(writer, words[i].Word)
	}

	writer.Flush()
	return nil
}

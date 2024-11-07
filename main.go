package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Проверяем аргументы командной строки
	var inputFile, outputFile string

	if len(os.Args) >= 2 {
		inputFile = os.Args[1]
	} else {
		inputFile = "input.txt"
	}

	if len(os.Args) >= 3 {
		outputFile = os.Args[2]
	} else {
		outputFile = "output.txt"
	}

	// Читаем содержимое файла с выражениями
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	// Открываем файл для записи результатов
	output, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Error opening output file: %v\n", err)
		return
	}
	defer output.Close()

	// Создаем буфер для записи результатов
	writer := bufio.NewWriter(output)

	// Разделяем содержимое файла на строки
	lines := strings.Split(string(content), "\n")

	// Создаем регулярное выражение для поиска математических выражений в строках
	regex := regexp.MustCompile(`(\d+)([+\-*\/])(\d+)=\?`)

	// Обработка каждой строки
	for _, line := range lines {
		// Проверяем, является ли строка математическим выражением
		if regex.MatchString(line) {
			// Извлекаем операнды и операцию из строки
			matches := regex.FindStringSubmatch(line)
			num1, _ := strconv.Atoi(matches[1])
			operator := matches[2]
			num2, _ := strconv.Atoi(matches[3])

			// Вычисляем результат
			result := 0
			switch operator {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "/":
				result = num1 / num2
			}

			// Формируем строку с результатом
			resultLine := fmt.Sprintf("%d%s%d=%d\n", num1, operator, num2, result)

			// Записываем строку в буфер
			writer.WriteString(resultLine)
		}
	}

	// Очищаем буфер и записываем результаты в файл
	writer.Flush()

	fmt.Println("Results written to", outputFile)
}

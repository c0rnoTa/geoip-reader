package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"io"
	"os"
	"strings"
)

func readFile(fileName string) ([]string, error) {
	// Открываем файл
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Читаем все содержимое файла
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Преобразуем содержимое в строки и разделяем их по новому строку
	return strings.Split(string(content), "\n"), nil
}

func writeFile(fileName string, geoipInfo map[string]*geoip2.City) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Записываем строки в файл
	for address, line := range geoipInfo {
		if _, err = io.WriteString(file, fmt.Sprintf("%s;%s;%s;%s\n", address, line.Country.IsoCode, line.Country.Names["ru"], line.City.Names["ru"])); err != nil {
			return err
		}
	}

	return nil
}

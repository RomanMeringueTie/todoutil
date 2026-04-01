package repository

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type TodoConfig struct {
	prefix string
	suffix string
}

func getTodoConfig(path string) *TodoConfig {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var prefix string
	var suffix string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		prefix = getConfigFieldValue(line, "PREFIX")
		suffix = getConfigFieldValue(line, "SUFFIX")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil
	}

	return &TodoConfig{prefix: prefix, suffix: suffix}
}

func getConfigFieldValue(line, field string) string {
	fieldPrefix := fmt.Sprintf("%s=", field)
	if strings.HasPrefix(line, fieldPrefix) {
		return strings.Split(line, "=")[1]
	}
	return ""
}

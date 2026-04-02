package repository

import (
	"bufio"
	"errors"
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
		if errors.Is(err, os.ErrNotExist) {
			err = fmt.Errorf("Not a todo directory (create .todoconfig)")
		}
		log.Fatal(err)
	}
	defer file.Close()

	var prefix string
	var suffix string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PREFIX=") {
			prefix = strings.Split(line, "=")[1]
		} else if strings.HasPrefix(line, "SUFFIX=") {
			suffix = strings.Split(line, "=")[1]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil
	}

	return &TodoConfig{prefix: prefix, suffix: suffix}
}

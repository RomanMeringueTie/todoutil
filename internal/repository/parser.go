package repository

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"todo/internal/model"
)

const (
	defaultTodoSuffix  = ".todo"
	todoConfigFilename = ".todoconfig"
)

func parseDirs(path string) []model.Todo {
	todoConfig := getTodoConfig(todoConfigFilename)
	if todoConfig == nil {
		return []model.Todo{}
	}

	todos := make([]model.Todo, 0)
	filepath.WalkDir(path, parseDirEntry(&todos, *todoConfig))
	return todos
}

func parseDirEntry(todos *[]model.Todo, config TodoConfig) func(string, fs.DirEntry, error) error {
	return func(path string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !dirEntry.IsDir() {
			if hasTodoSuffix(dirEntry.Name(), config.suffix) {
				parseFile(todos, path, config.prefix)
			}
		}

		return nil
	}
}

func hasTodoSuffix(filename, suffix string) bool {
	return strings.HasSuffix(filename, suffix) || strings.HasSuffix(filename, defaultTodoSuffix)
}

func parseFile(todos *[]model.Todo, path, prefix string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		todo := parseLine(path, scanner.Text(), prefix)
		if todo != nil {
			*todos = append(*todos, *todo)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func parseLine(path, line, prefix string) *model.Todo {
	// TODO: Read suffix and prefix from config file [x]
	if strings.Contains(line, prefix) {
		lineWithoutPrefix := eraseTodoPrefix(line, prefix)

		name, status := getTodoNameAndStatus(lineWithoutPrefix)
		return model.NewTodo(name, status, path)
	}

	return nil
}

func eraseTodoPrefix(line, prefix string) string {
	prefixIndex := strings.Index(line, prefix)
	lineWithoutPrefix := line[prefixIndex+len(prefix):]
	return lineWithoutPrefix
}

func getTodoNameAndStatus(line string) (string, model.Status) {
	splittedLine := strings.Split(line, " ")
	todoLineLastPart := splittedLine[len(splittedLine)-1]

	status, isFind := model.SymbolToStatus(todoLineLastPart)
	var name string
	if isFind {
		name = strings.Join(splittedLine[:len(splittedLine)-1], " ")
	} else {
		name = strings.Join(splittedLine, " ")
	}

	return name, status
}

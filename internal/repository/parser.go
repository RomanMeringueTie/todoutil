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

const defaultTodoSuffix = ".todo"

func parseDirs(path, suffix string) []model.Todo {
	todos := make([]model.Todo, 0)
	filepath.WalkDir(path, parseDirEntry(&todos, suffix))
	return todos
}

func parseDirEntry(todos *[]model.Todo, suffix string) func(string, fs.DirEntry, error) error {
	return func(path string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !dirEntry.IsDir() {
			if strings.HasSuffix(dirEntry.Name(), suffix) || strings.HasSuffix(dirEntry.Name(), defaultTodoSuffix) {
				parseFile(todos, path)
			}
		}

		return nil
	}
}

func parseFile(todos *[]model.Todo, path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		todo := parseLine(path, scanner.Text())
		if todo != nil {
			*todos = append(*todos, *todo)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func parseLine(path, line string) *model.Todo {
	if strings.Contains(line, "TODO: ") {
		prefixIndex := strings.Index(line, "TODO: ")
		lineWithoutPrefix := line[prefixIndex+len("TODO: "):]

		name, status := getTodoNameAndStatus(lineWithoutPrefix)
		return model.NewTodo(name, status, path)
	}

	return nil
}

func getTodoNameAndStatus(line string) (string, model.Status) {
	// TODO: Replace with regex [*]
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

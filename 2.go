package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var arrays = make(map[string][]int)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Пример инструкций
	instructions := []string{
		"Load A, in.txt;",
		"Load B, in2.txt;",
		"Print A, all;",
		"Sort A+;",
		"Stats A;",
	}

	for _, instruction := range instructions {
		err := executeInstruction(instruction)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func executeInstruction(instruction string) error {
	parts := strings.Split(instruction, " ")
	cmd := strings.ToLower(parts[0])
	arrayName := parts[1][:1] // Имя массива - первая буква после запятой

	switch cmd {
	case "load":
		filename := parts[2][:len(parts[2])-1] // Удаляем ";" из конца строки
		err := loadArray(arrayName, filename)
		if err != nil {
			return err
		}
	case "save":
		filename := parts[2][:len(parts[2])-1]
		err := saveArray(arrayName, filename)
		if err != nil {
			return err
		}
	case "rand":
		count, _ := strconv.Atoi(parts[2])
		lb, _ := strconv.Atoi(parts[3])
		rb, _ := strconv.Atoi(parts[4][:len(parts[4])-1])
		randArray(arrayName, count, lb, rb)
	case "concat":
		secondArray := parts[2][:1]
		err := concatArrays(arrayName, secondArray)
		if err != nil {
			return err
		}
	// Другие команды...
	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}

	return nil
}

func loadArray(name, filename string) error {   //Load A, in.txt
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }

    strValues := strings.Fields(string(data))   //преобразуем в строки и разбиваем на подстроки
    values := make([]int, 0, len(strValues))  
    for _, strValue := range strValues {
        value, err := strconv.Atoi(strValue)
        if err != nil {
            return fmt.Errorf("invalid value in file: %s", strValue)
        }
        values = append(values, value)
    }

    arrays[name] = values
    return nil
}

func saveArray(name, filename string) error { //Save A, out.txt
	values, ok := arrays[name]
	if !ok {
		return fmt.Errorf("array %s not found", name)
	}

	data := []byte(strings.Join(strings.Fields(fmt.Sprintf("%v", values)), "\n"))
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}


func randArray(name string, count, lb, rb int) {
	values := make([]int, count)
	for i := 0; i < count; i++ {
		values[i] = rand.Intn(rb-lb+1) + lb
	}

	arrays[name] = values
}

func concatArrays(first, second string) error {
	arr1, ok1 := arrays[first]
	arr2, ok2 := arrays[second]

	if !ok1 || !ok2 {
		return fmt.Errorf("one of the arrays not found")
	}

	arrays[first] = append(arr1, arr2...)
	return nil
}

func sortArray(name string, desc bool) error {
	values, ok := arrays[name]
	if !ok {
		return fmt.Errorf("array %s not found", name)
	}

	if desc {
		sort.Sort(sort.Reverse(sort.IntSlice(values)))
	} else {
		sort.Ints(values)
	}

	arrays[name] = values
	return nil
}


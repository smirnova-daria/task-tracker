package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Tracker struct {
	Tasks map[string][]string
}

func New() *Tracker {
	return &Tracker{
		Tasks: map[string][]string{},
	}
}

func (t *Tracker) Add(date, task string) {
	tasksByDate, ok := t.Tasks[date]
	if !ok {
		t.Tasks[date] = append(t.Tasks[date], task)
		return
	}

	isExist := false
	for _, t := range tasksByDate {
		if t == task {
			isExist = true
		}
	}
	if !isExist {
		t.Tasks[date] = append(t.Tasks[date], task)
	}
}

func (t *Tracker) DeleteTask(date, task string) {
	tasksByDate, ok := t.Tasks[date]
	if !ok {
		fmt.Println("Event not found")
		return
	}

	isExist := false
	idx := 0
	for i, t := range tasksByDate {
		if t == task {
			isExist = true
			idx = i
		}
	}
	if isExist {
		t.Tasks[date] = append(t.Tasks[date][:idx], t.Tasks[date][idx+1:]...)
		fmt.Println("Deleted successfully")
	} else {
		fmt.Println("Event not found")
	}
}

func (t *Tracker) DeleteAllTasks(date string) {
	tasksCount := 0

	tasksByDate, ok := t.Tasks[date]
	if ok {
		tasksCount = len(tasksByDate)
	}

	delete(t.Tasks, date)
	fmt.Printf("Deleted %d events\n", tasksCount)
}

func (t *Tracker) Find(date string) {
	tasksByDate, ok := t.Tasks[date]

	if ok {
		sort.Strings(tasksByDate)
		for _, v := range tasksByDate {
			fmt.Println(v)
		}
	}
}

func (t *Tracker) Print() {
	sortedDates := sortDates(t.Tasks)

	for _, d := range sortedDates {
		tasks := t.Tasks[d]
		sort.Strings(tasks)
		for _, tt := range tasks {
			formattedDate := formatDate(d)
			fmt.Printf("%s %s\n", formattedDate, tt)
		}
	}
}

func sortDates(tasks map[string][]string) []string {
	dates := make([]string, 0, len(tasks))
	for k := range tasks {
		dates = append(dates, k)
	}

	sort.Strings(dates)
	return dates
}

func formatDate(date string) string {
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return date
	}
	year := fmt.Sprintf("%04s", parts[0])
	month := fmt.Sprintf("%02s", parts[1])
	day := fmt.Sprintf("%02s", parts[2])
	return fmt.Sprintf("%s-%s-%s", year, month, day)
}

var tracker *Tracker

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		isQuit, err := AnalyzeInput(scanner.Text())

		if err != nil {
			fmt.Println(err)
		}

		if isQuit {
			return
		}
	}

}

func AnalyzeInput(input string) (isQuit bool, err error) {
	inputSlice := strings.Split(input, " ")

	cmd := inputSlice[0]

	if cmd == "Quit" {
		return true, nil
	}

	if tracker == nil && cmd != "StartApp" {
		return false, fmt.Errorf("please input 'StartApp' for start tracker")
	}

	switch cmd {
	case "StartApp":
		tracker = New()
	case "Add":
		tracker.Add(inputSlice[1], inputSlice[2])
	case "Find":
		tracker.Find(inputSlice[1])
	case "Del":
		if len(inputSlice) > 2 {
			tracker.DeleteTask(inputSlice[1], inputSlice[2])
		} else {
			tracker.DeleteAllTasks(inputSlice[1])
		}
	case "Print":
		tracker.Print()
	}
	return false, nil
}

package main

import (
	"net/url"
	"os"
	"strconv"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := strconv.Itoa(entryCount) + ": " + text
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

// separation of concerns
func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	// ...
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	// ...
}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

type Persistance struct {
	lineSeparator string
}

func (p *Persistance) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	println("Journal entries:")
	println(j.String())

	p := Persistance{"\r"}
	p.SaveToFile(&j, "journal.txt")
}

// Journal is responsible for adding and removing entries, but it is also responsible for saving and loading data.
// This is a violation of the Single Responsibility Principle.
// The Journal class should be responsible for managing entries, and a separate class should be responsible for persistence.

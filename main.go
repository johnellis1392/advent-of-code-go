package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	aoc "github.com/johnellis1392/advent-of-code-go/2022"
	common "github.com/johnellis1392/advent-of-code-go/common"
)

func DownloadInputForDay(d common.Day, sessionId string) error {
	var err error = nil

	client := &http.Client{}
	fmt.Printf("Downloading input for %s/day%s...\n", d.Year(), d.Day())
	filename := fmt.Sprintf("./input/%s/day%s.input.txt", d.Year(), d.Day())
	dir := path.Dir(filename)
	if err = os.MkdirAll(dir, 0777); err != nil {
		return err
	}

	day, _ := strconv.Atoi(d.Day())
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%v/input", d.Year(), day)
	var req *http.Request
	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return err
	}
	cookie := fmt.Sprintf("session=%s", sessionId)
	req.Header.Add("Cookie", cookie)

	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return err
	}
	defer res.Body.Close()

	var body []byte
	if body, err = io.ReadAll(res.Body); err != nil {
		return err
	}

	if err = os.WriteFile(filename, body, 0777); err != nil {
		return err
	}

	return nil
}

func ReadInputForDay(d common.Day, sessionId string) (string, error) {
	filename := fmt.Sprintf("./input/%s/day%s.input.txt", d.Year(), d.Day())
	_, err := os.Stat(filename)
	if err != nil {
		DownloadInputForDay(d, sessionId)
	}

	if data, err := os.ReadFile(filename); err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

func RunDay(d common.Day, sessionId string) {
	var err error
	var input string
	if input, err = ReadInputForDay(d, sessionId); err != nil {
		log.Panic(err)
	}

	err = d.Parse(input)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%s, Day %s, Part 1: %v\n", d.Year(), d.Day(), d.Part1())
	fmt.Printf("%s, Day %s, Part 2: %v\n", d.Year(), d.Day(), d.Part2())
}

func main() {
	d := new(aoc.Day01)
	sessionId := ""
	fmt.Println("Running...")
	RunDay(d, sessionId)
}

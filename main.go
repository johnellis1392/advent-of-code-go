package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	common "github.com/johnellis1392/advent-of-code-go/common"
)

func SessionId() string {
	return ""
}

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

func RunDay(d common.Day, sessionId string) (any, any, error) {
	var err error
	var input string
	if input, err = ReadInputForDay(d, sessionId); err != nil {
		return nil, nil, err
	}

	err = d.Parse(input)
	if err != nil {
		return nil, nil, err
	}

	part1 := d.Part1()
	part2 := d.Part2()
	fmt.Printf("%s, Day %s, Part 1: %v\n", d.Year(), d.Day(), part1)
	fmt.Printf("%s, Day %s, Part 2: %v\n", d.Year(), d.Day(), part2)
	return part1, part2, nil
}

func RunPart1(d common.Day, sessionId string) (any, error) {
	var err error
	var input string
	if input, err = ReadInputForDay(d, sessionId); err != nil {
		return nil, err
	}

	err = d.Parse(input)
	if err != nil {
		return nil, err
	}

	part1 := d.Part1()
	fmt.Printf("%s, Day %s, Part 1: %v\n", d.Year(), d.Day(), part1)
	return part1, nil
}

func RunPart2(d common.Day, sessionId string) (any, error) {
	var input string
	var err error
	if input, err = ReadInputForDay(d, sessionId); err != nil {
		return nil, err
	}

	err = d.Parse(input)
	if err != nil {
		return nil, err
	}

	part2 := d.Part2()
	fmt.Printf("%s, Day %s, Part 1: %v\n", d.Year(), d.Day(), part2)
	return part2, err
}

type ErrorResponse struct {
	Err error `json:"error"`
}

func writeErr(w http.ResponseWriter, err error) {
	e, err2 := json.Marshal(ErrorResponse{err})
	if err2 == nil {
		w.Write(e)
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func writeJson(w http.ResponseWriter, data any) {
	w.Header().Add("Content-Type", "application/json")
	s, err := json.Marshal(data)
	if err != nil {
		writeErr(w, err)
	} else {
		w.Write(s)
	}
}

func createRouter() *http.ServeMux {
	router := http.NewServeMux()

	yearsRouter := http.NewServeMux()

	// Fetch all years
	yearsRouter.HandleFunc("GET /", func(w http.ResponseWriter, req *http.Request) {
		n := 2024 - 2015 + 1
		years := make([]string, n)
		for i := range n {
			years[i] = fmt.Sprintf("%d", 2015+i)
		}
		type response struct {
			Years []string `json:"years"`
		}
		writeJson(w, response{years})
	})

	// Get available solutions for year
	yearsRouter.HandleFunc("GET /{year}", func(w http.ResponseWriter, req *http.Request) {
		year := req.PathValue("year")
		if !common.IsValidYear(year) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		type dayView struct {
			Part1 bool `json:"part1"`
			Part2 bool `json:"part2"`
		}

		type response struct {
			Year      string              `json:"year"`
			Solutions map[string]*dayView `json:"solutions"`
		}

		res := make(map[string]*dayView)
		for _, d := range common.AllDays() {
			solution := common.SolutionForDay(Solutions, year, d)
			if solution != nil {
				// TODO: Look up cached solutions to verify if something is actually complete
				res[d] = &dayView{true, true}
			} else {
				res[d] = nil
			}
		}
		writeJson(w, response{year, res})
	})

	daysRouter := http.NewServeMux()

	// Fetch day information for year
	daysRouter.HandleFunc("GET /", func(w http.ResponseWriter, req *http.Request) {})

	// Fetch information for specified day
	daysRouter.HandleFunc("GET /{day}", func(w http.ResponseWriter, req *http.Request) {
		// year := req.PathValue("year")
		// day := req.PathValue("day")
	})

	// Run part1 & part2 for day
	daysRouter.HandleFunc("POST /{day}", func(w http.ResponseWriter, req *http.Request) {
		year, day := req.PathValue("year"), req.PathValue("day")
		solution := common.SolutionForDay(Solutions, year, day)
		if solution == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		RunDay(solution, SessionId())
	})

	// Download input for day
	daysRouter.HandleFunc("POST /{day}/input", func(w http.ResponseWriter, req *http.Request) {
		year, day := req.PathValue("year"), req.PathValue("day")
		solution := common.SolutionForDay(Solutions, year, day)
		if solution == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		DownloadInputForDay(solution, SessionId())
	})

	// Get input for day
	daysRouter.HandleFunc("GET /{day}/input", func(w http.ResponseWriter, req *http.Request) {
		year, day := req.PathValue("year"), req.PathValue("day")
		solution := common.SolutionForDay(Solutions, year, day)
		if solution == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		input, err := ReadInputForDay(solution, SessionId())
		if err != nil {
			writeErr(w, err)
			return
		}
		type response struct {
			Year  string `json:"year"`
			Day   string `json:"day"`
			Input string `json:"input"`
		}
		writeJson(w, response{year, day, input})
	})

	// Get Part1 for day
	daysRouter.HandleFunc("GET /{day}/part1", func(w http.ResponseWriter, req *http.Request) {})

	// Run Part1 for day
	daysRouter.HandleFunc("POST /{day}/part1", func(w http.ResponseWriter, req *http.Request) {
		year, day := req.PathValue("year"), req.PathValue("day")
		solution := common.SolutionForDay(Solutions, year, day)
		if solution == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		RunPart1(solution, SessionId())
	})

	// Get Part2 for day
	daysRouter.HandleFunc("GET /{day}/part2", func(w http.ResponseWriter, req *http.Request) {})

	// Run Part2 for day
	daysRouter.HandleFunc("POST /{day}/part2", func(w http.ResponseWriter, req *http.Request) {
		year, day := req.PathValue("year"), req.PathValue("day")
		solution := common.SolutionForDay(Solutions, year, day)
		if solution == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		RunPart2(solution, SessionId())
	})

	yearsRouter.Handle("/days", daysRouter)
	router.Handle("/years", yearsRouter)

	return router
}

func main() {
	// d := new(aoc.Day01)
	// sessionId := ""
	// fmt.Println("Running...")
	// RunDay(d, sessionId)

	server := createRouter()

	address := ":8080"
	fmt.Printf("Listening on %s...\n", address)
	log.Fatal(http.ListenAndServe(address, server))
}

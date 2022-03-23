package pkg

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func (g *Graph) ReadFile() {

	var rooms []string
	var links [][]string
	var hasStart bool = false
	var hasEnd bool = false
	if len(os.Args) != 2 {
		fmt.Println("usage: go run main.go maps/example00.txt")
		os.Exit(0)
	}
	inputFile := os.Args[1]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	data = []byte(strings.ReplaceAll(string(data), "\r", ""))
	Task = string(data)

	lines := strings.Split(string(data), "\n")

	for i, v := range lines {
		if i == 0 {
			g.AntNumber, _ = strconv.Atoi(v)
			if g.AntNumber < 1 {
				fmt.Println("ERROR: invalid data, no ants in given map")
				os.Exit(1)
			}
			if g.AntNumber > 9999 {
				fmt.Println("ERROR: invalid data, too many ants in given map")
				os.Exit(1)
			}
		}
		if v == "##start" {
			hasStart = true
			g.StartRoom = getRoomName(lines[i+1])

		}
		if v == "##end" {
			hasEnd = true
			g.EndRoom = getRoomName(lines[i+1])
		}
		if strings.Contains(v, " ") {
			var roomName string
			room := strings.Split(v, " ")
			roomName = room[0]

			var match = regexp.MustCompile(`^[^L#]\s*`)
			result := match.MatchString(roomName)

			if result {
				rooms = append(rooms, roomName)
			} else {
				fmt.Println("ERROR: invalid data, room name should not start with an L or #")
				os.Exit(0)
			}
		}
		if strings.Contains(v, "-") {
			link := strings.Split(v, "-")
			for _, roomInLink := range link {
				if !matched(roomInLink, rooms) {
					fmt.Println("ERROR: invalid data, the link`s room doesn`t exist")
					os.Exit(1)
				}
			}
			links = append(links, link)
		}
	}
	if !hasStart {
		fmt.Println("ERROR: invalid data, no start room in file")
		os.Exit(1)
	}
	if !hasEnd {
		fmt.Println("ERROR: invalid data, no end room in file")
		os.Exit(1)
	}
	if len(links) < 1 {
		fmt.Println("ERROR: invalid data, no rooms are connected")
		os.Exit(1)
	}
	for i := range rooms {
		for j := i + 1; j < len(rooms); j++ {
			if rooms[i] == rooms[j] {
				fmt.Println("ERROR: invalid data, room duplicate found")
				os.Exit(1)
			}
		}
	}
	g.GetAdjList(rooms, links)

}

func matched(roomInLink string, rooms []string) bool {
	for _, room := range rooms {
		if room == roomInLink {
			return true
		}
	}
	return false
}

func getRoomName(ln string) string {
	splittedLine := strings.Split(ln, " ")
	return splittedLine[0]

}

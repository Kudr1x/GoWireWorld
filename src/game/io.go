package game

import (
	"fmt"
	"github.com/sqweek/dialog"
	"log"
	"os"
	"strconv"
	"strings"
)

func (g *Game) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for cell, state := range g.cells {
		line := fmt.Sprintf("%d %d %d\n", cell.X, cell.Y, state)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) LoadFromFile(filename string) error {
	g.cells = make(map[Cell]int)

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 3 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		state, _ := strconv.Atoi(parts[2])
		g.cells[Cell{x, y}] = state
	}
	return nil
}

func (g *Game) SaveWithDialog() {
	filename, err := dialog.File().Title("Save Game").Filter("Wireworld Save File", "wws").Save()
	if err != nil {
		log.Println("Save dialog canceled:", err)
		return
	}
	if filename == "" {
		return
	}
	if !strings.HasSuffix(filename, ".wws") {
		filename += ".wws"
	}
	if err := g.SaveToFile(filename); err != nil {
		log.Println("Failed to save game:", err)
	} else {
		log.Println("Game saved to:", filename)
	}
}

func (g *Game) LoadWithDialog() {
	filename, err := dialog.File().Title("Load Game").Filter("Wireworld Save File", "wws").Load()
	if err != nil {
		log.Println("Load dialog canceled:", err)
		return
	}
	if filename == "" {
		return
	}
	if err := g.LoadFromFile(filename); err != nil {
		log.Println("Failed to load game:", err)
	} else {
		log.Println("Game loaded from:", filename)
	}
}

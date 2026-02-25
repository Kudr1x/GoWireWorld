package game

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sqweek/dialog"
)

func (g *Game) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	g.mu.RLock()
	defer g.mu.RUnlock()

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
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	newCells := make(map[Cell]int)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 3 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		state, _ := strconv.Atoi(parts[2])
		newCells[Cell{x, y}] = state
	}

	g.mu.Lock()
	g.cells = newCells
	g.mu.Unlock()

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

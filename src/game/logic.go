package game

func CalculateNextState(cells map[Cell]int) map[Cell]int {
	newCells := make(map[Cell]int, len(cells))
	for cell, currentState := range cells {
		switch currentState {
		case ElectronHead:
			newCells[cell] = ElectronTail
		case ElectronTail:
			newCells[cell] = Conductor
		case Conductor:
			count := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 {
						continue
					}
					if cells[Cell{cell.X + dx, cell.Y + dy}] == ElectronHead {
						count++
					}
				}
			}
			if count == 1 || count == 2 {
				newCells[cell] = ElectronHead
			} else {
				newCells[cell] = Conductor
			}
		}
	}
	return newCells
}

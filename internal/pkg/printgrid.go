package pkg

import "fmt"

const cols = 5

func PrintGrid(data map[string]float64) {
	pairs := make([]struct{ ID, Text string}, 0, len(data))
	maxWidth := 0
	for id, rate := range data {
		text := fmt.Sprintf("'%s': %.2f", id, rate)
		if w := len(text); w > maxWidth {
			maxWidth = w
		}
		pairs = append(pairs, struct{ID string; Text string}{ID: id, Text: text})
	}

	for i, p := range pairs {
		fmt.Printf("%-*s", maxWidth, p.Text)

		if (i+1)%cols == 0 {
			fmt.Println()
		} else {
			fmt.Print("    ")
		}
	}
	fmt.Println()
}
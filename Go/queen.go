package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	queens := solve(N)
	prettyPrint(N, queens)
}

func solve(N int) map[string]bool {
	queens := make(map[string]bool)

	helper(queens, 0, N)
	return queens
}

func helper(queens map[string]bool, y int, N int) bool {
	if y >= N {
		return true
	}

	for i := 0; i < N; i++ {
		if isSafe(queens, i, y, N) {
			queens[k([]int{y, i})] = true
			if helper(queens, y+1, N) {
				return true
			} else {
				queens[k([]int{y, i})] = false
			}
		}
	}

	return false
}

func isSafe(queens map[string]bool, x int, y int, N int) bool {
	// horizontal
	for i := 0; i < N; i++ {
		if i != x && queens[k([]int{y, i})] {
			return false
		}
	}

	// vertical
	for i := 0; i < N; i++ {
		if i != y && queens[k([]int{i, x})] {
			return false
		}
	}

	// diagonal
	directions := [][]int{{-1, 1}, {1, -1}, {1, 1}, {-1, -1}}

	for _, tuple := range directions {
		diag_cnt_x := x
		diag_cnt_y := y

		for diag_cnt_x >= 0 && diag_cnt_x < N && diag_cnt_y >= 0 && diag_cnt_y < N {
			if inSet, ok := queens[k([]int{diag_cnt_y, diag_cnt_x})]; x != diag_cnt_x && y != diag_cnt_y && inSet && ok {
				return false
			}
			diag_cnt_x += tuple[0]
			diag_cnt_y += tuple[1]
		}
	}

	return true
}

func prettyPrint(N int, queens map[string]bool) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if inSet, ok := queens[k([]int{i, j})]; ok && inSet {
				fmt.Print("Q ")
			} else {
				fmt.Print("X ")
			}
		}
		fmt.Print("\n")
	}
}

func k(list []int) string {
	return fmt.Sprintf("%q", list)
}

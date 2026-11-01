package quests

import (
	"strconv"
	"strings"
)

type Quest5 struct{}

func (q *Quest5) getQuestId() int {
	return 5
}

type Point struct {
	row int
	col int
}

func (p *Point) String() string {
	return strconv.Itoa(p.row) + "," + strconv.Itoa(p.col)
}

func (q *Quest5) buildBoard(input string) (int, [][]rune) {
	board := make([][]rune, 0)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line != "" {
			board = append(board, []rune(line))
		}
	}

	startRow := len(board) / 2
	startCol := len(board[0]) / 2

	grandTotal := 0
	total := 1

	for total > 0 {
		total = 0
		stack := []Point{{row: startRow, col: startCol}}
		visited := make(map[string]bool)

		for len(stack) > 0 {
			if visited[stack[len(stack)-1].String()] {
				stack = stack[:len(stack)-1]
				continue
			}
			visited[stack[len(stack)-1].String()] = true

			r := stack[len(stack)-1].row
			c := stack[len(stack)-1].col
			stack = stack[:len(stack)-1]

			if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) {
				continue
			}

			if board[r][c] != '#' {
				count := 0
				if c+1 < len(board[0]) && (board[r][c+1] == '1' || board[r][c+1] == '#') {
					count++
				}
				if r+1 < len(board) && c+1 < len(board[0]) && (board[r+1][c+1] == '1' || board[r+1][c+1] == '#') {
					count++
				}
				if r+1 < len(board) && (board[r+1][c] == '1' || board[r+1][c] == '#') {
					count++
				}
				if r+1 < len(board) && c-1 >= 0 && (board[r+1][c-1] == '1' || board[r+1][c-1] == '#') {
					count++
				}
				if c-1 >= 0 && (board[r][c-1] == '1' || board[r][c-1] == '#') {
					count++
				}
				if r-1 >= 0 && c-1 >= 0 && (board[r-1][c-1] == '1' || board[r-1][c-1] == '#') {
					count++
				}
				if r-1 >= 0 && (board[r-1][c] == '1' || board[r-1][c] == '#') {
					count++
				}
				if r-1 >= 0 && c+1 < len(board[0]) && (board[r-1][c+1] == '1' || board[r-1][c+1] == '#') {
					count++
				}

				if count >= 5 {
					board[r][c] = '#'
					total++
				}
			}

			if board[r][c] == '#' {
				stack = append(stack, Point{row: r, col: c + 1})
				stack = append(stack, Point{row: r + 1, col: c + 1})
				stack = append(stack, Point{row: r + 1, col: c})
				stack = append(stack, Point{row: r + 1, col: c - 1})
				stack = append(stack, Point{row: r, col: c - 1})
				stack = append(stack, Point{row: r - 1, col: c - 1})
				stack = append(stack, Point{row: r - 1, col: c})
				stack = append(stack, Point{row: r - 1, col: c + 1})
			}
		}

		grandTotal += total
	}

	return grandTotal, board
}

func (q *Quest5) buildBeach(input string) (int, [][]rune) {
	_, board := q.buildBoard(input)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '#' {
				board[i][j] = '.'
			}
		}
	}

	grandTotal := 0
	total := 1

	for total > 0 {
		total = 0

		for r := 0; r < len(board); r++ {
			for c := 0; c < len(board[0]); c++ {
				if board[r][c] == '.' {
					continue
				}

				if board[r][c] == '#' {
					beach := false
					if c+1 >= len(board[0]) || board[r][c+1] == '.' {
						beach = true
					}
					if r+1 >= len(board) || c+1 >= len(board[0]) || board[r+1][c+1] == '.' {
						beach = true
					}
					if r+1 >= len(board) || board[r+1][c] == '.' {
						beach = true
					}
					if r+1 >= len(board) || c-1 < 0 || board[r+1][c-1] == '.' {
						beach = true
					}
					if c-1 < 0 || board[r][c-1] == '.' {
						beach = true
					}
					if r-1 < 0 || c-1 < 0 || board[r-1][c-1] == '.' {
						beach = true
					}
					if r-1 < 0 || board[r-1][c] == '.' {
						beach = true
					}
					if r-1 < 0 || c+1 >= len(board[0]) || board[r-1][c+1] == '.' {
						beach = true
					}

					count := 0

					if !beach {
						if c+1 < len(board[0]) && board[r][c+1] == '=' {
							count++
						}
						if r+1 < len(board) && c+1 < len(board[0]) && board[r+1][c+1] == '=' {
							count++
						}
						if r+1 < len(board) && board[r+1][c] == '=' {
							count++
						}
						if r+1 < len(board) && c-1 >= 0 && board[r+1][c-1] == '=' {
							count++
						}
						if c-1 >= 0 && board[r][c-1] == '=' {
							count++
						}
						if r-1 >= 0 && c-1 >= 0 && board[r-1][c-1] == '=' {
							count++
						}
						if r-1 >= 0 && board[r-1][c] == '=' {
							count++
						}
						if r-1 >= 0 && c+1 < len(board[0]) && board[r-1][c+1] == '=' {
							count++
						}
					}

					if beach || count >= 5 {
						board[r][c] = '='
						total++
					}
				}
			}
		}

		grandTotal += total
	}
	return grandTotal, board
}

func (q *Quest5) part1(input string) string {
	grandTotal, _ := q.buildBoard(input)

	return strconv.Itoa(grandTotal)
}

func (q *Quest5) part2(input string) string {
	grandTotal, _ := q.buildBeach(input)
	return strconv.Itoa(grandTotal)
}

func (q *Quest5) part3(input string) string {
	_, board := q.buildBeach(input)

	mask := [][]rune{
		{'.', '+', '+', '+', '.'},
		{'+', '+', '+', '+', '+'},
		{'+', '+', '+', '+', '+'},
		{'+', '+', '+', '+', '+'},
		{'.', '+', '+', '+', '.'},
	}

	grandTotal := 0

	for r := 2; r < len(board)-2; r++ {
		for c := 2; c < len(board[0])-2; c++ {

			valid := true

			for mr := -2; mr < len(mask)-2; mr++ {
				for mc := -2; mc < len(mask[0])-2; mc++ {
					if mask[mr+2][mc+2] == '.' {
						continue
					}
					if board[r+mr][c+mc] == '.' || board[r+mr][c+mc] == '=' {
						valid = false
						break
					}
				}
				if !valid {
					break
				}
			}

			if valid {
				board[r][c] = '@'
				grandTotal++
			}

		}
	}

	for i := 0; i < 2; i++ {
		newTrees := make([]Point, 0)
		for r := 2; r < len(board)-2; r++ {
			for c := 2; c < len(board[0])-2; c++ {
				if board[r][c] == '.' || board[r][c] == '=' || board[r][c] == '@' {
					continue
				}

				valid := true
				if c+1 < len(board[0]) && board[r][c+1] == '=' {
					valid = false
				}
				if r+1 < len(board) && c+1 < len(board[0]) && board[r+1][c+1] == '=' {
					valid = false
				}
				if r+1 < len(board) && board[r+1][c] == '=' {
					valid = false
				}
				if r+1 < len(board) && c-1 >= 0 && board[r+1][c-1] == '=' {
					valid = false
				}
				if c-1 >= 0 && board[r][c-1] == '=' {
					valid = false
				}
				if r-1 >= 0 && c-1 >= 0 && board[r-1][c-1] == '=' {
					valid = false
				}
				if r-1 >= 0 && board[r-1][c] == '=' {
					valid = false
				}
				if r-1 >= 0 && c+1 < len(board[0]) && board[r-1][c+1] == '=' {
					valid = false
				}

				if valid {
					count := 0

					if c+1 < len(board[0]) && board[r][c+1] == '@' {
						count++
					}
					if r+1 < len(board) && c+1 < len(board[0]) && board[r+1][c+1] == '@' {
						count++
					}
					if r+1 < len(board) && board[r+1][c] == '@' {
						count++
					}
					if r+1 < len(board) && c-1 >= 0 && board[r+1][c-1] == '@' {
						count++
					}
					if c-1 >= 0 && board[r][c-1] == '@' {
						count++
					}
					if r-1 >= 0 && c-1 >= 0 && board[r-1][c-1] == '@' {
						count++
					}
					if r-1 >= 0 && board[r-1][c] == '@' {
						count++
					}
					if r-1 >= 0 && c+1 < len(board[0]) && board[r-1][c+1] == '@' {
						count++
					}

					if count >= 4 {
						newTrees = append(newTrees, Point{row: r, col: c})
					}
				}
			}
		}

		for _, p := range newTrees {
			board[p.row][p.col] = '@'
			grandTotal++
		}
	}

	return strconv.Itoa(grandTotal)
}

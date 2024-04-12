package main

func exist(board [][]byte, word string) bool {
	pointer := 0
	var m, n = len(board), len(board[0])
	for i := 0; i < m; {
		for j := 0; j < n; {
			if board[i][j] == word[pointer] {
				pointer++
			} else {
				pointer = 0
			}

		}
	}
}

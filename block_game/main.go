package main

import "fmt"

// 坐标
type pos struct {
	row int
	col int
}

// 图形
type piece struct {
	pieceType pieceType
	rule      [][]byte
}

// 图形类别
type pieceType byte

const (
	L pieceType = iota + 1
	H
)

func main() {
	// 棋盘
	b := [][]byte{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	isFull := is_full(b, 1)
	fmt.Println(isFull)
	pos := pos{
		row: 1,
		col: 1,
	}
	lockDot := lock_dot(b, pos)
	fmt.Println(lockDot)
	// L 图形
	p := piece{pieceType: L, rule: [][]byte{{1, 0, 0}, {1, 0, 0}, {1, 1, 1}}}
	lockPiece := lock_piece(b, pos, p)
	fmt.Println(lockPiece)
}

func is_full(board [][]byte, row int) bool {
	if row >= len(board) || row < 0 {
		return false
	}
	for i := 0; i < len(board[row]); i++ {
		if board[row][i] == 0 {
			return false
		}
	}
	return true
}

func lock_dot(board [][]byte, pos pos) bool {
	if pos.row >= len(board) || pos.row < 0 || pos.col < 0 || pos.col >= len(board[0]) {
		return false
	}
	if board[pos.row][pos.col] == 0 {
		return true
	}
	return false
}

// lock_piece 逻辑类似，这里只实现了 L 图形
func lock_piece(board [][]byte, pos pos, piece piece) bool {
	switch piece.pieceType {
	case L:
		for i := 0; i < len(piece.rule); i++ {
			for j := 0; j < len(piece.rule[0]); j++ {
				// 找到图形在二维数组中的点，即[x,y]=1
				if piece.rule[i][j] == 0 {
					continue
				}
				// 判断与指定pos相对的位置坐标是否为0
				if pos.row+i >= len(board) || pos.col+j >= len(board[0]) || board[pos.row+i][pos.col+j] != 0 {
					return false
				}
			}
		}
		// 图形验证完毕返回true
		return true
	case H:

	}
	return false
}

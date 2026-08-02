package chessboard

type File []bool

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	column, exist := cb[file]
	count := 0

	if !exist {
		return count
	}

	for _, filled := range column {
		if filled {
			count++
		}
	}

	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	if rank > 8 || rank < 1 {
		return count
	}

	index := rank - 1

	for column := range cb {
		file := cb[column]
		filled := file[index]

		if filled {
			count++
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0

	for _, file := range cb {
		count += len(file)
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0

	for column := range cb {
		file := cb[column]

		for _, filled := range file {
			if filled {
				count++
			}
		}
	}

	return count
}

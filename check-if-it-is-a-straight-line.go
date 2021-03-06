package main

func checkStraightLine(coordinates [][]int) bool {
	for i := 0; i < len(coordinates)-2; i++ {
		if (coordinates[i+2][1]-coordinates[i+1][1])*(coordinates[i+1][0]-coordinates[i][0]) !=
			(coordinates[i+2][0]-coordinates[i+1][0])*(coordinates[i+1][1]-coordinates[i][1]) {
			return false
		}
	}

	return true
}

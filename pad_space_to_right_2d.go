package strc

// PadSpaceToRight2D takes a 2D Array of strings, iterates through them, and adds spaces to right of corresponding element of the inside array that would make it equal to the biggest element on that position in the sequence of arrays
func PadSpaceToRight2D(twoDArr [][]string) [][]string {
	// check the length of all the inner arrays and get the length of the biggest array
	innerArrayWithMostElemsLen := 0
	for _, innerArr := range twoDArr {
		if len(innerArr) > innerArrayWithMostElemsLen {
			innerArrayWithMostElemsLen = len(innerArr)
		}
	}
	// make a map with index as key and length of each element in the innerArrays as value
	lenMapAtIndex := make(map[int]int)

	// loop over each inner array and update lemMapIndex to highest length among all
	for _, innerArr := range twoDArr {
		for i, str := range innerArr {
			// update the value  at key if it is less. Do not worry if key exists or not
			if lenMapAtIndex[i] < len(str) {
				lenMapAtIndex[i] = len(str)
			}
		}
	}
	// run through all strings in all innerArr and pad each string according to map
	outArr := [][]string{}
	for _, innerArr := range twoDArr {
		outArrInnerArr := []string{}
		for i, str := range innerArr {
			outArrInnerArr = append(outArrInnerArr, PadSpaceToRight(str, lenMapAtIndex[i]))
		}
		outArr = append(outArr, outArrInnerArr)
	}
	return outArr
}

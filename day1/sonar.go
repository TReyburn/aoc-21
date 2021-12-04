package day1

func ScanIncreases(start *Node) int {
	var countIncreases int
	var curr, next *Node

	curr = start

	for {
		startDepth := curr.Value()
		next = curr.Next()
		if next == nil {
			break
		}
		nextDepth := next.Value()
		if startDepth < nextDepth {
			countIncreases++
		}

		curr = next
	}
	return countIncreases
}

func ScanExtendedIncreases(start *Node) int {
	var countIncreases int
	var curr, next *Node

	curr = start

	for {
		next = curr.Next()
		if next == nil {
			break
		}
		nn := next.Next()
		if nn == nil {
			break
		}
		nnn := nn.Next()
		if nnn == nil {
			break
		}
		startDepth := curr.Value() + next.Value() + nn.Value()
		nextDepth := next.Value() + nn.Value() + nnn.Value()
		if nextDepth > startDepth {
			countIncreases++
		}
		curr = next
	}
	return countIncreases
}

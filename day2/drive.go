package day2

const (
	forward = "forward"
	up      = "up"
	down    = "down"
)

func DriveSub(instructions []Instruction) (int, int) {
	var hpos, depth int

	controls := map[string]func(int){
		forward: func(x int) { hpos += x },
		up:      func(x int) { depth -= x },
		down:    func(x int) { depth += x },
	}

	for _, cmd := range instructions {
		proc, ok := controls[cmd.method]
		if ok {
			proc(cmd.value)
		}
	}

	return hpos, depth
}

func DriveSubWithAim(instructions []Instruction) (int, int) {
	var hpos, depth, aim int

	controls := map[string]func(int){
		forward: func(x int) { hpos += x; depth += x * aim },
		up:      func(x int) { aim -= x },
		down:    func(x int) { aim += x },
	}

	for _, cmd := range instructions {
		proc, ok := controls[cmd.method]
		if ok {
			proc(cmd.value)
		}
	}

	return hpos, depth
}

func DriveOutcome(x int, y int) int {
	return x * y
}

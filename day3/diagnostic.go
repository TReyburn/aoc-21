package day3

import "strconv"

func RunPowerDiagnostic(binArr []string) (string, string) {
	var gammaRate, epsilonRate string

	l := len(binArr[0])

	// O(n * m)
	for i := 0; i < l; i++ {
		var countOne, countZero int
		for _, bin := range binArr {
			if string(bin[i]) == "1" {
				countOne++
			} else {
				countZero++
			}
		}
		if countOne > countZero {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	return gammaRate, epsilonRate
}

func CalculatePower(gamma, epsilon string) int64 {
	gammDec, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return -1
	}

	epDec, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return -1
	}

	return gammDec * epDec
}

func FindLifeSupportRatings(binArr []string) (string, string) {
	var o2rating, co2rating string
	var o2searchItems, co2searchItems []string

	l := len(binArr[0])
	o2chan := make(chan string)
	co2chan := make(chan string)

	// O(n)
	if l > 0 {
		var countZero, countOne int
		zeroStart := make([]string, 0)
		oneStart := make([]string, 0)
		for _, bin := range binArr {
			if string(bin[0]) == "1" {
				countOne++
				oneStart = append(oneStart, bin)
			} else {
				countZero++
				zeroStart = append(zeroStart, bin)
			}
		}
		if countOne >= countZero {
			o2searchItems = oneStart
			co2searchItems = zeroStart
		} else {
			o2searchItems = zeroStart
			co2searchItems = oneStart
		}
	}

	go func() {
		// O(n * (m-1))
		for i := 1; i < l; i++ {
			if len(o2searchItems) == 1 {
				break
			}
			var countZero, countOne int
			zeroStart := make([]string, 0)
			oneStart := make([]string, 0)
			for _, bin := range o2searchItems {
				if string(bin[i]) == "1" {
					countOne++
					oneStart = append(oneStart, bin)
				} else {
					countZero++
					zeroStart = append(zeroStart, bin)
				}
			}
			if countOne >= countZero {
				o2searchItems = oneStart
			} else {
				o2searchItems = zeroStart
			}
		}
		o2chan <- o2searchItems[0]
		close(o2chan)
	}()

	go func() {
		for i := 1; i < l; i++ {
			if len(co2searchItems) == 1 {
				break
			}
			var countZero, countOne int
			zeroStart := make([]string, 0)
			oneStart := make([]string, 0)
			for _, bin := range co2searchItems {
				if string(bin[i]) == "1" {
					countOne++
					oneStart = append(oneStart, bin)
				} else {
					countZero++
					zeroStart = append(zeroStart, bin)
				}
			}
			if countOne >= countZero {
				co2searchItems = zeroStart
			} else {
				co2searchItems = oneStart
			}
		}
		co2chan <- co2searchItems[0]
		close(co2chan)
	}()

	o2rating = <-o2chan
	co2rating = <-co2chan
	return o2rating, co2rating
}

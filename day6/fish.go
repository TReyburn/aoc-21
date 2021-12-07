package day6

type School struct {
	initTimer  int
	resetTimer int
	school     map[int]int
}

func NewSchool(input []int) *School {
	s := make(map[int]int)
	for i := 0; i < 9; i++ {
		s[i] = 0
	}

	school := &School{
		initTimer:  8,
		resetTimer: 6,
		school:     s,
	}

	school.Seed(input...)
	return school
}

func (s *School) Seed(fish ...int) {
	for _, n := range fish {
		s.school[n]++
	}
}

func (s *School) Day() {
	var newFish int
	newSchool := make(map[int]int)
	for k, v := range s.school {
		k--
		if k < 0 {
			newFish = v
		} else {
			newSchool[k] = v
		}
	}
	newSchool[6] += newFish
	newSchool[8] = newFish
	s.school = newSchool
}

func (s *School) CountFish() int {
	var fishCount int
	for _, v := range s.school {
		fishCount += v
	}
	return fishCount
}

func (s *School) ProcessDays(dayCount int) {
	for i := 0; i < dayCount; i++ {
		s.Day()
	}
}

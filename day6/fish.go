package day6

type Fish struct {
	count int
}

func NewFish(v int) *Fish {
	return &Fish{count: v}
}

type School struct {
	initTimer  int
	resetTimer int
	school     map[int]*Fish
}

func NewSchool(input []int) *School {
	s := make(map[int]*Fish)
	for i := 0; i < 9; i++ {
		s[i] = NewFish(0)
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
		s.school[n].count++
	}
}

func (s *School) Day() {
	//tba := make([]int, 0)
	//
	//for idx, fish := range s.school {
	//	if fish == 0 {
	//		tba = append(tba, s.initTimer)
	//		s.school[idx] = s.resetTimer
	//	} else {
	//		s.school[idx]--
	//	}
	//}
	//
	//s.Seed(tba...)
	var newFish int
	newSchool := make(map[int]*Fish)
	for k, v := range s.school {
		k--
		if k < 0 {
			newFish = v.count
		} else {
			newSchool[k] = v
		}
	}
	newSchool[6].count += newFish
	newSchool[8] = NewFish(newFish)
	s.school = newSchool
}

func (s *School) CountFish() int {
	var fishCount int
	for _, v := range s.school {
		fishCount += v.count
	}
	return fishCount
}

func (s *School) ProcessDays(dayCount int) {
	for i := 0; i < dayCount; i++ {
		s.Day()
	}
}

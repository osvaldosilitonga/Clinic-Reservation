package main

func FindSequence(main, seq []int) bool {
	for i := 0; i < len(main); i++ {
		if i == len(main)-1 {
			return false
		}

		if main[i] == seq[0] && main[i+1] == seq[1] {
			return true
		}
	}

	return false
}

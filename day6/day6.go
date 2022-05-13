package day6

// CustomsGroup wraps operations on a customs form.
type CustomsGroup struct {
	rows       []string
	fields     map[string]bool
	occurences map[string]int
}

// CountAnyone counts the questions marked "Yes" in the customs form.
func (cg *CustomsGroup) CountAnyone() int {
	cg.fields = make(map[string]bool)
	for _, row := range cg.rows {
		for _, question := range row {
			cg.fields[string(question)] = true
		}
	}
	return len(cg.fields)
}

// CountEveryone counts the questions marked "Yes" in the customs form
// that everyone has answered positively.
func (cg *CustomsGroup) CountEveryone() int {
	cg.occurences = make(map[string]int)
	for _, row := range cg.rows {
		for _, question := range row {
			cg.occurences[string(question)]++
		}
	}
	var count int
	for _, v := range cg.occurences {
		if v == len(cg.rows) {
			count++
		}
	}
	return count
}

// Add every row from a group customs form questionnaire.
func (cg *CustomsGroup) Add(row string) {
	cg.rows = append(cg.rows, row)
}

// CountYes takes a list of inputs denoting an "yes"
// answer to a questionnaire (questions a-z).
func CountYes(input []string, kind string) int {
	groups := []int{}
	group := CustomsGroup{}
	for _, row := range input {
		if row == "" {
			if kind == "anyone" {
				groups = append(groups, group.CountAnyone())
			} else {
				groups = append(groups, group.CountEveryone())
			}
			group = CustomsGroup{}
		} else {
			group.Add(row)
		}
	}
	if kind == "anyone" {
		groups = append(groups, group.CountAnyone())
	} else {
		groups = append(groups, group.CountEveryone())
	}
	return sum(groups)
}

func sum(groups []int) int {
	var total int
	for _, group := range groups {
		total += group
	}
	return total
}

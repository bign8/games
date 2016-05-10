package games

import "sort"

// Utility gives a numeric value to a given state
type Utility func(State) int

// Sort orders a series of actions based on a utility function
func (u Utility) Sort(start State) []Action {
	list := start.Actions()
	values := make([]int, len(list))
	for i, action := range list {
		values[i] = u(start.Apply(action))
	}
	sorter := &utilitySorter{start, list, values}
	sort.Sort(sorter)
	return list
}

type utilitySorter struct {
	start   State
	actions []Action
	values  []int
}

func (us *utilitySorter) Len() int {
	return len(us.actions)
}

func (us *utilitySorter) Less(i, j int) bool {
	return us.values[i] < us.values[j]
}

func (us *utilitySorter) Swap(i, j int) {
	us.actions[i], us.actions[j] = us.actions[j], us.actions[i]
	us.values[i], us.values[j] = us.values[j], us.values[i]
}

// UtilityCascade breaks ties with next utility in the list
func UtilityCascade(list ...Utility) Utility {
	// TODO: implement this
	return func(x State) int {
		return 0
	}
}

// UtilityCombined orders states based on an overall ranking between utilities
func UtilityCombined(list ...Utility) Utility {
	// TODO: implement this
	return func(x State) int {
		return 0
	}
}

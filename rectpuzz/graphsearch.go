package rectpuzz

type noRouteError struct {
}

func (nr noRouteError) Error() string {
	return "could not find a route to target"
}

func SearchGraph(intGrid [][]uint) ([]Rect, error) {
	return linearSearch(NewState(&intGrid), new([]State))
}

func linearSearch(s State, visited *[]State) ([]Rect, error) {
	if s.IsSolved() {
		return s.Rects(), nil
	}

	for _, visitedState := range *visited {
		if StateEquals(&s, &visitedState) {
			return nil, noRouteError{}
		}
	}

	neighbours, err := s.Neighbours()
	if err != nil {
		return nil, err
	}

	for _, neighState := range neighbours {
		result, err := linearSearch(neighState, visited)

		if err != nil {
			switch err.(type) {
			case noRouteError:
				continue // Proceed...
			default:
				return nil, err
			}
		} else {
			return result, nil
		}
	}

	// 'neighbours' is either empty, or all neighbours were already visited.
	*visited = append(*visited, s)
	*visited = append(*visited, neighbours...)
	return nil, noRouteError{}
}
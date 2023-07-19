package mergeintervals

type Interval struct {
	start, end int
}

func MergeIntervals(intervals []Interval, newInterval Interval) []Interval {
	var result []Interval

	// i want to solve it, dont give suggestions
	for _, inter := range intervals {
		if inter.end < newInterval.start {
			result = append(result, inter)
		} else if newInterval.end < inter.start {
			result = append(result, newInterval)
			newInterval = inter
		} else {
			newInterval.start = min(newInterval.start, inter.start)
			newInterval.end = max(newInterval.end, inter.end)
		}
	}
	result = append(result, newInterval)
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

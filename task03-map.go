package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	sortedKeys := make([]int, 0, len(input))

	for k := range input {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)

	sortedValues := make([]string, 0, len(input))
	for k := range sortedKeys {
		sortedValues = append(sortedValues, input[k])
	}

	return sortedValues
}

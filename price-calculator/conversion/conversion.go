package conversion

import "strconv"

func StringsToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for index, s := range strings {
		var err error
		floats[index], err = strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
	}
	return floats, nil
}

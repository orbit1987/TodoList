package tools

func MapHasOnlyValidData(data map[string]interface{}, validation ...string) bool {
	validationMap := make(map[string]interface{})
	for k, v := range data {
		validationMap[k] = v
	}

	for _, v := range validation {
		delete(validationMap, v)
	}

	return len(validationMap) > 0
}

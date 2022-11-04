package catalogmanagement

func SIToSS(i []interface{}) []string {
	var ss []string
	for _, iface := range i {
		ss = append(ss, iface.(string))
	}
	return ss
}

// Contains checks if a string is present in a slice
func Contains(list []string, val string) (contains bool, index int) {
	for i, v := range list {
		if v == val {
			return true, i
		}
	}
	return false, -1
}

func AppendUnique(list []string, val string) []string {
	if list == nil {
		return []string{val}
	}

	// Is item in list?
	contains, _ := Contains(list, val)

	if contains {
		return list
	}

	return append(list, val)
}

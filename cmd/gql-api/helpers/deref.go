package helpers

func SafeDerefString(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

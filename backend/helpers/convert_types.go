package helpers

func String(v string) *string {
	return &v
}

func StringValue(v *string) string {
	if v != nil {
		return *v
	}

	return ""
}

func Int(v int) *int {
	return &v
}

func IntValue(v *int) int {
	if v != nil {
		return *v
	}

	return 0
}

func Int32(v int32) *int32 {
	return &v
}

func Int32Value(v *int32) int32 {
	if v != nil {
		return *v
	}
	return 0
}

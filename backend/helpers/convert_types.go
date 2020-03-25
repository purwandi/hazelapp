package helpers

// String is to convert string into string pointer
func String(v string) *string {
	return &v
}

// StringValue is to convert string pointer into string
func StringValue(v *string) string {
	if v != nil {
		return *v
	}

	return ""
}

// Int is to convert int into int pointer
func Int(v int) *int {
	return &v
}

// IntValue is to convert int pointer into int
func IntValue(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

// Int32ValueToInt is to convert int32 pointer into int pointer
func Int32ValueToInt(v *int32) *int {
	if v == nil {
		return nil
	}

	val := int(*v)
	return &val
}

// IntValueToInt32 is to convert int pointer into int pointer
func IntValueToInt32(v *int) *int32 {
	if v == nil {
		return nil
	}

	val := int32(*v)
	return &val
}

// IntToInt32 to convert int into in32
func IntToInt32(v int) int32 {
	return int32(v)
}

// Int32 is to convert int32 into int32 pointer
func Int32(v int32) *int32 {
	return &v
}

// Int32Value is to convert int32 pointer into int32
func Int32Value(v *int32) int32 {
	if v == nil {
		return 0
	}

	return *v
}

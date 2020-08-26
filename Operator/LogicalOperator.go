package operator

// && And 同时为True则为True，否则为False
// || Or 有一个为True则为True，否则为False
// ！ Not 值为True则为False， 值为False则为True

var True_ bool = true
var False_ bool = false

// And 解释And的作用
func And(x, y bool) string {
	if x && y {
		return "AND"
	} else {
		return ""
	}
}

// Or 解释or的运作
func Or(x, y bool) string {
	if x || y {
		return "OR"
	}
	return ""
}

// Not 解释Not的运作
func Not(x bool) string {
	if !x {
		return "Not"
	} else {
		return ""
	}
}

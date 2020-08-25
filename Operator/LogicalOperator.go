package operator

// && And 同时为True则为True，否则为False
// || Or 有一个为True则为True，否则为False
// ！ Not 值为True则为False， 值为False则为True

var True_ bool = true
var Fasle_ bool = false

func And(x, y bool) string {
	if x && y {
		return "AND"
	} else {
		return ""
	}
}

func Or(x, y bool) string {
	if x || y {
		return "OR"
	} else {
		return ""
	}
}

func Not(x bool) string {
	if !x {
		return "Not"
	} else {
		return ""
	}
}

// Code generated by "stringer -type=PositionAction"; DO NOT EDIT.

package xfunding

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PositionNoOp-0]
	_ = x[PositionOpening-1]
	_ = x[PositionClosing-2]
}

const _PositionAction_name = "PositionNoOpPositionOpeningPositionClosing"

var _PositionAction_index = [...]uint8{0, 12, 27, 42}

func (i PositionAction) String() string {
	if i < 0 || i >= PositionAction(len(_PositionAction_index)-1) {
		return "PositionAction(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PositionAction_name[_PositionAction_index[i]:_PositionAction_index[i+1]]
}

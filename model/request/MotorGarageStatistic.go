package request

type MotorGarageStatisticGet struct {
	From  int64  `form:"from"`
	To    int64  `form:"to"`
	Range string `form:"range"` // day, month
}

package common

type FeesModelName string

const (
	Mall    FeesModelName = "mall"
	Stadium FeesModelName = "stadium"
	Airport FeesModelName = "airport"
)

const (
	KeyPerHourRate = "per_hour"
	KeyPerDay      = "per_day"
	// Stadium
	Key_0_4_Hrs    = "0_4_hrs"
	Key_4_12_Hrs   = "4_12_hrs"
	Key_12_Inf_Hrs = "12_inf_hrs"
	// Airport
	Key_0_1_Hrs   = "0_1_hrs"
	Key_1_8_Hrs   = "1_8_hrs"
	Key_8_24_Hrs  = "8_24_hrs"
	Key_0_12_Hrs  = "0_12_hrs"
	Key_12_24_Hrs = "12_24_hrs"
)

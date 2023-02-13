package common

type FeesModelName string

const (
	Mall    FeesModelName = "mall"
	Stadium FeesModelName = "stadium"
	Airport FeesModelName = "airport"
)

type FeeSlabType string

const (
	PerHourSlab = "per_hour"
	PerDaySlab  = "per_day"
	FlatSlab    = "flat"
)

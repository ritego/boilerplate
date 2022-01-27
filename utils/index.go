package utils

const (
	ALL_CLASS          = "all"
	HAS_INTEREST_CLASS = "has_interest"
	NO_INTEREST_CLASS  = "no_interest"
)

const (
	DAILY_PERIOD    = "daily"
	WEEKLY_PERIOD   = "weekly"
	MONTHLY_PERIOD  = "monthly"
	QUATERLY_PERIOD = "quaterly"
	YEARLY_PERIOD   = "yearly"
)

func InSlice(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
		println(needle, v)
	}
	return false
}

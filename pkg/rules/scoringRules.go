package rules

const (
	Owner     = 4.5
	OpenSimi  = 26
	CloseSimi = 1.5
)

type claims struct {
	value      float64
	percentage float64
}

func ClaimsRulesDefinitions(percentage float64) float64 {
	rules := []claims{
		{
			value:      0,
			percentage: 0,
		},
		{
			value:      -20,
			percentage: 10,
		},
		{
			value:      -30,
			percentage: 20,
		},
		{
			value:      -50,
			percentage: 40,
		},
		{
			value:      -70,
			percentage: 60,
		},
		{
			value:      -90,
			percentage: 100,
		},
	}

	for _, item := range rules {
		if percentage <= item.percentage {
			return item.value
		}
	}
	return 0
}

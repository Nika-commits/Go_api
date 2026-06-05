package tools

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"pranish": {
		AuthToken: "1234AVD",
		Username:  "pranish123",
	},
	"Anushree": {
		AuthToken: "qw341@",
		Username:  "anu",
	},
	"Smith": {
		AuthToken: "qwerty",
		Username:  "smith",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"pranish": {
		Coins:    1000,
		Username: "pranish123",
	},
	"Anushree": {
		Coins:    500,
		Username: "anu",
	},
	"Smith": {
		Coins:    2000,
		Username: "smith",
	},
}



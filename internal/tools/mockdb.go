package tools

import "time"

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

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clienData = LoginDetails{}
	clienData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}
	return &clienData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var coinData = CoinDetails{}
	coinData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}
	return &coinData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

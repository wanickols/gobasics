package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"emma": {
		AuthToken: "456DEF",
		Username:  "emma",
	},
	"jack": {
		AuthToken: "789GHI",
		Username:  "jack",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"emma": {
		Coins:    50,
		Username: "emma",
	},
	"jack": {
		Coins:    200,
		Username: "jack",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var accountData = CoinDetails{}
	accountData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &accountData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

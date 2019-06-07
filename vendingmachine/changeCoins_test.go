package vendingmachine

import (
	"testing"
)

func TestTotalBalanceIs_1_ShouldChangeCoin_O(t *testing.T) {
	expectedResult := "O"
	v := NewVendingMachine()
	v.totalCoins = 1

	actualResult := v.changeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestTotalBalanceIs_2_ShouldChangeCoin_TW(t *testing.T) {
	expectedResult := "TW"
	v := NewVendingMachine()
	v.totalCoins = 2

	actualResult := v.changeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestTotalBalanceIs_5_ShouldChangeCoin_F(t *testing.T) {
	expectedResult := "F"
	v := NewVendingMachine()
	v.totalCoins = 5

	actualResult := v.changeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestTotalBalanceIs_10_ShouldChangeCoin_T(t *testing.T) {
	expectedResult := "T"
	v := NewVendingMachine()
	v.totalCoins = 10

	actualResult := v.changeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestTotalBalanceIs_7_ShouldChangeCoin_F_TW(t *testing.T) {
	expectedResult := "F, TW"
	v := NewVendingMachine()
	v.totalCoins = 7

	actualResult := v.changeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestTotalBalanceIs_4_ShouldChangeCoin_TW_TW(t *testing.T) {
	expectedResult := "TW, TW"
	v := NewVendingMachine()
	v.totalCoins = 4

	actualResult := v.changeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestTotalBalanceIs_25_ShouldChangeCoin_T_T_F(t *testing.T) {
	expectedResult := "T, T, F"
	v := NewVendingMachine()
	v.totalCoins = 25

	actualResult := v.changeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestTotalBalanceIs_27_ShouldChangeCoin_T_T_F_TW(t *testing.T) {
	expectedResult := "T, T, F, TW"
	v := NewVendingMachine()
	v.totalCoins = 27

	actualResult := v.changeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

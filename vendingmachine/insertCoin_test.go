package vendingmachine_test

import (
	"testing"
	vm "simple-web/vendingmachine"
)

func TestInsertCoinWith_T_F_TW_O_MachineShouldHave_18_Bath(t *testing.T) {
	v := vm.NewVendingMachine()
	expectedResult := 18

	v.InsertCoins("T")
	v.InsertCoins("F")
	v.InsertCoins("TW")
	v.InsertCoins("O")
	actualResult := v.ShowTotalBalance()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}

}

func TestInsertCoinWith_T_MachineShouldHave_10_Bath(t *testing.T) {
	v := vm.NewVendingMachine()
	expectedResult := 10

	v.InsertCoins("T")
	actualResult := v.ShowTotalBalance()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestInsertCoinWith_F_MachineShouldHave_5_Bath(t *testing.T) {
	v := vm.NewVendingMachine()
	expectedResult := 5

	v.InsertCoins("F")
	actualResult := v.ShowTotalBalance()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}

}

func TestInsertCoinWith_T_F_MachineShouldHave_15_Bath(t *testing.T) {
	v := vm.NewVendingMachine()
	expectedResult := 15

	v.InsertCoins("T")
	v.InsertCoins("F")
	actualResult := v.ShowTotalBalance()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

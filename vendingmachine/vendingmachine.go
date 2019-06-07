package vendingmachine

type VendingMachine struct {
	totalCoins int
	m          map[string]int
}

func (v VendingMachine) ShowTotalBalance() int {
	return v.totalCoins
}

func (v *VendingMachine) InsertCoins(coin string) {
	elem := v.m[coin]
	v.totalCoins += elem
}

func NewVendingMachine() VendingMachine {
	m := make(map[string]int)
	m["T"] = 10
	m["F"] = 5
	m["TW"] = 2
	m["O"] = 1
	return VendingMachine{m: m}
}

func (v *VendingMachine) BuyDrink(drink string) string {
	m := make(map[string]int)
	m["SD"] = 18
	m["CC"] = 12
	m["DW"] = 7

	price, ok := m[drink]

	if !ok {
		return "No Item"
	}
	if v.totalCoins < price {
		return "Add more money"
	}
	v.totalCoins -= price

	if v.totalCoins == 0 {
		return drink
	}
	return drink + ", " + v.changeCoins()
}

func (v *VendingMachine) changeCoins() string {
	var changeCoin string
	coinsValue := [4]int{10, 5, 2, 1}
	coinsText := [4]string{"T", "F", "TW", "O"}
	for counter := 0; counter < len(coinsValue); counter++ {
		for v.totalCoins >= coinsValue[counter] {
			changeCoin += coinsText[counter]
			v.totalCoins -= coinsValue[counter]
			if v.totalCoins != 0 {
				changeCoin += ", "
			}
		}
	}
	return changeCoin
}

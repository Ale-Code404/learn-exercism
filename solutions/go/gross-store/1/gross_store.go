package gross

var units = map[string]int{
	"quarter_of_a_dozen": 3,
	"half_of_a_dozen":    6,
	"dozen":              12,
	"small_gross":        120,
	"gross":              144,
	"great_gross":        1728,
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, exists := units[unit]
	if !exists {
		return false
	}

	current := bill[item]
	bill[item] = current + quantity

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemQuantity, existsInBill := bill[item]
	if !existsInBill {
		return false
	}

	unitQuantity, existsUnit := units[unit]
	if !existsUnit {
		return false
	}

	quantity := itemQuantity - unitQuantity
	if quantity < 0 {
		return false
	}

	if quantity == 0 {
		delete(bill, item)
		return true
	}

	bill[item] = quantity

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	if !exists {
		return 0, false
	}

	return quantity, true
}

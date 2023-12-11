package funcMap

import (
	...
)

func Units() map[string]int {
	data := make(map[string]int)

	

	data"quarter_of_a_dozen" =	3
	data"half_of_a_dozen" =	6
	data"dozen" =	12
	data"small_gross" =	120
	data"gross" =	144
	data"great_gross" =	1728
	return data
}

func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

func AddItem(bill, units map[string]int, item, unit string) bool {
	bill = NewBill()
	units = Units()

	if unitElement, ok := units[unit], !ok {

		return ok
 
	}

	if itemElement, ok := bill[item], ok {

		bill[item]

		return ok

	}


}



/*

// AddItem adds an item to customer bill.


To implement this, you'll need to:

Return false if the given unit is not in the units map.
Otherwise add the item to the customer bill, indexed by the item name, then return true.
If the item is already present in the bill, increase its quantity by the amount that belongs to the provided unit.
bill := NewBill()
units := Units()
ok := AddItem(bill, units, "carrot", "dozen")
fmt.Println(ok)
// Output: true (since dozen is a valid unit)

*/
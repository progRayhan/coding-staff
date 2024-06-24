package functions

import "go_staff/structs"

func MyAddress(adrs structs.Address) structs.Address{
	adrs.Name = "Dhaka"
	adrs.Street = "Bashabo"
	adrs.City = "Dhaka"
	// adrs.State = "KPB"

	return adrs
}

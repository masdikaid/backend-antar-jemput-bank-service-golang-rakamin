package seeds

import (
	"backend-a-antar-jemput/internal/databases"
	"backend-a-antar-jemput/internal/entities"
)

func RunSeeds() {
	logins := []*entities.Login{
		{},
		{},
		{},
		{},
	}

	databases.Load()
	databases.DBCon.Create(&logins)

	customers := []*entities.Customers{
		{Login: *logins[0], Users: entities.Users{Name: "Dika", PhoneNumber: "085771002550"}},
		{Login: *logins[1], Users: entities.Users{Name: "Rudi", PhoneNumber: "085771002551"}},
	}

	databases.DBCon.Create(&customers)

	locations := []*entities.Location{
		{Login: *logins[0], City: "jakarta"},
		{Login: *logins[1], City: "Bekasi"},
		{Login: *logins[2], City: "Bogor"},
		{Login: *logins[3], City: "Palembang"},
	}

	databases.DBCon.Create(&locations)

	agents := []*entities.Agents{
		{Login: *logins[2], Location: *locations[0], OutletName: "Berkah Link", Users: entities.Users{Name: "Firman", PhoneNumber: "085771002552"}},
		{Login: *logins[3], Location: *locations[1], OutletName: "Jaya ATM Link", Users: entities.Users{Name: "Bagong", PhoneNumber: "085771002553"}},
	}

	databases.DBCon.Create(&agents)

	// gorm.Model
	// 	CustomersID uint
	// 	Customers   Customers
	// 	AgentsID    uint
	// 	Agents      Agents
	// 	LocationID  uint
	// 	Location    Location
	// 	Tipe        string
	// 	Amount      int
	// 	Status      int
	// 	Rating      float64
	transactions := []*entities.Transaction{
		{Customers: *customers[0], Agents: *agents[1], Location: *locations[0], Tipe: "dummy", Amount: 10000, Status: 1},
		{Customers: *customers[1], Agents: *agents[0], Location: *locations[0], Tipe: "dummy 2", Amount: 30000, Status: 0},
		{Customers: *customers[1], Agents: *agents[1], Location: *locations[0], Tipe: "dummy 3", Amount: 300000, Status: 3},
	}
	databases.DBCon.Create(transactions)
}

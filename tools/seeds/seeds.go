package seeds

import (
	"backend-a-antar-jemput/internal/databases"
	"backend-a-antar-jemput/internal/entities"
)

func RunSeeds() {
	logins := []*entities.Login{
		{Username: "masdika", Password: "rakamin00", LoginAs: 1},
		{Username: "rudianto", Password: "rakamin00", LoginAs: 1},
		{Username: "firman", Password: "rakamin00", LoginAs: 2},
		{Username: "bagong", Password: "rakamin00", LoginAs: 2},
	}

	databases.DBCon.Create(&logins)

	customers := []*entities.Customers{
		{Login: *logins[0], Users: entities.Users{Name: "Dika", PhoneNumber: "085771002550"}},
		{Login: *logins[1], Users: entities.Users{Name: "Rudi", PhoneNumber: "085771002551"}},
	}

	databases.DBCon.Create(&customers)

	locations := []*entities.Location{
		{Login: *logins[2], City: "jakarta", District: "Jakarta Kecamatan"},
		{Login: *logins[3], City: "Bekasi", District: "Bekasi Kecamatan"},
	}

	databases.DBCon.Create(&locations)

	services := []*entities.Services{
		{ServiceName: "Laku Pandai", Description: "Tarik Tunai"},
		{ServiceName: "Mini ATM BRI", Description: "Setor-Pasti"},
		{ServiceName: "Tunai", Description: "Setoran Pinjaman"},
	}

	databases.DBCon.Create(&services)

	agents := []*entities.Agents{
		{Login: *logins[2], Location: *locations[0], OutletName: "Berkah Link", Users: entities.Users{Name: "Firman", PhoneNumber: "085771002552"}, Services: []*entities.Services{services[0], services[1], services[2]}, MaxTrx: 25000000, IsAvailable: true},
		{Login: *logins[3], Location: *locations[1], OutletName: "Jaya ATM Link", Users: entities.Users{Name: "Bagong", PhoneNumber: "085771002553"}, Services: []*entities.Services{services[2]}, MaxTrx: 25000000, IsAvailable: true},
	}

	databases.DBCon.Create(&agents)

	transactions := []*entities.Transaction{
		{Customers: *customers[0], Agents: *agents[1], Address: "jl asem", Province: "Jakarta", City: "Jakarta Pusat", District: "Mangga Besar", Latitude: 1.2, Longitude: 2.2, Services: *services[0], Type: "dummy", Amount: 10000, Status: 1},
		{Customers: *customers[1], Agents: *agents[0], Address: "jl asem", Province: "Jakarta", City: "Jakarta Pusat", District: "Mangga Besar", Latitude: 1.2, Longitude: 2.2, Services: *services[0], Type: "dummy 2", Amount: 30000, Status: 0},
		{Customers: *customers[1], Agents: *agents[1], Address: "jl asem", Province: "Jakarta", City: "Jakarta Pusat", District: "Mangga Besar", Latitude: 1.2, Longitude: 2.2, Services: *services[0], Type: "dummy 3", Amount: 300000, Status: 3},
	}
	databases.DBCon.Create(transactions)
}

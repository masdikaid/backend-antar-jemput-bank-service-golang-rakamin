package contract

type Location struct {
}

type Province struct {
	Id   string
	Name string
}

type City struct {
	Id          string
	Province_id string
	Name        string
}

type District struct {
	Id         string
	Regency_id string
	Name       string
}

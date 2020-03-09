package main

type userInfo struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  company `json:"company"`
}

type address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     geo    `json:"geo"`
}

type geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type comment struct {
	UserID int    `json:"userId"`
	ID     int    `json:"Id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type userFullInfo struct {
	UserInfo userInfo  `json:"userInfo"`
	Comments []comment `json:"comments"`
}

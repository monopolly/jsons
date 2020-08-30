package main

type domain struct {
	id           int
	user         int
	added        int64
	up           int64
	name         string   //compaby.com
	zone         string   //com
	start        []string //words
	contain      []string
	end          []string
	verify       bool //domain author, dns...
	len          int
	pattern      []string
	registered   int64
	expire       int64
	installments int  //12,24,36 month
	price        int  //1,000,000
	offer        bool //true
	status       int  //active, hidden, ban, transfer, sold
}

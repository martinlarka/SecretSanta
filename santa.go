package main

type Santa struct {
	Id		int       `json:"id"`
	Name      	string    `json:"name"`
	Phone	  	string    `json:"phone"`
	Selectable	[]int
	Selected	int
	Exclude   	[]int     `json:"exclude"`
	//Shit	bool	`json:"selected"`
}

type Santas []Santa

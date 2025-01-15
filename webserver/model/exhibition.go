package model

type Exhibition struct {
	Id          int
	Title       string
	Description string
}

// Counter to keep track of the next ID
var currentID = 3

func GetAll() []Exhibition {
	return list
}

func GetById(id int) Exhibition {
	for _, item := range list {
		if item.Id == id {
			return item
		}
	}
	return Exhibition{}
}

func AddData(exe Exhibition) {
	// Increment the ID counter
	currentID++
	exe.Id = currentID
	list = append(list, exe)
}

var list = []Exhibition{
	{
		Id:          1,
		Title:       "Coding in Go",
		Description: "Go is blazingly fast",
	},
	{
		Id:          2,
		Title:       "Coding in Java",
		Description: "Java is Amazing",
	},
	{
		Id:          3,
		Title:       "Learning DSA",
		Description: "DSA is a must know but it's weird",
	},
}

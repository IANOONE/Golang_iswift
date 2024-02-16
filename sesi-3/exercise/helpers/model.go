package helpers

func Students() []Student {
	var data = []Student{
		{Name: "Alex", Address: "Tangerang Selatan", Profession: "UX Researcher", Reason: "Interesting with learning programming"},
		{Name: "Alex", Address: "Jakarta Barat", Profession: "BackEnd", Reason: "Interesting with learning GO"},
		{Name: "Sam", Address: "Depok", Profession: "FrontEnd", Reason: "Interesting with learning GO especially BackEnd"},
		{Name: "John", Address: "Bekasi", Profession: "FullStack", Reason: "Interesting with learning GO"},
		{Name: "Roy", Address: "Jakarta Timur", Profession: "FrontEnd", Reason: "Interesting with learning GO especially BackEnd"},
	}

	return data
}
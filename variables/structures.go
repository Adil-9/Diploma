package variables

type PersonWebData struct {
	PersonName string
	BirthDate  string
	Gender     string
	City       string
	VkID       string
}

type AllVariables struct {
	Request        bool
	Name           string
	Surname        string
	Country        string
	StringVar      []string
	IntegerVar     []string
	SpecialCharAll bool
}

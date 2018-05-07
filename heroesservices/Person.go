package heroesservices

type Person struct {
	FirstName string
	LastName  string
	HeroName  string
}

func (person *Person) SetHeroName() {
	person.HeroName = GetFirstName(person.FirstName) + " " + GetLastName(person.LastName)
}

func GetFirstName(name string) string {
	return "Hata"
}

func GetLastName(name string) string {
	return "Mahata"
}

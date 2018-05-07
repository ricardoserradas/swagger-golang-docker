package heroesservices

import (
	"strings"
)

// Person represents the person that will be given a Hero name
//
// swagger:model Person
type Person struct {
	// The Person's first name
	//
	// required: true
	FirstName string `json:"firstName"`

	// The Person's last name
	//
	// required: true
	LastName string `json:"lastName"`

	// The Person's given Hero name
	//
	// required: false
	HeroName string `json:"heroName"`
}

func (person *Person) SetHeroName() {
	person.HeroName = GetFirstName(person.FirstName) + " " + GetLastName(person.LastName)
}

func GetLastName(name string) string {
	m := make(map[string]string)

	m["A"] = "Lightning"
	m["B"] = "Knight"
	m["C"] = "Hulk"
	m["D"] = "Centurion"
	m["E"] = "Surfer"
	m["F"] = "Girl"
	m["G"] = "Warrior"
	m["H"] = "Man"
	m["I"] = "Ghost"
	m["J"] = "Master"
	m["K"] = "Hornet"
	m["L"] = "Phantom"
	m["M"] = "Crusader"
	m["N"] = "Daredevil"
	m["O"] = "Machine"
	m["P"] = "America"
	m["Q"] = "X"
	m["R"] = "Doom"
	m["S"] = "First"
	m["T"] = "Shadow"
	m["U"] = "Patriot"
	m["V"] = "Claw"
	m["W"] = "Panther"
	m["X"] = "Hawk"
	m["Y"] = "Storm"
	m["Z"] = "Thunder"

	lastHeroName := m[strings.ToUpper(string(name[0]))]

	return lastHeroName
}

func GetFirstName(name string) string {
	m := make(map[string]string)

	m["A"] = "Captain"
	m["B"] = "Night"
	m["C"] = "The"
	m["D"] = "Ancient"
	m["E"] = "Spider"
	m["F"] = "Invisible"
	m["G"] = "Master"
	m["H"] = "Mr"
	m["I"] = "Silver"
	m["J"] = "Dark"
	m["K"] = "Professor"
	m["L"] = "Radioactive"
	m["M"] = "Incredible"
	m["N"] = "Impossible"
	m["O"] = "Iron"
	m["P"] = "Rocket"
	m["Q"] = "Human"
	m["R"] = "Power"
	m["S"] = "Green"
	m["T"] = "Super"
	m["U"] = "Wonder"
	m["V"] = "Metal"
	m["W"] = "Doctor"
	m["X"] = "Masked"
	m["Y"] = "Crimson"
	m["Z"] = "Omega"

	firstHeroName := m[strings.ToUpper(string(name[0]))]

	return firstHeroName
}

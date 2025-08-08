package pokeapi

type User struct {
	Pokedex map[string]Pokemon
}

func NewUser() User {
	return User{
		Pokedex: make(map[string]Pokemon),
	}
}

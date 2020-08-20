package main

import "fmt"

type clubs struct {
	clubname string
	players  int
}

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltl bool
}

type szemely struct {
	eloNev       string
	utoNev       string
	kedvencFagyi []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle bool
	luxury  bool
}

func main() {

	// sa1 := secretAgent{
	// 	person: person{first: "James", last: "Bond"},
	// 	ltl:    true,
	// }

	// p2 := person{
	// 	first: "Miss",
	// 	last:  "Moneypenny",
	// }

	// cl := clubs{
	// 	clubname: "FC Barcelona",
	// 	players:  22,
	// }

	// cl2 := struct {
	// 	clubname string
	// 	players  int
	// }{
	// 	clubname: "FC Barcelona",
	// 	players:  22,
	// }

	// fmt.Printf("%v %v\n", sa1.first, sa1.last, sa1.ltl)
	// fmt.Printf("%v %v\n", p2.first, p2.last)

	// fmt.Printf("Clubname: %v\nPlayers: %v\n", cl.clubname, cl.players)
	// fmt.Printf(cl2.clubname)

	// sz1 := szemely{
	// 	eloNev:       "Bob",
	// 	utoNev:       "M.",
	// 	kedvencFagyi: []string{"choco", "bana", "juce"},
	// }

	// sz2 := szemely{
	// 	eloNev:       "John",
	// 	utoNev:       "S.",
	// 	kedvencFagyi: []string{"choco", "bana", "juce"},
	// }

	// m := map[string]szemely{
	// 	sz1.utoNev: sz1,
	// 	sz2.utoNev: sz2,
	// }

	// fmt.Printf("First name: %v\nLast name: %v\nFav. ice cream: %v\n", sz1.eloNev, sz1.utoNev, sz1.kedvencFagyi)
	// fmt.Printf("\nFirst name: %v\nLast name: %v\nFav. ice cream: %v\n", sz2.eloNev, sz2.utoNev, sz2.kedvencFagyi)

	// for _, v := range sz1.kedvencFagyi {
	// 	fmt.Println(v)
	// }

	// for _, v := range sz2.kedvencFagyi {
	// 	fmt.Println(v)
	// }

	// for i, v := range m {
	// 	fmt.Println(i, v.eloNev)
	// }

	Vehicle()
	Anon()
}

func Vehicle() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
	}
	fmt.Println(t.color)
}

func Anon() {
	ano := struct {
		car   string
		doors int
	}{
		car:   "Mercedes",
		doors: 4,
	}

	fmt.Println(ano.car, ano.doors)
}

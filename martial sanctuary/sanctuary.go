package main

import (
	"fmt"
	"math/rand"
)

type animal interface{
	Eat() string
	Move() string
	Sleep(bool) bool
}

//Defining daily schedule
const (
	day 		 = 24
	sunset  	 = 23
	sunrise 	 = 6
	daysToTravel = 3
)

//Lion
type lion struct{
	Name string
	Asleep bool
}

func (l lion) String() string{
	return "Lion: " + l.Name
}

func (l lion) Eat() string{
	switch i := rand.Intn(3); i{
	case 0:
		return "rabbit"
	case 1:
		return "meat"
	case 2:
		return "water"
	}

	return "nothing"
}

func (l lion) Move() string{
	switch i := rand.Intn(3); i{
	case 0:
		return "to the right"
	case 1:
		return "to the left"
	case 2:
		return "in the same place"
	}

	return "nowhere"
}

func (l lion) Sleep(state bool) bool{
	if state{
		l.Asleep = true
		return l.Asleep
	}else{
		l.Asleep = false
		return l.Asleep
	}
}

//Zebra
type zebra struct{
	Name string
	Asleep bool
}

func (z zebra) String() string{
	return "Zebra: " + z.Name
}

func (z zebra) Eat() string{
	switch i := rand.Intn(3); i{
	case 0:
		return "grass"
	case 1:
		return "cabage"
	case 2:
		return "water"
	}

	return "nothing"
}

func (z zebra) Move() string{
	switch i := rand.Intn(3); i{
	case 0:
		return "to the right"
	case 1:
		return "to the left"
	case 2:
		return "in the same place"
	}

	return "nowhere"
}

func (z zebra) Sleep(state bool) bool{
	if state{
		z.Asleep = true
		return z.Asleep
	}else{
		z.Asleep = false
		return z.Asleep
	}
}
//Elephant
type elephant struct{
	Name string
	Asleep bool
}

func (e elephant) String() string{
	return "Elephant: " + e.Name
}

func (e elephant) Eat() string{
	switch i := rand.Intn(3); i{
	case 0:
		return "fruit"
	case 1:
		return "grass"
	case 2:
		return "water"
	}
	
	return "nothing"
}

func (e elephant) Move() string{
	switch i := rand.Intn(3); i{
	case 0:
		return "to the right"
	case 1:
		return "to the left"
	case 2:
		return "in the same place"
	}

	return "nowhere"
}

func (e elephant) Sleep(state bool) bool{
	if state{
		e.Asleep = true
		return e.Asleep
	}else{
		e.Asleep = false
		return e.Asleep
	}
}
//Doggy
type dog struct{
	Name string
	Asleep bool
}

func (d dog) String() string{
	return "Dog: " + d.Name
}

func (d dog) Eat() string{
	switch i := rand.Intn(3); i{
	case 0:
		return "rat"
	case 1:
		return "meat"
	case 2:
		return "water"
	}

	return "nothing"
}

func (d dog) Move() string{
	switch i := rand.Intn(3); i{
	case 0:
		return "to the deck"
	case 1:
		return "to the cockpit"
	case 2:
		return "in the same place"
	}

	return "nowhere"
}

func (d dog) Sleep(state bool) bool{
	if state{
		d.Asleep = true
		return d.Asleep
	}else{
		d.Asleep = false
		return d.Asleep
	}
}

//Monkey D Luffy
type monkey struct{
	Name string
	Asleep bool
}

func (m monkey) String() string{
	return "Monkey: " + m.Name
}

func (m monkey) Eat() string{
	switch i := rand.Intn(3); i{
	case 0:
		return "banana"
	case 1:
		return "apple"
	case 2:
		return "water"
	}

	return "nothing"
}

func (m monkey) Move() string{
	switch i := rand.Intn(3); i{
	case 0:
		return "to the top"
	case 1:
		return "to the cargo"
	case 2:
		return "in the same place"
	}

	return "nowhere"
}

func (m monkey) Sleep(state bool) bool{
	if state{
		m.Asleep = true
		return m.Asleep
	}else{
		m.Asleep = false
		return m.Asleep
	}
}

func main(){
	sanctuary := make([]animal,0,5)
	sanctuary = append(sanctuary, lion{Name: "Jeorge"})
	sanctuary = append(sanctuary, zebra{Name: "Juliy"})
	sanctuary = append(sanctuary, elephant{Name: "Enkel"})
	sanctuary = append(sanctuary, dog{Name: "Jessi"})
	sanctuary = append(sanctuary, monkey{Name: "Bob"})

	runSimulation(sanctuary)
}

//Starts from the morning of the first day
func runSimulation(sanctuary []animal){
	for d := 1; d <= daysToTravel; d++{
		fmt.Printf("%18s %v\n", "Day:", d)
		fmt.Println("====================================")
		for h := sunrise; h <= day; h++{
			if h == sunset{
				for _,a := range sanctuary{
					a.Sleep(true)
					fmt.Printf("%v has went to sleep\n",a)
				}
			}else if h == sunrise{
				for _,a := range sanctuary{
					a.Sleep(false)
					fmt.Printf("%v has risen\n",a)
				}
			}else if h < sunset && h > sunrise{
				a := sanctuary[rand.Intn(len(sanctuary))]
				switch r := rand.Intn(2); r{
				case 0:
					fmt.Printf("%s has consumed: %s\n",a, a.Eat())
				case 1:
					fmt.Printf("%s has moved: %s\n",a, a.Move())
				}
			}
		}
		fmt.Println("====================================")
	}
}
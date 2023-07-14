package main

//func (s *Saiyan) Super(amount int) {
//	s.Power += amount
//}

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//goku := &Saiyan{"Goku", 9001}
	//goku.Super(1000)
	//fmt.Println(goku.Power)
	//cards := newDeckFromFile("mycards")
	////card := cards.toString()
	//cards.print()
}

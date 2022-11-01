package main

import "fmt"

type Goods struct {
	name  string
	Price int
}

type Book struct {
	Goods
	writer string
}

func (s Goods) getName() string {
	return s.name
}

func (b Book) getWriter() string {
	return b.writer
}

func main() {
	g := Goods{
		name:  "药水",
		Price: 18,
	}
	fmt.Println(g.getName())

	b := Book{
		writer: "是",
		Goods: Goods{
			name:  "水晶",
			Price: 2,
		},
	}
	b.Goods.name = "sss"
	b.name = "b.Goods.name可以简化成b.name "
	// b.Goods = Goods{
	// 	name:  "水晶",
	// 	Price: 2,
	// }

	fmt.Println(b.getName())
	fmt.Println(b.getWriter())

}

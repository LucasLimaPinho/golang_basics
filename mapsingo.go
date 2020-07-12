package main

import "fmt"

var x [5]int
func main(){
	arr := [...]string{"a","b","c","d","e","f","g"}
	s1 := arr[1:3]
	s2 := arr[2:5]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(cap(s1))
	fmt.Println(cap(s2))
	sli := []int{1,2,3}
	fmt.Println(sli)
	fmt.Println(len(sli))
	fmt.Println(cap(sli))
	slice := make([]int,10)
	fmt.Println(slice)
	slice2 := make([]int,10,15)
	fmt.Println(slice2)
	slice3 := make([]int, 0, 3)
	fmt.Println(len(slice3))
	slice3 = append(slice3,100)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	var idMap map [string]int
	idMap = make(map[string]int)
	fmt.Println(idMap)
	var idMap2 = map[string]int{"joe": 123,
		"maria": 7561,
		"jose":  1234}
	fmt.Println(idMap2)
	fmt.Println(idMap2["maria"])
	idMap2["jane"] = 456
	fmt.Println(idMap2)
	id,p := idMap2["joe"]
	fmt.Println(id)
	fmt.Println(p)
	fmt.Println(len(idMap2))
	for key, val := range idMap2{
		fmt.Println(key,val)
	}

}


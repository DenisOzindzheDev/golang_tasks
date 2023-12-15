package main

import "log"

func main() {
	test := destCity([][]string{{"London", "NewYork"}, {"NewYork", "Lima"}, {"Lima", "SaoPaulo"}})
	_ = test

}
func destCity(paths [][]string) string {
	cityMap := make(map[string]string)
	for _, path := range paths {
		cityMap[path[0]] = path[1]
	} // создаем карту ключ город
	log.Printf("map destination %s", cityMap)
	for _, path := range paths {
		if _, ok := cityMap[path[1]]; !ok {
			return cityMap[path[0]]
		}
	}
	return ""
}

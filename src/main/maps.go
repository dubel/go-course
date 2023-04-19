package main

func main() {
	idMap := map[string]int{
		"1": 1133342,
	}
	println(idMap["1"])
	idMap["2"] = 2233342

	//deleting
	delete(idMap, "2")
	println(idMap["2"])
	println(idMap["2234234342343"])
	id, presence := idMap["2"]
	println(id, presence)
	println(len(idMap))

	for key, value := range idMap {
		println(key, value)
	}
}

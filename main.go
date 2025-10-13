package main

// A package is a way to group functions. main is a special package that tells the Go compiler that the package should compile as an executable program instead of a shared library.

// formatting and console logging package
// "math/Add"

// func perf() {
// 	curTime := time.Now().Unix()
// 	fmt.Println((curTime))
// 	k := 0
// 	for i := 0; i < 100000000000; i++ {
// 		// fmt.Printf("%d", arr[i])
// 		k += 1
// 	}
// 	afterLoopTime := time.Now().Unix()
// 	fmt.Println(afterLoopTime)
// 	fmt.Println(afterLoopTime - curTime)
// }

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type elecEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) calculateMilesLeft() uint32 {
	return uint32(e.mpg) * uint32(e.gallons)
}

func (e elecEngine) calculateMilesLeft() uint32 {
	return uint32(e.kwh) * uint32(e.mpkwh)
}

type engines interface {
	calculateMilesLeft() uint32
}

func canMakeIt(e engines) uint32 {
	return e.calculateMilesLeft()
}

func main() {
	// var arr = [10]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	// fmt.Println(arr)
	// var countMap = make(map[int]int) // make return an object of map|slice|chan with given type
	// for i := range 10 {
	// 	countMap[arr[i]] += 1
	// }
	// keys := make([]int, 0, len(countMap))
	// for k := range countMap {
	// 	keys = append(keys, k)
	// }

	// fmt.Println(keys)

	// var myStr string = "abhinav"
	// fmt.Println(utf8.RuneCountInString(myStr))
	// div, err := math.Division(5, 5)
	// if err != nil {
	// fmt.Println(err.Error())
	// return
	// } else {
	// 	fmt.Println(div)
	// }
	// fmt.Println(math.Add(5, 6))

	// array and slices sizes and capacities

	// var intArr []int
	// fmt.Printf("len of array: %v, capacity of array: %v", len(intArr), cap(intArr))
	// println()
	// intArr = append(intArr, 10) //append to array
	// fmt.Printf("len of array: %v, capacity of array: %v", len(intArr), cap(intArr))
	// println()
	// intArr = append(intArr, 10) //append to array
	// fmt.Printf("len of array: %v, capacity of array: %v", len(intArr), cap(intArr))
	// println()
	// intArr = append(intArr, 10) //append to array
	// fmt.Printf("len of array: %v, capacity of array: %v", len(intArr), cap(intArr))
	// println()
	// intArr = append(intArr, 10) //append to array
	// fmt.Printf("len of array: %v, capacity of array: %v", len(intArr), cap(intArr))
	// println()
	// intArr = append(intArr, 10) //append to array
	// fmt.Printf("len of array: %v, capacity of array: %v", len(intArr), cap(intArr))

	//REALLOCATION OF ARRAY
	// intArr := make([]int, 10000)
	// println("Initial elem addreess", &intArr)
	// println("Second elem addreess", &intArr[1])
	// fmt.Printf("len of array: %v, capacity of array: %v", len(intArr), cap(intArr))
	// println()
	// intArr = append(intArr, 10)
	// println("Initial elem addreess after append", &intArr)
	// println("Second elem addreess", &intArr[1])
	// fmt.Printf("len of array: %v, capacity of array: %v", len(intArr), cap(intArr))

	// var myStr string = "abhinavndfjleknjkbkjb"
	// var strBuild strings.Builder // internally is an array.

	// for k := range myStr {
	// 	strBuild.WriteString(string(myStr[k]))
	// 	println("str builder length", strBuild.Len(), strBuild.Cap())
	// }
	// println(strBuild.String()) // here the array converted to a string.

	// structs and interfaces
	var gas = gasEngine{10, 10}
	var elec = elecEngine{10, 15}
	println(gas.calculateMilesLeft(), canMakeIt(gas))
	println(elec.calculateMilesLeft(), canMakeIt(elec))
	var p *int // referencing and deferencing the pointer
	intVal := 10
	println(p, &p, &intVal)
	p = &intVal
	println(*p, &p, &intVal) // *p is dereferencing the pointer. or getting the value
}

//args: -Eexportloopref
package testdata

func dummyFunction() {
	var array [4]*int
	var slice []*int
	var ref *int
	var str struct{ x *int }

	println("loop expecting 10, 11, 12, 13")
	for i, p := range []int{10, 11, 12, 13} {
		printp(&p)
		slice = append(slice, &p) // ERROR : exporting a pointer for the loop variable p
		array[i] = &p             // ERROR : exporting a pointer for the loop variable p
		if i%2 == 0 {
			ref = &p   // ERROR : exporting a pointer for the loop variable p
			str.x = &p // ERROR : exporting a pointer for the loop variable p
		}
		var vStr struct{ x *int }
		var vArray [4]*int
		var v *int
		if i%2 == 0 {
			v = &p
			vArray[1] = &p
			vStr.x = &p
		}
		_ = v
	}

	println(`slice expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range slice {
		printp(p)
	}
	println(`array expecting "10, 11, 12, 13" but "13, 13, 13, 13"`)
	for _, p := range array {
		printp(p)
	}
	println(`captured value expecting "12" but "13"`)
	printp(ref)
}

func printp(p *int) {
	println(*p)
}

package main

func myFunc() {
	if num := 15; num == 10 {
		println("num is ten")
	} else {
		println("num is not ten")
	}
}

func mySwitchFunc() {
	switch num := 3; num {
	case 1:
		println("I am one")
	case 2:
		println("I am Two")
	case 3:
		println("I am Three")
		fallthrough
	case 4:
		println("I am Fore")
	}
}

func myForFunc() {
	for i := 1; i < 10; i++ {
		if i == 5 {
			println("i am 5 out")
		}
		println(i)
	}
}

func myDoubleForFunc() {
flag:
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				break flag
			}
		}
	}
}

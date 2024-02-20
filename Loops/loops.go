package main

func main() {

	//if else in go
	count := 45

	if count > 60 {
		print("count is greater than 60")
	} else {
		print("count is less than 60")
	}

	if num := 10; num%2 == 0 {
		print("\n", num, "is even")
	} else {
		print("\n", num, "is odd")
	}

	//loops in go
	//for loop
	for i := 0; i < 5; i++ {
		print("\n", i)
	}

	//while loop
	j := 0
	for j < 5 {
		print("\n", j)
		j++
	}

}

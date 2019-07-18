package main

func CheckIfError(err error) {
	if err != nil {
		panic(err)
	}
}

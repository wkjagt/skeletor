package main

func main() {
	db, err := openDatabase("./data.db")
	if err != nil {
		panic(err)
	}

	err = createDemoTable(db)
	if err != nil && err.Error() != "table File already exists" {
		panic(err)
	}
}

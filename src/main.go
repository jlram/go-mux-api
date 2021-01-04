package main

func main() {
	a := App{}

	// TODO: use env variables
	a.Initialize(
		"openpg",
		"openpgpwd",
		"postgres")

	a.Run(":8010")
}

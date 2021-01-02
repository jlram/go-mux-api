package main

func main() {
	a := App{}
	a.Initialize(
		"openpg",
		"openpgpwd",
		"postgres")

	a.Run(":8010")
}

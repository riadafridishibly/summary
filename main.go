package main

func main() {
	dc := new(DataCollector)
	dc.Populate(".")
	dc.Print()
}

package main

func main() {
	dir := flag.String("d", "", "instance directory path")
	flag.Parse()
	conf, err := readConfig(*dir)
	if err != nil {
		log.Fatal(err)
	}
	var s Server
	s.path = *dir
	log.Print(s.Start(conf))
}

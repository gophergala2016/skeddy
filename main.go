package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	cronfile = flag.String("i", "", "import from crontab file.")
	host     = flag.String("s", "0.0.0.0", "bind to ip address.")
	port     = flag.Int("p", 8080, "port to listen.")
	dbfile   = flag.String("d", "skeddy.db", "use existing storage.")
)

var Store *Storage

func main() {
	flag.Parse()

	var err error
	Store, err = NewStorage(*dbfile)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer Store.Close()

	if len(*cronfile) != 0 {
		err := ImportFile(*cronfile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
		} else {
			fmt.Fprintf(os.Stdout, "Import of %s complete ...\n", *cronfile)
			go StartScheduler()
			StartAdminInterface(*host, *port)
		}
	} else {
		go StartScheduler()
		StartAdminInterface(*host, *port)
	}
}

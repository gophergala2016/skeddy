package main

import(
  "fmt"
  "flag"
  "os"
)

var(
	cronfile = flag.String("i", "", "import from crontab file.")
  host     = flag.String("s", "0.0.0.0", "bind to ip address.")
	port     = flag.Int("p", 8080, "port to listen.")
)

func main(){
  flag.Parse()
  if len(*cronfile) != 0 {
		err := ImportFile(*cronfile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
		} else {
			fmt.Fprintf(os.Stdout, "Import of %s complete ...\n", *cronfile)
		}
	} else {
		flag.Usage()
    os.Exit(1)
	}
  StartAdminInterface(*host, *port)
}

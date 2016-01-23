package main

import(
  "fmt"
  "flag"
  "os"
)

var(
	cronfile = flag.String("i", "", "import from crontab file.")
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
}

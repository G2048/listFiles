package cli

import "flag"

type Args struct {
	Colour bool
}

func Get() Args {
	colour := flag.Bool("colour", true, "colour output")
	flag.Parse()
	return Args{Colour: *colour}
}

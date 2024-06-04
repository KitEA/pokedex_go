package main

type config struct {
	previousLocationAreaURL *string
	nextLocationAreaURL     *string
}

func main() {
	cfg := config{}

	startRepl(&cfg)
}

package fetcher

type Fetcher interface {
	Fetch(resource string) (output FetchOutput)
}

type FetchOutput struct {
	Operand1  float64
	Operand2  float64
	Operation string
}

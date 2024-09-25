Writing tests

1. Writing a test is just like writing a function, with a few rules
2. It needs to be in a file with a name like xxx_test.go
3. The test function must start with the word Test
4. The test function takes one argument only t *testing.T

- go test
- go test -v    // run examples and more details about test results
- To run the benchmarks do `go test -bench=.`

// the benchmark framework will determine what is a "good" value for that to let you have some decent results.

goos: darwin
goarch: amd64
pkg: integerrs/03iterations
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkIterations-16          10572080               102.6 ns/op
PASS
ok      integerrs/03iterations  1.524s

----- 102.6 ns/op - means is our function takes on average 102 nanoseconds to run (on my computer). To test this it ran it 10572080 times.

# Captor

Capture calls to builtin functions and the standard library by using **Captor**. This provides a collection of aliases that can be used to intercept arguments during tests.

`Exit` is an alias for `os.Exit`. Capture the exit code using `ExpectExitCode`.

`Panic` is an alias for the builtin `panic` function. Capture the cause of a panicking goroutine using `ExpectPanic`
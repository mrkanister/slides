package os

// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicates success, non-zero an error.
// The program terminates immediately; deferred functions are not run. // HL
//
// For portability, the status code should be in the range [0, 125].
func Exit(code int)

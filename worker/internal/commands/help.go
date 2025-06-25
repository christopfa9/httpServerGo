package commands

import (
	"strings"
)

// Help returns a readable list of all available endpoints.
func Help() (string, error) {
	lines := []string{
		"/fibonacci?num={N}                     : Calculates the N-th Fibonacci number",
		"/createfile?name={name}&content={text}&repeat={times} : Creates or truncates a file with repeated content",
		"/deletefile?name={name}               : Deletes an existing file",
		"/reverse?text={text}                  : Reverses the given text string",
		"/toupper?text={text}                  : Converts the text to uppercase",
		"/random?count={c}&min={min}&max={max} : Generates an array of random numbers",
		"/timestamp                            : Returns the current time in ISO-8601 format",
		"/hash?text={text}                     : Computes the SHA-256 hash of the text",
		"/simulate?seconds={s}&task={name}     : Simulates a task by sleeping for X seconds",
		"/sleep?seconds={s}                    : Suspends execution for X seconds",
		"/loadtest?tasks={n}&sleep={s}         : Runs N concurrent tasks sleeping S seconds each",
		"/status                               : Shows server metrics (uptime, connections, processes)",
		"/help                                 : Displays this help message",
	}
	return strings.Join(lines, "\r\n"), nil
}

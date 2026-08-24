package echoer

// Echo returns the first element of args, or "" if args is empty or nil.
// It performs no I/O, relies on no global state, and preserves internal whitespace unmodified.
func Echo(args []string) string {
	if len(args) == 0 {
		return ""
	}
	return args[0]
}

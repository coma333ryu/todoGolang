package exceptions

//ProcError : proc error
func ProcError(err error) {
	if err != nil {
		panic(err)
	}
}

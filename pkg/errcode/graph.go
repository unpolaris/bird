package errcode

func IsErrNoRecord(err error) bool {
	return err.Error() == "Result contains no more records"
}

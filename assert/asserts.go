package assert

func NilReceiverCheck(r *interface{}) string {
	if r == nil {
		return nil
	}
	return strconv.Itoa(r.n)
}

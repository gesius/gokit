package assert

func NilReceiverCheck(r *interface{}) interface{} {
	if r == nil {
		return nil
	}
	return r
}

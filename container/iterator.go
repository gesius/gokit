


//MapIterator type alias

//Example: (use of type assertion for type conversion)
//	m := MapIterator{ "key1":"val1", "key2":"val2" };
//	m.Foreach(func(key interface{}, val interface{}) {
//		printMapValues(key.(string), val.(string))
//	})
type MapIterator map[interface{}]interface{}

func (iter *MapIterator) Foreach(f func(interface{}, interface{})) {
	for k, v := range *iter {
		f(k, v)
	}
}

// N returns a slice of n 0-sized elements, suitable for ranging over.
//
// For example:
//
//    for i := range To(10) {
//        fmt.Println(i)
//    }
//
// ... will print 0 to 9, inclusive.
//
// It does not cause any allocations.
func To(n int) []struct{} {
	return make([]struct{}, n)
}
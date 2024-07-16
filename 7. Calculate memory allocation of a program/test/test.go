package test

func Mytest(cycle int) {
	var c interface{}
	for i := range cycle {
		c = i
		_ = c
	}

}

package monkey

func networkCompute(a, b int) (int, error) {
	c := a + b
	return c, nil
}

func Compute(a, b int) (int, error) {
	sum, err := networkCompute(a, b)
	/*
		业务逻辑
	*/
	return sum, err
}

type Computer struct {
}

func (c *Computer) NetworkCompute(a, b int) (int, error) {
	sum := a + b
	return sum, nil
}

func (c *Computer) Compute(a, b int) (int, error) {
	sum, err := c.NetworkCompute(a, b)
	/*
		业务逻辑
	*/
	return sum, err
}

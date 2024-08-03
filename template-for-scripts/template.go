package template_for_scripts

import "errors"

type Item struct {
	Price float64
}

type ExampleType struct {
	Items []Item
}

func New() *ExampleType {
	return &ExampleType{
		Items: make([]Item, 0),
	}
}

func (c *ExampleType) ApplyDiscount(discount float64) error {
	if discount < 0.1 || discount > 1.0 {
		return errors.New("invalid discount")
	}

	for i := range c.Items {
		c.Items[i].Price = c.Items[i].Price * (1 - discount)
	}

	return nil
}

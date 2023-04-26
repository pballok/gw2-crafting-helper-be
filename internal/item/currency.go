package item

type Currency struct {
	Id   int
	Name string
	Icon string
}

func NewCurrency(id int) *Currency {
	return &Currency{Id: id}
}

func (c *Currency) FetchAll() error {
	err := c.fetchBase()
	if err != nil {
		return err
	}

	return nil
}

func (c *Currency) fetchBase() error {
	return nil
}

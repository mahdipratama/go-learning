package books

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

func (client *Client) GetBook(ID string) (Book, error) {
	return Book{}, nil
}

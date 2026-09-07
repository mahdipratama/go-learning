package books

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

func (client *Client) GetBook(ID string) (Book, error) {

	resp, err := http.Get("http://" + client.addr + "/v1/find/" + ID)

	if err != nil {
		return Book{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status %d", resp.StatusCode)
	}

	book := Book{}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &book)
	if err != nil {
		fmt.Printf("%v in %q ", err, data)
	}

	fmt.Println(book)

	return book, nil
}

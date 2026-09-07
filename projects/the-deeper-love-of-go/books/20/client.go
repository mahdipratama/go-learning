package books

import (
	"encoding/json"
	"errors"
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

	// if resp.StatusCode != http.StatusOK {
	// 	fmt.Printf("Unexpected status %d", resp.StatusCode)
	// }

	if resp.StatusCode == http.StatusNotFound {
		return Book{}, fmt.Errorf("%q not found", ID)
	}

	book := Book{}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Book{}, err
	}

	err = json.Unmarshal(data, &book)
	if err != nil {
		return Book{}, fmt.Errorf("%v in %q ", err, data)
	}

	return book, nil
}

func (client *Client) GetAllBook() ([]Book, error) {
	resp, err := http.Get("http://" + client.addr + "/v1/list")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	bookList := []Book{}
	err = json.Unmarshal(data, &bookList)
	if err != nil {
		return nil, fmt.Errorf("%v in %q ", err, data)
	}

	return bookList, nil
}

func (client *Client) MakeAPIRequest(URI string, result any) error {
	resp, err := http.Get("http://" + client.addr + "/v1" + URI)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return errors.New("not found")
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return fmt.Errorf("%v in %q", err, data)
	}

	return nil
}

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

func (client *Client) GetCopies(ID string) (int, error) {
	copies := 0
	err := client.MakeAPIRequest("/getcopies/"+ID, &copies)
	if err != nil {
		return 0, err
	}

	return copies, nil
}

func (client *Client) AddCopies(ID string, copies int) (int, error) {
	// URI := "/addcopies/" + ID + "/" + strconv.Itoa(copies)
	URI := fmt.Sprintf("/addcopies/%s/%d", ID, copies)
	stock := 0

	err := client.MakeAPIRequest(URI, &stock)
	if err != nil {
		return 0, err
	}

	return stock, nil
}

func (client *Client) SubCopies(ID string, copies int) (int, error) {
	// URI := "/addcopies/" + ID + "/" + strconv.Itoa(copies)
	URI := fmt.Sprintf("/subcopies/%s/%d", ID, copies)
	stock := 0

	err := client.MakeAPIRequest(URI, &stock)
	if err != nil {
		return 0, err
	}

	return stock, nil
}

func (client *Client) MakeAPIRequest(URI string, result any) error {
	resp, err := http.Get("http://" + client.addr + "/v1" + URI)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		if len(data) > 0 {
			return errors.New(string(data))
		}
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

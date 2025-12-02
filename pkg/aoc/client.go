package aoc

import (
	"fmt"
	"io"
	"net/http"

	"github.com/anaskhan96/soup"
)

const (
	defaultURL       = "https://adventofcode.com/%d/day/%d"
	defaultCachePath = "./inputs/%d/day_%d/%s.txt"
)

type AoCClient interface {
	FetchInput(year int, day int, simple bool) (string, error)
}

type aocClient struct {
	url       string
	token     string
	cachePath string
}

func NewAOCClient(opts ...AoCOption) (AoCClient, error) {
	client := &aocClient{
		url:       defaultURL,
		cachePath: defaultCachePath,
	}

	for _, opt := range opts {
		opt(client)
	}

	if client.token == "" {
		fmt.Print("Input token: ")
		_, err := fmt.Scan(&client.token)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

type AoCOption func(*aocClient)

func WithBaseURL(url string) AoCOption {
	return func(c *aocClient) {
		c.url = url
	}
}

func WithCachePath(path string) AoCOption {
	return func(c *aocClient) {
		c.cachePath = path
	}
}

func WithToken(token string) AoCOption {
	return func(c *aocClient) {
		c.token = token
	}
}

func (c *aocClient) FetchInput(year int, day int, simple bool) (string, error) {
	if simple {
		return c.fetchSimpleInput(year, day)
	}

	return c.fetchActualInput(year, day)
}

func (c *aocClient) fetchSimpleInput(year int, day int) (string, error) {
	resp, err := soup.Get(fmt.Sprintf(c.url, year, day))
	if err != nil {
		return "", err
	}

	doc := soup.HTMLParse(resp)
	blocks := doc.FindAll("code")

	if len(blocks) == 0 {
		return "", fmt.Errorf("could not find example input on the page")
	}

	var example string
	for _, block := range blocks {
		text := block.Text()
		if len(text) > len(example) {
			example = text
		}
	}

	return example, nil
}

func (c *aocClient) fetchActualInput(year int, day int) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(c.url+"/input", year, day), nil)
	if err != nil {
		return "", err
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: c.token,
	})

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("failed to close response body: %v\n", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch input: status code %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

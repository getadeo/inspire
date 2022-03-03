package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/charmbracelet/lipgloss"
)

const endpoint = "https://api.quotable.io/random"

type Quote struct {
	ID           string   `json:"_id"`
	Tags         []string `json:"tags"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	AuthorSlug   string   `json:"authorSlug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
}

func main() {
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var quote Quote

	if err := json.Unmarshal(body, &quote); err != nil {
		fmt.Println("Can't unmarshal JSON")
	}

	text := "\""
	text += quote.Content
	text += "\""
	text += " - "
	text += quote.Author

	var style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#2E3440")).
		Background(lipgloss.Color("#88C0D0")).
		Width(len(text))

	styledQuote := style.Render(text)

	fmt.Println(style.Render(styledQuote))

}

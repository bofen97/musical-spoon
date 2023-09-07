package arxiv

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

const (
	arxiv = "http://export.arxiv.org/api/query?search_query="
)

type Result struct {
	Entry []*ArxivJson `xml:"entry"`
}
type ArxivJson struct {
	Id        string      `xml:"id"`
	Updated   string      `xml:"updated"`
	Published string      `xml:"published"`
	Title     string      `xml:"title"`
	Summary   string      `xml:"summary"`
	Authors   []*Author   `xml:"author"`
	Categorys []*Category `xml:"category"`
}

type Author struct {
	Name string `xml:"name"`
}

type Category struct {
	Term string `xml:"term,attr"`
}

func QuerySubmittedDateTopic(cat string) ([]byte, error) {

	queryParam := "cat:" + cat + "&sortBy=submittedDate&sortOrder=descending&max_results=50"
	queryURL := arxiv + queryParam
	resp, err := http.Get(queryURL)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		//resp.Body.Close()
		return nil, fmt.Errorf("Getting %s: %s", queryURL, resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func PrintNewInfo(categorys ...string) {
	for _, categoryStr := range categorys {

		res := new(Result)
		//cs.ai ..etl
		data, err := QuerySubmittedDateTopic(categoryStr)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		err = xml.Unmarshal([]byte(data), &res)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		for _, entry := range res.Entry {
			fmt.Printf("*********************************\n")
			fmt.Printf("Id\t%s\n", entry.Id)
			fmt.Printf("Update\t%s\n", entry.Updated)
			fmt.Printf("Published\t%s\n", entry.Published)
			fmt.Printf("Title\t%s\n", entry.Title)
			fmt.Printf("Summary\t%s\n", entry.Summary)
			for _, author := range entry.Authors {
				fmt.Printf("author\t%s\n", author.Name)

			}
			for _, cat := range entry.Categorys {
				fmt.Printf("Cat Term\t%s\n", cat.Term)
			}

		}

	}
}

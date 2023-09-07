package arxiv

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
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
func (res *Result) MakeResultFromCate(cat string) error {

	data, err := QuerySubmittedDateTopic(cat)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	err = xml.Unmarshal([]byte(data), res)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}

	return nil

}
func (res *Result) PrintEntrys(w io.Writer) {
	for _, entry := range res.Entry {
		fmt.Fprintf(w, "*********************************\n")
		fmt.Fprintf(w, "Id\t%s\n", entry.Id)
		fmt.Fprintf(w, "Update\t%s\n", entry.Updated)
		fmt.Fprintf(w, "Published\t%s\n", entry.Published)
		fmt.Fprintf(w, "Title\t%s\n", entry.Title)
		fmt.Fprintf(w, "Summary\t%s\n", entry.Summary)
		for _, author := range entry.Authors {
			fmt.Fprintf(w, "author\t%s\n", author.Name)
		}
		for _, cat := range entry.Categorys {
			fmt.Fprintf(w, "Cat Term\t%s\n", cat.Term)
		}

	}

}

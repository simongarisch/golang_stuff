package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	url = "https://github.com/EbookFoundation/free-programming-books/blob/master/free-programming-books.md"
)

func getHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(html), nil
}

func getLinks(url string) ([]string, error) {
	links := []string{}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return links, err
	}

	doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		if strings.HasSuffix(href, ".pdf") {
			links = append(links, href)
		}
	})
	return links, nil
}

func getPdfName(link string) (string, error) {
	var name string

	if !strings.HasSuffix(link, ".pdf") {
		return name, fmt.Errorf("link %q does point to a pdf file", link)
	}
	parts := strings.Split(link, "/")
	return parts[len(parts)-1], nil
}

func fileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err != nil, err
}

func DownloadFile(filePath string, link string) error {

	exists, _ := fileExists(filePath)
	if exists {
		fmt.Printf("File %q already exists", filePath)
		return nil // don't download the file again
	}

	// read link
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	fmt.Println("creating file", filePath)
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func DownloadAll(folder string, url string) error {

	links, err := getLinks(url)
	if err != nil {
		return err
	}

	folderPath := filepath.Join(".", folder)
	os.MkdirAll(folderPath, os.ModePerm) // returns nil if success, else err

	var name string
	for _, link := range links {
		name, err = getPdfName(link)
		if err != nil {
			return err
		}
		filePath := filepath.Join(folderPath, name)
		err := DownloadFile(filePath, link)
		if err != nil {
			fmt.Printf("Unable to download link %q: %q\n", link, err.Error())
		}
	}
	return nil
}

func main() {
	/*
		html, err := getHTML(url)
		if err != nil {
			panic(err)
		}

		links, err := getLinks(url)
		if err != nil {
			panic(err)
		}

		var name string
		for _, href := range links {
			name, err = getPdfName(href)
			if err != nil {
				panic(err)
			}
			fmt.Println(name, "-->", href)
		}
	*/
	DownloadAll("Ebooks", url)
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

type ScrapeResult struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	Date     string `json:"date"`
	Desc     string `json:"desc"`
	Content  string `json:"content"`
	ImgURL   string `json:"img_url"` // Menambahkan URL gambar
}

func fetchPage(url string, headers map[string]string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP status error: %d", resp.StatusCode)
	}

	// Using io.Copy to read the response body into a strings.Builder
	var sb strings.Builder
	_, err = io.Copy(&sb, resp.Body)
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}

func parseContent(url string) (string, error) {
	html, err := fetchPage(url, nil)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}

	var paragraphs []string
	doc.Find("div.detail__body-text > p").Each(func(i int, s *goquery.Selection) {
		paragraphs = append(paragraphs, s.Text())
	})

	if len(paragraphs) == 0 {
		return "No content available.", nil
	}

	return strings.Join(paragraphs, "\n"), nil
}

func parseItem(result *goquery.Selection) (*ScrapeResult, error) {
	// Mengambil title, date, url, dan desc
	title := result.Find("h3.media__title").Text()
	date := result.Find(".media__date > span").AttrOr("title", "")
	url := result.Find("a").AttrOr("href", "")
	desc := result.Find("div.media__desc").Text()

	// Mencari gambar pertama pada halaman
	imgURL, _ := result.Find("img").Attr("src") // Mengambil src dari img pertama

	// Menangkap konten dari halaman terkait
	content, err := parseContent(url)
	if err != nil {
		return nil, err
	}

	// Mengembalikan hasil dalam ScrapeResult
	return &ScrapeResult{
		Title:   strings.TrimSpace(title),
		URL:     url,
		Date:    date,
		Desc:    strings.TrimSpace(desc),
		Content: content,
		ImgURL: imgURL, // Menyertakan URL gambar
	}, nil
}

func parse(url string, params map[string]string, headers map[string]string) ([]*ScrapeResult, error) {
	// Building query parameters
	reqURL := fmt.Sprintf("%s?%s", url, buildQuery(params))

	html, err := fetchPage(reqURL, headers)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var results []*ScrapeResult
	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		result, err := parseItem(s)
		if err == nil {
			results = append(results, result)
		}
	})

	return results, nil
}

func buildQuery(params map[string]string) string {
	var queryParts []string
	for key, value := range params {
		queryParts = append(queryParts, fmt.Sprintf("%s=%s", key, value))
	}
	return strings.Join(queryParts, "&")
}

func fetchJSON(url string, headers map[string]string) (map[string]interface{}, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status error: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getTrendingKeywords(apiURL string, headers map[string]string) ([]string, error) {
	jsonData, err := fetchJSON(apiURL, headers)
	if err != nil {
		return nil, err
	}

	body, ok := jsonData["body"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	topKeywordSearch, ok := body["topKeywordSearch"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("missing 'topKeywordSearch' in response")
	}

	var keywords []string
	for _, item := range topKeywordSearch {
		if keywordItem, ok := item.(map[string]interface{}); ok {
			if keyword, ok := keywordItem["keyword"].(string); ok {
				keywords = append(keywords, keyword)
			}
		}
	}

	return keywords, nil
}

func trendingKeywords(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/244.178.44.111 Safari/537.36",
	}
	apiURL := "https://explore-api.detik.com/trending"

	keywords, err := getTrendingKeywords(apiURL, headers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(keywords)
}

func scrape(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	keyword := params.Get("keyword")
	if keyword == "" {
		http.Error(w, "Keyword is required", http.StatusBadRequest)
		return
	}

	pages := 1
	if pagesStr := params.Get("pages"); pagesStr != "" {
		fmt.Sscanf(pagesStr, "%d", &pages)
	}

	searchURL := "https://www.detik.com/search/searchall"
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/244.178.44.111 Safari/537.36",
	}

	var allItems []*ScrapeResult
	for page := 1; page <= pages; page++ {
		params := map[string]string{
			"query": keyword,
			"page":  fmt.Sprintf("%d", page),
		}
		items, err := parse(searchURL, params, headers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		allItems = append(allItems, items...)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allItems)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Welcome to the DETIKScraper API!",
		"description": "An API for scraping search results and fetching trending keywords from Detik.",
		"endpoints": map[string]interface{}{
			"/trending_keywords": map[string]interface{}{
				"method":      "GET",
				"description": "Retrieve a list of trending keywords.",
			},
			"/scrape": map[string]interface{}{
				"method": "GET",
				"description": "Scrape search results for a specific keyword.",
				"parameters": map[string]interface{}{
					"keyword": "str (required) - The search term to scrape.",
					"pages":   "int (optional) - The number of pages to scrape, defaults to 1.",
				},
			},
		},
		
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", root).Methods("GET")
	r.HandleFunc("/trending_keywords", trendingKeywords).Methods("GET")
	r.HandleFunc("/scrape", scrape).Methods("GET")

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}

# DetikScraper API

DetikScraper API is a Go-based web scraper designed to fetch trending keywords and scrape search results from the [Detik](https://www.detik.com) website. This project exposes two main API endpoints:

1. **`/trending_keywords`**: Retrieve a list of trending keywords from Detik.
2. **`/scrape`**: Scrape search results for a specific keyword, including title, URL, date, description, content, and image.

## Features

- Scrapes articles from Detik based on a keyword search.
- Retrieves trending keywords from Detik.
- Extracts article content and images (if available).
- Supports pagination for search results.

## Requirements

- **Go 1.18** or higher
- An internet connection to fetch web data

## Installation

Follow these steps to set up the project locally.

### Step 1: Clone the Repository

Clone the repository to your local machine:


### Step 2: Install Dependencies

git clone https://github.com/yourusername/detikscraper.git
cd detikscraper

Step 2: Install Dependencies
Run the following commands to install the required dependencies:

### Step 3: Run the Go Server

go get -u github.com/PuerkitoBio/goquery
go get -u github.com/gorilla/mux

Step 3: Run the Go Server
To run the Go server locally, execute:

go run main.go

### API Usage
1. /trending_keywords [GET]
Description:
This endpoint retrieves a list of trending keywords from Detik.

Example Request:
bash
Copy code
curl http://localhost:8080/trending_keywords
Example Response:
json
Copy code
[
  "golang",
  "cloud computing",
  "artificial intelligence",
  "cryptocurrency"
]
2. /scrape [GET]
Description:
This endpoint scrapes search results for a specific keyword, including title, URL, date, description, content, and image.

Query Parameters:
keyword (required): The search term to scrape.
pages (optional): The number of pages to scrape (default is 1).
Example Request:
bash
Copy code
curl "http://localhost:8080/scrape?keyword=golang&pages=3"
Example Response:
json
Copy code
[
  {
    "title": "Exploring the Latest in Golang Development",
    "url": "https://www.detik.com/technology/golang-development",
    "date": "2024-12-14",
    "desc": "Golang has seen significant growth in the developer community...",
    "content": "In this article, we explore the latest updates in Golang...",
    "img_url": "https://www.detik.com/images/golang_thumbnail.jpg"
  },
  {
    "title": "Why Golang is the Future of Backend Development",
    "url": "https://www.detik.com/technology/golang-future",
    "date": "2024-12-15",
    "desc": "Golang continues to emerge as a leading language in backend...",
    "content": "With new features and growing community support, Golang is set to dominate backend development...",
    "img_url": "https://www.detik.com/images/golang_backend_future.jpg"
  }
]
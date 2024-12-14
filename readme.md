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

curl http://localhost:8080/trending_keywords
Example Response:
```bash
[
  "golang",
  "cloud computing",
  "artificial intelligence",
  "cryptocurrency"
]
```
2. /scrape [GET]
Description:
This endpoint scrapes search results for a specific keyword, including title, URL, date, description, content, and image.

Query Parameters:
keyword (required): The search term to scrape.
pages (optional): The number of pages to scrape (default is 1).

Example Request:
```bash
curl "http://localhost:8080/scrape?keyword=desa&pages=3"
Example Response:

[
  {
        "title": "Warga Dua Desa Bentrok Imbas Penikaman Pemuda di Bima",
        "url": "https://www.detik.com/bali/hukum-dan-kriminal/d-7685628/warga-dua-desa-bentrok-imbas-penikaman-pemuda-di-bima",
        "date": "Jumat, 13 Des 2024 19:23 WIB",
        "desc": "Bentrokan antara warga Desa Samili dan Dadibou di Bima dipicu penikaman. Polisi amankan delapan orang, tiga ditahan. Situasi kini kondusif.",
        "content": "Warga Desa Samili dan Desa Dadibou, Kecamatan Woha, Kabupaten Bima, Nusa Tenggara Barat (NTB), bentrok. Bentrokan terjadi imbas penikaman seorang pemuda Desa Samili beberapa waktu lalu.\n\"Warga dua desa saling serang karena saling provokasi. Namun, sejauh ini sudah kondusif,\" kata Kepala Bagian Operasi (Kabag Ops) Polres Bima, AKP Iwan Sugianto, kepada detikBali, Jumat (13/12/2024).\nIwan mengatakan warga dua desa saling serang pada Kamis (12/12/2024). Hal itu dipicu meninggalnya, Rahmansyah (20). Warga Desa Samili itu meninggal diduga ditikam sejumlah warga Dadibou, Sabtu (7/12/2024).\nPolisi langsung mengamankan delapan orang dari kasus penikaman itu. Setelah diklarifikasi dan diperiksa, hanya tiga orang yang ditahan dan diproses hukum. Sedangkan lima orang lainnya dilepas.\n\"Penahanan tiga orang dan dilepas lima orang ini tidak diterima oleh pihak keluarga korban dan warga. Mereka ingin delapan orang yang diamankan ditahan semuanya,\" kata Iwan.\nIwan menegaskan Polres Bima tidak bisa menahan semua orang tersebut. Mengingat, hasil klarifikasi dan pemeriksaan, hanya lima orang dinyatakan tak terlibat dalam kasus penikaman. Selain itu, keterlibatan mereka juga tidak cukup bukti.\n\"Karena persoalan ini, akhirnya sebagian warga terprovokasi dan saling serang,\" beber Iwan.\nIwan mengimbau dua warga desa, terutama keluarga korban, agar menahan diri. Polisi sudah maksimal untuk menuntaskan persoalan kasus penikaman. Buktinya, dengan bergerak cepat menangkap terduga pelaku.\n\"Kami imbau warga tidak terpancing dan mudah terprovokasi. Percayakan kasus ditangani polisi, apalagi terduga pelakunya sudah ditangkap,\" jelas Iwan.",
        "img_url": "https://akcdn.detik.net.id/visual/2015/11/18/12cde27d-e92e-46a3-bb52-710f6d9c547c_43.jpg?w=250&q=90"
    }, {
        ...
    }
]
```
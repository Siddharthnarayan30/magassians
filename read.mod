#flow of the API calls

    RESTAPI call --> middleware --> unmarshal urls --> crawl the web --> marshal response -> send response


#steps to start the server
--checkout th project
go mod download
go run .

#sample RESTAPI body 

POST http://localhost:9000/api/crawl
Content-Type: application/json

{
"urls": ["https://google.com", "https://github.com"]
}

# Uptime Monitoring Service.
- Moniters URLs or websides by doing HTTP GET request
- The system accept a list of URL, Crawl Timeout(In seconds), Frequency(In Seconds), Failure Threshold
- The system iterates over all the urls in the system and try to do a HTTP GET on the URL, and wait for the crawl_timeout seconds if responce is not 200 OK, it marks as failure and increases failure count in the database
- Once failure count reaches to failure threshold it marks url as inactive and stops crawling

## Tech Stack Used:
- Golang - gin (microframework)
- Mysql
  - Gorm as orm library
- Docker


## API
### Base URL
```
http://localhost:8080
```


### To Register A New URL For Monitoring 
Use `POST /urls/` to add a URL to the service and provide following data as json string.
- crawl_timeout: Time for which system wait before giving up on URL.
- frequency :  It determines that how frequenct URL will be pinnged.
- failure_threshold  : As failure count reaches the threshold, System mark is inactive and stops crawling it.
#### Provided information is stored into database and an UUID is Assigned with each URL 
- Request: `POST /urls`
```
{
   "url":                        ”www.example.com”,
   "crawl_timeout":              10,
   “frequency”:                  15, 
   “failure_threshold” :         10,  
}
```
Response:
```
{
  "id":"                        b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "url":                        ”www.example.com”,
  "crawl_timeout":              10,
  “frequency”:                  15,
  “failure_threshold” :         10,
  “status”:                     “active”,
  “failure_count”:               0
}

```

- Curl Command Example:
```
curl -i -X POST http://localhost:8080/urls --header "Content-Type: application/json" --data '{"url": "https://www.google.com", "frequency": 10,"crawl_timeout" : 8, "failure_threshold":3 }'
```


### To GET A Already Stored Url Information 
Use `GET /urls/:id` to fetch monitoring information of a URL which has corresponding id.


- Request: `GET /urls/:id`
- Response:
```
{
  "id":"                        b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "url":                        ”www.example.com”,
  "crawl_timeout":              10,
  “frequency”:                  15,
  “failure_threshold” :         10,
  “status”:                     “active”,
  “failure_count”:               0,
}

```

- Curl Command Example:
```
curl -i -X GET http://localhost:8080/urls/b7f32a21-b863-4dd1-bd86-e99e8961ffc6
```




### To Update URL Information 
Use `PATCH /urls/:id` and Provide following data as json string.
- crawl_timeout 
- frequency
- failure_threshold
#### Note :
- Provide any/all above parameters, only provided parameters will be updated rest will remain same
- Only above parameters can be updated     
- Request: `POST /urls/:id`
```
{
   "crawl_timeout":              10,
   “frequency”:                  15, 
   “failure_threshold” :         10,  
}
```
Response:
```
{
  "id":"                        "b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "url":                        ”www.example.com”,
  "crawl_timeout":              10,
  “frequency”:                  15,
  “failure_threshold” :         10,
  “status”:                     “active”,
  “failure_count”:               0
}

```

- Curl Command Example:
```
curl -i -X PATCH http://localhost:8080/urls/b7f32a21-b863-4dd1-bd86-e99e8961ffc6 --header "Content-Type: application/json" --data '{"frequency": 10,"crawl_timeout" : 8, "failure_threshold":3 }'
```


### To Activate Monitoring for a URL

- Request: `POST /urls/:id/activate`

Response:
```
Status is StatusOk if URL was deactivated before
Status is StatusNotAcceptable if alreday activated
```

- Curl Command Example:
```
curl -i -X POST http://localhost:8080/urls/b7f32a21-b863-4dd1-bd86-e99e8961ffc6/activate
```

### To Deactivate Monitoring for a URL

- Request: `POST /urls/:id/deactivate`

Response:
```
Status is StatusOk if URL was activated before
Status is StatusNotAcceptable if alreday deactivated
```

- Curl Command Example:
```
curl -i -X POST http://localhost:8080/urls/b7f32a21-b863-4dd1-bd86-e99e8961ffc6/deactivate
```




### To Delete a URL

- Request: `DELETE /urls/:id`

Response:
```
Status is StatusNoContent
```

- Curl Command Example:
```
curl -i -X DELETE http://localhost:8080/urls/b7f32a21-b863-4dd1-bd86-e99e8961ffc6
```
## How to Run
1. Install Mysql server on your local machine or use any remote mysql server
2. Create a database
3. Clone this repository : 
  ```git clone https://github.com/shubhamiitmandi123/UpTimeMonitor.git```
4. cd UpTimeMonitor
5. If you have golang installed in your machine and want to run on your local machine then
-  Edit .env file and write your mysql host, password, port and database name you created.
```
export DB_USER="your_mysql_user_name"
export DB_PASS="your_mysql_password"
export DB_NAME="database_name"
export DB_HOST="mysql_host_IP_Address"
export DB_PORT="Port_of_mysql_server"
```
- Note: if your mysql server is installed on host machine then change host to host.docker.internal

- Install modules
``` 
go mod download
```
- Build executable
```
go build .
```
- Run executable
```
./main
```
7. if You wish to run in docker container
- Build image 
```
docker build -t schoudhary2608/uptimemonitor .
```
- Run container
```
docker run -p 8080:8080 -e DOCKER=true -e DB_USER='user' -e DB_PASS='password' -e DB_HOST='host' -e DB_PORT='3306' -e DB_NAME='database_name' schoudhary2608/uptimemonitor

```
- Note: if your mysql server is installed on host machine then change host to host.docker.internal
```
export DB_HOST="host.docker.internal"
```
- Docker image is also available on docker hub
```
schoudhary2608/uptimemonitor 
```

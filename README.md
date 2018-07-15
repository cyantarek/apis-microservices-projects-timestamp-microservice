# API Project: Timestamp Microservice

### Reference: https://curse-arrow.glitch.me/

### Live Deployed on Heroku: https://hidden-tor-86938.herokuapp.com/api/timestamp/

## Guideline

**User Stories (WIP):**

    1. The API endpoint is GET [project_url]/api/timestamp/:date_string?
    
    2. A date string is valid if can be successfully parsed by new Date(date_string).
    Note that the unix timestamp needs to be an integer (not a string) specifying milliseconds.
    In our test we will use date strings compliant with ISO-8601 (e.g. "2016-11-20") because this will ensure an UTC timestamp.
    
    3. If the date string is empty it should be equivalent to trigger new Date(), i.e. the service uses the current timestamp.
    
    4. If the date string is valid the api returns a JSON having the structure
    {"unix": <date.getTime()>, "utc" : <date.toUTCString()> }
    e.g. {"unix": 1479663089000 ,"utc": "Sun, 20 Nov 2016 17:31:29 GMT"}
    
    5. If the date string is invalid the api returns a JSON having the structure
    {"error" : "Invalid Date" }.

**Example Usage:**

    [project url]/api/timestamp/2015-12-25
    [project url]/api/timestamp/1450137600

**Example Output:**

{"unix":1451001600000, "utc":"Fri, 25 Dec 2015 00:00:00 GMT"}

# Load Testing Results

 1. Bombarding http://127.0.0.1:14759/api/timestamp/2015-12-25 with 1000000 request(s) using 125 connection(s)
 2. 1000000 / 1000000 [===========================>] 100.00% 9s
 3.Done!
 4. Statistics        Avg      Stdev        Max
 5. Reqs/sec    104514.28   14997.61  142855.76
 6. Latency        1.19ms   658.57us    80.17ms
 7. HTTP codes:
 8. 1xx - 0, 2xx - 1000000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
 9. Throughput:    25.14MB/s%      

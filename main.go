package main

import (
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"errors"
	"strings"
	"strconv"
)

type Response struct {
	UNIX int64  `json:"unix"`
	UTC  string `json:"utc"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/timestamp/", APIHandler)

	if err := http.ListenAndServe(":80", mux); err != nil {
		fmt.Println(err.Error())
	}
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if timestamp := r.URL.Path[len("/api/timestamp/"):]; timestamp != "" { //checks if timestamp presents in the url
			if resp, err := TimeResponseGenerator(timestamp); err == nil { //sends non empty data to function
				w.Header().Set("Content-Type", "application/json") //if no errors, function will return response data
				w.WriteHeader(http.StatusOK) //sets status 200 OK
				if err := json.NewEncoder(w).Encode(resp); err != nil { //we set json header and returns json response
					w.WriteHeader(http.StatusInternalServerError) //if encoding error occurs
					return
				}

			} else {
				w.Header().Set("Content-Type", "application/json") //if errors found from the function
				w.WriteHeader(http.StatusInternalServerError) //sets status 500
				json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()}) //returns error as json
				return
			}
		} else { //if timestamp is not given
			if resp, err := TimeResponseGenerator(timestamp); err == nil {
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(resp); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

			} else { //if timestamp wrong format given
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
				return
			}
		}
	}
}

func TimeResponseGenerator(timestamp string) (Response, error) {
	if timestamp == "" { //if timestamp is not given
		timestamp = time.Now().Format("2006-01-02") // sets timestamp to current time
	}

	if check := strings.Contains(timestamp, "-"); check { //checks if UTC format
		layout := "2006-01-02" //set layout
		if t, err := time.Parse(layout, timestamp); err == nil { //parses the input timestamp
			var response Response
			response.UNIX = t.Unix() //sets unix response data
			response.UTC = t.UTC().Format("Mon, 02 Jan 2006 3:04:05 GMT") //sets utc response data
			return response, nil //returns response and no error
		} else {
			var err = errors.New("invalid time format") //if parsing error
			var response Response
			return response, err //empty response and error
		}
	} else { //for UNIX format timestamp given
		timeInt, _ := strconv.Atoi(timestamp) //converts the string to int
		utcTime := time.Unix(int64(timeInt), 0) //converts the unix to utc
		var response Response
		response.UTC = utcTime.Format("Mon, 02 Jan 2006 3:04:05 GMT") //sets utc response data
		response.UNIX = int64(timeInt) //sets unix response data
		return response, nil //returns response and no error
	}
}

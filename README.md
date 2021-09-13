# Restful-api

Created a RESTful API with two endpoints. The code is written in Golang without using framework (includes mux and
routers). It uses “net/http” package, which provides http client and server implementations.

### Main Structure

main.go => Controllers => Services => Repositories => Database

#### First Part: Fetch data from mongodb

- Endpoint: https://getir-go.herokuapp.com/filtered-records

Sample request:

```
{
    "startDate": "2016-01-26",
    "endDate": "2018-02-02",
    "minCount": 2700,
    "maxCount": 3000
}
```

Sample success response:

```
{
    "code": 0,
    "msg": "Success",
    "records": [
        {
            "key": "ibfRLaFT",
            "createdAt": "2016-12-25T16:43:27.909Z",
            "totalCount": 2892
        },
        {
            "key": "pxClAvll",
            "createdAt": "2016-12-19T10:00:40.05Z",
            "totalCount": 2772
        },
        .
        .
        .
    ]
}
```

Unsuccessful Requests and Their Codes:

- code 1: If the request structure is invalid
- code 2: If there is/are missing field(s)
- code 3: If minCount is greater than maxCount
- code 4: If the start date format is not correct
- code 5: If the end date format is not correct
- code 6: If start date is greater than end date in request
- code 7: If there are no records found after filtering

#### Second Part: In memory endpoints

a) POST endpoint: https://getir-go.herokuapp.com/in-memory/set

- The request payload of POST endpoint includes a JSON with 2 fields (key and value). The response returns the echo of
  the request.

Sample request and response:

```
{
    "key": "active-tabs",
    "value": "getir"
}
```

- If key or value is missing in the request payload, the response would be “Error: Please provide the key” or “Error:
  Please provide the value”.
- If the request structure is invalid; the response would be “Invalid request structure: ”, and then the request error.

b) GET endpoint: https://getir-go.herokuapp.com/in-memory?key=(provide-key-here)

- The request payload of GET endpoint includes 1 query parameter, which is the “key” parameter. The response returns the
  key and value.

Sample: https://getir-go.herokuapp.com/in-memory?key=active-tabs

- If the key parameter is missing, the response returns “Error: Missing key name in query string!”.
- If the key does not exist, the response returns “This key does not exist”.

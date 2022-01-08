# Go Getir Case

Go Getir Study Case

## Installation

Run:

+ Standard
   ```
   go run main.go
  ```
+ Docker Container
    ``` 
    docker build . -t go-getir-challenge
    ```
    
    ```
    docker run -p 4444:4444 go-getir-challenge
    ```
# Heroku Link
https://lit-mountain-77398.herokuapp.com/

# Endpoints
+ set
+ get/{key}
+ fetch

# What is added?
+ Docker Container
+ Postman Tests
+ Unit Tests
+ Error Handling


# Fetch Data

**URL** : `/fetch`

**Method** : `POST`


**Example Data**

```json
{
    "startDate": "2016-01-26",
    "endDate": "2018-02-02",
    "minCount": 2700,
    "maxCount": 3000
}
```

+ “startDate” and “endDate” fields will contain the date in a “YYYY-MM-DD” format.
+ "minCount" and "maxCount" should be positive and "maxCount" should be greater than "minCount"


## Success Response

**Code** : `200`

**Success Response**

```json
{
"code":0,
    "msg":"Success",
    "records": [
        {
            "key":"TAKwGc6Jr4i8Z487",
            "createdAt":"2017-01-28T01:22:14.398Z",
            "totalCount":2800
}, {
} ]
}
```


# Set Data
```
- The request payload of POST endpoint will include a JSON with 2 fields.
- “key” fields holds the key (any key in string type)
- “value” fields holds the value (any value in string type)
```

**URL** : `/set`

**Method** : `POST`

**Example Data**

```json
{
   "key":"active-tabs",
   "value":"getir"
}
```

## Success Response

**Code** : `200`


```json
{
   "key":"active-tabs",
   "value":"getir"
}
```

# Get Data
```
- The request payload of GET endpoint will include 1 query parameter. That is “key” param holds the key
```
**URL** : `/get`

**Method** : `GET`

**Example Request**

```
 http://localhost/get?key=active-tabs
```

## Success Response

**Code** : `200`


```json
{
   "key":"active-tabs",
   "value":"getir"
}
```


# Unit Test

+ All Run :
 
 ```
go test
 ```
 
+ Related Folder Run:
 ``` 
go test -v
 ```
 

+ Coverage Report:
 ```
go test -cover
 ```
 
+ Coverege Report on Browser:
```
go test --coverprofile=coverage.out && go tool cover -html=coverage.out  --> we can see coverage report on html
# Currency-exchange

### Things to know inother to get started with this project.

#### Ensure you have the [Go](https://go.dev/doc/install) compiler installed.

- This task can be tested both locally or using the docs page on the live [url](https://currency-exchange-g8p8.onrender.com/api/v1/docs/index.html) with which [swagger](github.com/swaggo/gin-swagger) was used for it's documentation, it also provides an environment to test API endpoint(s) live.

- Golang is used with the gin framework for http handling. APIs are in REST.

- This task was hosted on the [render.com](https://render/com) server to enable remote access. With the live base URL as [https://currency-exchange-g8p8.onrender.com/api/v1](https://currency-exchange-g8p8.onrender.com/api/v1)

### To run this project locally
- Clone the [github repo](https://github.com/iqquee/currency-exchange) or unzip the project.
- Open the project folder and run this command `go get ./...` to download all project dependencies.
- After project dependencies are downloaded, run the command `go run main.go` to start up the local server which runs on port `5002`
- This is the base-url for testing locally `127.0.0.1:5002/api/v1`

#### API endpoint
- This is the API endpoint to test for currency exchange rate locally `127.0.0.1:5002/api/v1/exchange/rates`
- This is a `POST` request method.
- It takes in a request body of `JSON` with payload as 
```json
{
    "currency-pair": "USD-EUR"
}
```
- The server response is in JSON and the response looks like this
```json
{
    "code": 200,
    "data": {
        "USD-EUR": 0.92
    },
    "message": "successful currency conversion rate",
    "error": null
}
```
### Technicalities
- There is a validation function written to validate the currency pair privided.
- A currency.json file was created to validate the incoming currency type. Below are the list of currencies in the ./pkg/currency/currency.json file which you can use to test as a valid currency type,
```json
[
    {
        "ID":"1",
        "Currency": "USD" 
    },
    {
        "ID":"2",
        "Currency": "EUR"
    },
    {
        "ID":"3",
        "Currency": "AUD" 
    },
    {
        "ID":"4",
        "Currency": "RUB" 
    },
    {
        "ID":"5",
        "Currency": "QAR" 
    },
    {
        "ID":"6",
        "Currency": "PLN" 
    },
    {
        "ID":"7",
        "Currency": "GBP" 
    },
    {
        "ID":"8",
        "Currency": "NGN" 
    }
]
```
- A `go routine` was wraped around the `for` loop in the`ExchangeRateReq{}.validate()` method to validate both currencies against the `currency.json file` concurrently inother to speed up the response process.
- A `go routine` was also used in `Controller{}.ServerResponse()` in the `controller package` to make concurrent request to the CurrencyServerA() and CurrencyServerB() which are our mock server. 
- - The `go routine` makes a request to both servers concurrently and returns a response from the first server to respond with a success. 
- - If eventually the first server encountered an error, it will wait for the second server for response and return the response if successful. But there is a case in which both servers could return errors and only in this situation will it return an error.

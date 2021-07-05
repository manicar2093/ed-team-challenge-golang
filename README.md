# Cryptocurrencies Charter

This is the code for ed.tead challenge. It has many errors, but I make this in one nigth xD

## How to run?

It is needed the next Environment variables:
- **NOMICS_KEY**: The key that will be used to communicate with Nomics API

If this does not exist a **panic** will be raised indicating this missing env variable

**NOTE:** The port to use will be 5000 so make sure you're not using it already.

To run just type in yot terminal:
```
make run
```
or
```
go run cmd/api/main.go
```
And that will be enough

## How to make a request?

The JSON is like this:
```
{
    "cryptos": ["BTC", "ETH"],
    "from_date": "2021-01-01T00:00:00-05:00"
}
```

## IMPORTANT CONSIDERATIONS: 
- Dates must have the RFC3339 format to be accepted. In the opposite case a 400 Bad Request will be responded with a JSON like this:

```
{
    "message": "request does not satisfy requirements. check documentation"
}
```
- If you request more than one currency chart might not look pretty well due currencies change amount. Consider request just one currency at a time.

## Comments
This issue:

```
Yes, this has failures. I don't know why but the chart image can't be downloaded from Postman. Besides that I feel the code organizations is great and I'd like to have your review :D
```
Has been solved. :D
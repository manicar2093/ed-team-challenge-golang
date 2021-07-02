# Cryptocurrencies Charter

This is the code for ed.tead challenge. It has many errors, but I make this in one nigth xD

## How to run?

It is needed the next Environment variables:
- **FILES_PATH**: Indicates were charts will be saved in the server.
- **NOMICS_KEY**: The key that will be used to communicate with Nomics API

If one of this does not exist a **panic** will be raised indicating wich env variable is missing

**NOTE:** The port to use will be 5000 so make sure you're not using it already.

To run just type in yot terminal:
```
make run
```
or
```
go run main.go
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

**IMPORTANT:** It is importan date has the RFC3339 format to be accepted. In the opposite case a 400 Bad Request will be responded with a JSON like this:

```
{
    "message": "request does not satisfy requirements. check documentation"
}
```

## Comments

Yes, this has failures. I don't know why but the chart image can't be downloaded from Postman. Besides that I feel the code organizations is great and I'd like to have your review :D

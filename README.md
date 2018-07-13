# Request / Response Recorder

This service allows user to setup predefined response for given request. Which allows stubbing 3rd party services in f.e. e2e testing.

To setup simply run 
```
docker-compose build
docker-compose up
```
This will setup every needed service.

### How To

NOTE: Currently only `LawPay - recurring charge` endpoint is implemented. 

### Recording Response

#### 1.CREATE
In order to record response one needs to `POST localhost:8080/record/create` with correct `response` data which can be found [here](https://developers.affinipay.com/reference/api.html#RecurringCharges)

Example data:
```json
{
    "ID": "adfadfasgasgasfaf",
    "Status": "Active",
    "account_id": "TSTTSTTSTSTST",
    "Method": {
        "Type": "card",
        "Number": "4242424242424242",
        "Fingerprint": "GunPelYVthifNV63LEw1",
        "card_type": "VISA",
        "exp_month": 10,
        "exp_year": 2022,
        "Name": "Test Customer"
    },
    "Schedule": {
        "Start": "2016-01-01",
        "interval_unit": "MONTH",
        "interval_delay": 1
    },
    "Description": "Monthly recurring charge",
    "Amount": 1250,
    "Currency": "USD",
    "total_occurrences": 0,
    "total_amount": 0,
    "next_payment": "2016-01-01",
    "Occurrences": [
        {
            "ID": "_LIG1tsDQZ21oBgPYTRJdQ",
            "Amount": 1250,
            "Status": "PENDING",
            "due_date": "2016-01-01",
            "Attempts": 0
        }
    ]
}
```

When response is recorded you can start using this service as a mock. By hitting it with `POST localhost:8080/lawpay` request with valid create recurring charge data.
```json
{
	"description": "Monthly recurring charge",
	"account_id": "TSTTSTTSTSTST",
	"amount": "1250",
	"method": "OtmNJP6YRpKrcJ0RdZxGcw",
	"schedule": {
		"start": "2016-01-01",
		"interval_unit": "MONTH",
		"interval_delay": 1
	}
}
```

and you receive previously predefined response.

#### 2.GET

In order to mimics `GET recurrnig charge` one needs to hit `GET localhost:8080/lawpay/<ID>`, Where `<ID>` is one of fields which we get as a response for `POST create` action. 
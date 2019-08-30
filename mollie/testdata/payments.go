package testdata

//GetPaymentResponseExample describes how a response should
//look like when calling the Payments API with any of the
//methods expecting a single object response.
var GetPaymentResponseExample = `{
    "resource": "payment",
    "id": "tr_7UhSN1zuXS",
    "mode": "test",
    "createdAt": "2018-03-20T09:13:37+00:00",
    "amount": {
        "value": "10.00",
        "currency": "EUR"
    },
    "description": "Order #12345",
    "method": null,
    "metadata": {
        "order_id": "12345"
    },
    "status": "open",
    "isCancelable": false,
    "expiresAt": "2018-03-20T09:28:37+00:00",
    "details": null,
    "profileId": "pfl_QkEhN94Ba",
    "sequenceType": "oneoff",
    "redirectUrl": "https://webshop.example.org/order/12345/",
    "webhookUrl": "https://webshop.example.org/payments/webhook/",
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments/tr_7UhSN1zuXS",
            "type": "application/json"
        },
        "checkout": {
            "href": "https://www.mollie.com/payscreen/select-method/7UhSN1zuXS",
            "type": "text/html"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/payments-api/create-payment",
            "type": "text/html"
        }
    }
}`

//ListPaymentsExample for mocking
var ListPaymentsExample = `{
    "count": 1,
    "_embedded": {
        "payments": [
            {
                "resource": "payment",
                "id": "tr_7UhSN1zuXS",
                "mode": "test",
                "createdAt": "2018-02-12T11:58:35.0Z",
                "expiresAt": "2018-02-12T12:13:35.0Z",
                "status": "open",
                "isCancelable": false,
                "amount": {
                    "value": "75.00",
                    "currency": "GBP"
                },
                "description": "Order #12345",
                "method": "ideal",
                "metadata": null,
                "details": null,
                "profileId": "pfl_QkEhN94Ba",
                "redirectUrl": "https://webshop.example.org/order/12345/",
                "_links": {
                    "checkout": {
                        "href": "https://www.mollie.com/paymentscreen/issuer/select/ideal/7UhSN1zuXS",
                        "type": "text/html"
                    },
                    "self": {
                        "href": "https://api.mollie.com/v2/payments/tr_7UhSN1zuXS",
                        "type": "application/hal+json"
                    }
                }
            }
        ]
    },
    "_links": {
        "self": {
            "href": "https://api.mollie.com/v2/payments?limit=5",
            "type": "application/hal+json"
        },
        "previous": null,
        "next": {
            "href": "https://api.mollie.com/v2/payments?from=tr_SDkzMggpvx&limit=5",
            "type": "application/hal+json"
        },
        "documentation": {
            "href": "https://docs.mollie.com/reference/v2/payments-api/list-payments",
            "type": "text/html"
        }
    }
}`

# Mollie API Golang client 

[![Build Status](https://travis-ci.org/VictorAvelar/mollie-api-go.svg?branch=master)](https://travis-ci.org/VictorAvelar/mollie-api-go)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-85%25-brightgreen.svg?longCache=true&style=flat-square)</a>

Accepting [iDEAL](https://www.mollie.com/payments/ideal/), [Apple Pay](https://www.mollie.com/payments/apple-pay), [Bancontact](https://www.mollie.com/payments/bancontact/), [SOFORT Banking](https://www.mollie.com/payments/sofort/), [Creditcard](https://www.mollie.com/payments/credit-card/), [SEPA Bank transfer](https://www.mollie.com/payments/bank-transfer/), [SEPA Direct debit](https://www.mollie.com/payments/direct-debit/), [PayPal](https://www.mollie.com/payments/paypal/), [Belfius Direct Net](https://www.mollie.com/payments/belfius/), [KBC/CBC](https://www.mollie.com/payments/kbc-cbc/), [paysafecard](https://www.mollie.com/payments/paysafecard/), [ING Home'Pay](https://www.mollie.com/payments/ing-homepay/), [Giftcards](https://www.mollie.com/payments/gift-cards/), [Giropay](https://www.mollie.com/payments/giropay/), [EPS](https://www.mollie.com/payments/eps/) and [Przelewy24](https://www.mollie.com/payments/przelewy24/) online payments without fixed monthly costs or any punishing registration procedures. Just use the Mollie API to receive payments directly on your website or easily refund transactions to your customers.

## Requirements ##
To use the Mollie API client, the following things are required:

+ Get yourself a free [Mollie account](https://www.mollie.com/signup). No sign up costs.
+ Now you're ready to use the Mollie API client in test mode.
+ Follow [a few steps](https://www.mollie.com/dashboard/?modal=onboarding) to enable payment methods in live mode, and let us handle the rest.
+ Up-to-date OpenSSL (or other SSL/TLS toolkit)

For leveraging [Mollie Connect](https://docs.mollie.com/oauth/overview) (advanced use cases only), it is recommended to be familiar with the OAuth2 protocol.

## Roadmap
- [x] Project setup
- [x] Base HTTP Client
- [x] Authentication
    - [x] API token
    - [x] Organization token
- [ ] Payments
    - [ ] Recurring payments
    - [ ] Multicurrency payments
    - [ ] QR Codes
    - [ ] Gift cards
    - [ ] Webhook (Status changes)
- [ ] Orders
    - [ ] Discounts
- [ ] Wallets
    - [ ] Apple pay
- [ ] Mollie Connect
    - [ ] App fees
    - [ ] Onboarding
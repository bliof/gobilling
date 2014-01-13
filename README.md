# GoBilling

Provides a unified API for different payment gateways.

Inspired by [active_merchant](https://github.com/Shopify/active_merchant)

## Installation

    go get github.com/bliof/gobilling

## How should work

    gateway := gateway.PayPal{
        User:      "TestMerchant",
        Password:  "password",
        Signature: "ashjdfasdkf",
    }

    amount := 20.0

    // The verificationValue is also known as CVV2, CVC2, or CID
    creditCard := gobilling.CreditCard{
        FirstName:         "Rose",
        LastName:          "Tyler",
        Number:            "4222222222222",
        Month:             "9",
        Year:              time.Now().Year() + 1,
        VerificationValue: "000",
    }

    // When validating the credit card, the Brand will be automaticly filled
    err := creditCard.Validate()

    if err == nil {
        response := gateway.Purchase(amount, creditCard)

        if response.IsSuccessful() {
            fmt.Printf("Charged %.2f to the credit card %s", amount, creditCard.DisplayNumber())
        }
    }

    // Output:
    // Charged 20.00 to the credit card XXXX-XXXX-XXXX-2222

## TODO

Design the structure of the lib.

Provide a working gateway plugin (most likely with PayPal Payments Pro).

Provide a working integration (when redirecting to perform the payment process on another site) plugin.

## License

The MIT License (MIT)

Copyright (c) 2014 Aleksandar Ivanov

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

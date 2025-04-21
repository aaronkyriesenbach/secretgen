# secretgen

Lightweight Go API capable of generating multiple passwords at a time as well as ECDSA private keys. Intended for use as
an External Secrets Operator webhook generator.

## Supported endpoints:

### /password

Returns a set of generated secrets with specified names.

Parameters:

| Name        | Type     | Description                                    | Default       |
|-------------|----------|------------------------------------------------|---------------|
| names       | string[] | List of secret names to generate               | ["secretKey"] |
| length      | test     | Length of each secret                          | 64            |
| numDigits   | int      | Number of digits to be present in each secret  | 8             |
| numSymbols  | int      | Number of symbols to be present in each secret | 8             |
| noUpper     | bool     | Disallow uppercase characters                  | false         |
| allowRepeat | bool     | Allow repeated characters                      | true          |

### /ecdsa

Returns a base64-encoded PEM format ECDSA P-384 private key.
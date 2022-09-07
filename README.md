# TechCamp 2022 - JWT101
Notes, slides and sources of "TechCamp 2022 - JSON Web Token 101" talk.

https://techcamp.hamburg/

## Disclaimer

THE CONTENT OF THIS REPOSITORY INCLUDING ANY REFERENCED EXTERNAL CONTENT IS FOR EDUCATIONAL PURPOSE ONLY. 

IT COMES WITHOUT WARRANTY OF ANY KIND, CORRECTNESS AND/OR FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.

USE AT YOUR OWN RISK!

Any source code included is licensed under the MIT License (see the repositories `LICENSE` file)

## Slides
* [Google Docs](https://docs.google.com/presentation/d/1ojRa7PQOXPtTF2CuPESm0B3LS9fpKelZHpL-R41guX0)
* [PDF](./TC2022-JWT101.pdf)

## RFCs
* [RFC 7518 - JSON Web Algorithms (JWA)](https://www.rfc-editor.org/rfc/rfc7518)
* [RFC 7519 - JSON Web Token (JWT)](https://www.rfc-editor.org/rfc/rfc7519)
* [RFC 2104 - HMAC: Keyed-Hashing for Message Authentication](https://www.rfc-editor.org/rfc/rfc2104)
* [RFC 4648 - Base 64 Encoding with URL and Filename Safe Alphabet](https://www.rfc-editor.org/rfc/rfc4648#section-5)

## Signing algorithms

### Code snippets
* [HS256 code example](examples/signing-algorithms/HS256-code-example.go)
* [RS256 code example](examples/signing-algorithms/RS256-code-example.go)

## Common vulnerabilities 

### Code snippets
* [Vulnerable verification example - GoLang](examples/common-vulnerabilities/vulnerable-code-example.go)
* [Constant-time algorithm - GoLang](examples/common-vulnerabilities/constant-time-algorithm.go)
* [Constant-time algorithm - PHP](examples/common-vulnerabilities/constant-time-algorithm.php)
* [Constant-time algorithm - Python](examples/common-vulnerabilities/constant-time-algorithm.py)

### Examples
* [CVE-2016-5431 - Key Confusion Attack](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2016-5431)
* [CVE-2015-9235 - alg:none Attack](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2015-9235)
* [CVE-2020-28042 - Null Signature Attack](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-28042)
 

## Hands on

### Six-digit passphrase recovery (`secret`)

```bash
hashcat --increment-min=4 --increment-max=8 --increment -m16500 \
  'eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJ0ZWNoY2FtcC5oYW1idXJnIiwianRpIjoiOTZmNzBkODkiLCJpYXQiOjE2NjEzMzAzMDgsIm5iZiI6MTY2MTMzMDMwOCwiZXhwIjoxNjYxMzM3NTA4LCJzdWIiOiJtLnJlaWNoZWwiLCJpc3MiOiJpZC50ZWNoY2FtcC5oYW1idXJnIn0.mKdydmAO5Mh6bHFBtguwLAdLtxIR3oczRl7hCjsiK0w' \
   -a3 -1 "?l" "?1?1?1?1?1?1?1?1" -D 2 -d 5 -w 3
```

### Seven-digit passphrase recovery (`hamburg`)

```bash
hashcat --increment-min=4 --increment-max=8 --increment -m16500 \
  'eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJ0ZWNoY2FtcC5oYW1idXJnIiwianRpIjoiOTZmNzBkODkiLCJpYXQiOjE2NjEzMzAzMDgsIm5iZiI6MTY2MTMzMDMwOCwiZXhwIjoxNjYxMzM3NTA4LCJzdWIiOiJtLnJlaWNoZWwiLCJpc3MiOiJpZC50ZWNoY2FtcC5oYW1idXJnIn0.tVtDVw5BlIYEQt1lVdo0YFdlS9yrNvQR0JnGU81DYQA' \
   -a3 -1 "?l" "?1?1?1?1?1?1?1?1" -D 2 -d 5 -w 3
```

## Sources, references and further readings

Big thanks and kudos going to:

* https://techcamp.hamburg/
* https://silpion.de/
* https://jwt.io/
* https://github.com/hashcat/hashcat
* https://github.com/ticarpi/jwt_tool
* https://cheatsheetseries.owasp.org/cheatsheets/JSON_Web_Token_for_Java_Cheat_Sheet.html
* https://book.hacktricks.xyz/pentesting-web/hacking-jwt-json-web-tokens
* https://auth0.com/blog/json-web-token-signing-algorithms-overview/
* https://0xn3va.gitbook.io/cheat-sheets/web-application/json-web-token-vulnerabilities
* https://codahale.com/a-lesson-in-timing-attacks/

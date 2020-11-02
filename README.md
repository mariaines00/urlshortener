# urlshortener-technical-test

### Description

The challenge is to build a HTTP-based RESTful API for managing Short URLs and redirecting
clients similar to bit.ly or goo.gl. (Be thoughtful that the system must eventually support millions
of short urls.
Please include a README with documentation on how to build, and run and test the system. 

A Short Url:
1. Has one long url
2. Permanent; Once created
3. Is Unique; If a long url is added twice it should result in two different short urls.
4. Not easily discoverable; incrementing an already existing short url should have a low
probability of finding a working short url.

Your solution must support:
1. Generating a short url from a long url
2. Redirecting a short url to a long url within 10 ms.
3. Listing the number of times a short url has been accessed in the last 24 hours, past
week and all time.
4. Persistence (data must survive computer restarts)

Shortcuts
1. No authentication is required
2. No html, web UI is required
3. Transport/Serialization format is your choice, but the solution should be testable via curl
4. Anything left unspecified is left to your discretion
---

## How to run:

## TDB:
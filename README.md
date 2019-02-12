# selitra
[![Go Report Card](https://goreportcard.com/badge/github.com/saromanov/selitra)](https://goreportcard.com/report/github.com/saromanov/selitra)
[![Build Status](https://travis-ci.org/saromanov/selitra.svg?branch=master)](https://travis-ci.org/saromanov/selitra)

Log processing tool

Example of request
```
curl -X POST -H "Content-Type: application/json" -d '{"level": "LOW", 
  "message":"Awesome log output",
  "name":"test service",
  "labels":["dev","test"],
  "entry":"service",
  "service": "first"
}' http://127.0.0.1:6320/api/selitra/stats
```

## Query
```
!service=first;date=today;level=ERROR"
```

## Server stat

Return current server stat
```
http://127.0.0.1:6320/api/selitra/server
```
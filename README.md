# selitra

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
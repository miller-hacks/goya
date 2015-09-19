# goya
Our brand-new bleeding-edge server

## How to use

POST data with image url

```
POST / HTTP/1.1
Accept: application/json
Accept-Encoding: gzip, deflate
Content-Length: 53
Content-Type: application/json; charset=utf-8
Host: 127.0.0.1:8001
User-Agent: HTTPie/0.8.0

{
    "url": "http://example.com/image.png"
}
```

response:

```
HTTP/1.1 200 OK
Content-Length: 43
Content-Type: application/json
Date: Sat, 19 Sep 2015 10:45:59 GMT

[
    {
        "height": 264,
        "width": 264,
        "x": 39,
        "y": 39
    }
]
```

### Amazing web interface

in progress...

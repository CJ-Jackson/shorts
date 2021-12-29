# Shorts

A simply engineered URL shortener that can be manage by terminal.
No databases or caches just files and that it.

# Usage

```
# Let's run shorts in the background
$ shorts -a :8080 -u /home/user/shorts &

# Let's go to that directory
$ cd /home/user/shorts

# Add a url, must have .txt extension or it won't work.
$ echo "https://www.example.com" > example.txt

# Let's test it with curl.
$ curl -i http://127.0.0.1:8080/example
HTTP/1.1 308 Permanent Redirect
Location: https://www.example.com
Date: Wed, 29 Dec 2021 20:46:43 GMT
Content-Length: 0

# Let's remove it
rm example.txt

$ curl -i http://127.0.0.1:8080/example
HTTP/1.1 404 Not Found
Date: Wed, 29 Dec 2021 20:48:31 GMT
Content-Length: 10
Content-Type: text/plain; charset=utf-8

Not Found
```
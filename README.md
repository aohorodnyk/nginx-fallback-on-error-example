## Nginx Reverse Proxy fallback URL Example
There is a simple example with fallback URL in the nginx as a reverse proxy.

### Start Instructions
For using this example you should have installed [docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/).

For start this example you should run command in console located in the project directory:
```bash
$ docker-compose up
```

As a result you'll have started project on the port 8080. For checking rules you should request the url with http code:
```bash
$ curl http://127.0.0.1:8080/404
{"Parsed Code":"200","Source URL":"/200","Status Text":"OK"}
$ curl http://127.0.0.1:8080/400
{"Parsed Code":"400","Source URL":"/400","Status Text":"Bad Request"}
```

Check requests to the golang applciations you could see in the console with started `docker-compose`:
```bash
codes    | 2018/02/25 12:34:20 Listening...
codes    | {"Parsed Code":"404","Source URL":"/404","Status Text":"Not Found"}
codes    | {"Parsed Code":"200","Source URL":"/200","Status Text":"OK"}
nginx    | 172.18.0.1 - - [25/Feb/2018:12:34:22 +0000] "GET /404 HTTP/1.1" 200 60 "-" "curl/7.54.0" "-"
codes    | {"Parsed Code":"400","Source URL":"/400","Status Text":"Bad Request"}
nginx    | 172.18.0.1 - - [25/Feb/2018:12:34:24 +0000] "GET /400 HTTP/1.1" 400 69 "-" "curl/7.54.0" "-"
```
Follow logs from the docker console we can see that:
* First request `curl http://127.0.0.1:8080/404` to the nginx triggered two requets from the nginx to the codes application
* Second request `curl http://127.0.0.1:8080/400` to the nginx triggered only one request from the nginx to the codes application

### Change codes with fallback requests
If you want to check some other codes with fallback requests you should open `./conf/www.conf` in the project directory and change list of codes in error page line.
For example by default was set three codes (403, 404, 503) `error_page 403 404 503 = @outage;` you could add code 400 `error_page 400 403 404 503 = @outage;`.
After all changes you should save file and restart `docker-compose up`.

### Restrictions
Application codes supporting only valid codes in other cases you'll get 500 response with ERROR message.

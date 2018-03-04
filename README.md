## Nginx Reverse Proxy fallback URL Example
There is a simple example with fallback URL in the nginx as a reverse proxy.

### Start Instructions
For using this example you should have installed [docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/).

For starting this project you should run the command in a console, you should be located in the project directory:
```bash
$ docker-compose up
```

As a result you'll have the started project on the port 8080. For checking rules you should request the url with a http code, for exaple:
```bash
$ curl http://127.0.0.1:8080/404
{"Parsed Code":"200","Source URL":"/200","Status Text":"OK"}
$ curl http://127.0.0.1:8080/400
{"Parsed Code":"400","Source URL":"/400","Status Text":"Bad Request"}
```

Check result of your previous requests in the console where you started `docker-compose`:
```bash
codes    | 2018/02/25 12:34:20 Listening...
codes    | {"Parsed Code":"404","Source URL":"/404","Status Text":"Not Found"}
codes    | {"Parsed Code":"200","Source URL":"/200","Status Text":"OK"}
nginx    | 172.18.0.1 - - [25/Feb/2018:12:34:22 +0000] "GET /404 HTTP/1.1" 200 60 "-" "curl/7.54.0" "-"
codes    | {"Parsed Code":"400","Source URL":"/400","Status Text":"Bad Request"}
nginx    | 172.18.0.1 - - [25/Feb/2018:12:34:24 +0000] "GET /400 HTTP/1.1" 400 69 "-" "curl/7.54.0" "-"
```
Following the logs from the docker-compose console we can see that:
* The first command `curl http://127.0.0.1:8080/404` with the request to the nginx triggered two requests from the nginx to the application of the codes. So we can see the triggered fallback to the 404 code response.
* The second command `curl http://127.0.0.1:8080/400` with the request to the nginx triggered only one request from the nginx to the application of the codes.

### Change codes which triggering fallback requests
If you want to check some other codes with fallback requests you should open `./conf/www.conf` in the project directory and change list of codes in the error page line.
For example was set three codes (403, 404, 503) by default in the `error_page 403 404 503 = @outage;` line, so you could add code 400 `error_page 400 403 404 503 = @outage;` in this line.
After finished changes, you have to save file and restart `docker-compose up`.

### Restrictions
* Application codes supporting only valid codes in other cases you'll get 500 response with the error message.

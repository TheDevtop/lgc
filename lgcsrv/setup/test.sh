#!/bin/sh

echo 'Test: /sys/time' && curl http://localhost:8080/sys/time
echo 'Test: /sys/pipe' && curl -X PUT http://localhost:8080/sys/pipe -d 'This is sparta' && curl http://localhost:8080/sys/pipe
echo 'Test: /sys/log' && curl -X PUT http://localhost:8080/sys/log?module=curl -d 'Doing work...' && curl http://localhost:8080/sys/log
echo 'Test: /env' && curl -X PUT 'http://localhost:8080/env?key=hotel&val=trivago' && curl http://localhost:8080/env
echo 'Test: /api' && curl -X PUT --header "route: /weather" --header "dest: http://wttr.in/Amsterdam" http://localhost:8080/api && curl -X GET http://localhost:8080/api
echo 'Test API:' && curl --insecure https://localhost/weather
exit 0

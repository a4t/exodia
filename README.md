# Exodia

Server status check and Run script.

## Example yaml

```
wait: 30
interval: 5
checkretry: 5
scripts:
  pre:
    - service nginx stop
    - service td-agent stop
  check:
    - ps aux | grep [t]d-agent && exit 1 || exit 0
    - ps aux | grep [n]ginx && exit 1 || exit 0
  post:
    - shutdown -h now
```

## yaml spec
- wait
  - pre -> wait(Second)-> check
- interval
  - Retry interval on miss check command.
- checkretry
  - Retry max count.
- scripts
  - Run at each step.

## Run

```
$ ./exodia -f ./script.yml
$ ./exodia -h
Usage of ./exodia:
  -d  Debug
  -f string
      Setting yml file (default "./script.yml")```

## LICENSE
The MIT License (MIT)

Copyright (c) 2016- Shigure Onishi
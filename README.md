# skeddy

A lightweight cloud savvy time based scheduler which does nothing but handoff the heavy lifting to a bunch of http endpoints. Optional payload can be sent to these http endpoints.

# Basic functionality

[![ScreenCast](https://github.com/gophergala2016/skeddy/blob/master/images/video.jpg)](https://github.com/gophergala2016/skeddy/blob/master/screencast/skeddy.gif)

# Installation

- Install go packages as follows :

    ```
      $ go get github.com/pborman/uuid
      $ go get github.com/robfig/cron
      $ go get github.com/syndtr/goleveldb/leveldb
    ```
- Install the project and build it

  ```
  $ go get https://github.com/gophergala2016/skeddy
  $ go build
  ```
- Start the server

  ```
    $ skeddy [-d] [-i] [-p] [-s]
  ```

### Options

  ``` -d dbname.db ```
    use existing storage (default "skeddy.db")

  ``` -i cronTab.txt ```
    import from crontab file

  ``` -p port ```
    port to listen (default 8080)

  ``` -s host-name ```
    bind to ip address (default "0.0.0.0")


### CRON Table Format

```
  <cron_expression> <http_endpoint> <payload>
```

cron_expression should be in accordance to the cron expression format given [here](https://godoc.org/github.com/robfig/cron), payload can be raw json or json file from where payload will be read. JSON file should be given as @filename. While using UI interface, you can also upload a file whose content will be read and send as payload.

# Usage

skeddy can be used to send periodic notifications to http endpoints. It has a user friendly UI interface with which anyone can hit my server and add their events in my skeddy. This can be used cross-platform.

# License

  MIT

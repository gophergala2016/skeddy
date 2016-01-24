# skeddy

Skeddy is a lightweight scheduler that can be configured to initiate http requests in fire-and-forget mode. Configuration options are inspired from cron. It is a cloud savvy time based scheduler.

# Basic functionality

![skeddy](https://github.com/gophergala2016/skeddy/blob/master/screencast/skeddy.gif)

# Installation

- Install go packages as follows :

    ```
      $ go get github.com/pborman/uuid
      $ go get github.com/robfig/cron
      $ go get github.com/syndtr/goleveldb/leveldb
    ```
- Clone the project and build it

  ```
  $ git clone https://github.com/gophergala2016/skeddy.git
  $ go build
  ```
- Start the server

  ```
    $ ./skeddy [-d] [-i] [-p] [-s]
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

cron_expression should be in accordance to the cron expression format given [here](https://godoc.org/github.com/robfig/cron), payload is optional and can be raw json or json file from where payload will be read. JSON file should be given as @filename. While using UI interface, you can also upload a file whose content will be read and send as payload.

# Usage

Skeddy can be used to send periodic notifications to http endpoints. It has a user friendly UI interface with which anyone can hit my server and add their events in my skeddy. This can be used cross-platform unlike cron (used only in UNIX). Skeddy uses the same format as cron scheduler.

# Motivation

Motivation to build skeddy came because I wanted to use the efficient cron scheduler cross-platform and wanted it to be cloud savvy, so that I can add entries to my scheduler from anywhere.

# License

  MIT

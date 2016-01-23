# skeddy

A time based scheduler which does nothing but handoff the heavy lifting to a bunch of http endpoints. Optional payload can be sent to these http endpoints.

# Installation

```
$ go get https://github.com/gophergala2016/skeddy
```
### CRON Table Format

```
<cron_expression> <http_endpoint> <payload>
```

cron_expression should be in accordance to the cron expression format, payload can be raw json or json file from where payload will be read. JSON file should be given as @filename.json.


### CRON Expression Format:

A cron expression has 6 space-separated fields.

Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
Note: Month and Day-of-week field values are case insensitive. "SUN", "Sun", and "sun" are equally accepted.

#### Special Characters
**$ Asterisk ( * )**

The asterisk indicates that the cron expression will match for all values of the field; e.g., using an asterisk in the 5th field (month) would indicate every month.

**$ Slash ( / )**

Slashes are used to describe increments of ranges. For example 3-59/15 in the 1st field (minutes) would indicate the 3rd minute of the hour and every 15 minutes thereafter. The form **"*\/..."** is equivalent to the form "first-last/...", that is, an increment over the largest possible range of the field. The form "N/..." is accepted as meaning "N-MAX/...", that is, starting at N, use the increment until the end of that specific range. It does not wrap around.

**$ Comma ( , )**

Commas are used to separate items of a list. For example, using "MON,WED,FRI" in the 5th field (day of week) would mean Mondays, Wednesdays and Fridays.

**$ Hyphen ( - )**

Hyphens are used to define ranges. For example, 9-17 would indicate every hour between 9am and 5pm inclusive.

**$ Question mark ( ? )**

Question mark may be used instead of **'*'** for leaving either day-of-month or day-of-week blank.

**$ Pre-defined schedules**
Cron Expression may also take following pre-defined schedules

Entry                  | Description                                | Equivalent To
-----                  | -----------                                | -------------
@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
@monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
@weekly                | Run once a week, midnight on Sunday        | 0 0 0 * * 0
@daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
@hourly                | Run once an hour, beginning of hour        | 0 0 * * * *


### TODO

Provide a view to add, edit and delete entries in the scheduler. Functionality to restart the scheduler when entries are added and deleted.

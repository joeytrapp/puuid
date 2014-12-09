# puuid

Print a new uuid in the console.

## Options

* `-n <number>`: Print that many uuids spread across newlines.
* `-x <number>`: Specify the UUID version to use. Currently only version 1 and 4 are supported.
* `-s '<string>'`: String to be used as the separator when joining the list of uuids.
* `-f '<string>'`: Format string applied to each uuid (%s replaced with uuid).
* `-b '<string>'`: String added before the list of uuids.
* `-a '<string>'`: String added after the list of uuids.
* `-t`: Trim whitespace off of final string before printing.
* `-v`: Show version
* `-h`: Show help

**Note**

Options that take a string value may contain '\n' or '\t', and they will translate to a newline and tab respectively.

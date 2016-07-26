# Go echo log

## Presentation

That project offers a logging module that can be used by backend project written in go. That module
configures echo server to use logrus to log data. In development phase the logs will be printed in
standard output whereas in production the logs will be written to a file. The log file could then be
served using a tool such as filebeat.

## Dependencies

See `vendor/vendor.json`.

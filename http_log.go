package echo_log

import (
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

// create an http handler that logs message using logrus
// inspired by middleware/logger.go of echo
// use log.withFields instead of wrapping json into a string (default behaviour)
func getLogrusMiddlewareHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			fields := make(map[string]interface{})
			fields["request_time"] = time.Now().Format(time.RFC3339)

			fields["remote_ip"] = c.RealIP()
			fields["host"] = req.Host

			fields["method"] = req.Method
			fields["uri"] = req.RequestURI
			fields["status"] = res.Status

			path := req.URL.Path
			if path == "" {
				path = "/"
			}
			fields["path"] = path

			latency := stop.Sub(start)
			fields["request_latency"] = int64(latency.Nanoseconds() / 1000)
			fields["request_latency_human"] = latency.String()

			rx := req.Header.Get(echo.HeaderContentLength)
			if rx == "" {
				rx = "0"
			}
			fields["bytes_in"], _ = strconv.ParseInt(rx, 10, 64)
			fields["bytes_out"] = int64(res.Size)

			log.WithFields(log.Fields(fields)).Info("Http request received")

			return
		}
	}
}

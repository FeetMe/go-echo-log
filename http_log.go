package log

import (
	"net"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

// create an http handler that logs message using logrus
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
			fields["time"] = time.Now().Format(time.RFC3339)

			ra := req.RemoteAddress()
			if ip := req.Header().Get(echo.HeaderXRealIP); ip != "" {
				ra = ip
			} else if ip = req.Header().Get(echo.HeaderXForwardedFor); ip != "" {
				ra = ip
			} else {
				ra, _, _ = net.SplitHostPort(ra)
			}
			fields["remote_ip"] = ra

			fields["method"] = req.Method()
			fields["uri"] = req.URI()
			fields["status"] = res.Status()

			latency := stop.Sub(start)
			fields["latency"] = int64(latency.Nanoseconds() / 1000)
			fields["latency_human"] = latency.String()

			rx := req.Header().Get(echo.HeaderContentLength)
			if rx == "" {
				rx = "0"
			}
			fields["rx_bytes"], _ = strconv.ParseInt(rx, 10, 64)
			fields["tx_bytes"] = int64(res.Size())

			log.WithFields(log.Fields(fields)).Info("Http request received")

			return
		}
	}
}

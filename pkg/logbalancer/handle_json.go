package logbalancer

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
)

func HandleJson(lb *LogBalancer, c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to read request body",
		})
	}

	var requestData pb_logs.RuntimeLogs
	err = protojson.Unmarshal(body, &requestData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Invalid JSON: %v", err),
		})
	}
	err = lb.HandleLog(&requestData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Invalid JSON: %v", err),
		})
	}
	return c.JSONBlob(http.StatusOK, []byte(`{}`))
}

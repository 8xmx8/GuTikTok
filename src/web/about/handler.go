package about

import (
	"GuTikTok/src/constant/strings"
	"GuTikTok/src/extra/tracing"
	"GuTikTok/src/utils/logging"
	"GuTikTok/src/web/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Handle(c *gin.Context) {
	_, span := tracing.Tracer.Start(c.Request.Context(), "AboutHandler")
	defer span.End()
	logging.SetSpanWithHostname(span)

	var req models.AboutReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": strings.GateWayErrorCode,
			"status_msg":  strings.GateWayError,
		})
	}
	res := models.AboutRes{
		Echo:      req.Echo,
		TimeStamp: time.Now().Unix(),
	}
	c.JSON(http.StatusOK, res)
}

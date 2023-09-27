package jsons

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"
)

// CustomJSON 是一个自定义的 JSON 渲染器类型，实现了 gin 框架的 Render 接口。
type CustomJSON struct {
	Data    proto.Message // 要序列化为 JSON 的数据
	Context *gin.Context  // gin 上下文对象
}

// m 是一个 protojson.MarshalOptions 类型的变量，用于配置序列化选项。
var m = protojson.MarshalOptions{
	EmitUnpopulated: true, // 序列化未设置值的字段
	UseProtoNames:   true, // 使用 Proto 定义中的字段名作为 JSON 键名
}

// Render 实现了 gin 框架的 Render 接口，用于渲染 HTTP 响应。
func (r CustomJSON) Render(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)       // 设置响应的内容类型为 application/json; charset=utf-8
	res, _ := m.Marshal(r.Data) // 将 r.Data 序列化为 JSON 格式
	_, err = w.Write(res)       // 将序列化后的 JSON 数据写入 http.ResponseWriter
	return
}

// WriteContentType 设置响应的内容类型为 application/json; charset=utf-8。
func (r CustomJSON) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

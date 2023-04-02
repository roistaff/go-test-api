package main

import (
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/gin-gonic/gin"
)
type INFO struct {
	Total uint64 `json:"total"`
	Free uint64 `json:"free"`
	Use uint64 `json:"use"`
}
func info_get() INFO {
	v, _ := mem.VirtualMemory()
	info := INFO{
		Total: v.Total /1024 /1024,
		Free:  v.Free /1024 /1024,
		Use:   v.Used /1024 /1024,
	}
	return info
}
func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        c.Next()
    })
	r.GET("/info", func(c *gin.Context) {
		info := info_get()
		c.JSON(200,info)
	})
	r.GET("/welcome",func(c *gin.Context){
		message := c.DefaultQuery("message","unknown")
		c.String(200,"You said : %s",message)
	})
	r.Run(":8080")
}


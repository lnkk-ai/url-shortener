package internal

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// Debug is used to print debug/error staments
func Debug(bp string, e error) {
	if e != nil {
		fmt.Println(fmt.Sprintf(" --> %s:%s", bp, e.Error()))
	}
}

// RequestLogger is used to print the GIN raw request message body
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		fmt.Println(readBody(rdr1)) // Print request body

		c.Request.Body = rdr2
		c.Next()
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}

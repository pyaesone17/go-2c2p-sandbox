package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Payment(c *gin.Context) {
	redirectURL := c.PostForm("result_url_1")
	redirectURL2 := c.PostForm("result_url_2")
	ginData := gin.H{
		"title":          "2c2p sandbox in Go lang",
		"redirectURL":    redirectURL,
		"payment_status": "000",
		"payment_scheme": "MPU",
		"currency":       "104",
		"user_defined_5": c.PostForm("user_defined_5"),
		"user_defined_1": c.PostForm("user_defined_1"),
		"user_defined_2": c.PostForm("user_defined_2"),
		"user_defined_3": c.PostForm("user_defined_3"),
		"user_defined_4": c.PostForm("user_defined_4"),
		"amount":         c.PostForm("amount"),
		"customer_email": c.PostForm("customer_email"),
	}

	data, _ := json.Marshal(ginData)
	sendPostRequest(redirectURL2, data)

	fmt.Println(redirectURL)
	c.HTML(http.StatusOK, "index.tmpl", ginData)
}

func PaymentJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"payment_status": "000",
		"payment_scheme": "MPU",
		"user_defined_5": c.PostForm("user_defined_5"),
		"user_defined_1": c.PostForm("user_defined_1"),
		"user_defined_2": c.PostForm("user_defined_2"),
		"customer_email": c.PostForm("customer_email"),
	})
}

func sendPostRequest(redirectURL string, data []byte) {
	url := redirectURL
	fmt.Println("URL:>", url)

	fmt.Println(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

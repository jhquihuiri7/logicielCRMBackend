package request

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RequestStandard struct {
	ClientName string `json:"clientName"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Mail       string `json:"mail"`
	Message    string `json:"message"`
}
type RequestBulk struct {
	ClientName string            `json:"clientName"`
	Template   string            `json:"template"`
	Tos        []RequestStandard `json:"tos"`
	Limits     []int             `json:"limits"`
}
type RequestResponse struct {
	Success string `json:"success"`
	Error   string `json:"error"`
}

func (r *RequestStandard) ParseRequestStandardData(c *gin.Context) {
	err := json.NewDecoder(c.Request.Body).Decode(&r)
	if err != nil {
		log.Fatal(err)
	}
}
func (r *RequestBulk) ParseRequestBulkData(c *gin.Context) {
	err := json.NewDecoder(c.Request.Body).Decode(&r)
	if err != nil {
		log.Fatal(err)
	}
}

func (resp *RequestResponse) Marshal() string {
	JSONresponse, _ := json.Marshal(resp)
	return string(JSONresponse)
}
func (r *RequestResponse) ParseResponse(response *http.Response) {
	err := json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		log.Fatal(err)
	}
}

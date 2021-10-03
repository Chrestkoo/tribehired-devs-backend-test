package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
)

/**

This is a file to place helper functions unrelated to database models
Please do not import models in here so it can be imported to models file

*/

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func GetCurrentTime() string {
	//init the loc
	loc, _ := time.LoadLocation("Asia/Kuala_Lumpur")

	//set timezone
	now := time.Now().In(loc).Format("2006-01-02 15:04:05.000000")

	return now
}

// ValueToInt convert value to Int
func ValueToInt(value string) (int, error) {
	data, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return data, nil
}

// ValueToDuration convert value to duration
func ValueToDuration(value string) (time.Duration, error) {
	data, err := ValueToInt(value)
	if err != nil {
		return 0, err
	}
	return time.Duration(data), nil
}

// StringInSlice verify if string in slice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Paginate function
func Paginate(pageNum int, pageSize int, sliceLength int) (int, int) {
	start := pageNum * pageSize

	if start > sliceLength {
		start = sliceLength
	}

	end := start + pageSize
	if end > sliceLength {
		end = sliceLength
	}

	return start, end
}

// IntInSlice verify if int in slice
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// func NumberFormatPhp. this is work like php number_format
func NumberFormatPhp(number float64, decimals uint, decPoint, thousandsSep string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimals)
	// Will round off
	decimalFormat := "%." + strconv.Itoa(dec) + "f"
	// fmt.Println("decimalFormat:", decimalFormat)
	str := fmt.Sprintf(decimalFormat, number)
	// str := strconv.FormatFloat(number, 'f', dec, 64)
	// fmt.Println("str:", str)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousandsSep)
	n, l1, l2 := 0, len(prefix), len(sep)
	// thousands sep num
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	s := string(tmp)
	if dec > 0 {
		s += decPoint + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

// APIResponse struct
type APIResponse struct {
	Status     string
	StatusCode int
	Header     http.Header
	Body       string
}

// RequestAPI func
func RequestAPI(method, url string, header map[string]string, body map[string]interface{}) (*APIResponse, error) {

	var (
		req     *http.Request
		err     error
		reqBody []byte
	)

	switch method {
	case "GET":
		if body != nil {
			reqBody, err = json.Marshal(body)
			if err != nil {
				// models.ErrorLog("RequestAPI_GET_failed_in_json_Marshal_body", err.Error(), body)
				return nil, err
			}
		} else {
			reqBody = nil
		}

		req, err = http.NewRequest(method, url, bytes.NewBuffer(reqBody))

		// data add
		q := req.URL.Query()
		if body != nil {
			for k, d := range body {
				bd, ok := d.(string)
				if !ok {
					// models.ErrorLog("RequestAPI_GET_failed_in_add_query_data_string_assertion", err.Error(), d)
					return nil, errors.New("invalid param " + k)
				}
				q.Add(k, bd)
			}
		}

		// query encode
		req.URL.RawQuery = q.Encode()

	case "POST":
		if body != nil {
			reqBody, err = json.Marshal(body)

			if err != nil {
				// models.ErrorLog("RequestAPI_POST_failed_in_json_Marshal_body", err.Error(), body)
				return nil, err
			}
		} else {
			reqBody = nil
		}

		req, err = http.NewRequest(method, url, bytes.NewBuffer(reqBody))

	default:
		return nil, errors.New("invalid method for api " + url)
	}

	for k, h := range header {
		req.Header.Set(k, h)
	}

	if err != nil {
		// models.ErrorLog("RequestAPI_failed_in_NewRequest", err.Error(), nil)
		return nil, err
	}

	// client := &http.Client{}
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		// models.ErrorLog("RequestAPI_failed_in_client_Do", err.Error(), req)
		return nil, err
	}

	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// models.ErrorLog("RequestAPI_failed_in_ioutil_ReadAll", err.Error(), resp.Body)
		return nil, err
	}

	response := &APIResponse{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(resBody),
	}

	return response, nil
}

type CustomError struct {
	HTTPCode     int
	Code         int
	Msg          string
	Data         interface{}
	TemplateData map[string]interface{}
}

// Error func
func (e *CustomError) Error() string {
	if e.Msg == "" {
		return "please_try_again_later"
	}
	return e.Msg
}

type MsgStruct struct {
	Msg      string
	LangCode string
	Params   map[string]string
}

type Response struct {
	Rst  int         `json:"rst"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

type Gin struct {
	C *gin.Context
}

func (g Gin) ReturnResponse(errCode int, httpCode int, message string, data interface{}) {

	g.C.JSON(httpCode, Response{
		Rst:  errCode,
		Msg:  message,
		Data: data,
	})
	return
}

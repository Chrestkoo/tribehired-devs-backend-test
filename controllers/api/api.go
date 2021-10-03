package api

import (
	"regexp"
	"strconv"

	"github.com/chrestkoo/project/tribehired-devs-backend-test/helpers"
	"github.com/chrestkoo/project/tribehired-devs-backend-test/service/post_service"
	"github.com/gin-gonic/gin"
)

func GetTopPostListing(c *gin.Context) {

	var (
		helper = helpers.Gin{C: c}
	)

	arrDataReturn, err := post_service.ProcessGetPostListing()

	if err != nil {
		helper.ReturnResponse(0, 200, err.Error(), nil)
		return
	}
	helper.ReturnResponse(1, 200, "success", arrDataReturn)
}

func SearchPostListing(c *gin.Context) {

	var (
		helper = helpers.Gin{C: c}
	)

	var arrData post_service.PostListingSearchingField

	postIDString := c.Query("post_id")
	if postIDString != "" {
		postIDInt, err := strconv.Atoi(postIDString)
		if err != nil {
			helper.ReturnResponse(1, 200, "success", nil)
			return
		}
		if postIDInt > 0 {
			arrData.PostID = postIDInt
		}
	}

	totNumCmtString := c.Query("total_number_of_comments")
	if totNumCmtString != "" {
		totNumCmt, err := strconv.Atoi(totNumCmtString)
		if err != nil {
			helper.ReturnResponse(1, 200, "success", nil)
			return
		}
		if totNumCmt > 0 {
			arrData.TotNumOfComments = totNumCmt
		}
	}

	strRegexp := "^[a-zA-Z0-9_]*$" // only allow alpha numeric
	postTitleString := c.Query("post_title")
	if postTitleString != "" {
		re := regexp.MustCompile(strRegexp)
		if !re.MatchString(postTitleString) {
			helper.ReturnResponse(1, 200, "success", nil)
			return
		}
		arrData.PostTitle = postTitleString
	}

	postBodyString := c.Query("post_body")
	if postBodyString != "" {
		re := regexp.MustCompile(strRegexp)
		if !re.MatchString(postBodyString) {
			helper.ReturnResponse(1, 200, "success", nil)
			return
		}
		arrData.PostBody = postBodyString
	}

	arrDataReturn, err := post_service.ProcessSearchPostListing(arrData)

	if err != nil {
		helper.ReturnResponse(0, 200, err.Error(), nil)
		return
	}

	helper.ReturnResponse(1, 200, "success", arrDataReturn)
}

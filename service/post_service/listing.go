package post_service

import (
	"encoding/json"
	"errors"
	"sort"
	"strings"

	"github.com/chrestkoo/project/tribehired-devs-backend-test/helpers"
	"github.com/chrestkoo/project/tribehired-devs-backend-test/models"
)

type ProcessGetPostListingRst struct {
	PostId           int    `json:"post_id"`
	PostTitle        string `json:"post_title"`
	PostBody         string `json:"post_body"`
	TotNumOfComments int    `json:"total_number_of_comments"`
}

func ProcessGetPostListing() ([]ProcessGetPostListingRst, error) {
	arrPost, err := GetPostListingViaApi()
	if err != nil {
		return nil, err
	}
	// start debug
	// arrPost := make([]*models.PostApi, 0)
	// arrPost = append(arrPost,
	// 	&models.PostApi{ID: 1},
	// 	&models.PostApi{ID: 2},
	// 	&models.PostApi{ID: 3},
	// 	&models.PostApi{ID: 4},
	// 	&models.PostApi{ID: 5},
	// 	&models.PostApi{ID: 6},
	// 	&models.PostApi{ID: 7},
	// )
	// end debug

	arrComment, err := GetCommentListingViaApi()
	// start debug
	// arrComment := make([]*models.CommentApi, 0)

	// arrComment = append(arrComment,
	// 	&models.CommentApi{PostId: 1, ID: 1},
	// 	&models.CommentApi{PostId: 2, ID: 2},
	// 	&models.CommentApi{PostId: 2, ID: 3},
	// 	&models.CommentApi{PostId: 2, ID: 4},
	// 	&models.CommentApi{PostId: 3, ID: 5},
	// 	&models.CommentApi{PostId: 7, ID: 6},
	// 	&models.CommentApi{PostId: 101, ID: 7},
	// )
	// end debug
	if err != nil {
		return nil, err
	}
	arrDataReturn := make([]ProcessGetPostListingRst, len(arrPost))
	arrTotComment := make(map[int]int)
	if len(arrComment) > 0 {
		for _, arrCommentV := range arrComment {
			if arrTotComment[arrCommentV.PostId] > 0 {
				arrTotComment[arrCommentV.PostId] = arrTotComment[arrCommentV.PostId] + 1
			} else {
				arrTotComment[arrCommentV.PostId] = 1
			}
		}
	}

	for arrPostK, arrPostV := range arrPost {
		arrDataReturn[arrPostK].PostId = arrPostV.ID
		arrDataReturn[arrPostK].PostTitle = arrPostV.Title
		arrDataReturn[arrPostK].PostBody = arrPostV.Body
		if arrTotComment[arrPostV.ID] > 0 {
			arrDataReturn[arrPostK].TotNumOfComments = arrTotComment[arrPostV.ID]
		}
	}

	sort.Slice(arrDataReturn, func(i, j int) bool {
		return arrDataReturn[i].TotNumOfComments > arrDataReturn[j].TotNumOfComments
	})

	return arrDataReturn, nil
}

func GetPostListingViaApi() ([]models.PostApi, error) {
	var (
		method        string = "GET"
		url           string = "https://jsonplaceholder.typicode.com/posts"
		header        map[string]string
		body          map[string]interface{}
		arrDataReturn []models.PostApi
	)

	apiRst, err := helpers.RequestAPI(method, url, header, body)

	if err != nil {
		return nil, err
	}

	if apiRst.StatusCode != 200 {
		return nil, errors.New("api error result")
	}

	err = json.Unmarshal([]byte(apiRst.Body), &arrDataReturn)
	if err != nil {
		// models.ErrorLog("RequestAPI_failed_in_json_Unmarshal_resBody", err.Error(), resBody)
		return nil, err
	}

	return arrDataReturn, nil
}

func GetCommentListingViaApi() ([]models.CommentApi, error) {
	var (
		method        string = "GET"
		url           string = "https://jsonplaceholder.typicode.com/comments"
		header        map[string]string
		body          map[string]interface{}
		arrDataReturn []models.CommentApi
	)

	apiRst, err := helpers.RequestAPI(method, url, header, body)

	if err != nil {
		return nil, err
	}

	if apiRst.StatusCode != 200 {
		return nil, errors.New("api error result")
	}

	err = json.Unmarshal([]byte(apiRst.Body), &arrDataReturn)
	if err != nil {
		// models.ErrorLog("RequestAPI_failed_in_json_Unmarshal_resBody", err.Error(), resBody)
		return nil, err
	}

	return arrDataReturn, nil
}

type PostListingSearchingField struct {
	PostID           int
	PostTitle        string
	PostBody         string
	TotNumOfComments int
}

func ProcessSearchPostListing(arrData PostListingSearchingField) ([]ProcessGetPostListingRst, error) {
	arrPost, err := GetPostListingViaApi()
	if err != nil {
		return nil, err
	}
	// start debug
	// arrPost := make([]*models.PostApi, 0)
	// arrPost = append(arrPost,
	// 	&models.PostApi{ID: 1},
	// 	&models.PostApi{ID: 2},
	// 	&models.PostApi{ID: 3},
	// 	&models.PostApi{ID: 4},
	// 	&models.PostApi{ID: 5},
	// 	&models.PostApi{ID: 6},
	// 	&models.PostApi{ID: 7},
	// )
	// end debug

	arrDataReturn := make([]ProcessGetPostListingRst, 0)

	postIDStatus := true
	postTitleStatus := true
	postBodyStatus := true

	// i := 0
	if len(arrPost) > 0 {
		for _, arrPostV := range arrPost {
			if arrData.PostID > 0 {
				postIDStatus = false
				if arrData.PostID == arrPostV.ID {
					postIDStatus = true
				} else {
					continue
				}
			}

			if arrData.PostTitle != "" {
				postTitleStatus = false
				if strings.Contains(arrPostV.Title, arrData.PostTitle) {
					postTitleStatus = true
				} else {
					continue
				}
			}

			if arrData.PostBody != "" {
				postBodyStatus = false
				if strings.Contains(arrPostV.Title, arrData.PostBody) {
					postBodyStatus = true
				} else {
					continue
				}
			}

			if postIDStatus && postTitleStatus && postBodyStatus {
				arrDataReturn = append(arrDataReturn,
					ProcessGetPostListingRst{
						PostId:    arrPostV.ID,
						PostTitle: arrPostV.Title,
						PostBody:  arrPostV.Body,
					},
				)
			}
		}
	}

	arrComment, err := GetCommentListingViaApi()
	// start debug
	// arrComment := make([]*models.CommentApi, 0)

	// arrComment = append(arrComment,
	// 	&models.CommentApi{PostId: 1, ID: 1},
	// 	&models.CommentApi{PostId: 2, ID: 2},
	// 	&models.CommentApi{PostId: 2, ID: 3},
	// 	&models.CommentApi{PostId: 2, ID: 4},
	// 	&models.CommentApi{PostId: 3, ID: 5},
	// 	&models.CommentApi{PostId: 7, ID: 6},
	// 	&models.CommentApi{PostId: 101, ID: 7},
	// )
	// end debug
	if err != nil {
		return nil, err
	}

	arrTotComment := make(map[int]int)
	if len(arrComment) > 0 {
		for _, arrCommentV := range arrComment {
			if arrTotComment[arrCommentV.PostId] > 0 {
				arrTotComment[arrCommentV.PostId] = arrTotComment[arrCommentV.PostId] + 1
			} else {
				arrTotComment[arrCommentV.PostId] = 1
			}
		}
	}

	for k, v := range arrTotComment {
		for arrDataReturnK, arrDataReturnV := range arrDataReturn {
			if k == arrDataReturnV.PostId {
				arrDataReturn[arrDataReturnK].TotNumOfComments = v
			}
		}
	}

	if arrData.TotNumOfComments <= 0 {
		return arrDataReturn, nil
	}

	arrNewDataReturn := make([]ProcessGetPostListingRst, 0)
	// fmt.Println("arrDataReturn:", arrDataReturn)
	if len(arrDataReturn) > 0 {
		for _, arrDataReturnV := range arrDataReturn {
			if arrData.TotNumOfComments == arrDataReturnV.TotNumOfComments {
				arrNewDataReturn = append(arrNewDataReturn,
					ProcessGetPostListingRst{
						PostId:           arrDataReturnV.PostId,
						PostTitle:        arrDataReturnV.PostTitle,
						PostBody:         arrDataReturnV.PostBody,
						TotNumOfComments: arrDataReturnV.TotNumOfComments,
					},
				)
			}
		}
	}
	return arrNewDataReturn, nil
}

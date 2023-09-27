package controller

import (
	"GoAPI/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"net/http"
	"strconv"
	"time"
)

// api tim 1 item bang id
func GetPersonById(ctx *gin.Context) {
	id := ctx.Param("id")
	result := model.ModelGet(id)
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func ListPerson(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	search := ctx.DefaultQuery("search", "")
	results := model.ModelList(search)
	if err := ctx.ShouldBind(&results); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	total := len(results)
	// Số dòng trên mỗi trang
	rowsPerPage := 6
	// Tính vị trí bắt đầu và kết thúc của dữ liệu trên trang hiện tại
	startIndex := (page - 1) * rowsPerPage
	endIndex := startIndex + rowsPerPage
	if endIndex > total {
		endIndex = total
	}
	//thêm total
	// ???

	// Lấy dữ liệu trên trang hiện tại
	currentPageData := results[startIndex:endIndex]

	// Tính tổng số trang
	totalPages := int(math.Ceil(float64(len(results)) / float64(rowsPerPage)))

	// Tạo danh sách các trang
	var pages []int
	for i := 1; i <= totalPages; i++ {
		pages = append(pages, i)
	}

	var nextPage int
	var isLastPage bool

	if page < totalPages {
		nextPage = page + 1
	} else {
		isLastPage = true
	}
	// Render template
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"data":        currentPageData,
		"prevPage":    page - 1,
		"currentPage": page,
		"total":       total,
		"nextPage":    nextPage,
		"startIndex":  startIndex + 1,
		"endIndex":    endIndex,
		"isLastPage":  isLastPage,
		"pages":       pages,
		"search":      search,
	})
}

func AddPerson(ctx *gin.Context) {
	var person model.Person
	if err := ctx.ShouldBind(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	person.ID = primitive.NewObjectID()
	person.Appearance = false
	now := time.Now()
	person.CreatedAt = &now
	person.UpdatedAt = &now
	message := model.ModelCreate(person)
	fmt.Println(message)
	ctx.Redirect(http.StatusFound, "/person/info")
}

func DeletePersonById(ctx *gin.Context) {
	id := ctx.Query("id")
	page := ctx.PostForm("page")
	search := ctx.PostForm("search")
	message := model.ModelDelete(id)
	fmt.Println(message)
	redirectURL := fmt.Sprintf("/person/info?page=%s&search=%s", page, search)
	ctx.Redirect(http.StatusFound, redirectURL)
}

func UpdatePersonById(ctx *gin.Context) {
	id := ctx.Query("id")
	var person model.Person
	if err := ctx.ShouldBind(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	page := ctx.PostForm("page")
	search := ctx.PostForm("search")
	person.ID, _ = primitive.ObjectIDFromHex(id)
	now := time.Now()
	person.UpdatedAt = &now
	message := model.ModelUpdate(person)
	fmt.Println(message)
	fmt.Printf("page is %s and search is %s", page, search)
	redirectURL := fmt.Sprintf("/person/info?page=%s&search=%s", page, search)
	ctx.Redirect(http.StatusFound, redirectURL)
}

func ToggleAppearance(ctx *gin.Context) {
	id := ctx.Query("id")
	var person model.Person
	if err := ctx.ShouldBind(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	page := ctx.PostForm("page")
	search := ctx.PostForm("search")
	person.ID, _ = primitive.ObjectIDFromHex(id)
	//update true sang false
	person.Appearance = !person.Appearance
	person.Name = ctx.PostForm("name")
	person.Major = ctx.PostForm("major")
	message := model.ModelUpdate(person)
	fmt.Println(message)
	fmt.Printf("page is %s and search is %s", page, search)
	redirectURL := fmt.Sprintf("/person/info?page=%s&search=%s", page, search)
	ctx.Redirect(http.StatusFound, redirectURL)
}

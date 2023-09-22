package controller

import (
	"GoAPI/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
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
	search := ctx.DefaultQuery("search", "")
	results := model.ModelList(search)
	if err := ctx.ShouldBind(&results); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	//render template
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"data": results,
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

	message := model.ModelDelete(id)
	fmt.Println(message)
	ctx.Redirect(http.StatusFound, "/person/info")
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
	person.ID, _ = primitive.ObjectIDFromHex(id)
	now := time.Now()
	person.UpdatedAt = &now
	message := model.ModelUpdate(person)
	fmt.Println(message)
	ctx.Redirect(http.StatusFound, "/person/info")
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
	person.ID, _ = primitive.ObjectIDFromHex(id)

	//update true sang false
	person.Appearance = !person.Appearance
	person.Name = ctx.PostForm("name")
	person.Major = ctx.PostForm("major")
	message := model.ModelUpdate(person)
	fmt.Println(message)

	ctx.Redirect(http.StatusFound, "/person/info")
}

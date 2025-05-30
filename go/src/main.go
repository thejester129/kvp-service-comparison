package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonItem map[string]any

func main() {
	loadTable()

	router := gin.Default()
	router.GET("/", getItems)
	router.GET("/:key", getByKey)
	router.PUT("/:key", putItem)
	router.PATCH("/:key", patchItem)
	router.DELETE("/:key", deleteItem)

	router.Run("localhost:8080")
}

func getItems(c *gin.Context) {
	items, err := getAllTableItems(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, items)
}

func getByKey(c *gin.Context) {
	key := c.Param("key")
	item, err := getTableItem(c, key)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	respondWithItem(c, http.StatusOK, item)
}

func putItem(c *gin.Context) {
	key := c.Param("key")

	var newItem JsonItem

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	newItem["key"] = key
	err := putTableItem(c, newItem)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	respondWithItem(c, http.StatusOK, newItem)
}

func patchItem(c *gin.Context) {
	key := c.Param("key")

	var patchItem JsonItem

	if err := c.BindJSON(&patchItem); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	patchItem["key"] = key
	updated, err := updateTableItem(c, patchItem)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	respondWithItem(c, http.StatusCreated, updated)
}

func deleteItem(c *gin.Context) {
	key := c.Param("key")

	err := deleteTableItem(c, key)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	respondWithItem(c, http.StatusOK, JsonItem{})
}

func respondWithItem(c *gin.Context, status int, item JsonItem) {
	delete(item, "key")
	c.IndentedJSON(status, item)
}

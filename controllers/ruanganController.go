package controllers

import (
	"example/inventaris-api-v1/models"
	"example/inventaris-api-v1/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RuangController struct {
	RuangService *services.RuangService
}

func NewRuangController(ruangService *services.RuangService) *RuangController {
	return &RuangController{RuangService: ruangService}
}

func (rc *RuangController) Create(c *gin.Context) {
	var data models.Ruangan
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := rc.RuangService.CreateRuang(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK,"success": true, "message":"success", "data": data})
}

func (rc *RuangController) GetAll(c *gin.Context) {
	data, err := rc.RuangService.GetAllRuang()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": http.StatusInternalServerError})
		return
	}

	if data == nil {
		c.JSON(http.StatusNoContent, gin.H{"message":"no data found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK,"success": true, "message":"success", "data": data})
}

func (rc *RuangController) GetOne(c *gin.Context){
	id := c.Param("id")

	data, err := rc.RuangService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": err.Error(), "message": "no data found by id: " + id})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "success": true, "message": "success", "data": data})
}

func (rc *RuangController) Update(c *gin.Context){
	id := c.Param("id")

	ruang, err := rc.RuangService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": err.Error(), "message": "no ruang found by id: " + id}) 
	}

	var updatedAttributes map[string]interface{}
	if err := c.ShouldBindJSON(&updatedAttributes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := rc.RuangService.UpdateRuang(ruang, updatedAttributes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "success": true, "message":"success", "data": ruang})
}

func (rc *RuangController) Delete(c *gin.Context) {
	id := c.Param("id")

	_, err := rc.RuangService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": err.Error(), "message": "no data found by id: " + id})
		return 
	}

	if err := rc.RuangService.DeleteRuang(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "success": true, "message": "deleted successfully"})
}
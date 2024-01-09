package controllers

import (
	"example/inventaris-api-v1/models"
	"example/inventaris-api-v1/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventarisController struct {
	InventService *services.InventarisService
}

func NewInventController(inventService *services.InventarisService) *InventarisController {
	return &InventarisController{InventService: inventService}
}

func (ic *InventarisController) Create(c *gin.Context) {
	var data models.Inventaris
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ic.InventService.CreateInventaris(&data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK,"success": true, "message":"success", "data": data})
}

func (ic *InventarisController) GetAll(c *gin.Context) {

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	kondisiBarang := c.Query("kondisi_barang")
	pengembalianLaptop := c.Query("pengembalian_laptop")

	invent, err := ic.InventService.GetAllInventaris(startDate, endDate, kondisiBarang, pengembalianLaptop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": http.StatusInternalServerError})
		return
	}

	if invent == nil {
		c.JSON(http.StatusNoContent, gin.H{"message":"no data found"})
		return
	}

	var result []map[string]interface{}
	for _, inv := range invent{
		data := map[string]interface{}{
			"id": inv.ID,
			"nama_barang": inv.NamaBarang,
			"tipe": inv.Tipe,
			"tanggal_peroleh": inv.TanggalPeroleh,
			"kondisi_barang": inv.KondisiBarang,
			"sumber_dana": inv.SumberDana,
			"kode_barang": inv.KodeBarang,
			"harga_satuan": inv.HargaSatuan,
			"kode_inventaris": inv.KodeInventaris,
			"ruangan": inv.Ruangan.NamaRuangan,
			"nama_pengguna": inv.NamaPengguna,
			"unit": inv.Unit,
			"keterangan": inv.Keterangan,
			"digunakan": inv.Digunakan,
			"tidak_digunakan": inv.TidakDigunakan,
			"pengembalian_laptop_lama": inv.PengembalianLaptopLama,
		}
		result = append(result, data)
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK,"success": true, "message":"success", "data": result})
}

func (ic *InventarisController) GetOne(c *gin.Context){
	id := c.Param("id")

	invent, err := ic.InventService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": err.Error(), "message": "no data found by id: " + id}) 
		return
	}
	data := map[string]interface{}{
		"id": invent.ID,
		"nama_barang": invent.NamaBarang,
		"tipe": invent.Tipe,
		"tanggal_peroleh": invent.TanggalPeroleh,
		"kondisi_barang": invent.KondisiBarang,
		"sumber_dana": invent.SumberDana,
		"kode_barang": invent.KodeBarang,
		"harga_satuan": invent.HargaSatuan,
		"kode_inventaris": invent.KodeInventaris,
		"ruangan": invent.Ruangan.NamaRuangan,
		"nama_pengguna": invent.NamaPengguna,
		"unit": invent.Unit,
		"keterangan": invent.Keterangan,
		"digunakan": invent.Digunakan,
		"tidak_digunakan": invent.TidakDigunakan,
		"pengembalian_laptop_lama": invent.PengembalianLaptopLama,
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "success": true, "message": "success", "data": data})
}

func (ic *InventarisController) Search(c *gin.Context) {

	keywords := c.Query("keywords")

	invent, err := ic.InventService.SearchData(keywords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": http.StatusInternalServerError})
		return
	}

	if invent == nil {
		c.JSON(http.StatusNoContent, gin.H{"message":"no data found"})
		return
	}

	var result []map[string]interface{}
	for _, inv := range invent{
		data := map[string]interface{}{
			"id": inv.ID,
			"nama_barang": inv.NamaBarang,
			"tipe": inv.Tipe,
			"tanggal_peroleh": inv.TanggalPeroleh,
			"kondisi_barang": inv.KondisiBarang,
			"sumber_dana": inv.SumberDana,
			"kode_barang": inv.KodeBarang,
			"harga_satuan": inv.HargaSatuan,
			"kode_inventaris": inv.KodeInventaris,
			"ruangan": inv.Ruangan.NamaRuangan,
			"nama_pengguna": inv.NamaPengguna,
			"unit": inv.Unit,
			"keterangan": inv.Keterangan,
			"digunakan": inv.Digunakan,
			"tidak_digunakan": inv.TidakDigunakan,
			"pengembalian_laptop_lama": inv.PengembalianLaptopLama,
		}
		result = append(result, data)
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK,"success": true, "message":"success", "data": result})
}


func (ic *InventarisController) Update(c *gin.Context){
	id := c.Param("id")

	data, err := ic.InventService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": err.Error(), "message": "no data found by id: " + id})
		return 
	}

	var updatedAttributes map[string]interface{}
	if err := c.ShouldBindJSON(&updatedAttributes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ic.InventService.UpdateInventaris(data, updatedAttributes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "success": true, "message":"success", "data": data})
}

func (ic *InventarisController) Delete(c *gin.Context) {
	id := c.Param("id")

	_, err := ic.InventService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": err.Error(), "message": "no data found by id: " + id}) 
		return
	}

	if err := ic.InventService.DeleteInventaris(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "success": true, "message": "deleted successfully"})
}
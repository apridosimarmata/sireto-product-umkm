package Controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"umkm_service/Models"
	"umkm_service/Utils"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Result string `json:"result"`
}

func CreateUMKMApi(c *gin.Context) {
	userId := c.Params.ByName("user_id")
	var umkmRequest Models.UMKMRequest

	Utils.LoadEnv()
	sess := Utils.ConnectAws()
	c.BindJSON(&umkmRequest)

	photosUrl := make([]string, len(umkmRequest.Photos))

	for index, item := range umkmRequest.Photos {
		photosUrl[index] = Utils.RandStringRunes(20)
		go Utils.AddFileToS3(sess, item, photosUrl[index])
	}

	facilityIds := strings.Join(umkmRequest.FacilityIds, ",")

	subcategoryIds := strings.Join(umkmRequest.SubCategoryIds, ",")

	photosString := strings.Join(photosUrl, ",")

	var err error

	var umkm Models.UMKM
	umkm.Name = umkmRequest.Name
	umkm.Address = umkmRequest.Address
	umkm.Phone = umkmRequest.Phone
	umkm.CategoryId = 1
	umkm.ProvinceId = umkmRequest.ProvinceId
	umkm.LocationId = umkmRequest.LocationId
	umkm.Photos = photosString
	umkm.SubcategoryIds = subcategoryIds
	umkm.FacilityIds = facilityIds
	umkm.OwnerId, err = strconv.Atoi(userId)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(umkmRequest)
	err = Models.CreateUMKM(&umkm)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(`{"result" : "success"}`))
}

func CloseBusiness(c *gin.Context) {
	var umkm Models.UMKM

	err := Models.GetUMKMById(&umkm, c.Params.ByName("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	umkm.Active = false

	err = Models.UpdateUMKM(&umkm)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, umkm)

}

func OpenBusiness(c *gin.Context) {
	var umkm Models.UMKM

	err := Models.GetUMKMById(&umkm, c.Params.ByName("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	umkm.Active = true

	err = Models.UpdateUMKM(&umkm)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, umkm)

}

func CreateUMKM(c *gin.Context) {
	var umkmRequest Models.UMKMRequest
	Utils.LoadEnv()
	sess := Utils.ConnectAws()
	c.BindJSON(&umkmRequest)

	token, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//body := map[string]string{"token": token, "refresh_token": refreshToken}

	//postBody, _ := json.Marshal(body)

	//resp, err := http.Post("https://account.sireto.id/api/v1/user/profile", "application/json", bytes.NewBuffer(postBody), http.Cookie{})
	req, err := http.NewRequest("GET", "https://account.sireto.id/api/v1/user/profile", nil)

	if err != nil {
		var response Response
		response.Result = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	req.AddCookie(&http.Cookie{Name: "token", Value: token})
	req.AddCookie(&http.Cookie{Name: "refresh_token", Value: refreshToken})

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		c.AbortWithStatus(http.StatusBadGateway)
	}

	var user Models.User

	responseBody, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(responseBody, &user)

	fmt.Println(user)
	photosUrl := make([]string, len(umkmRequest.Photos))

	for index, item := range umkmRequest.Photos {
		photosUrl[index] = Utils.RandStringRunes(20)
		go Utils.AddFileToS3(sess, item, photosUrl[index])
	}

	facilityIds := strings.Join(umkmRequest.FacilityIds, ",")

	subcategoryIds := strings.Join(umkmRequest.SubCategoryIds, ",")

	photosString := strings.Join(photosUrl, ",")

	var umkm Models.UMKM
	umkm.Name = umkmRequest.Name
	umkm.Address = umkmRequest.Address
	umkm.Phone = umkmRequest.Phone
	umkm.CategoryId = umkmRequest.CategoryId
	umkm.ProvinceId = umkmRequest.ProvinceId
	umkm.LocationId = umkmRequest.LocationId
	umkm.Photos = photosString
	umkm.SubcategoryIds = subcategoryIds
	umkm.FacilityIds = facilityIds
	umkm.OwnerId = user.Id

	err = Models.CreateUMKM(&umkm)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(`{"result" : "success"}`))
}

func CreateUMKMCategory(c *gin.Context) {
	var umkmCategory Models.UMKMCategory
	c.BindJSON(&umkmCategory)

	err := Models.CreateUMKMCategory(&umkmCategory)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(`{"result" : "success"}`))
}

func GetAllUMKMCategories(c *gin.Context) {
	var umkmCategories []Models.UMKMCategory

	err := Models.GetAllUMKMCategories(&umkmCategories)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, umkmCategories)
}

func CreateUMKMSubCategory(c *gin.Context) {
	var umkmSubCategory Models.UMKMSubCategory
	c.BindJSON(&umkmSubCategory)

	err := Models.CreateUMKMSubCategory(&umkmSubCategory)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(`{"result" : "success"}`))
}

func GetUMKMSubCategoriesByCategoryId(c *gin.Context) {
	umkmCategoryId := c.Params.ByName("category_id")
	var umkmSubCategories []Models.UMKMSubCategory

	err := Models.GetUMKMSubCategoriesByCategoryId(&umkmSubCategories, umkmCategoryId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, umkmSubCategories)
}

func GetUMKMByLocationId(c *gin.Context) {
	locationId := c.Params.ByName("location_id")
	var umkm []Models.UMKM

	err := Models.GetUMKMByLocationId(&umkm, locationId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, umkm)
}

func GetUMKMById(c *gin.Context) {
	umkmId := c.Params.ByName("id")
	var umkm Models.UMKM

	err := Models.GetUMKMById(&umkm, umkmId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, umkm)
}

func AddUMKMScore(c *gin.Context) {
	var newScore Models.UMKMNewScore
	c.BindJSON(&newScore)
	fmt.Println(newScore)
	var umkm Models.UMKM
	err := Models.GetUMKMById(&umkm, newScore.BusinessId)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	umkm.ReviewsNumber = umkm.ReviewsNumber + 1
	umkm.Score = umkm.Score + newScore.Score

	Models.UpdateUMKM(&umkm)

	c.JSON(http.StatusOK, umkm)
}

func GetUMKMSByOwnerId(c *gin.Context) {
	ownerId := c.Params.ByName("owner_id")
	var umkm []Models.UMKM

	err := Models.GetUMKMSByOwnerId(&umkm, ownerId)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, umkm)
}

func GetFavoriteUMKM(c *gin.Context) {
	var umkm []Models.UMKM

	err := Models.GetFavoriteUMKM(&umkm)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, umkm)
}

func SearchUMKMByName(c *gin.Context) {
	var umkms []Models.UMKM

	err := Models.GetUMKMSByName(&umkms, c.Params.ByName("name"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, umkms)
}

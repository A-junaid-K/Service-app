package controllers

import (
	"fmt"
	"service-api/database"
	"service-api/helpers"
	"service-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ServicerSignup(c *gin.Context) {
	var input struct {
		UserName    string `json:"username"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone"`
		Password    string `json:"password"`
	}

	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}
	var servicer1 models.Servicer
	if err := database.DB.Where("email = ?", input.Email).First(&servicer1).Error; err == nil {
		c.JSON(400, gin.H{
			"error": "email already exist",
		})
		return
	}

	otp := helpers.GenerateOtp()

	otpstr := strconv.Itoa(otp)

	helpers.SendOtp(otpstr, input.Email)

	password, _ := helpers.HashPassword(input.Password)

	if err := database.DB.Create(&models.Servicer{
		UserName:    input.UserName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    string(password),
		OTP:         otpstr,
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var servicer models.Servicer

	database.DB.Last(&servicer)

	c.JSON(200, gin.H{
		"id":  servicer.ID,
		"otp": servicer.OTP,
	})
}

func AddDocuments(c *gin.Context) {
	var input struct {
		FullName             string `json:"fullname"`
		Description          string `json:"description"`
		ServiceCatagory      string `json:"servicecatagory"`
		VerificationDocument string `json:"verificationdocument"`
		Amount               int    `json:"amount"`
		Location             string `json:"location"`
		ServicerImage        string `json:"servicerimage"`
		ServicerDocument     string `json:"servicerdocument"`
	}

	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Binding error",
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DB.Model(&models.Servicer{}).Where("id=?", id).Updates(map[string]interface{}{"full_name": input.FullName,
		"description": input.Description, "service_catagory": input.ServiceCatagory,
		"verification_document": input.VerificationDocument, "amount": input.Amount, "location": input.Location,
		"servicer_image": input.ServicerImage, "servicer_document": input.ServicerDocument, "status": "Pending"}).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "updation error",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": "created new servicer",
	})

}

func GetServicerDetails(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))
	var servicer models.Servicer
	if err := database.DB.First(&servicer, id).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to find servicer",
		})
		return
	}
	c.JSON(200, gin.H{
		"servicer": servicer,
	})
}

func ServicerLogin(c *gin.Context) {
	var login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&login); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}
	var servicer models.Servicer
	if err := database.DB.Where("email = ?", login.Email).First(&servicer).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "incorrect eamail and password",
		})
		return
	}
	if err := helpers.CheckPassword(servicer.Password, login.Password); err != nil {
		c.JSON(400, gin.H{
			"error": "incorrect eamail and password",
		})
		return
	}
	c.JSON(200, gin.H{
		"id": servicer.ID,
	})

}

type BookingDetails struct {
	Id            int    `json:"id"`
	BuildingName  string `json:"buildingname"`
	City          string `json:"city"`
	Road          string `json:"road"`
	Phone         string `json:"phone"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	Description   string `json:"description"`
	ServiceAmount int    `json:"serviceamount"`
	Status        string `json:"status"`
	User_Name     string `json:"username"`
}

func GetAllBookings(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))

	var booking []BookingDetails

	if err := database.DB.Table("bookings").Select("bookings.id,bookings.building_name,bookings.city,bookings.road,bookings.phone,bookings.date,bookings.time,bookings.description,bookings.service_amount,bookings.status,users.user_name").
		Joins("INNER JOIN users ON users.id=bookings.user_id").Where("bookings.servicer_id=?", id).Scan(&booking).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get data",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": booking,
	})
}

func ChangeBookingStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("servicer_id"))
	var input struct {
		BookingID int    `json:"bookingid"`
		Status    string `json:"status"`
	}
	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get body",
		})
		return
	}

	fmt.Println(input, id)

	var booking models.Booking
	if err := database.DB.Where("id = ? AND servicer_id = ?", input.BookingID, id).Find(&booking).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "this servicer not access this booking",
		})
		return
	}
	if err := database.DB.Model(&models.Booking{}).Where("id = ?", input.BookingID).Update("status", input.Status).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "failed to update data",
		})
		return
	}
	c.JSON(200, gin.H{
		"data": "successfully updated status",
	})
}

//http://localhost:5000/user/dsflk?user_id=1&servicer_id=5

package services

import (
	"errors"
	"fmt"
	"interview/pkg/entity"

	repo "interview/internal/repository"
	"log"
	"strconv"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

var ItemPriceMapping = map[string]float64{
	"shoe":  100,
	"purse": 200,
	"bag":   300,
	"watch": 300,
}

type (
	CartServiceInterface interface {
		GetCartData(c *gin.Context)
		AddItemToCart(c *gin.Context)
		DeleteCartItem(c *gin.Context)
	}

	CartItemForm struct {
		Product  string `form:"product"   binding:"required"`
		Quantity string `form:"quantity"  binding:"required"`
	}
)

func (s *service) GetCartData(c *gin.Context) {
	data := map[string]interface{}{
		"Error": c.Query("error"),
	}

	cookie, err := c.Request.Cookie("ice_session_id")
	if err == nil {
		data["CartItems"] = GetCartItemData(s.repository, cookie.Value)
	}

	html, err := RenderTemplate(data)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(500)
		return
	}

	c.Header("Content-Type", "text/html")
	c.String(200, html)
}

func (s *service) AddItemToCart(c *gin.Context) {
	cookie, _ := c.Request.Cookie("ice_session_id")

	var isCartNew bool
	cartEntity, err := s.repository.GetCartBySessionID(cookie.Value)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.Redirect(302, "/")
			return
		}
		isCartNew = true
		cartEntity = &entity.CartEntity{
			SessionID: cookie.Value,
			Status:    entity.CartOpen,
		}
		if err := s.repository.CreateCart(cartEntity); err != nil {
			c.Redirect(302, "/")
			return
		}
	}

	addItemForm, err := GetCartItemForm(c)
	if err != nil {
		c.Redirect(302, "/?error="+err.Error())
		return
	}

	item, ok := ItemPriceMapping[addItemForm.Product]
	if !ok {
		c.Redirect(302, "/?error=invalid item name")
		return
	}

	quantity, err := strconv.ParseInt(addItemForm.Quantity, 10, 0)
	if err != nil {
		c.Redirect(302, "/?error=invalid quantity")
		return
	}

	var cartItemEntity *entity.CartItem
	if isCartNew {
		cartItemEntity = &entity.CartItem{
			CartID:      cartEntity.ID,
			ProductName: addItemForm.Product,
			Quantity:    int(quantity),
			Price:       item * float64(quantity),
		}
		if err := s.repository.CreateCartItem(cartItemEntity); err != nil {
			c.Redirect(302, "/")
			return
		}
	} else {
		cartItemEntity, err = s.repository.GetCartItemByCartIDAndProductName(cartEntity.ID, addItemForm.Product)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				c.Redirect(302, "/")
				return
			}
			cartItemEntity = &entity.CartItem{
				CartID:      cartEntity.ID,
				ProductName: addItemForm.Product,
				Quantity:    int(quantity),
				Price:       item * float64(quantity),
			}
			if err := s.repository.CreateCartItem(cartItemEntity); err != nil {
				c.Redirect(302, "/")
				return
			}
		} else {
			cartItemEntity.Quantity += int(quantity)
			cartItemEntity.Price += item * float64(quantity)
			if err := s.repository.UpdateCartItem(cartItemEntity); err != nil {
				c.Redirect(302, "/")
				return
			}
		}
	}

	c.Redirect(302, "/")
}

func (s *service) DeleteCartItem(c *gin.Context) {
	cartItemIDString := c.Query("cart_item_id")
	if cartItemIDString == "" {
		c.Redirect(302, "/")
		return
	}

	cookie, _ := c.Request.Cookie("ice_session_id")

	cartEntity, err := s.repository.GetCartBySessionID(cookie.Value)
	if err != nil {
		c.Redirect(302, "/")
		return
	}

	cartItemID, err := strconv.Atoi(cartItemIDString)
	if err != nil {
		c.Redirect(302, "/")
		return
	}

	if err := s.repository.DeleteCartItem(cartItemID, cartEntity.ID); err != nil {
		c.Redirect(302, "/")
		return
	}

	c.Redirect(302, "/")
}

func GetCartItemForm(c *gin.Context) (*CartItemForm, error) {
	if c.Request.Body == nil {
		return nil, fmt.Errorf("body cannot be nil")
	}

	form := &CartItemForm{}

	if err := binding.FormPost.Bind(c.Request, form); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return form, nil
}

func GetCartItemData(repo repo.CartRepositoryInterface, sessionID string) (items []map[string]interface{}) {
	cartEntity, err := repo.GetCartBySessionID(sessionID)
	if err != nil {
		return
	}

	cartItems, err := repo.GetCartItemsByCartID(cartEntity.ID)
	if err != nil {
		return
	}

	for _, cartItem := range cartItems {
		item := map[string]interface{}{
			"ID":       cartItem.ID,
			"Quantity": cartItem.Quantity,
			"Price":    cartItem.Price,
			"Product":  cartItem.ProductName,
		}

		items = append(items, item)
	}
	return items
}

func RenderTemplate(pageData interface{}) (string, error) {
	tmpl, err := template.ParseFiles("../../static/add_item_form.html")
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v ", err)
	}

	var renderedTemplate strings.Builder

	err = tmpl.Execute(&renderedTemplate, pageData)
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v ", err)
	}

	return renderedTemplate.String(), nil
}

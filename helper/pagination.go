package helper

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit        int         `json:"limit,omitempty;query:limit"`
	CurrentPage  int         `json:"currentPage,omitempty;query:page"`
	Sort         string      `json:"sort,omitempty;query:sort"`
	TotalElement int64       `json:"totalElement"`
	TotalPage    int         `json:"totalPage"`
	NextPage     int         `json:"nextPage"`
	PrevPage     int         `json:"prevPage"`
	Content      interface{} `json:"-"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.CurrentPage == 0 {
		p.CurrentPage = 1
	}
	return p.CurrentPage
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id_deposit desc"
	}
	return p.Sort
}

func PaginateByUserId(value interface{}, userId uint64, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Where("id_cust = ?", userId).Count(&totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPage = totalPages
	pagination.NextPage = pagination.GetPage() + 1
	pagination.PrevPage = pagination.GetPage() - 1
	pagination.Content = value
	pagination.TotalElement = totalRows

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

//GeneratePaginationFromRequest
func GeneratePaginationFromRequest(c *gin.Context) Pagination {
	// Initializing default
	//	var mode string
	limit := 2
	page := 1
	sort := "id_deposit desc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	return Pagination{
		Limit:       limit,
		CurrentPage: page,
		Sort:        sort,
	}
}

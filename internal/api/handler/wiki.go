package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"smolathon/internal/entity"
	"smolathon/pkg/shttp"
	"smolathon/pkg/slogger"
)

type getEncyclopediaItemUri struct {
	EncyclopediaItemId uuid.UUID `form:"encyclopedia_item_id"`
}

type getEncyclopediaItemResponse struct {
	Id           uuid.UUID
	SettlementId uuid.UUID
	Title        string
	Description  string
}

func newGetEncyclopediaItemResponse(item entity.EncyclopediaItem) getEncyclopediaItemResponse {
	response := getEncyclopediaItemResponse{
		Id:           item.Id,
		SettlementId: item.SettlementId,
		Title:        item.Title,
		Description:  item.Description,
	}
	return response
}

// GetEncyclopediaItem godoc
// @Summary     Получение элемента энциклопедии
// @Tags		Wiki
// @Accept      json
// @Produce     json
// @Param       EncyclopediaItemId path string true "Идентификатор элемента энциклопедии"
// @Success     200 {object} shttp.ResponseWithDetails[getEncyclopediaItemResponse]
// @Failure     400 {object} shttp.ResponseError "Bad request"
// @Failure     500 {object} shttp.ResponseError "Internal server error"
// @Router      /api/wiki/{EncyclopediaItemId} [get]
func GetEncyclopediaItem(logger slogger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getEncyclopediaItemUri
		if err := c.ShouldBindUri(&req); err != nil {
			shttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		//test data
		item := entity.EncyclopediaItem{
			Id:           uuid.MustParse("506e21ab-a0b3-4208-94c6-3fd2e996180c"),
			SettlementId: uuid.MustParse("6543e3fa-7ccf-11ee-b962-0242ac120002"),
			Title:        "Музей «Смоленск — щит России»",
			Description:  "Один из смоленских музеев, посвящённый боевой истории Смоленска и его роли в истории России. Располагается в Громовой башне Смоленской крепостной стены",
		}

		var err error
		switch {
		case err == nil:
		default:
			shttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		}
		response := newGetEncyclopediaItemResponse(item)
		shttp.OkResponseWithResult(c, response)
	}
}

type listEncyclopediaItemRequest struct {
	SettlementId uuid.UUID `json:"settlement_id"`
}
type listEncyclopediaItemResponse []listEncyclopediaItemElement

type listEncyclopediaItemElement struct {
	Id           uuid.UUID `json:"id"`
	SettlementId uuid.UUID `json:"settlement_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
}

func newListEncyclopediaItemResponse(items []entity.EncyclopediaItem) listEncyclopediaItemResponse {
	var response = make(listEncyclopediaItemResponse, 0, len(items))
	for _, item := range items {
		response = append(response, listEncyclopediaItemElement{
			Id:           item.Id,
			SettlementId: item.SettlementId,
			Title:        item.Title,
			Description:  item.Description,
		})
	}
	return response
}

// ListEncyclopediaItem godoc
// @Summary     Получение элементов энциклопедии
// @Tags		Wiki
// @Accept      json
// @Produce     json
// @Param       Body body listEncyclopediaItemRequest true "Параметры выборки"
// @Success     200 {object} shttp.ResponseWithDetails[listEncyclopediaItemResponse]
// @Failure     400 {object} shttp.ResponseError "Bad request"
// @Failure     500 {object} shttp.ResponseError "Internal server error"
// @Router      /api/wiki/list [post]
func ListWiki(logger slogger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req listEncyclopediaItemRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			shttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}
		encyclopediaItems := []entity.EncyclopediaItem{
			{
				Id:           uuid.MustParse("861e8f4c-7d99-11ee-b962-0242ac120002"),
				SettlementId: uuid.MustParse("6ed83b6e-7ccf-11ee-b962-0242ac120002"),
				Title:        "Художественная галерея Cмоленского государственного музея-заповедника",
				Description:  "В галерее большая коллекция редких и интересных картин, которые стоит увидеть и повнимательнее рассмотреть",
			},
			{
				Id:           uuid.MustParse("89ad0d8c-7d99-11ee-b962-0242ac120002"),
				SettlementId: uuid.MustParse("6ed83b6e-7ccf-11ee-b962-0242ac120002"),
				Title:        "Музей скульптуры С.Т. Коненкова",
				Description:  "Музей находится в восхитительно красивом здании 19 века, которое является объектом культурного наследия регионального значения. Экспозиция также оригинальна и интересна.",
			},
		}
		var err error
		switch {
		case err == nil:
		default:
			shttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		}

		response := newListEncyclopediaItemResponse(encyclopediaItems)
		shttp.OkResponseWithResult(c, response)
	}
}

// CreateEncyclopediaItem godoc
// @Summary     Создание элемента энциклопедии
// @Tags		Wiki
// @Accept      json
// @Produce     json
// @Success     200 {object} shttp.ResponseOk "Ok" //хз как тут сделать
// @Failure     400 {object} shttp.ResponseError "Bad request"
// @Failure     500 {object} shttp.ResponseError "Internal server error"
// @Router      /api/wiki/create [post]
func CreateEncyclopediaItem(logger slogger.Logger, item entity.EncyclopediaItem) gin.HandlerFunc {
	return func(c *gin.Context) {

		encyclopediaItem := entity.EncyclopediaItem{

			Id:           uuid.MustParse("89ad0d8c-7d99-11ee-b962-0242ac120002"),
			SettlementId: uuid.MustParse("6ed83b6e-7ccf-11ee-b962-0242ac120002"),
			Title:        "Музей скульптуры С.Т. Коненкова",
			Description:  "Музей находится в восхитительно красивом здании 19 века, которое является объектом культурного наследия регионального значения. Экспозиция также оригинальна и интересна.",
		}

		var err error
		switch {
		case err == nil:
		default:
			shttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		}
		shttp.OkResponseWithResult(c, encyclopediaItem)
	}
}

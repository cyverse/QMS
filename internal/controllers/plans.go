package controllers

import (
	"database/sql"
	"net/http"

	"github.com/cyverse/QMS/internal/model"
	"github.com/labstack/echo"
)

// swagger:route GET /plans plans listPlans
// Returns a List all the plans
// responses:
//   200: plansResponse
//   404: RootResponse

func (s Server) GetAllPlans(ctx echo.Context) error {
	data := []model.Plan{}
	err := s.GORMDB.Debug().Find(&data).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse(err.Error(), http.StatusInternalServerError))
	}
	return ctx.JSON(http.StatusOK, model.SuccessResponse(data, http.StatusOK))
}

// swagger:route GET /plans/{PlanID} plans listPlansByID
// Returns a List all the plans
// responses:
//   200: plansResponse
//   500: RootResponse

func (s Server) GetPlansForID(ctx echo.Context) error {
	plan_id := ctx.Param("plan_id")
	if plan_id == "" {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse("Invalid PlanID", http.StatusInternalServerError))
	}
	data := model.Plan{}
	err := s.GORMDB.Debug().Where("id=@id", sql.Named("id", plan_id)).Find(&data).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse(err.Error(), http.StatusInternalServerError))
	}
	if data.Name == "" || data.Description == "" {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse("Invalid PlanID", http.StatusInternalServerError))
	}

	return ctx.JSON(http.StatusOK, model.SuccessResponse(data, http.StatusOK))
}

type Plan struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s Server) AddPlans(ctx echo.Context) error {
	id := "2e146110-7bf1-11ec-90d6-0242ac120003"
	var req = model.Plan{ID: &id, Name: "test1", Description: "Basic Plan"}
	err := s.GORMDB.Debug().Create(&req).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse(err.Error(), http.StatusInternalServerError))
	}
	return ctx.JSON(http.StatusOK, model.SuccessResponse("Success", http.StatusOK))
}
func (s Server) AddResourceType(ctx echo.Context) error {
	id := "230d8bd2-7cc5-11ec-90d6-0242ac120003"
	var req = model.ResourceTypes{ID: &id, Name: "STORAGE", Unit: "Terabytes"}
	err := s.GORMDB.Debug().Create(&req).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse(err.Error(), http.StatusInternalServerError))
	}
	return ctx.JSON(http.StatusOK, model.SuccessResponse("Success", http.StatusOK))
}
func (s Server) AddPlanQuotaDefault(ctx echo.Context) error {
	id := "230d8bd2-7cc5-11ec-90d6-0242ac120003"
	planID := "2e146110-7bf1-11ec-90d6-0242ac120003"
	resourceTypeID := "1783e71c-7cb5-11ec-90d6-0242ac120003"
	var req = model.PlanQuotaDefaults{ID: &id, PlanID: &planID, ResourceTypeID: &resourceTypeID, QuotaValue: 1000}
	err := s.GORMDB.Debug().Create(&req).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse(err.Error(), http.StatusInternalServerError))
	}
	return ctx.JSON(http.StatusOK, model.SuccessResponse("Success", http.StatusOK))
}

func (s Server) AddUserPlanDetails(ctx echo.Context) error {
	id := "230d8bd2-7cc5-11ec-90d6-0242ac120003"
	planID := "2e146110-7bf1-11ec-90d6-0242ac120003"
	userID := "fbdd1f8c-52c1-11ec-bf63-0242ac130002"

	var req = model.UserPlans{ID: &id, UserID: &userID, PlanID: &planID}
	err := s.GORMDB.Debug().Create(&req).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ErrorResponse(err.Error(), http.StatusInternalServerError))
	}
	return ctx.JSON(http.StatusOK, model.SuccessResponse("Success", http.StatusOK))
}

package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/csguojin/reserve/dal/mocks"
	"github.com/csguojin/reserve/http/handler"
	"github.com/csguojin/reserve/model"
	"github.com/csguojin/reserve/service"
)

func TestHandlerStruct_GetAllRoomsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	testService := service.NewService(mockDAL)

	httpHandler := handler.NewHandler(testService)

	router := gin.Default()

	router.GET("/rooms", httpHandler.GetAllRoomsHandler)

	expectedRooms := []*model.Room{
		{
			ID:       1,
			Name:     "test room 1",
			Capacity: 10,
		},
		{
			ID:       2,
			Name:     "test room 2",
			Capacity: 20,
		},
		{
			ID:       3,
			Name:     "test room 3",
			Capacity: 30,
		},
	}

	mockDAL.EXPECT().GetAllRooms(&model.Pager{Page: 1, PerPage: 3}).Return(expectedRooms, nil).Times(1)

	req, err := http.NewRequest("GET", "/rooms?page=1&per_page=3", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)

	var response []*model.Room
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedRooms, response)
}

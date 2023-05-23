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

func TestHandlerStruct_GetAllSeatsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDAL := mocks.NewMockDal(ctrl)

	testService := service.NewService(mockDAL)

	httpHandler := handler.NewHandler(testService)

	router := gin.Default()

	router.GET("/rooms/:room_id/seats", httpHandler.GetAllSeatsHandler)

	var roomID int = 1

	seats := []*model.Seat{
		{
			ID:     1,
			RoomID: 1,
			Name:   "Test-R1-S1",
		},
		{
			ID:     2,
			RoomID: 1,
			Name:   "Test-R1-S2",
		},
		{
			ID:     3,
			RoomID: 1,
			Name:   "Test-R1-S3",
		},
	}

	mockDAL.EXPECT().GetAllSeats(roomID).Return(seats, nil).Times(1)

	req, err := http.NewRequest("GET", "/rooms/1/seats", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response []*model.Seat
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, seats, response)
}

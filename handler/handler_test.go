package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"up_time_monitor/database"
	mockdb "up_time_monitor/database/mocks"
	"up_time_monitor/handler"
	mt "up_time_monitor/monitor"
	mockmt "up_time_monitor/monitor/mocks"
	st "up_time_monitor/structures"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

func TestDeleteHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDataBase := mockdb.NewMockDatabase(mockCtrl)
	database.SetDatabase(mockDataBase)

	mockMonitor := mockmt.NewMockMonitor(mockCtrl)
	mt.SetMonitor(mockMonitor)

	id := StringToUUID("afcec707-0bb4-4251-9f34-c07744935e6d")

	router := getRouter()
	router.DELETE("/urls/:id", handler.DeleteURL)

	req, _ := http.NewRequest("DELETE", "http://localhost:8080/urls/afcec707-0bb4-4251-9f34-c07744935e6d", nil)
	w := httptest.NewRecorder()
	urlinfo := st.URLInfo{
		Base:     st.Base{ID: id},
		Crawling: true,
	}
	mockDataBase.EXPECT().GetURLByID(id).Return(urlinfo, nil)
	mockMonitor.EXPECT().StopMonitoring(id)
	mockDataBase.EXPECT().DeleteFromDatabase(id)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fail()
	}
	_, err1 := ioutil.ReadAll(w.Body)
	if err1 != nil {
		t.Fail()
	}
}

func TestFethInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDataBase := mockdb.NewMockDatabase(mockCtrl)
	database.SetDatabase(mockDataBase)

	id := StringToUUID("afcec707-0bb4-4251-9f34-c07744935e6d")
	router := getRouter()
	router.GET("/urls/:id", handler.FetchInfo)

	req, _ := http.NewRequest("GET", "http://localhost:8080/urls/afcec707-0bb4-4251-9f34-c07744935e6d", nil)
	w := httptest.NewRecorder()

	urlinfo := st.URLInfo{
		Base:             st.Base{ID: id},
		URL:              "this is test url",
		CrawlTimeout:     1,
		Frequency:        2,
		FailureThreshold: 3,
		FailureCount:     4,
		Status:           "active",
		Crawling:         true,
	}

	mockDataBase.EXPECT().GetURLByID(id).Return(urlinfo, nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}
	v, err1 := ioutil.ReadAll(w.Body)
	if err1 != nil {
		fmt.Println(err1.Error())
		t.Fail()
	}
	var info struct {
		URL              string    `json:"url"`
		CrawlTimeout     int       `json:"crawl_timeout"`
		FailureCount     int       `json:"failure_count"`
		Frequency        int       `json:"frequency"`
		ID               uuid.UUID `json:"id"`
		Status           string    `json:"status"`
		FailureThreshold int       `json:"failure_threshold"`
	}

	err := json.Unmarshal(v, &info)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	checkEquality(t, urlinfo.Frequency, info.Frequency)
	checkEquality(t, urlinfo.FailureCount, info.FailureCount)
	checkEquality(t, urlinfo.FailureThreshold, info.FailureThreshold)
	checkEquality(t, urlinfo.URL, info.URL)
	checkEquality(t, urlinfo.CrawlTimeout, info.CrawlTimeout)
	checkEquality(t, urlinfo.ID, info.ID)
	checkEquality(t, urlinfo.Status, info.Status)
}

func TestActivater(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDataBase := mockdb.NewMockDatabase(mockCtrl)
	database.SetDatabase(mockDataBase)

	mockMonitor := mockmt.NewMockMonitor(mockCtrl)
	mt.SetMonitor(mockMonitor)

	id := StringToUUID("afcec707-0bb4-4251-9f34-c07744935e6d")

	router := getRouter()
	router.POST("/urls/:id/activate", handler.Activater)

	req, _ := http.NewRequest("POST", "http://localhost:8080/urls/afcec707-0bb4-4251-9f34-c07744935e6d/activate", nil)
	w := httptest.NewRecorder()
	urlinfo := st.URLInfo{
		Base:     st.Base{ID: id},
		Crawling: false,
	}
	urlinfoUpdate := st.URLInfo{
		Base:         st.Base{ID: id},
		Crawling:     true,
		Status:       "active",
		FailureCount: 0,
	}
	mockDataBase.EXPECT().GetURLByID(id).Return(urlinfo, nil)
	mockDataBase.EXPECT().UpdateDatabase(urlinfoUpdate)
	mockMonitor.EXPECT().StartMonitoring(urlinfoUpdate)

	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
}

func TestDeactivater(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDataBase := mockdb.NewMockDatabase(mockCtrl)
	database.SetDatabase(mockDataBase)

	mockMonitor := mockmt.NewMockMonitor(mockCtrl)
	mt.SetMonitor(mockMonitor)

	id := StringToUUID("afcec707-0bb4-4251-9f34-c07744935e6d")

	router := getRouter()
	router.POST("/urls/:id/deactivate", handler.Deactivater)

	req, _ := http.NewRequest("POST", "http://localhost:8080/urls/afcec707-0bb4-4251-9f34-c07744935e6d/deactivate", nil)
	w := httptest.NewRecorder()
	urlinfo := st.URLInfo{
		Base:     st.Base{ID: id},
		Crawling: true,
	}
	mockDataBase.EXPECT().GetURLByID(id).Return(urlinfo, nil)
	mockDataBase.EXPECT().UpdateColumnInDatabase(id, "crawling", false)
	mockMonitor.EXPECT().StopMonitoring(id)

	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	_, err1 := ioutil.ReadAll(w.Body)
	if err1 != nil {
		t.Fail()
	}

}

func TestPatchInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDataBase := mockdb.NewMockDatabase(mockCtrl)
	database.SetDatabase(mockDataBase)

	mockMonitor := mockmt.NewMockMonitor(mockCtrl)
	mt.SetMonitor(mockMonitor)

	id := StringToUUID("afcec707-0bb4-4251-9f34-c07744935e6d")

	router := getRouter()
	router.PATCH("/urls/:id", handler.PatchInfo)

	var jsonStr = []byte(`{"frequency":2, "crawl_timeout" : 1, "failure_threshold" : 3}`)

	req, _ := http.NewRequest("PATCH", "http://localhost:8080/urls/afcec707-0bb4-4251-9f34-c07744935e6d", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	urlinfo := st.URLInfo{
		Base:     st.Base{ID: id},
		Crawling: true,
		URL:      "www.test.com",
	}

	urlinfoUpdate := st.URLInfo{
		URL:              "www.test.com",
		Base:             st.Base{ID: id},
		Crawling:         true,
		Status:           "active",
		FailureCount:     0,
		Frequency:        2,
		FailureThreshold: 3,
		CrawlTimeout:     1,
	}

	mockDataBase.EXPECT().GetURLByID(id).Return(urlinfo, nil)
	mockDataBase.EXPECT().UpdateDatabase(urlinfoUpdate)
	mockMonitor.EXPECT().StopMonitoring(id)
	mockMonitor.EXPECT().StartMonitoring(urlinfoUpdate)

	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	data, err1 := ioutil.ReadAll(w.Body)
	if err1 != nil {
		t.Fail()
	}

	var info struct {
		URL              string    `json:"url"`
		CrawlTimeout     int       `json:"crawl_timeout"`
		FailureCount     int       `json:"failure_count"`
		Frequency        int       `json:"frequency"`
		ID               uuid.UUID `json:"id"`
		Status           string    `json:"status"`
		FailureThreshold int       `json:"failure_threshold"`
	}

	err := json.Unmarshal(data, &info)
	if err != nil {
		fmt.Println("here : ", err.Error())
		t.Fail()
	}
	checkEquality(t, urlinfoUpdate.Frequency, info.Frequency)
	checkEquality(t, urlinfoUpdate.FailureCount, info.FailureCount)
	checkEquality(t, urlinfoUpdate.FailureThreshold, info.FailureThreshold)
	checkEquality(t, urlinfoUpdate.URL, info.URL)
	checkEquality(t, urlinfoUpdate.CrawlTimeout, info.CrawlTimeout)
	checkEquality(t, urlinfoUpdate.ID, info.ID)
	checkEquality(t, urlinfoUpdate.Status, info.Status)
}

func TestUrlRegister(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDataBase := mockdb.NewMockDatabase(mockCtrl)
	database.SetDatabase(mockDataBase)

	mockMonitor := mockmt.NewMockMonitor(mockCtrl)
	mt.SetMonitor(mockMonitor)

	id := StringToUUID("afcec707-0bb4-4251-9f34-c07744935e6d")

	router := getRouter()
	router.POST("/urls", handler.URLRegister)
	var jsonStr = []byte(`{"frequency":2, "crawl_timeout" : 1, "failure_threshold" : 3, "url":"www.test.com"}`)

	req, _ := http.NewRequest("POST", "http://localhost:8080/urls", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()

	urlinfoUpdate := st.URLInfo{
		URL:              "www.test.com",
		Base:             st.Base{ID: id},
		Crawling:         true,
		Status:           "active",
		FailureCount:     0,
		Frequency:        2,
		FailureThreshold: 3,
		CrawlTimeout:     1,
	}

	mockDataBase.EXPECT().CreateDataBase(gomock.Any()).Return(id, nil)
	mockMonitor.EXPECT().StartMonitoring(urlinfoUpdate)

	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	data, err1 := ioutil.ReadAll(w.Body)
	if err1 != nil {
		t.Fail()
	}

	var info struct {
		URL              string    `json:"url"`
		CrawlTimeout     int       `json:"crawl_timeout"`
		FailureCount     int       `json:"failure_count"`
		Frequency        int       `json:"frequency"`
		ID               uuid.UUID `json:"id"`
		Status           string    `json:"status"`
		FailureThreshold int       `json:"failure_threshold"`
	}

	err := json.Unmarshal(data, &info)
	if err != nil {
		fmt.Println("here : ", err.Error())
		t.Fail()
	}
	checkEquality(t, urlinfoUpdate.Frequency, info.Frequency)
	checkEquality(t, urlinfoUpdate.FailureCount, info.FailureCount)
	checkEquality(t, urlinfoUpdate.FailureThreshold, info.FailureThreshold)
	checkEquality(t, urlinfoUpdate.URL, info.URL)
	checkEquality(t, urlinfoUpdate.CrawlTimeout, info.CrawlTimeout)
	checkEquality(t, urlinfoUpdate.ID, info.ID)
	checkEquality(t, urlinfoUpdate.Status, info.Status)

}

func checkEquality(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		fmt.Println("Expected: ", a, ",Got: ", b)
		t.Fail()
	}
}

func getRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	return router
}

func StringToUUID(st string) uuid.UUID {
	id, err := uuid.FromString(st)
	if err != nil {
		fmt.Println("Error while Converting UID", err.Error())
	}
	return id
}

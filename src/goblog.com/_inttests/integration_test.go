package _inttests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"goblog.com/api/db"
	"goblog.com/api/utils"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goblog API Integration Test Suite")
}

// Helper methods for running GET/POST/PUT methods without boiler plate
func get(route string) *http.Response {
	return callServer(route, http.MethodGet, nil)
}

func post(route string, payload interface{}) *http.Response {
	return callServer(route, http.MethodPost, payload)
}

func put(route string, payload interface{}) *http.Response {
	return callServer(route, http.MethodPut, payload)
}

func delete(route string) *http.Response {
	return callServer(route, http.MethodDelete, nil)
}

func callServer(route string, method string, payload interface{}) *http.Response {
	if !strings.HasPrefix(route, "/") {
		GinkgoT().Fatalf("Routes must start with a forward slash. Offender: %v", route)
	}

	port := os.Getenv("PORT")
	body, err := json.Marshal(payload)
	if err != nil {
		GinkgoT().Fatalf("Received an error when marshalling JSON for body: %v", err.Error())
	}

	request, err := http.NewRequest(method, "http://localhost:"+port+route, bytes.NewReader(body))
	if err != nil {
		GinkgoT().Fatalf("Received an error when creating the request: %v", err.Error())
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		GinkgoT().Fatalf("Received an error when sending the request: %v", err.Error())
	}

	return response
}

// Lambda based function to allow for less verbose post validations
type bodyValidator func(bodyToValidate []byte)

func ensureAndValidatePayload(response *http.Response, expectedStatusCode int, fn bodyValidator) {
	Expect(response.StatusCode).To(Equal(expectedStatusCode))

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	// Should ship an error packet
	Expect(err).To(BeNil(), "Reading body from response failed")
	Expect(body).ToNot(BeNil())

	fn(body)
}

func expectEmptyObjectAnd200(response *http.Response) {
	var null utils.EmptyObject
	Expect(response.StatusCode).To(Equal(http.StatusOK))

	body, err := ioutil.ReadAll(response.Body)
	Expect(string(body)).To(MatchJSON("{}"))
	defer response.Body.Close()

	err = json.NewDecoder(bytes.NewReader(body)).Decode(&null)
	Expect(err).To(BeNil())
}

func expectEmptyListAnd200(response *http.Response) {
	var null []utils.EmptyObject
	Expect(response.StatusCode).To(Equal(http.StatusOK))

	body, err := ioutil.ReadAll(response.Body)
	Expect(string(body)).To(MatchJSON("[]"))
	defer response.Body.Close()

	err = json.NewDecoder(bytes.NewReader(body)).Decode(&null)
	Expect(err).To(BeNil())
}

func getAndPingDB() *sql.DB {
	creds, err := db.GetCredsFromEnv()
	if err != nil {
		GinkgoT().Fatalf("Caught error while getting DB creds: %v", err.Error())
	}

	db, err := db.NewDBConnection(creds)
	if err != nil {
		GinkgoT().Fatalf("Caught error while creating DB instance: %v", err.Error())
	}

	err = db.Ping()
	if err != nil {
		GinkgoT().Fatalf("Caught error while attempting to ping database: %v", err.Error())
	}

	return db
}

func clearTable(tableName string, db *sql.DB) {
	results, err := db.Exec("DELETE FROM " + tableName)
	if err != nil {
		GinkgoT().Fatalf("Caught error while attempting to clear target table: %v", err.Error())
	}

	if count, _ := results.RowsAffected(); count == 0 {
		GinkgoT().Logf("Unneeded table clear ran. No rows affected")
	}
}

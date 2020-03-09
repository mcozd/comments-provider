package user

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("user full info handler is successful", func() {

		req, err := http.NewRequest("GET", "/user/1", nil)
		if err != nil {
			GinkgoT().Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(UserFullInfoHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			GinkgoT().Errorf("handler returned wrong status code: got %v wanted %v",
				status, http.StatusOK)
		}
	})
})

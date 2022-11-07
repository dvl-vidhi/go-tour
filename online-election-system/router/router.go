package router

import (
	"net/http"
	"online-election-system/controller"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/api/user/add", controller.AddUser)
	router.HandleFunc("/api/user/verify/", controller.VerifyUser)
	router.HandleFunc("/api/user/update/", controller.UpdateUser)
	router.HandleFunc("/api/user/search/", controller.SearchOneUser)
	router.HandleFunc("/api/user/search-by-filter", controller.SearchMultipleUser)
	router.HandleFunc("/api/user/delete/", controller.DeleteUser)
	// router.HandleFunc("/api/user/deactivate", controller.)
	router.HandleFunc("/api/election/add", controller.AddElection)
	router.HandleFunc("/api/candidate/add", controller.AddCandidate)
	router.HandleFunc("/api/candidate/verify/", controller.VerifyCandidate)
	router.HandleFunc("/api/election/update/", controller.UpdateElection)
	router.HandleFunc("/api/election/search/", controller.SearchOneElection)
	router.HandleFunc("/api/election/search-by-filter", controller.SearchMultipleElection)
	router.HandleFunc("/api/election/deactivate", controller.DeactivateElection)

	router.HandleFunc("/api/login", controller.Login)

	return router
}

package router

import (
	"net/http"

	"github.com/VladRomanciuc/Go-classes/api/models"
	"github.com/gorilla/mux"

)

//Mux router methods

var muxRouter = mux.NewRouter()

type routerMux struct{}

func NewRouterMux() models.Router{
	return &routerMux{}
}

func (*routerMux) GET(url string, f func(w http.ResponseWriter, r *http.Request)){
	muxRouter.HandleFunc(url, f).Methods("GET")
}
func (*routerMux) POST(url string, f func(w http.ResponseWriter, r *http.Request)){
	muxRouter.HandleFunc(url, f).Methods("POST")
}
func (*routerMux) DELETE(url string, f func(w http.ResponseWriter, r *http.Request)){
	muxRouter.HandleFunc(url, f).Methods("DELETE")
}
func (*routerMux)	SERVE(port string){
	http.ListenAndServe(port, muxRouter)
}
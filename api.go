package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func WriteJson(w http.ResponseWriter , status int , v any) {
	jData , err := json.Marshal(v);
	if err != nil {
      fmt.Println("unable to respond with json");
	}
    w.WriteHeader(status)
    w.Header().Set("Content-Type", "application/json")
    w.Write(jData);
}

type apiFunc func(http.ResponseWriter , *http.Request) error

type ApiError struct{
   Error string
}

type ApiServer struct {
	 listenAddr string
}

func makeHttpHandlerFunc(f apiFunc)http.HandlerFunc {

  return func(w http.ResponseWriter , r *http.Request){
  	     if err := f(w,r); err!=nil {
  	     	 WriteJson(w,http.StatusBadRequest , ApiError{
  	     	 	Error: err.Error(),
  	     	 })
  	     }
  }

}



func NewApiServer(listenAddr string) *ApiServer {
   return &ApiServer{
   	 listenAddr: listenAddr,
   } 
}

func(s *ApiServer) Run(){
	 router := mux.NewRouter();
	 router.HandleFunc("/account" , makeHttpHandlerFunc(s.handleAccount))
	 router.HandleFunc("/account/{id}" , makeHttpHandlerFunc(s.handleGetAccount))
	 log.Println("JSON API Running at port")
	 http.ListenAndServe(s.listenAddr ,router);
}

func (s *ApiServer) handleAccount(w http.ResponseWriter,r *http.Request )error{
  if r.Method =="GET"{
  	 return s.handleGetAccount(w,r)
  } 
  if r.Method=="POST"{
  	 return s.handleCreateAccount(w,r)
  }
  if r.Method == "DELETE" {
  	 return s.handleDeleteAccount(w,r)
  }

  return fmt.Errorf("Method not allowed %v" , r.Method);
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request )error{
  id := mux.Vars(r)["id"];
  //account := NewAccount("Eswar" , "Mamidi");
  log.Println(id)
  WriteJson(w,http.StatusOK , &Account{});
  return nil;
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter,r *http.Request )error{
  return nil;
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter,r *http.Request )error{
  return nil;
}

func (s *ApiServer) handleTransferAccount(w http.ResponseWriter,r *http.Request )error{
  return nil;
}


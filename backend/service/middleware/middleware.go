package middleware

import (
    "net/http"

    "github.com/DamoFD/virtual-business/types"
)

type MiddlewareService struct {}

func NewMiddlewareService() types.Middleware {
    return &MiddlewareService{}
}

func (m *MiddlewareService) RateLimit(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        next.ServeHTTP(w, r)
    })
}

func (m *MiddlewareService) WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        //
    })
}

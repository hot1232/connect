package api

import (
	"github.com/emicklei/go-restful"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
)

func JwtAuthMiddleware(r *restful.Request, w *restful.Response,chain *restful.FilterChain){
	if r.Request.URL.Path == "/getToken" {

	} else if r.Request.URL.Path == "/verifyToken"{

	} else {
		jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte("My Secret"), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})

		if err := jwtMiddleware.CheckJWT(w.ResponseWriter, r.Request); err != nil {
			//w.WriteError(400,err)
			w.Write([]byte("errrr"))
			glog.Errorf("Authentication error: %v", err)
			return
		}
	}

	chain.ProcessFilter(r, w)
}


//cas token auth
func CasJwtTokenAuthMiddleware(r *restful.Request,w *restful.Response,chain *restful.FilterChain){
	if r.Request.URL.Path == "getToken" {

	} else if r.Request.URL.Path == "/verifyToken" {

	} else {
		//cas token auth
	}

	chain.ProcessFilter(r, w)
}
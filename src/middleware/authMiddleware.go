package middleware

import (
	"context"
	"fmt"
	//"log"
	"net/http"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"


)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {

            keyset, _ := jwk.Fetch(context.Background(),"https://cognito-idp.ap-south-1.amazonaws.com/ap-south-1_l2HIVx6A2/.well-known/jwks.json")
            token:= []byte(r.Header["Token"][0])
            parsedToken, err := jwt.Parse(
                token, //token is a []byte
                 jwt.WithKeySet(keyset),
                 jwt.WithValidate(true),
                 jwt.WithIssuer("https://cognito-idp.ap-south-1.amazonaws.com/ap-south-1_l2HIVx6A2"),
                 jwt.WithClaimValue("email_verified", true),
             )
        
             if(err != nil){
                fmt.Println(err)
                fmt.Print("Error")
            }else {
                claims := parsedToken.PrivateClaims()
                
                //TODO : 
        
                for key,val := range claims{
                    fmt.Printf("Key: %v, value: %v\n", key, val)
                }
        
				next.ServeHTTP(w, r)
            }
        
        } else {

            fmt.Fprintf(w, "Not Authorized")
        }
       
    })
}


func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        if r.Header["Token"] != nil {

            keyset, _ := jwk.Fetch(context.Background(),"https://cognito-idp.ap-south-1.amazonaws.com/ap-south-1_l2HIVx6A2/.well-known/jwks.json")
            token:= []byte(r.Header["Token"][0])
            parsedToken, err := jwt.Parse(
                token, //token is a []byte
                 jwt.WithKeySet(keyset),
                 jwt.WithValidate(true),
                 jwt.WithIssuer("https://cognito-idp.ap-south-1.amazonaws.com/ap-south-1_l2HIVx6A2"),
                 jwt.WithClaimValue("email_verified", true),
             )
        
             if(err != nil){
                fmt.Println(err)
                fmt.Print("Error")
            }else {
                claims := parsedToken.PrivateClaims()
                
        
                for key,val := range claims{
                    fmt.Printf("Key: %v, value: %v\n", key, val)
                }
        
                endpoint(w, r)
            }
        
        } else {

            fmt.Fprintf(w, "Not Authorized")
        }
    })
}
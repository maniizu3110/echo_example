package snippets

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
)


func AdminUser(c echo.Context) (err error) {
	cookie, err := c.Cookie("auth")
	if err != nil {
		if strings.Contains(err.Error(), "named cookie not present") {
			return c.String(http.StatusUnauthorized, "you dont have any cookie")
		}
		log.Println(err)
		return err
	}
	firebaseApp := InitializeAppDefault()
	token := cookie.Value
	VerifyIDToken(c.Request().Context(), firebaseApp, token)
	return nil

}


func InitializeAppDefault() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app
}


func VerifyIDToken(ctx context.Context, app *firebase.App, idToken string) *auth.Token {
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)
	return token
}
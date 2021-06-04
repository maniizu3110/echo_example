package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	//hello
)
//tokenはfirebaseからidToken(jwt)がcookieに登録されている

//CheckAdminMiddlewares 適切なTokenを持っているかチェック
func CheckAdminMiddlewares(g *echo.Group){
	g.Use(checkToken)
}

func checkToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("auth")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you dont have any cookie")
			}
			log.Println(err)
			return err
		}
		firebaseApp := initializeAppDefault()
		token := cookie.Value
		verifyIDToken(c.Request().Context(), firebaseApp, token)
		return next(c)
	}
}

func initializeAppDefault() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app
}

func verifyIDToken(ctx context.Context, app *firebase.App, idToken string) *auth.Token {
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	return token
}
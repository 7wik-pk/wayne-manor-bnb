package web

// // NoSurf is the csrf protection middleware
// // DOESNT WORK WITH GIN :/ NEED AN ALTERNATIVE
// func NoSurf(ctx *gin.Context) http.Handler {

// 	csrfHandler := nosurf.New(ctx.Handler())

// 	csrfHandler.SetBaseCookie(http.Cookie{
// 		HttpOnly: true,
// 		Path:     "/",
// 		Secure:   app.InProduction,
// 		SameSite: http.SameSiteLaxMode,
// 	})
// 	return csrfHandler
// }

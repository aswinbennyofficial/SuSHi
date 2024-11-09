package routes

import (
	"net/http"
	// "os"
	"strings"

	// "path/filepath"

	
	"github.com/aswinbennyofficial/SuSHi/models"
	"github.com/go-chi/chi/v5"
	
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/cors"
)

func Load(config models.Config) {
	r:=config.Router
	// Serve static files (e.g., xterm.js frontend)
	// workDir, _ := os.Getwd()
	// filesDir := http.Dir(filepath.Join(workDir, "static"))
	// FileServer(r, "/", filesDir)

	corsOptions := cors.Options{
		// Allowing all origins for demonstration. You can replace "*" with specific origins for stricter policies.
		AllowedOrigins: []string{"*"}, // You can specify your front-end origin here, e.g., "http://localhost:3000"
		AllowedMethods: []string{
			http.MethodGet, 
			http.MethodPost, 
			http.MethodPut, 
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
			"X-Requested-With",
			"Accept",
			"Origin",
		},
		ExposedHeaders: []string{
			"Content-Length",
		},
		MaxAge: 3600,
	}

	// Apply the CORS middleware
	r.Use(cors.Handler(corsOptions))

	r.Route("/api/v1/auth", func(r chi.Router){
		loadoAuthRoutes(r,config)
	})

	
	// JWT authentication middleware
	tokenAuth := jwtauth.New("HS256", []byte(config.JWTSecret), nil)

	r.Group(func(r chi.Router) {
		
		r.Use(jwtauth.Verifier(tokenAuth))
        r.Use(jwtauth.Authenticator(tokenAuth))

		// Define API version base path
		r.Route("/api/v1", func(r chi.Router) {
			loadMachineRoutes(r, config)
		})
	})

	// Load routes
	loadSSHRoutes(r, config)
	



	
	
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
// SPDX-FileCopyrightText: © 2020 Brett Smith <xbcsmith@gmail.com>
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/fatih/structs"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/oklog/ulid"

	"github.com/xbcsmith/mechserv/templates"
)

// GetEnv returns an env variable value or a default
// GetEnv func takes no as input and returns key, fallback string string
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// NewULID returns a ULID as a string.
// NewULID func takes no as input and returns string
func NewULID() string {
	newid, _ := ulid.New(ulid.Timestamp(time.Now()), rand.Reader)
	return newid.String()
}

// maketmpl helper function
func maketmpl(data map[string]interface{}, tmpl string) (string, error) {
	builder := &strings.Builder{}
	t := template.Must(template.New("new").Parse(tmpl))
	if err := t.Execute(builder, data); err != nil {
		return ``, err
	}
	s := builder.String()
	return s, nil
}

// Mechs struct for mechs
type Mechs struct {
	Mechs []Mech `json:"mechs"`
}

// Mech struct for mech
type Mech struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Release     string `json:"release"`
	Description string `json:"description"`
}

// NewMech func takes no as input and returns *Mech
func NewMech() *Mech {
	return &Mech{
		ID: NewULID(),
	}
}

// DecodeMechFromJSON func takes reader io.Reader as input and returns *Mech, error
func DecodeMechFromJSON(reader io.Reader) (*Mech, error) {
	mech := &Mech{}
	err := json.NewDecoder(reader).Decode(mech)
	if err != nil {
		return nil, err
	}
	mech.ID = NewULID()
	return mech, nil
}

// mechs variable for mechs
var mechs []Mech

// Config struct for config
type Config struct {
	Host string
	Port string
}

// NewConfig func takes no as input and returns host, port string *Config
func NewConfig(host, port string) *Config {
	return &Config{
		Host: host,
		Port: port,
	}
}

// Resp struct for resp
type Resp struct {
	ID    string `json:"id"`
	Mech  *Mech  `json:"mech"`
	Error string `json:"error"`
}

// NewAPI func takes no as input and returns cfg *Config *chi.Mux
// NewAPI create the routes for server
func NewAPI(cfg *Config) *chi.Mux {
	router := chi.NewRouter()

	s := NewServer(cfg)

	router.Group(func(r chi.Router) {
		cors := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "DELETE"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		})
		// Add some middleware to our router
		r.Use(cors.Handler,
			render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
			middleware.Logger,       // log api request calls
			middleware.Compress(5),  // compress results, mostly gzipping assets and json
			middleware.StripSlashes, // match paths with a trailing slash, strip it, and continue routing through the mux
			middleware.Recoverer,    // recover from panics without crashing server
		)

		r.Route("/api/v1", func(r chi.Router) {
			{
				r.Get("/mechs", s.GetMechs())
				r.Post("/mechs", s.CreateMech())
				r.Get("/mechs/{id}", s.GetMech())
				r.Delete("/mechs/{id}", s.DeleteMech())
				r.Get("/style.css", s.GetCSS())
			}
		})
		r.Route("/", func(r chi.Router) {
			{
				r.Get("/", s.GetHomeHTML())
				r.Get("/mechs", s.GetMechsHTML())
				r.Get("/mechs/{id}", s.GetMechHTML())
			}
		})
	})

	return router
}

// Server struct for server
type Server struct {
	cfg *Config
}

// GetHomeHTML func takes no as input and returns http.HandlerFunc
func (s *Server) GetHomeHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := templates.MustAsset("index.html")
		render.HTML(w, r, string(content))
	}
}

// GetCSS returns the CSS for HTML interfaces
func (s *Server) GetCSS() http.HandlerFunc {
	content := templates.MustAsset("mech.css")
	return func(w http.ResponseWriter, r *http.Request) {
		render.PlainText(w, r, string(content))
	}
}

// GetMechs func takes no as input and returns http.HandlerFunc
func (s *Server) GetMechs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, mechs)
	}
}

// GetMechsHTML func takes no as input and returns http.HandlerFunc
func (s *Server) GetMechsHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mechsHTML := string(templates.MustAsset("mechs.html"))
		index := &Mechs{
			Mechs: mechs,
		}
		content, err := maketmpl(structs.Map(index), mechsHTML)
		if err != nil {
			render.Status(r, http.StatusExpectationFailed)
			render.JSON(w, r, &Resp{ID: "", Mech: nil, Error: "Template failed to render"})
			return
		}
		render.HTML(w, r, content)
	}
}

// CreateMech func takes no as input and returns http.HandlerFunc
func (s *Server) CreateMech() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mech, err := DecodeMechFromJSON(r.Body)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, &Resp{ID: "", Mech: nil, Error: "Invalid mech JSON"})
			return
		}
		fmt.Printf("Creating new mech : %s\n", mech.ID)
		mechs = append(mechs, *mech)
		render.JSON(w, r, &Resp{ID: mech.ID, Mech: nil, Error: "null"})
	}
}

// GetMech func takes no as input and returns http.HandlerFunc
func (s *Server) GetMech() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimSpace(chi.URLParam(r, "id"))
		fmt.Printf("GET mech : %s\n", id)
		for i, item := range mechs {
			if item.ID == id {
				render.JSON(w, r, &Resp{ID: item.ID, Mech: &mechs[i], Error: "null"})
				return
			}
		}

		fmt.Printf("Did not find Mech : %s\n", id)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Resp{ID: "", Mech: nil, Error: "missing id"})
	}
}

// GetMechHTML func takes no as input and returns http.HandlerFunc
func (s *Server) GetMechHTML() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mechHTML := string(templates.MustAsset("mech.html"))
		id := strings.TrimSpace(chi.URLParam(r, "id"))
		fmt.Printf("GET mech : %s\n", id)
		for _, item := range mechs {
			if item.ID == id {
				content, err := maketmpl(structs.Map(item), mechHTML)
				if err != nil {
					render.Status(r, http.StatusExpectationFailed)
					render.JSON(w, r, &Resp{ID: "", Mech: nil, Error: "Template failed to render"})
					return
				}
				render.HTML(w, r, content)
				return
			}
		}

		fmt.Printf("Did not find Mech : %s\n", id)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Resp{ID: "", Mech: nil, Error: "missing id"})
	}
}

// DeleteMech func takes no as input and returns http.HandlerFunc
func (s *Server) DeleteMech() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimSpace(chi.URLParam(r, "id"))
		fmt.Printf("Delete ID : %s\n", id)
		for index, item := range mechs {
			if item.ID == id {
				mechs = append(mechs[:index], mechs[index+1:]...)
				render.Status(r, http.StatusOK)
				render.JSON(w, r, mechs)
				return
			}
		}

		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Resp{ID: "", Mech: nil, Error: "Invalid ID"})
	}
}

// NewServer func takes no as input and returns cfg *Config *Server
func NewServer(cfg *Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

// main func takes no as input and returns
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	defer wg.Wait()

	host := GetEnv("MECHSERV_HOST", "")
	port := GetEnv("MECHSERV_PORT", "9999")
	cert := GetEnv("MECHSERV_CERT", "")
	key := GetEnv("MECHSERV_KEY", "")
	cfg := NewConfig(host, port)
	router := NewAPI(cfg)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		panic(err)
	}

	if cert != "" && key != "" {
		servTLSCert, err := tls.LoadX509KeyPair(cert, key)
		if err != nil {
			fmt.Printf("invalid key pair: %v", err)
			panic(err)
		}

		// Create the TLS Config with the CA pool and enable Client certificate validation
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{servTLSCert},
			MinVersion:   tls.VersionTLS13,
		}

		// Create a Server instance to listen on port 8443 with the TLS config
		server.TLSConfig = tlsConfig

		listener = tls.NewListener(listener, tlsConfig)
		fmt.Printf("Serving https://%s:%s/api/v1/mechs\n", host, port)
	} else {
		// Run insecure if certs are not provided.
		fmt.Printf("Serving http://%s:%s/api/v1/mechs\n", host, port)
		fmt.Printf("WARING: TLS not enabled\n")
	}

	fmt.Printf("Serving http://%s:%s/api/v1/mechs\n", host, port)

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.Serve(listener)
		if err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
			fmt.Printf("listener closed")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		fmt.Printf("shutting down server")
		err := server.Shutdown(context.Background())
		if err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
		fmt.Printf("server shut down")
	}()

	wg.Wait()
}

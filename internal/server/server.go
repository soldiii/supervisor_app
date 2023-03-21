package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/soldiii/supervisor_app/internal/models"
	"github.com/soldiii/supervisor_app/internal/store"
)

type Server struct {
	config *ServerConfig
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func NewServer(config *ServerConfig) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) RunServer() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.ConfigureRouter()

	if err := s.ConfigureStore(); err != nil {
		return err
	}
	s.logger.Info("Сервер успешно запущен")

	return http.ListenAndServe(s.config.Addr, s.router)
}

func (s *Server) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

func (s *Server) ConfigureRouter() {
	s.router.HandleFunc("/users", s.HandleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/users/{email}", s.HandleGetUserFromEmail()).Methods(("GET"))
}

// TODO
func (s *Server) HandleGetUserFromEmail() http.HandlerFunc {
	return nil
}

// TODO: вроде добавляет, но ничего не выдает, + добавляет данные при ошибке
func (s *Server) HandleUsersCreate() http.HandlerFunc {
	type request struct {
		Email         string    `json:"email"`
		Name          string    `json:"name"`
		Surname       string    `json:"surname"`
		Patronymic    string    `json:"patronymic"`
		Reg_date_time time.Time `json:"reg_date_time"`
		Password      string    `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &models.User{
			Email:         req.Email,
			Name:          req.Name,
			Surname:       req.Surname,
			Patronymic:    req.Patronymic,
			Reg_date_time: req.Reg_date_time,
			Password:      req.Password,
		}
		s.store.User().Create(u)
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}

}

func (s *Server) ConfigureStore() error {
	st := store.NewStore(s.config.Store)
	if err := st.OpenStore(); err != nil {
		return err
	}
	s.store = st
	return nil
}

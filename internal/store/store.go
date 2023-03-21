package store

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	config         *StoreConfig
	db             *sqlx.DB
	userRepository *UserRepository
}

func NewStore(config *StoreConfig) *Store {

	return &Store{
		config: config,
	}
}

func (st *Store) OpenStore() error {
	db, err := sqlx.Open("pgx", st.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	st.db = db

	return nil
}

func (st *Store) CloseStore() {
	st.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

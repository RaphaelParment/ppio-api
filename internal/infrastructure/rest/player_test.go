package rest

//
//import (
//	"database/sql"
//	"encoding/json"
//	"fmt"
//	"github.com/RaphaelParment/ppio-api/internal/domain/player/model"
//	storage2 "github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence"
//	"log"
//	"net/http"
//	"net/http/httptest"
//	"os"
//	"testing"
//
//	"github.com/gorilla/mux"
//	_ "github.com/lib/pq"
//)
//
//// TODO set up tests with docker
//
//func setup() *server {
//	cfg := storage2.Config{
//		User:       "ppio",
//		Password:   "dummy",
//		Host:       "0.0.0.0",
//		Name:       "ppio_tests",
//		DisableTLS: false,
//	}
//	db, tidy, err := storage2.SetupDB(&cfg)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer tidy()
//	l := log.New(os.Stdout, "ppio-tests: ", log.LstdFlags)
//
//	srv := server{
//		DB:     db,
//		Logger: l,
//		Router: mux.NewRouter(),
//	}
//	srv.routes()
//
//	l.Println("removing items")
//	if err := storage2.RemoveAllPlayers(db); err != nil {
//		l.Printf("could not remove all players; %v", err)
//		return nil
//	}
//
//	l.Println("inserting dummy players")
//	if err := storage2.InsertDummyData(db); err != nil {
//		l.Printf("could not add dummy players; %v", err)
//		return nil
//	}
//
//	return &srv
//}
//
//func TestHandlePlayersGet(t *testing.T) {
//	srv := setup()
//	if srv == nil {
//		t.Fatalf("could not setup server / db")
//	}
//
//	defer t.Cleanup(func() {
//		srv.Logger.Println("removing players")
//		storage2.RemoveAllPlayers(srv.DB)
//		srv.Logger.Println("inserting dummy players")
//		storage2.InsertDummyData(srv.DB)
//		srv.Logger.Println("closing db")
//		srv.DB.Close()
//	})
//
//	tt := []struct {
//		name      string
//		operation func(*sql.DB) error
//		status    int
//		body      []model.Player
//		err       string
//	}{
//		{name: "Regular", operation: nil, body: []model.Player{
//			{ID: 1, FirstName: "Alice", LastName: "David", Email: "alice.david@brol.com", Points: 10},
//			{ID: 2, FirstName: "Bob", LastName: "Raymon", Email: "bob.raymon@brol.com", Points: 0}},
//			status: 200},
//		{name: "Empty", operation: storage2.RemoveAllPlayers, body: nil, status: 200},
//	}
//
//	for _, tc := range tt {
//		t.Run(tc.name, func(t *testing.T) {
//			if tc.operation != nil {
//				if err := tc.operation(srv.DB); err != nil {
//					t.Errorf("could not perform operation; %v", err)
//				}
//			}
//
//			req, err := http.NewRequest("GET", "localhost:9001/players", nil)
//			if err != nil {
//				t.Errorf("could not create request: %v", err)
//			}
//			rec := httptest.NewRecorder()
//			srv.handlePlayersGet()(rec, req)
//
//			res := rec.Result()
//			if res.StatusCode != http.StatusOK {
//				t.Errorf("wrong status code; want %v got %v", http.StatusOK, res.StatusCode)
//			}
//
//			if res.Body != nil {
//
//				var players []model.Player
//				if err := json.NewDecoder(res.Body).Decode(&players); err != nil {
//					t.Errorf("could not convert players to JSON; %v", err)
//				}
//
//				for i, p := range tc.body {
//					if p != players[i] {
//						t.Errorf("failed, expected player: %v got: %v", players[i], p)
//					}
//				}
//			}
//		})
//	}
//}
//
//func TestRouting(t *testing.T) {
//	srv := setup()
//	defer t.Cleanup(func() {
//		srv.Logger.Println("removing players")
//		storage2.RemoveAllPlayers(srv.DB)
//		srv.Logger.Println("inserting dummy players")
//		storage2.InsertDummyData(srv.DB)
//		srv.Logger.Println("closing db")
//		srv.DB.Close()
//	})
//	server := httptest.NewServer(srv)
//	defer server.Close()
//
//	res, err := http.Get(fmt.Sprintf("%s/players", server.URL))
//	if err != nil {
//		t.Fatalf("could not send GET request: %v", err)
//	}
//	if res.StatusCode != http.StatusOK {
//		t.Fatalf("wrong status code; want %v got %v", http.StatusOK, res.StatusCode)
//	}
//
//	if res.Body != nil {
//		results := []model.Player{
//			{ID: 1, FirstName: "Alice", LastName: "David", Email: "alice.david@brol.com", Points: 10},
//			{ID: 2, FirstName: "Bob", LastName: "Raymon", Email: "bob.raymon@brol.com", Points: 0},
//		}
//		var players []model.Player
//		if err := json.NewDecoder(res.Body).Decode(&players); err != nil {
//			t.Fatalf("could not convert players to JSON; %v", err)
//		}
//
//		for i, p := range results {
//			if p != players[i] {
//				t.Errorf("failed, expected player: %v got: %v", players[i], p)
//			}
//		}
//	}
//
//}

package main

import (
	"context"
	"log"
	"net/http"

	"strings"

	"github.com/topfreegames/pitaya/v3/pkg/acceptor"
	"github.com/topfreegames/pitaya/v3/pkg/component"
	"github.com/topfreegames/pitaya/v3/pkg/config"
	"github.com/topfreegames/pitaya/v3/pkg/groups"
	"github.com/topfreegames/pitaya/v3/pkg/timer"
)

type (
	// Room represents a component that contains a bundle of room related handler
	// like Join/Message
	Room struct {
		component.Base
		timer *timer.Timer
		app   pitaya.Pitaya
	}

	// UserMessage represents a message that user sent
	UserMessage struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	// NewUser message will be received when new user join room
	NewUser struct {
		Content string `json:"content"`
	}

	// AllMembers contains all members uid
	AllMembers struct {
		Members []string `json:"members"`
	}

	// JoinResponse represents the result of joining room
	JoinResponse struct {
		Code   int    `json:"code"`
		Result string `json:"result"`
	}
)

// NewRoom returns a Handler Base implementation
func NewRoom(app pitaya.Pitaya) *Room { _ = "STUB: not implemented"; return nil }

// AfterInit component lifetime callback
func (r *Room) AfterInit() { _ = "STUB: not implemented"; return }

// Join room
func (r *Room) Join(ctx context.Context, msg []byte) (*JoinResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// just use s.ID as uid !!!
// binding session uid

// notify others

// new user join group
// add session to group

// on session close, remove it from group

// Message sync last message to all members
func (r *Room) Message(ctx context.Context, msg *UserMessage) { _ = "STUB: not implemented"; return }

var app pitaya.Pitaya

func main() {
	conf := configApp()
	builder := pitaya.NewDefaultBuilder(true, "chat", pitaya.Cluster, map[string]string{}, *conf)
	builder.AddAcceptor(acceptor.NewWSAcceptor(":3250"))
	builder.Groups = groups.NewMemoryGroupService(builder.Config.Groups.Memory)
	app = builder.Build()

	defer app.Shutdown()

	err := app.GroupCreate(context.Background(), "room")
	if err != nil {
		panic(err)
	}

	// rewrite component and handler name
	room := NewRoom(app)
	app.Register(room,
		component.WithName("room"),
		component.WithNameFunc(strings.ToLower),
	)

	log.SetFlags(log.LstdFlags | log.Llongfile)

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))

	go http.ListenAndServe(":3251", nil)

	app.Start()
}

func configApp() *config.PitayaConfig { _ = "STUB: not implemented"; return nil }

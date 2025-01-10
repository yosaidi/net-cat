package clients

import (
	"fmt"
	"strings"
	"sync"
)

type Chat struct {
	mu              *sync.Mutex
	Clients         *AllClients
	PreviousHistory []string
}

type ConnectionStatus struct {
	IsConnected bool
	Name        string
}

type BroadcastMessage struct {
	Message string
	Name    string
	Client  *EachClient
}

func NewChat(mu *sync.Mutex, clients *AllClients) *Chat {
	return &Chat{
		mu:      mu,
		Clients: clients,
	}
}

func (ch *Chat) HandleChatRoutine(statusch chan ConnectionStatus, messagech chan BroadcastMessage) {
	for {
		select {
		case status := <-statusch:
			message := connectionMessage(status.Name, status.IsConnected)
			ch.Broadcast(message, status.Name, "status", statusch)
		case msg := <-messagech:
			ch.mu.Lock()
			ch.PreviousHistory = append(ch.PreviousHistory,FormatText(msg.Name,msg.Message) )
			ch.mu.Unlock()
			ch.Broadcast(msg.Message, msg.Name, "message", statusch)

		}
	}
}
func (ch *Chat) Broadcast(message, pseudo, msgtype string, statusch chan ConnectionStatus) {

	ch.Clients.Mu.Lock()
	defer ch.Clients.Mu.Unlock()
	for client, Conn := range ch.Clients.Clients {
		if client != pseudo {
			switch {
			case msgtype == "status":
				_, err := fmt.Fprint(Conn, message+FormatText(client,"\n"))
				if err != nil {
					delete(ch.Clients.Clients, client)
					ConnectionConfig("offline", client, statusch)
				}

			case msgtype == "message":
				_, err := fmt.Fprint(Conn, "\n"+FormatText(pseudo, message)+"\n"+FormatText(client,""))
				if err != nil {
					delete(ch.Clients.Clients, client)
					ConnectionConfig("offline", client, statusch)

				}
			}

		}
	}

}



func (c *Chat) History(client *EachClient) error {
	c.mu.Lock()
	history := strings.Join(c.PreviousHistory, "\n")
	c.mu.Unlock()

	_, err := fmt.Fprint(client.Conn, history+"\n")
	return err
}

func connectionMessage(name string, isConnected bool) string {
	if isConnected {
		return fmt.Sprintf("\n%s has joined the groupchat \n", name)
	}
	return fmt.Sprintf("\n%s has left the groupchat\n", name)
}

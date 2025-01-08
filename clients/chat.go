package clients

import (
	"fmt"
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
			message := fmt.Sprintf("%s has left the groupchat", status.Name)
			if status.IsConnected {
				message = fmt.Sprintf("%s has joined the groupchat", status.Name)
			}
			ch.Broadcast(message, status.Name, statusch)
		case msg := <-messagech:
			ch.Broadcast(msg.Message, msg.Name, statusch)
		}
	}
}

func (ch *Chat) Broadcast(message, pseudo string, statusch chan ConnectionStatus) {
	ch.Clients.Mu.Lock()
	defer ch.Clients.Mu.Unlock()

	for client, conn := range ch.Clients.Clients {
		if client != pseudo {
			_, err := fmt.Fprint(conn, "\n"+message+"\n"+FormatText(client, ""))
			if err != nil {
				delete(ch.Clients.Clients, client)
				statusch <- ConnectionStatus{IsConnected: false, Name: client}
				continue
			}
		}
	}

}

func (c *Chat) History(client *EachClient) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, history := range c.PreviousHistory {
		_, err := fmt.Fprint(client.Conn, history+"\n")
		if err != nil {
			return err
		}
	}
	return nil
}

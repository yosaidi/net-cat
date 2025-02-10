package clients

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net-cat/logo"
	"strings"
	"sync"
)

type AllClients struct {
	Mu      *sync.Mutex
	Clients map[string]net.Conn
}

type EachClient struct {
	Name string
	Conn net.Conn
}


func NewClients(mu *sync.Mutex) *AllClients {
	return &AllClients{
		Mu:      mu,
		Clients: make(map[string]net.Conn, 10),
	}
}

func NewClient(conn net.Conn) *EachClient {
	return &EachClient{
		Conn: conn,
	}
}

func (c *AllClients) AddClient(pseudo string, client net.Conn) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Clients[pseudo] = client
}

func (c *EachClient) HandleClient(statusch chan ConnectionStatus, messagech chan BroadcastMessage, chat *Chat) {
	defer c.Conn.Close()


	welcomemsg := logo.Logo() + "\n[ENTER YOUR NAME]:"

	_, err := fmt.Fprint(c.Conn, welcomemsg)
	if err != nil {
		log.Println("Error sending welcome message:", err)
		return
	}

	if err := c.AddPseudo(chat.Clients); err != nil {
		fmt.Println("Error adding pseudo:", err)
		return
	}

	chat.Clients.AddClient(c.Name, c.Conn)
	fmt.Printf("%s joined the chat\n", c.Name)

	if err := chat.History(c); err != nil {
		RemoveClient(chat, c.Name, statusch)
		fmt.Println(err)
		return
	}

	ConnectionConfig("online", c.Name, statusch)

	reader := bufio.NewReader(c.Conn)

	for {

		_, err := fmt.Fprint(c.Conn, FormatText(c.Name, ""))
		if err != nil {
			fmt.Println(err)
			RemoveClient(chat, c.Name, statusch)
			return
		}

		message, err := reader.ReadString('\n')

		if err == io.EOF {
			fmt.Printf("\n%s left the chat ", c.Name)
			RemoveClient(chat, c.Name, statusch)
			return
		}

		if err != nil {
			fmt.Println(err)
			RemoveClient(chat, c.Name, statusch)
			return
		}

		message = strings.TrimSpace(message)
		if ValidText(message) {
			messagech <- BroadcastMessage{
				Name:    c.Name,
				Message: message,
				Client:  c,
			}
		}
	}
}

func (c *EachClient) AddPseudo(clients *AllClients) error {

	reader := bufio.NewReader(c.Conn)

	for {
		pseudo, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		pseudo = strings.TrimSpace(pseudo)

		if ok, err := ValidPseudo(c, pseudo, clients); ok {
			c.Name = pseudo
			return nil
		} else if err != nil {
			return fmt.Errorf("invalid pseudo: %w", err)
		}
	}

}

func RemoveClient(chat *Chat, client string, statusch chan ConnectionStatus) {
	chat.Clients.Mu.Lock()
	defer chat.Clients.Mu.Unlock()
	delete(chat.Clients.Clients, client)
	ConnectionConfig("offline",client,statusch)
}

func ConnectionConfig(state, client string, statuch chan ConnectionStatus) chan ConnectionStatus {
	switch {
	case state == "online":
		statuch <- ConnectionStatus{
			IsConnected: true,
			Name:        client,
		}
	case state == "offline":
		statuch <- ConnectionStatus{
			IsConnected: false,
			Name:        client,
		}
	}
	return statuch
}

package clients

import (
	"fmt"
	"strings"
	"time"
)

func FormatText(pseudo, message string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	input := fmt.Sprintf("[%s][%s]:%s", timestamp, pseudo, message)
	return strings.TrimSpace(input)

}



func (c *EachClient) FullGroup() {
	_, err := fmt.Fprintf(c.Conn, "the groupchat is at full capacity, try again later!")
	if err != nil {
		fmt.Println(err)
	}
}

func ValidText(text string) bool {
	if text == "" {
		return false
	}
	return text[0] > 31
}

func ValidPseudo(client *EachClient, pseudo string, clients *AllClients) (bool, error) {
	for _, ch := range pseudo {
		if !ValidText(string(ch)) {
			_, err := fmt.Fprint(client.Conn, "Unsopported name format\nPlease try another name...\n[ENTER YOUR NAME]:")
			if err != nil {
				return false, err
			}
			return false, nil
		}
	}
	if !ValidText(pseudo) {
		_, err := fmt.Fprint(client.Conn, "Unsopported name format\nPlease try another name...\n[ENTER YOUR NAME]:")
		if err != nil {
			return false, err
		}
		return false, nil
	}

	clients.Mu.Lock()
	defer clients.Mu.Unlock()
	if _, exist := clients.Clients[pseudo]; exist {
		_, err := fmt.Fprint(client.Conn, "Specified name has already been taken\nPlease try another name...\n[ENTER YOUR NAME]:")
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

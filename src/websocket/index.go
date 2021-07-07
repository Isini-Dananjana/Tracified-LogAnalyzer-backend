package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"


	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

func init(){
    go manager.start()
}


type ClientManager struct{
    clients map[*Client]bool
    broadcast chan []byte
    register chan *Client
    unregister chan *Client
    unicast chan Message
}

type Client struct {
    id string
    socket *websocket.Conn
    send chan []byte
}

type Message struct{
    Sender string `json:"sender,omitempty"`
    Recipient string `json:"recipient,omitempty"`
    Content   string `json:"content,omitempty"`
}

var manager = ClientManager{
    broadcast:  make(chan []byte),
    register:   make(chan *Client),
    unregister: make(chan *Client),
    clients:    make(map[*Client]bool),
    unicast: make(chan Message),
}

func (manager *ClientManager) start() {
    fmt.Println("Socket started")
    for {
        select {
        case conn := <-manager.register:
            manager.clients[conn] = true
            jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
            manager.send(jsonMessage, conn)
        case conn := <-manager.unregister:
            if _, ok := manager.clients[conn]; ok {
                close(conn.send)
                delete(manager.clients, conn)
                jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
                manager.send(jsonMessage, conn)
            }
        case message := <-manager.broadcast:
            for conn := range manager.clients {
                select {
                case conn.send <- message:
                default:
                    close(conn.send)
                    delete(manager.clients, conn)
                }
            }
        case uniMessage := <-manager.unicast:
            id := uniMessage.Recipient
            fmt.Println("I AM HERE WITH ID" + id)
            
            for conn := range manager.clients{
                if(conn.id == id){
                    fmt.Println("NOW I AM TRUE")
                    fmt.Println(uniMessage.Content)
                    jsonMessage, _ := json.Marshal(&Message{Sender: "TADA", Content: uniMessage.Content})

                    select{
                    case conn.send <- jsonMessage:
                        fmt.Println(string(jsonMessage))
                    default:
                        close(conn.send)
                        delete(manager.clients,conn)
                    }
                }
            }
        }
    }
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
    for conn := range manager.clients {
        if conn != ignore {
            conn.send <- message
        }
    }
}


//Read Whatever Write
//Send To OTHERS
func (c *Client) read() {
    defer func() {
        manager.unregister <- c
        c.socket.Close()
    }()

    for {
        _, message, err := c.socket.ReadMessage()
        fmt.Println(string(message))
        if err != nil {
            manager.unregister <- c
            c.socket.Close()
            break
        }
        //mesageDetail :=Message{}
         //json.Unmarshal(message,&mesageDetail)
         //fmt.Println(string(mesageDetail.Sender))
        jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
        //manager.unicast <- mesageDetail
        manager.broadcast <-jsonMessage
    }
}



//Diplay In Each User
func (c *Client) write() {
    defer func() {
        c.socket.Close()
    }()

    for {
        select {
        case message, ok := <-c.send:
            if !ok {
                c.socket.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }

            c.socket.WriteMessage(websocket.TextMessage, message)
        }
    }
}


func WSPage(res http.ResponseWriter, req *http.Request) {


	//fmt.Fprintf(res,"WebSocket Working")
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	client := &Client{id: uuid.NewV4().String(), socket: conn, send: make(chan []byte)}

	fmt.Println(client.id)
	manager.register <- client

	go client.read()
	go client.write()


}
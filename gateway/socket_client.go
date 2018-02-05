/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package gateway

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type SocketClient struct {
	node *SocketNode

	conn *websocket.Conn

	send chan []byte
}

func (c *SocketClient) read() {
	defer func() {
		c.node.unregister <- c
		c.conn.Close()
	}()
	fmt.Print("read ------->1")
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	fmt.Print("read ------->2")
	for {
		fmt.Print("read ------->33")
		req := &WebsocketRequest{}
		err := c.conn.ReadJSON(&req)
		fmt.Print(req)
		fmt.Print(err)
		fmt.Print("read ------->3333")
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		fmt.Print("read ------->4")
		resp, err := c.handler(*req)
		if err == nil {
			c.send <- resp
		}
	}
}

func (c *SocketClient) handler(req WebsocketRequest) ([]byte, error) {
	walletService := WalletServiceImpl{}
	fmt.Print("read ------->5")
	input := req["ping"]
	fmt.Print("read ------->6")
	inputInt, _ := strconv.Atoi(input)
	return walletService.TestPing(inputInt)
}

func (c *SocketClient) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		fmt.Print("read ------->7")
		select {
		case message, ok := <-c.send:
			fmt.Println("read ------->8")
			fmt.Println(message)
			fmt.Println(ok)
			fmt.Println(string(message))
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := c.conn.WriteJSON(&message)
			if err != nil {
				fmt.Println("write error occurs")
				fmt.Println(err)
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/tobiasprima/kitchen/services/common/genproto/orders"
)

type httpServer struct{
	addr string
}

func NewHttpServer(addr string) *httpServer{
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)

		ctx , cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 24,
			ProductID: 3123,
			Quantity: 2,
		})

		if err != nil {
			log.Fatalln("client error: %v", err)
		}
	})

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Kitchen Orders</title>
</head>
<body>
	<h1>Orders List</h1>
	<table border="1">
		<tr>
			<th>Order Id</th>
			<th>Customer ID</th>
			<th>Quantity</th>
		</tr>
		{{range .}}
		<tr>
			<td>{{.OrderId}}</td>
			<td>{{.CustomerID}}</td>
			<td>{{.Quantity}}</td>
		</tr>
		{{end}}
	</table>
</body>
</html>
`
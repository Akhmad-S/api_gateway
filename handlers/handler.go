package handlers

import "github.com/uacademy/blogpost/api_gateway/clients"

type Handler struct {
	grpcClients *clients.GrpcClients
}

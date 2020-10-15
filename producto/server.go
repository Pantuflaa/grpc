package main
import (
        "log"
        "golang.org/x/net/context"
        "net"
        pb "github.com/Pantuflaa/grpc/producto/pb"
)

type productoServer struct {
	pb.UnimplementedProductoServer

}
func (s *productoServer) Peticion(ctx context.Context, objeto *Objeto) (*Serie, error) {
	log.Printf("Recibi el objeto con la siguiente id: %s", objeto.id)
	return &Serie{serie: 1}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
			log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductoServer(grpcServer, &productoServer{})

	if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}
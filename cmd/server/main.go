package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/shibu1x/ur_v3/pkg/db"
	pb "github.com/shibu1x/ur_v3/pkg/grpc"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRoomServiceServer
}

type RoomWithHouse struct {
	PrefCode  string
	HouseCode string
	RoomCode  string
	PrefName  string
	HouseName string
	Status    string
	Price     int32
	Fee       int32
	Type      string
	Space     int32
	Floor     int32
	LayoutUrl string
	Url       string
}

func (s *server) GetRoomsByPrefCode(ctx context.Context, req *pb.GetRoomsByPrefCodeRequest) (*pb.GetRoomsByPrefCodeResponse, error) {
	query := `
		SELECT
			p.code AS pref_code
			, h.code AS house_code
			, r.room_code
			, p."name" AS pref_name
			, h."name" AS house_name
			, r.status
			, r.price
			, r.fee
			, r.type
			, r.space
			, r.floor
			, r.layout_url
			, r.url
		FROM
			rooms r
			INNER JOIN houses h ON r.house_code = h.code
			INNER JOIN prefs p ON h.pref_code = p.code
		WHERE
			h.pref_code = $1
    `
	var results []RoomWithHouse
	if err := queries.Raw(query, req.PrefCode).BindG(ctx, &results); err != nil {
		return nil, err
	}

	response := &pb.GetRoomsByPrefCodeResponse{}
	for _, result := range results {
		response.Rooms = append(response.Rooms, &pb.Room{
			PrefCode:  result.PrefCode,
			HouseCode: result.HouseCode,
			RoomCode:  result.RoomCode,
			PrefName:  result.PrefCode,
			HouseName: result.HouseName,
			Status:    result.Status,
			Price:     result.Price,
			Fee:       result.Fee,
			Type:      result.Type,
			Space:     result.Space,
			Floor:     result.Floor,
			LayoutUrl: result.LayoutUrl,
			Url:       result.Url,
		})
	}
	return response, nil
}

func main() {
	// データベース接続を確立
	database := db.ConnectDB()
	defer database.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRoomServiceServer(s, &server{})

	fmt.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

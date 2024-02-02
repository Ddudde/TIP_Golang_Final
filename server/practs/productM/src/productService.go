package src

import (
	"context"
	"net/http"

	"practs/productM/db"
	"practs/productM/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedProductServiceServer
}

func (s *Server) GetProduct(ctx context.Context, prof *pb.Empty) (*pb.GetProductResponse, error) {
	// var product models.Prodct

	// result := []*pb.Prodct{}

	// result = append(result, &pb.Prodct{Zag: "Пломбир Из Лавки cолёная карамель", Price: 62, Gram: 75, Img: "caram"},
	// 	&pb.Prodct{Zag: "Фруктовый лёд Вкусовые сосочки", Price: 115, Gram: 75, Img: "lyod"},
	// 	&pb.Prodct{Zag: "Пломбир с клубникой в глазури Кактус", Price: 115, Gram: 80, Img: "kaktyz"},
	// 	&pb.Prodct{Zag: "Смузи клубника", Price: 1100, Gram: 1000, Img: "smyzi"},
	// 	&pb.Prodct{Zag: "Джелато сливочное «Из Лавки» банан", Price: 111, Gram: 90, Img: "dzhelato"},
	// 	&pb.Prodct{Zag: "Мороженое фисташковое с миндалем", Price: 249, Gram: 450, Img: "mindal"},
	// 	&pb.Prodct{Zag: "Пломбир Банан трюфель 12% пакет", Price: 249, Gram: 500, Img: "banana"},
	// 	&pb.Prodct{Zag: "Чудеса света мороженое йогуртно-вишневое", Price: 329, Gram: 450, Img: "chudo"},
	// 	&pb.Prodct{Zag: "Эскимо Ассорти «Монстры»", Price: 549, Gram: 620, Img: "monster"},
	// 	&pb.Prodct{Zag: "Мороженое «Кокосовый мусс с шоколадом»", Price: 65, Gram: 75, Img: "muss"})

	var products []*pb.Prodct

	if result := s.H.DB.Find(&products); result.Error != nil {
		return &pb.GetProductResponse{
			Status: http.StatusNotFound,
			Error:  "Products not found",
		}, nil
	}

	return &pb.GetProductResponse{
		Status:   http.StatusOK,
		Products: products,
	}, nil
}

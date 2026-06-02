package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/itsektionen/mimer/cfg"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/repository"
	"github.com/itsektionen/mimer/service"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/stdlib"
)

func stringPtr(s string) *string {
	return &s
}

func main() {
	mockGroups := []db.CreateGroupParams{
		{
			Name:        "QlubbMästeriet IN-Sektionen Kista",
			Slug:        "qmisk",
			ShortName:   "QMISK",
			Description: stringPtr("Vi anordnar fester"),
			Color:       "#800000",
			ImageUrl:    nil,
			WebsiteUrl:  stringPtr("https://qmisk.se"),
			EstablishedAt: pgtype.Date{
				Time:  time.Date(1994, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
		},
		{
			Name:        "TraditionsMEsterIT",
			Slug:        "tmeit",
			ShortName:   "TMEIT",
			Description: stringPtr("Vi anordnar också fester"),
			Color:       "#000067",
			ImageUrl:    nil,
			WebsiteUrl:  stringPtr("https://tmeit.se"),
			EstablishedAt: pgtype.Date{
				Time:  time.Date(2004, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
		},
		{
			Name:        "ITerativa Klubben",
			Slug:        "itk",
			ShortName:   "ITK",
			Description: stringPtr("Vi anordnar inte fester"),
			Color:       "#006700",
			ImageUrl:    nil,
			WebsiteUrl:  stringPtr("https://itk.gg"),
			EstablishedAt: pgtype.Date{
				Time:  time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
		},
		{
			Name:        "QLAN",
			Slug:        "qlan",
			ShortName:   "QLAN",
			Description: stringPtr("Vi anordnar inte heller fester"),
			Color:       "#800000",
			ImageUrl:    nil,
			WebsiteUrl:  stringPtr("https://qlan.se"),
			EstablishedAt: pgtype.Date{
				Time:  time.Date(1995, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
			DissolvedAt: pgtype.Date{
				Time:  time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
		},
	}

	mockUsers := []db.CreateUserParams{
		{
			FirstName: "John",
			LastName:  "Qmisk",
		},
		{
			FirstName: "Jonny",
			LastName:  "Tmeit",
		},
		{
			FirstName: "Joan",
			LastName:  "Itk",
		},
	}

	env := os.Getenv("MIMER_ENV")
	if env == "" {
		env = "development"
	}
	fmt.Println("env:", env)

	config := cfg.Load()

	ctx := context.Background()

	conn, err := db.SetupPostgresDB(ctx, config.Database.URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(ctx)

	// Create *sql.DB from pgx connection for migrations
	sqlDB := stdlib.OpenDB(*conn.Config())
	defer sqlDB.Close()

	queries := db.New(conn)

	groupRepo := repository.NewGroupRepository(queries)
	userRepo := repository.NewUserRepository(queries)
	positionRepo := repository.NewPositionRepository(queries)
	trusteeRepo := repository.NewTrusteeRepository(queries)

	groupService := service.NewGroupService(groupRepo, trusteeRepo)
	userService := service.NewUserService(userRepo)
	positionService := service.NewPositionService(positionRepo, trusteeRepo)

	groups := make([]model.Group, 0, len(mockGroups))
	users := make([]model.User, 0, len(mockUsers))

	for _, params := range mockGroups {
		c, err := groupService.Create(context.TODO(), params)
		if err != nil {
			log.Fatal(err)
		}
		groups = append(groups, c)
	}

	for _, params := range mockUsers {
		p, err := userService.Create(context.TODO(), params)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, p)
	}

	mockPositions := []db.CreatePositionParams{
		{
			Name:    "QlubbMästare",
			Email:   "qm@qmisk.com",
			GroupID: groups[0].ID,
			EstablishedAt: pgtype.Date{
				Time:  time.Date(1995, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
		},
		{
			Name:    "TraditionsMästare",
			Email:   "tm@tmeit.se",
			GroupID: groups[1].ID,
			EstablishedAt: pgtype.Date{
				Time:  time.Date(2004, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
		},
		{
			Name:    "root",
			Email:   "root@itk.gg",
			GroupID: groups[2].ID,
			EstablishedAt: pgtype.Date{
				Time:  time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
		},
	}

	positions := make([]model.Position, 0, len(mockPositions))

	for _, params := range mockPositions {
		p, err := positionService.Create(context.TODO(), params)
		if err != nil {
			log.Fatal(err)
		}
		positions = append(positions, p)
	}

	for i, p := range positions {
		positionService.Assign(context.TODO(), p.ID, users[i].ID)
	}
}

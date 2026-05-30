package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/itsektionen/mimer/cfg"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/model"
	"github.com/itsektionen/mimer/repository"
	"github.com/itsektionen/mimer/service"
	"github.com/jackc/pgx/v5/stdlib"
)

func stringPtr(s string) *string {
	return &s
}

func main() {
	mockCommittees := []db.CreateCommitteeParams{
		db.CreateCommitteeParams{
			Name:        "QlubbMästeriet IN-Sektionen Kista",
			Slug:        "qmisk",
			ShortName:   "QMISK",
			Description: stringPtr("Vi anordnar fester"),
			Color:       "#800000",
			ImageUrl:    nil,
			WebsiteUrl:  stringPtr("https://qmisk.se"),
		},
		db.CreateCommitteeParams{
			Name:        "TraditionsMEsterIT",
			Slug:        "tmeit",
			ShortName:   "TMEIT",
			Description: stringPtr("Vi anordnar också fester"),
			Color:       "#000067",
			ImageUrl:    nil,
			WebsiteUrl:  stringPtr("https://tmeit.se"),
		},
		db.CreateCommitteeParams{
			Name:        "ITerativa Klubben",
			Slug:        "itk",
			ShortName:   "ITK",
			Description: stringPtr("Vi anordnar inte fester"),
			Color:       "#006700",
			ImageUrl:    nil,
			WebsiteUrl:  stringPtr("https://itk.gg"),
		},
	}

	mockPeople := []db.CreatePersonParams{
		db.CreatePersonParams{
			FirstName: "John",
			LastName:  "Qmisk",
		},
		db.CreatePersonParams{
			FirstName: "Jonny",
			LastName:  "Tmeit",
		},
		db.CreatePersonParams{
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

	committeeRepo := repository.NewCommitteeRepository(queries)
	personRepo := repository.NewPersonRepository(queries)
	positionRepo := repository.NewPositionRepository(queries)
	apiKeyRepo := repository.NewApiKeyRepository(queries)
	trusteeRepo := repository.NewTrusteeRepository(queries)

	committeeService := service.NewCommitteeService(committeeRepo, trusteeRepo)
	personService := service.NewPersonService(personRepo)
	positionService := service.NewPositionService(positionRepo, trusteeRepo)
	apiKeyService := service.NewApiKeyService(apiKeyRepo)

	committees := make([]model.Committee, 0, len(mockCommittees))
	people := make([]model.Person, 0, len(mockPeople))

	for _, params := range mockCommittees {
		c, err := committeeService.CreateCommittee(context.TODO(), params)
		if err != nil {
			log.Fatal(err)
		}
		committees = append(committees, c)
	}

	for _, params := range mockPeople {
		p, err := personService.CreatePerson(context.TODO(), params)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, p)
	}

	mockPositions := []db.CreatePositionParams{
		db.CreatePositionParams{
			Name:        "QlubbMästare",
			Email:       "qm@qmisk.com",
			CommitteeID: committees[0].ID,
		},
		db.CreatePositionParams{
			Name:        "TraditionsMästare",
			Email:       "tm@tmeit.se",
			CommitteeID: committees[1].ID,
		},
		db.CreatePositionParams{
			Name:        "root",
			Email:       "root@itk.gg",
			CommitteeID: committees[2].ID,
		},
	}

	positions := make([]model.Position, 0, len(mockPositions))

	for _, params := range mockPositions {
		p, err := positionService.CreatePosition(context.TODO(), params)
		if err != nil {
			log.Fatal(err)
		}
		positions = append(positions, p)
	}

	for i, p := range positions {
		positionService.AssignPosition(context.TODO(), p.ID, people[i].ID)
	}

	key, err := apiKeyService.CreateApiKey(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(key)
}

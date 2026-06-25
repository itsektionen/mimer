package main

// NOTE: This file is SUPER vibe coded.
// ============== What is this? ==============
// A utility that imports group and position data from:
// https://github.com/itsektionen/trustee-positions
// (assuming it is cloned to .data/)

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/itsektionen/mimer/config"
	"github.com/itsektionen/mimer/db"
	"github.com/itsektionen/mimer/repository"
	"github.com/itsektionen/mimer/service"
	"github.com/jackc/pgx/v5/pgtype"
)

type GroupMetadata struct {
	Name      string
	ShortName string
	Color     string
}

var groupMetadata = map[string]GroupMetadata{
	"board":               {Name: "Sektionsstyrelsen", ShortName: "Styrelsen", Color: "#cc99ff"},
	"business_relations":  {Name: "Näringslivsnämnden", ShortName: "NN", Color: "#cc99ff"},
	"communication":       {Name: "Kommunikationsnämnden", ShortName: "KommN", Color: "#cc99ff"},
	"education":           {Name: "Studienämnden", ShortName: "SN", Color: "#cc99ff"},
	"elections":           {Name: "Valberedningen", ShortName: "VB", Color: "#cc99ff"},
	"init":                {Name: "init", ShortName: "init", Color: "#000000"},
	"itk":                 {Name: "ITerativa Klubben", ShortName: "ITK", Color: "#ADFF5B"},
	"jml":                 {Name: "JML-nämnden", ShortName: "JML", Color: "#cc99ff"},
	"qmisk":               {Name: "QlubbMästeriet IN-Sektionen Kista", ShortName: "QMISK", Color: "#800000"},
	"reception":           {Name: "Mottagningen", ShortName: "Mottagningen", Color: "#cc99ff"},
	"sports":              {Name: "Idrottsnämnden", ShortName: "Idrottsnämnden", Color: "#cc99ff"},
	"strängteoretiquerna": {Name: "Strängteoretiquerna", ShortName: "Strängteoretiquerna", Color: "#cc99ff"},
	"study_environment":   {Name: "Studiemiljönämnden", ShortName: "SMN", Color: "#cc99ff"},
	"tmeit":               {Name: "TraditionsMEsterIT", ShortName: "TMEIT", Color: "#436B7D"},
	"other":               {Name: "Övriga förtroendevalda", ShortName: "Övriga", Color: "#cc99ff"},
}

type ParsedPosition struct {
	Name        string
	Description string
	History     []HistoryRow
}

type HistoryRow struct {
	Year   int
	Spring []string
	Fall   []string
}

func parseMD(path string) (*ParsedPosition, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pos := &ParsedPosition{}

	reTitle := regexp.MustCompile(`^##\s+(.+)$`)
	reTableRow := regexp.MustCompile(`^\|?\s*(\d{4})\s*\|\s*([^|]*)\|\s*([^|]*)\|?$`)

	inDescription := false
	inHistory := false
	var descLines []string

	reIgnore := regexp.MustCompile(`(?i)^\s*[\-\*]\s*\*\*(Access Privileges|Time Requirement):?\*\*:?`)

	for scanner.Scan() {
		line := scanner.Text()

		if pos.Name == "" {
			if matches := reTitle.FindStringSubmatch(line); len(matches) > 1 {
				pos.Name = strings.TrimSpace(matches[1])
			}
		}

		if strings.HasPrefix(line, "### Description") {
			inDescription = true
			inHistory = false
			continue
		}

		if strings.HasPrefix(line, "### History") {
			inHistory = true
			inDescription = false
			continue
		}

		if strings.HasPrefix(line, "###") {
			inDescription = false
			continue
		}

		if inDescription {
			if reIgnore.MatchString(line) {
				continue
			}
			descLines = append(descLines, line)
		}

		if inHistory {
			if matches := reTableRow.FindStringSubmatch(line); len(matches) > 3 {
				yearStr := matches[1]
				springNames := parseNames(matches[2])
				fallNames := parseNames(matches[3])

				var year int
				if _, err := fmt.Sscanf(yearStr, "%d", &year); err == nil {
					pos.History = append(pos.History, HistoryRow{
						Year:   year,
						Spring: springNames,
						Fall:   fallNames,
					})
				}
			}
		}
	}

	pos.Description = strings.TrimSpace(strings.Join(descLines, "\n"))

	return pos, nil
}

func parseNames(s string) []string {
	s = strings.ReplaceAll(s, "\u200b", "")
	parts := strings.Split(s, ",")
	var names []string
	for _, p := range parts {
		name := strings.TrimSpace(p)
		if name != "" && !strings.Contains(strings.ToLower(name), "holder of position") {
			names = append(names, name)
		}
	}
	return names
}

func main() {
	cfg := config.Load()
	ctx := context.Background()

	pool, err := db.SetupPostgresPool(ctx, cfg.Database.URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	// Clean up existing data
	_, err = pool.Exec(ctx, "TRUNCATE trustees, positions, users, groups CASCADE")
	if err != nil {
		log.Printf("failed to truncate tables: %v", err)
	}

	queries := db.New(pool)

	groupRepo := repository.NewGroupRepository(queries)
	userRepo := repository.NewUserRepository(queries)
	positionRepo := repository.NewPositionRepository(queries)
	trusteeRepo := repository.NewTrusteeRepository(queries)

	groupService := service.NewGroupService(groupRepo, trusteeRepo, positionRepo)
	userService := service.NewUserService(userRepo)
	positionService := service.NewPositionService(positionRepo, trusteeRepo)

	users := make(map[string]uuid.UUID)

	dataDir := ".data"
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if !entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		slug := entry.Name()
		log.Printf("Importing group: %s", slug)
		meta, ok := groupMetadata[slug]
		if !ok {
			meta = GroupMetadata{
				Name:      titleCase(strings.ReplaceAll(slug, "_", " ")),
				ShortName: strings.ToUpper(slug),
				Color:     "#808080",
			}
		}

		posFiles, _ := os.ReadDir(filepath.Join(dataDir, slug))
		var parsedPositions []*ParsedPosition
		earliestYear := 2026

		for _, posFile := range posFiles {
			if posFile.IsDir() || !strings.HasSuffix(posFile.Name(), ".md") || posFile.Name() == "README.md" || posFile.Name() == "template.md" {
				continue
			}

			parsed, err := parseMD(filepath.Join(dataDir, slug, posFile.Name()))
			if err != nil {
				log.Printf("Failed to parse %s: %v", posFile.Name(), err)
				continue
			}

			if parsed.Name == "" {
				parsed.Name = titleCase(strings.ReplaceAll(strings.TrimSuffix(posFile.Name(), ".md"), "_", " "))
			}

			parsedPositions = append(parsedPositions, parsed)
			for _, row := range parsed.History {
				if row.Year < earliestYear {
					earliestYear = row.Year
				}
			}
		}

		g, err := groupService.Create(ctx, db.CreateGroupParams{
			Name:        meta.Name,
			Slug:        slug,
			ShortName:   meta.ShortName,
			Description: nil,
			Color:       meta.Color,
			EstablishedAt: pgtype.Date{
				Time:  time.Date(earliestYear, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
		})
		if err != nil {
			log.Printf("Failed to create group %s: %v", slug, err)
			continue
		}

		for _, parsed := range parsedPositions {
			log.Printf("  Importing position: %s", parsed.Name)
			pEarliest := 2026
			for _, row := range parsed.History {
				if row.Year < pEarliest {
					pEarliest = row.Year
				}
			}

			var description *string
			if parsed.Description != "" {
				description = &parsed.Description
			}

			p, err := positionService.Create(ctx, db.CreatePositionParams{
				Name:        parsed.Name,
				Email:       strings.ToLower(strings.ReplaceAll(parsed.Name, " ", ".")) + "@kth.it",
				Description: description,
				GroupID:     g.ID,
				EstablishedAt: pgtype.Date{
					Time:  time.Date(pEarliest, time.January, 1, 0, 0, 0, 0, time.UTC),
					Valid: true,
				},
			})
			if err != nil {
				log.Printf("Failed to create position %s: %v", parsed.Name, err)
				continue
			}

			for _, row := range parsed.History {
				springSet := make(map[string]bool)
				for _, n := range row.Spring {
					springSet[titleCase(n)] = true
				}
				fallSet := make(map[string]bool)
				for _, n := range row.Fall {
					fallSet[titleCase(n)] = true
				}

				allNames := make(map[string]bool)
				for n := range springSet {
					allNames[n] = true
				}
				for n := range fallSet {
					allNames[n] = true
				}

				for name := range allNames {
					uID := getOrCreateUser(ctx, userService, users, name)
					if uID == uuid.Nil {
						continue
					}

					startDate := time.Date(row.Year, time.January, 1, 0, 0, 0, 0, time.UTC)
					endDate := time.Date(row.Year, time.December, 31, 0, 0, 0, 0, time.UTC)

					if springSet[name] && !fallSet[name] {
						endDate = time.Date(row.Year, time.June, 30, 0, 0, 0, 0, time.UTC)
					} else if !springSet[name] && fallSet[name] {
						startDate = time.Date(row.Year, time.July, 1, 0, 0, 0, 0, time.UTC)
					}

					_, _ = trusteeRepo.Create(ctx, p.ID, uID, startDate, endDate)
				}
			}
		}
	}
	log.Println("Seeding completed successfully!")
}

func titleCase(s string) string {
	parts := strings.Fields(s)
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + strings.ToLower(p[1:])
		}
	}
	return strings.Join(parts, " ")
}

func getOrCreateUser(ctx context.Context, s service.UserService, cache map[string]uuid.UUID, fullName string) uuid.UUID {
	fullName = titleCase(fullName)
	if id, ok := cache[fullName]; ok {
		return id
	}

	parts := strings.Split(fullName, " ")
	firstName := parts[0]
	lastName := ""
	if len(parts) > 1 {
		lastName = strings.Join(parts[1:], " ")
	} else {
		lastName = "-"
	}

	u, err := s.Create(ctx, db.CreateUserParams{
		FirstName: firstName,
		LastName:  lastName,
	})
	if err != nil {
		log.Printf("Failed to create user %s: %v", fullName, err)
		return uuid.Nil
	}

	cache[fullName] = u.ID
	return u.ID
}

package router

import (
	"net/http"

	"github.com/itsektionen/mimer/internal/app/v1/handler"
	"github.com/itsektionen/mimer/internal/app/v1/service"
)

func SetupV1Router(
	committeeService service.CommitteeService,
	personService service.PersonService,
	positionService service.PositionService,
) http.Handler {
	mux := http.NewServeMux()

	committeeHandler := handler.NewCommitteeHandler(committeeService)
	personHandler := handler.NewPersonHandler(personService)
	positionHandler := handler.NewPositionHandler(positionService)

	mux.HandleFunc("/people", personHandler.HandlePeople)
	mux.HandleFunc("/people/", personHandler.HandlePersonById)

	mux.HandleFunc("/positions", positionHandler.HandlePositions)
	mux.HandleFunc("/positions/", positionHandler.HandlePositionById)

	mux.HandleFunc("/committees", committeeHandler.HandleCommittees)
	mux.HandleFunc("/committees/", committeeHandler.HandleCommitteeById)

	mux.HandleFunc("GET /health", handler.GetHealth)
	mux.HandleFunc("/", handler.GetIndex)

	return mux
}

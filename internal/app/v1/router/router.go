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

	mux.HandleFunc("GET /people", personHandler.HandleGetAllPeople)
	mux.HandleFunc("POST /people", personHandler.HandleCreatePerson)
	mux.HandleFunc("GET /people/", personHandler.HandleGetPersonById)

	mux.HandleFunc("GET /positions", positionHandler.HandleGetAllPositions)
	mux.HandleFunc("POST /positions", positionHandler.HandleCreatePosition)
	mux.HandleFunc("GET /positions/", positionHandler.HandleGetPositionById)

	mux.HandleFunc("GET /committees", committeeHandler.HandleGetAllCommittees)
	mux.HandleFunc("POST /committees", committeeHandler.HandleCreateCommittee)
	mux.HandleFunc("GET /committees/", committeeHandler.HandleGetCommitteeById)

	mux.HandleFunc("GET /health", handler.GetHealth)
	mux.HandleFunc("GET /", handler.GetIndex)

	return mux
}

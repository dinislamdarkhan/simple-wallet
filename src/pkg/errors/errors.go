package errors

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	OK                    = &Error{200, "OK"}
	NotFound              = &Error{404, "not found"}
	BadRequest            = &Error{400, "bad request"}
	CassandraSaveError    = &Error{409, "cassandra write error"}
	CassandraReadError    = &Error{409, "cassandra read error"}
	CassandraConnectError = &Error{503, "cassandra connection error"}
	DeserializeError      = &Error{415, "deserialization error"}
)

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d %s", e.Status, e.Message)
}

func (e *Error) SetMessage(Message string) *Error {
	e.Message = Message
	return e
}

func EncodeErrorJSON(_ context.Context, err error, w http.ResponseWriter) {
	switch err {
	case OK:
		w.WriteHeader(http.StatusOK)
	case CassandraSaveError, CassandraReadError:
		w.WriteHeader(http.StatusConflict)
	case NotFound:
		w.WriteHeader(http.StatusNotFound)
	case CassandraConnectError:
		w.WriteHeader(http.StatusServiceUnavailable)
	case DeserializeError:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	case BadRequest:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(err)
}

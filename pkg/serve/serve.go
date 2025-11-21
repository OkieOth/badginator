package serve

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

// --- Data Models ---

type PutBadgePayload struct {
	StateID string `json:"stateId"`
	Caption string `json:"caption,omitempty"`
}

func writeSvgToResponse(w http.ResponseWriter, svg, handler string) {
	w.Header().Set("Content-Type", "image/svg+xml")
	if _, err := io.WriteString(w, svg); err != nil {
		slog.Error("failed writing response",
			"err", err,
			"handler", handler)
		return
	}
}

// --- Handlers ---

// GET /badge/{OBJECT_NAME}/{BADGE_NAME}
func GetBadge(w http.ResponseWriter, r *http.Request) {
	object := chi.URLParam(r, "OBJECT_NAME")
	badge := chi.URLParam(r, "BADGE_NAME")

	slog.Info("GetBadge", "object", object, "badge", badge)

	// TODO: fetch the SVG badge data
	svg := fmt.Sprintf("<svg><!-- badge for %s/%s --></svg>", object, badge)

	w.WriteHeader(http.StatusOK)
	writeSvgToResponse(w, svg, "GetBadge")
}

// PUT /badge/{OBJECT_NAME}/{BADGE_NAME}
func PutBadge(w http.ResponseWriter, r *http.Request) {
	object := chi.URLParam(r, "OBJECT_NAME")
	badge := chi.URLParam(r, "BADGE_NAME")

	slog.Info("PutBadge", "object", object, "badge", badge)

	var payload PutBadgePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// TODO: create or update badge
	created := false // set true if newly created

	if created {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// GET /badge/{OBJECT_NAME}/{REVISION}/{BADGE_NAME}
func GetBadgeRevision(w http.ResponseWriter, r *http.Request) {
	object := chi.URLParam(r, "OBJECT_NAME")
	revision := chi.URLParam(r, "REVISION")
	badge := chi.URLParam(r, "BADGE_NAME")

	slog.Info("GetBadgeRevision", "object", object, "revision", revision, "badge", badge)

	// TODO
	svg := fmt.Sprintf("<svg><!-- badge for %s/%s/%s --></svg>", object, revision, badge)

	w.WriteHeader(http.StatusOK)
	writeSvgToResponse(w, svg, "GetBadgeRevision")
}

// PUT /badge/{OBJECT_NAME}/{REVISION}/{BADGE_NAME}
func PutBadgeRevision(w http.ResponseWriter, r *http.Request) {
	object := chi.URLParam(r, "OBJECT_NAME")
	revision := chi.URLParam(r, "REVISION")
	badge := chi.URLParam(r, "BADGE_NAME")

	slog.Info("PutBadgeRevision", "object", object, "revision", revision, "badge", badge)

	var payload PutBadgePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// TODO
	created := false

	if created {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// --- Server Setup ---

func Start(port int) {
	// initialize slog logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	r := chi.NewRouter()

	r.Get("/badge/{OBJECT_NAME}/{BADGE_NAME}", GetBadge)
	r.Put("/badge/{OBJECT_NAME}/{BADGE_NAME}", PutBadge)

	r.Get("/badge/{OBJECT_NAME}/{REVISION}/{BADGE_NAME}", GetBadgeRevision)
	r.Put("/badge/{OBJECT_NAME}/{REVISION}/{BADGE_NAME}", PutBadgeRevision)

	slog.Info("Badge Service API running", "port", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

package contactsedit

import (
	"log/slog"
	"net/http"
	"templ-crud/db"
	"templ-crud/layout"

	"github.com/gorilla/schema"
	"github.com/segmentio/ksuid"
)

type Handler struct {
	Log *slog.Logger
	DB  *db.DB
}

func NewHandler(log *slog.Logger, db *db.DB) http.Handler {
	return &Handler{
		Log: log,
		DB:  db,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Get(w, r)
	case http.MethodPost:
		h.Post(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	// Read the ID from the URL.
	id := r.PathValue("id")
	model := NewModel()
	if id != "" {
		// Get the existing contact from the database and populate the form.
		contact, ok, err := h.DB.Get(r.Context(), id)
		if err != nil {
			h.Log.Error("Failed to get contact", slog.String("id", id), slog.Any("error", err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !ok {
			http.Redirect(w, r, "/contacts/edit", http.StatusSeeOther)
			return
		}
		model = ModelFromContact(contact)
	}
	h.DisplayForm(w, r, model)
}

func (h *Handler) DisplayForm(w http.ResponseWriter, r *http.Request, m Model) {
	layout.Handler(View(m)).ServeHTTP(w, r)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var model Model

	//Decode the from
	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)
	err = dec.Decode(&model, r.PostForm)
	if err != nil {
		h.Log.Warn("Failed to decode form", slog.Any("error", err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(model.Validate()) > 0 {
		h.DisplayForm(w, r, model)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		id = ksuid.New().String()
	}

	contact := db.NewContact(id, model.Name, model.Email)
	if err = h.DB.Save(r.Context(), contact); err != nil {
		h.Log.Error("Failed to save contact", slog.String("id", id), slog.Any("error", err))
		model.Error = "Failed to save the contact. Please try again."
		h.DisplayForm(w, r, model)
		return
	}

	// Redirect back to the contact list.
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

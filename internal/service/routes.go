package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zhetkerbaevan/green-api/internal/models"
)
var (
	tmpl = template.Must(template.ParseFiles("tmpl/index.html"))
)
type Handler struct {
	data map[string]interface{}
}

func NewHandler() *Handler {
	return &Handler{data: make(map[string]interface{}),}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./tmpl/static"))))
	router.HandleFunc("/", h.homeHandler).Methods("GET")
	router.HandleFunc("/get-settings", h.handleGetSettings).Methods("GET")
	router.HandleFunc("/get-state-instance", h.handleGetStateInstance).Methods("GET")
	router.HandleFunc("/send-message", h.handleSendMessage).Methods("POST")
	router.HandleFunc("/send-file-by-url", h.handleSendFileByUrl).Methods("POST")
}

func (h *Handler) homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, h.data)
}

func (h *Handler) handleGetSettings(w http.ResponseWriter, r *http.Request) {
	//Get data from HTML form
	idInstance := r.FormValue("idInstance")
	apiTokenInstance := r.FormValue("apiTokenInstance")
	apiUrl := idInstance[:4]

	//Form url 
	url := fmt.Sprintf("https://%s.api.greenapi.com/waInstance%s/getSettings/%s", apiUrl, idInstance, apiTokenInstance)

	//Make a request to get data from GREEN-API
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	//Read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, string(body), resp.StatusCode)
		return
	}

	//Get data to structure Settings
	var settings models.Settings
	if err := json.Unmarshal(body, &settings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	settingsJSON, err := json.MarshalIndent(settings, "", "  ") //Convert to JSON
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the data map with the result
	h.data["Result"] = string(settingsJSON) 
	tmpl.Execute(w, h.data)
}


func (h *Handler) handleGetStateInstance(w http.ResponseWriter, r *http.Request) {
	idInstance := r.FormValue("idInstance")
	apiTokenInstance := r.FormValue("apiTokenInstance")
	apiUrl := idInstance[:4]

	//Form url 
	url := fmt.Sprintf("https://%s.api.greenapi.com/waInstance%s/getStateInstance/%s", apiUrl, idInstance, apiTokenInstance)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	//Read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var state models.State
	if err := json.Unmarshal(body, &state); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stateJSON, err := json.MarshalIndent(state, "", "  ") //Convert to JSON
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the data map with the result
	h.data["Result"] = string(stateJSON) 
	tmpl.Execute(w, h.data)
}

func (h *Handler) handleSendMessage(w http.ResponseWriter, r *http.Request) {
	idInstance := r.FormValue("idInstance")
	apiTokenInstance := r.FormValue("apiTokenInstance")
	apiUrl := idInstance[:4]

	url := fmt.Sprintf("https://%s.api.greenapi.com/waInstance%s/sendMessage/%s", apiUrl, idInstance, apiTokenInstance)

	chatId := r.FormValue("chatId1") + "@c.us" 
	message := r.FormValue("message")

	var messagePayload models.MessagePayload
	messagePayload.ChatID = chatId
	messagePayload.Message = message

	jsonData, err := json.Marshal(messagePayload) //Convert data to JSON
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	//Read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Get data from body to structure MessageResponse
	var messageResponse models.MessageResponse
	if err := json.Unmarshal(body, &messageResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	messageResponseJSON, err := json.MarshalIndent(messageResponse, "", "  ") //Convert to JSON
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the data map with the result
	h.data["Result"] = string(messageResponseJSON) 
	tmpl.Execute(w, h.data)
}

func (h *Handler) handleSendFileByUrl(w http.ResponseWriter, r *http.Request) {

}
package router

import (
	"apigateway/config"
	"apigateway/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Router struct {
	mux.Router
}

func NewRouter() Router {
	return Router{}
}

func (r *Router) InitHandlers() {
	r.HandleFunc("/news/filter", r.GetNewsWithFilter).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/news", r.GetNewsList).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/news/{id:[0-9]+}", r.GetNewsById).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/comments", r.GetComments).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/comments", r.AddComment).Methods(http.MethodPost, http.MethodOptions)
	r.Use(
		r.CheckRequestId,
		r.RequestLogger,
		r.SetHeaders,
	)
}

// GetNewsList показывает список новостей
func (r *Router) GetNewsList(w http.ResponseWriter, req *http.Request) {
	// Меняем хост на целевой
	req.URL.Host = config.Configs.NewsHost

	// Если схема не указана, устанавливаем по умолчанию "http"
	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}

	// Создаем новый запрос на основе оригинального
	newReq, err := http.NewRequest(req.Method, req.URL.String(), req.Body)
	if err != nil {
		http.Error(w, "Failed to create new request", http.StatusInternalServerError)
		return
	}

	// Копируем заголовки из оригинального запроса в новый запрос
	for key, values := range req.Header {
		for _, value := range values {
			newReq.Header.Add(key, value)
		}
	}

	// Выполняем новый запрос
	client := &http.Client{}
	resp, err := client.Do(newReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Копируем статус-код из ответа сервиса в ответ клиенту
	w.WriteHeader(resp.StatusCode)

	// Копируем заголовки из ответа сервиса в ответ клиенту
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Копируем тело ответа из ответа сервиса в ответ клиенту
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to copy response body", http.StatusInternalServerError)
		return
	}
}

// GetNewsWithFilter показывает новости с возможностью фильтрации
func (r *Router) GetNewsWithFilter(w http.ResponseWriter, req *http.Request) {
	// Меняем хост на целевой
	req.URL.Host = config.Configs.NewsHost

	// Если схема не указана, устанавливаем по умолчанию "http"
	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}

	// Создаем новый запрос на основе оригинального
	newReq, err := http.NewRequest(req.Method, req.URL.String(), req.Body)
	if err != nil {
		http.Error(w, "Failed to create new request", http.StatusInternalServerError)
		return
	}

	// Копируем заголовки из оригинального запроса в новый запрос
	for key, values := range req.Header {
		for _, value := range values {
			newReq.Header.Add(key, value)
		}
	}

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(newReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Копируем статус-код из ответа сервиса в ответ клиенту
	w.WriteHeader(resp.StatusCode)

	// Копируем заголовки из ответа сервиса в ответ клиенту
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Копируем тело ответа из ответа сервиса в ответ клиенту
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to copy response body", http.StatusInternalServerError)
		return
	}
}

func (r *Router) GetComments(w http.ResponseWriter, req *http.Request) {
	// Меняем хост на целевой
	req.URL.Host = config.Configs.CommentsHost

	// Если схема не указана, устанавливаем по умолчанию "http"
	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}

	newRequest, err := http.NewRequest(req.Method, req.URL.String(), req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Копируем заголовки из оригинального запроса
	newRequest.Header = req.Header
	// Копируем query-параметры в новый реквест
	newRequest.URL.RawQuery = req.URL.Query().Encode()

	// Если схема не указана, устанавливаем по умолчанию "http"
	if newRequest.URL.Scheme == "" {
		newRequest.URL.Scheme = "http"
	}

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(newRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Копируем статус-код из ответа сервиса в ответ клиенту
	w.WriteHeader(resp.StatusCode)

	// Копируем заголовки из ответа сервиса в ответ клиенту
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Копируем тело ответа из ответа сервиса в ответ клиенту
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to copy response body", http.StatusInternalServerError)
		return
	}
}

func (r *Router) AddComment(w http.ResponseWriter, req *http.Request) {
	// Декодируем тело запроса в структуру Comment
	var comment entity.Comment
	if err := json.NewDecoder(req.Body).Decode(&comment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Проверяем обязательные поля
	if comment.NewsId == "" || comment.Content == "" {
		http.Error(w, "NewsId and Content are required", http.StatusBadRequest)
		return
	}

	// Устанавливаем время создания комментария
	comment.Created = time.Now()

	// Кодируем комментарий в JSON
	commentJSON, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "Failed to encode comment", http.StatusInternalServerError)
		return
	}

	ok := censor(commentJSON)
	if !ok {
		http.Error(w, "Comment rejected by censor", http.StatusBadRequest)
		return
	}

	// Создаем новый запрос на сервис комментариев
	newReq, err := http.NewRequest(http.MethodPost, "http://"+config.Configs.CommentsHost, bytes.NewBuffer(commentJSON))
	if err != nil {
		http.Error(w, "Failed to create new request", http.StatusInternalServerError)
		return
	}

	// Копируем заголовки из оригинального запроса в новый запрос
	for key, values := range req.Header {
		for _, value := range values {
			newReq.Header.Add(key, value)
		}
	}

	// Выполняем новый запрос
	client := &http.Client{}
	resp, err := client.Do(newReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Копируем заголовки из ответа сервиса в ответ клиенту
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	// Копируем тело ответа из ответа сервиса в ответ клиенту
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to copy response body", http.StatusInternalServerError)
		return
	}
}

func (r *Router) GetNewsById(w http.ResponseWriter, req *http.Request) {
	// Извлекаем id из URL
	id := mux.Vars(req)["id"]
	idNum, err := strconv.Atoi(id)
	if err != nil || idNum < 1 {
		http.Error(w, "wrong news ID", http.StatusBadRequest)
		return
	}

	// Каналы для результатов и ошибок
	newsChan := make(chan *entity.NewsFullDetailed, 1)
	commentsChan := make(chan []entity.Comment, 1)
	errChan := make(chan error, 2)

	// Запускаем горутины для параллельных запросов
	go func() {
		news, err := r.fetchNews(id)
		if err != nil {
			errChan <- err
			return
		}
		newsChan <- news
	}()

	go func() {
		comments, err := r.fetchComments(id)
		if err != nil {
			errChan <- err
			return
		}
		commentsChan <- comments
	}()

	// Ожидаем завершения всех горутин
	var news *entity.NewsFullDetailed
	var comments []entity.Comment
	for i := 0; i < 2; i++ {
		select {
		case news = <-newsChan:
		case comments = <-commentsChan:
		case err := <-errChan:
			log.Printf("Failed to fetch data: %v", err)
			http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
			return
		}
	}

	// Собираем полную структуру
	newsFull := entity.NewsFullDetailed{
		ID:      news.ID,
		Title:   news.Title,
		Content: news.Content,
		Link:    news.Link,
		PubTime: news.PubTime,
		Comment: comments,
	}

	// Сериализуем результат в JSON
	jsonData, err := json.Marshal(newsFull)
	if err != nil {
		log.Printf("Failed to serialize news: %v", err)
		http.Error(w, "Failed to serialize news", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовки
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	// Отправляем JSON в ответ
	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Failed to write response: %v", err)
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

// fetchNews запрашивает новость по id с сервиса новостей
func (r *Router) fetchNews(id string) (*entity.NewsFullDetailed, error) {
	// Формируем URL для запроса новости
	newsURL := fmt.Sprintf("http://%s/news/%s", config.Configs.NewsHost, id)

	// Создаем новый запрос
	req, err := http.NewRequest(http.MethodGet, newsURL, nil)
	if err != nil {
		return nil, err
	}
	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Декодируем ответ
	var news entity.NewsFullDetailed
	if err := json.NewDecoder(resp.Body).Decode(&news); err != nil {
		return nil, err
	}

	return &news, nil
}

// fetchComments запрашивает комментарии по id новости с сервиса комментариев
func (r *Router) fetchComments(id string) ([]entity.Comment, error) {
	// Формируем URL для запроса комментариев
	commentsURL := fmt.Sprintf("http://%s/comments?id=%s", config.Configs.CommentsHost, id)

	// Создаем новый запрос
	req, err := http.NewRequest(http.MethodGet, commentsURL, nil)
	if err != nil {
		return nil, err
	}

	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Декодируем ответ
	var comments []entity.Comment
	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		return nil, err
	}

	return comments, nil
}

func censor(commentJSON []byte) bool {
	// Создаем новый запрос на сервис цензуры
	url := "http://" + config.Configs.CommentCensorHost
	newReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(commentJSON))
	if err != nil {
		return false
	}

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(newReq)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

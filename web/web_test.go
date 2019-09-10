package web

import (
	"encoding/json"
	"github.com/slimjim777/imagebuild/config"
	"github.com/slimjim777/imagebuild/service"
	"io"
	"net/http"
	"net/http/httptest"
	"path"
)

func sendRequest(method, url string, data io.Reader, srv *Web) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, url, data)

	srv.Router().ServeHTTP(w, r)
	return w
}

func defaultsService() (*config.Settings, *service.BoardService) {
	settings := &config.Settings{
		DocRoot:    "../static",
		BoardsPath: path.Join("..", config.DefaultBoardsPath),
	}
	brdService := service.NewBoardService(settings)
	return settings, brdService
}

func parseBoardsResponse(r io.Reader) (BoardsResponse, error) {
	// Parse the response
	result := BoardsResponse{}
	err := json.NewDecoder(r).Decode(&result)
	return result, err
}

package web

import "net/http"

// BoardsList list the available boards for enablement
func (srv Web) BoardsList(w http.ResponseWriter, r *http.Request) {
	boards := srv.BoardSrv.List()
	formatBoardsResponse(boards, w)
}

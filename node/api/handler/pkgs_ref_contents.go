package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

func (hdlr _handler) pkgs_ref_contents(
	httpResp http.ResponseWriter,
	httpReq *http.Request,
) {
	vars := mux.Vars(httpReq)

	pkgRef, err := url.PathUnescape(vars["ref"])
	if nil != err {
		http.Error(httpResp, err.Error(), http.StatusInternalServerError)
		return
	}

	pkgContents, err := hdlr.core.ListPkgContents(pkgRef)
	if nil != err {
		http.Error(httpResp, err.Error(), http.StatusInternalServerError)
		return
	}

	pkgContentsBytes, err := hdlr.json.Marshal(pkgContents)
	if nil != err {
		http.Error(httpResp, err.Error(), http.StatusInternalServerError)
		return
	}

	httpResp.WriteHeader(http.StatusOK)
	httpResp.Header().Set("Content-Type", "application/json")
	httpResp.Write(pkgContentsBytes)
}

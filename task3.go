// Handler
package main

import (
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) getNextPageBankByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bankID, err := strconv.ParseInt(ps.ByName("bankId"), 10, 64)
	if err != nil {
		response.Error(http.StatusBadRequest, errors.ErrBadRequest).WriteTo(w)
		return
	}

	if bankID == 0 {
		response.Error(http.StatusBadRequest, bank.ErrIDInvalid).WriteTo(w)
		return
	}

	result, err := s.Dependencies.BankSvc.NextPage(r.Context(), bankID)
	if err != nil {
		if strings.Contains(err.Error(), "no more pages") {
			response.Error(http.StatusNotFound, bank.ErrorNoMoreNextPages)
			return
		}
		response.Error(http.StatusInternalServerError, err).WriteTo(w)
		return
	}
	response.JSON(http.StatusOK, result).WriteTo(w)
}

func (s *Server) getPreviousPageBankByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bankID, err := strconv.ParseInt(ps.ByName("bankId"), 10, 64)
	if err != nil {
		response.Error(http.StatusBadRequest, errors.ErrBadRequest).WriteTo(w)
		return
	}

	if bankID == 0 {
		response.Error(http.StatusBadRequest, bank.ErrIDInvalid).WriteTo(w)
		return
	}

	result, err := s.Dependencies.BankSvc.PreviousPage(r.Context(), bankID)
	if err != nil {
		if strings.Contains(err.Error(), "no more pages") {
			response.Error(http.StatusNotFound, bank.ErrorNoMorePreviosPages)
			return
		}
		response.Error(http.StatusInternalServerError, err).WriteTo(w)
		return
	}

	response.JSON(http.StatusOK, result).WriteTo(w)
}

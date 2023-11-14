func (s *Server) NewBankAccountHandler(r *routing.Router) {
	groupBO.GET("/page/:page", s.getNextPageBankAccountByID, reqOnlyUser, reqBankView)
}

func (s *Server) getNextPageBankAccountByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := strconv.ParseInt(ps.ByName("page"), 10, 64)
	if err != nil {
		response.Error(http.StatusBadRequest, errors.ErrBadRequest).WriteTo(w)
		return
	}

	if page <= 0 {
		response.Error(http.StatusBadRequest, bankacc.ErrorPageInvalid).WriteTo(w)
		return
	}

	result, err := s.Dependencies.BankAccountSvc.NextPrevPage(r.Context(), page)
	if err != nil {
		response.Error(http.StatusInternalServerError, err).WriteTo(w)
		return
	}

	response.JSON(http.StatusOK, result).WriteTo(w)
}

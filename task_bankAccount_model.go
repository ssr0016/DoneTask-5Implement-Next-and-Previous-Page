

var (
	ErrorPageInvalid = errors.New("page-invalid", "Page is invalid")
)

type Service struct {
	NextPrevPage(ctx context.Context, next int64) (*BankAccountDTO, error)
}
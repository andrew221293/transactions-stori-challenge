package usecase

type (
	StoriUseCase struct {
		Store StoriStore
	}
)

//StoriStore Implement all database methods
type StoriStore interface {
}

package simple

type fooRepository struct {
}

func NewFooRepository() *fooRepository {
	return &fooRepository{}
}

type FooService struct {
	fooRepository *fooRepository
}

func NewFooService(fooRepository *fooRepository) *FooService{
	return &FooService{
		fooRepository: fooRepository,
	}
}
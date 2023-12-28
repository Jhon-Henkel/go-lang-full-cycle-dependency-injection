package product

type ProductUseCase struct {
	repository *ProductRepository
}

func NewProductUseCase(repository *ProductRepository) *ProductUseCase {
	return &ProductUseCase{repository}
}

// GetProduct returns a Product by id.
// This Product was not supported to be returned. We should a DTO instead.
// However, we will return it for now to keep the example simple.
func (uc *ProductUseCase) GetProduct(id int) (Product, error) {
	return uc.repository.GetProduct(id)
}

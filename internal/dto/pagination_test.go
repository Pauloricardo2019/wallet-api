package dto_test

import (
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"testing"
	"wallet-api/internal/dto"
)

func TestPagination_ConvertToVO(t *testing.T) {
	fakerGen := faker.New()

	paginationDto := dto.Pagination{
		Limit: fakerGen.Int(),
		Page:  fakerGen.Int(),
		Sort:  "ascendente",
	}

	paginationModel := paginationDto.ConvertToVO()

	assert.Equal(t, paginationDto.Limit, paginationModel.Limit)
	assert.Equal(t, paginationDto.Page, paginationModel.Page)
	assert.Equal(t, paginationDto.Sort, paginationModel.Sort)
}

package utils

import (
	dto "app/src/shared/dto"
)

func GeneratePaginationResult[T comparable](total int, data []T, options *dto.PageOptionsDto) (*dto.PageDto[T]) {
	pageMetaDto := dto.NewPageMetaDto(options, total)
	return dto.NewPageDto[T](data, pageMetaDto)
}

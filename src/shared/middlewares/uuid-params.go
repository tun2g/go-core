package middlewares

import (
	"errors"
	"app/src/shared/exception"
	httpContext "app/src/shared/http-context"
	"sync"

	"github.com/google/uuid"
)

func UUIDParamsMiddleware(fields ...string) func(ctx *httpContext.CustomContext) {
	if len(fields) == 0 {
		fields = []string{"id"}
	}

	return func(ctx *httpContext.CustomContext) {
		var wg sync.WaitGroup
		errCh := make(chan error, 1) // Buffer of 1 to ensure we can send an error without blocking

		for _, field := range fields {
			wg.Add(1)
			go func(field string) {
				defer wg.Done()
				id := ctx.Param(field)
				if _, err := uuid.Parse(id); err != nil {
					select {
					case errCh <- errors.New("Invalid UUID Params for field: " + field):
					default:
					}
				}
			}(field)
		}

		wg.Wait()
		close(errCh)

		if err := <-errCh; err != nil {
			exception.ThrowUnprocessableEntityException(ctx, err)
			return
		}

		ctx.Next()
	}
}
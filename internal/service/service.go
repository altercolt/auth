package service

func toPtr[T any](obj T) *T {
	if obj != nil {
		return &obj
	}

	return nil
}

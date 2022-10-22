package web

type Middleware func(Handler) Handler

func wrapMiddleware(handler Handler, middlewares []Middleware) Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		if handler != nil {
			handler = middlewares[i](handler)
		}
	}

	return handler
}

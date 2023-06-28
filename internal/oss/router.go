package oss

import (
	"github.com/go-chi/chi"
	"golang.org/x/exp/slog"
)

func NewRouter(logger *slog.Logger, container *features.ServiceContainer) *chi.Mux {
	router := chi.NewRouter()
	// router.Use(middleware.Logger)
	// router.Use(middleware.Recoverer)
	// router.Use(middleware.RequestID)
	// router.Use(shared.CorsPolicy)
	// router.Use(shared.JsonContentType)
	// router.Use(middleware.AllowContentType("application/json"))

	// router = usersApi.MakeUserRoutes(logger, router, container.UsersService)
	// router = profilesApi.MakeProfileRoutes(logger, router, container.ProfilesService)
	// router = articlesApi.MakeArticlesRoutes(logger, router, container.ArticlesService)
	// router = commentsApi.MakeCommentsRoutes(logger, router, container.CommentsService)
	router.Mount("/api", router)

	return router
}

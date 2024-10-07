package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)              // Создание списка
			lists.GET("/", h.getAllLists)              // Получение всех списков
			lists.GET("/:id", h.getListById)           // Получение списка по ID
			lists.PUT("/:id", h.updateList)            // Обновление списка по ID
			lists.DELETE("/:id", h.deleteList)         // Удаление списка по ID
		}

		items := api.Group("/items")
		{
			items.POST("/:id", h.createItem)           // Создание элемента в списке
			items.GET("/", h.getAllItems)              // Получение всех элементов
			items.GET("/:item_id", h.getItemById)      // Получение элемента по ID
			items.PUT("/:item_id", h.updateItem)       // Обновление элемента по ID
			items.DELETE("/:item_id", h.deleteItem)    // Удаление элемента по ID
		}

		filters := api.Group("/filters")
		{
			filters.POST("/", h.postFilters)           // Создание фильтра
			filters.GET("/", h.getFilters)             // Получение всех фильтров
			filters.PUT("/:id", h.updateFilter)        // Обновление фильтра по ID
			filters.DELETE("/:id", h.deleteFilter)     // Удаление фильтра по ID
		}

		favorites := api.Group("/favorites")
		{
			favorites.POST("/", h.postFavorite)        // Добавление в избранное
			favorites.GET("/", h.getFavorites)         // Получение избранного
			favorites.DELETE("/:id", h.deleteFavorite) // Удаление из избранного по ID
		}

		stats := api.Group("/stats")
		{
			stats.GET("/regions", h.getStatsByRegion)  // Получение статистики по регионам
			stats.GET("/cities", h.getStatsByCity)     // Получение статистики по городам
			stats.GET("/:property_id", h.getStatsByProperty) // Получение статистики по конкретному объекту
			stats.GET("/districts", h.getStatsByDistrict)  // Получение статистики по районам
		}
	}

	return router
}

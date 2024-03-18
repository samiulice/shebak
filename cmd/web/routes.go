package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samiulru/shebak/internal/config"
	"github.com/samiulru/shebak/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/user/login", handlers.Repo.AdminLogin)
	mux.Post("/user/login", handlers.Repo.PostAdminLogin)
	mux.Get("/user/logout", handlers.Repo.AdminLogout)

	//publicFileServer serves files for public site
	publicFileServer := http.FileServer(http.Dir("./static/public"))
	mux.Handle("/static/public/*", http.StripPrefix("/static/public", publicFileServer))

	/*
		// AccessLevel := 1 means SuperAdmin
		// AccessLevel := 2 means Admin
		// AccessLevel := 3 means Employee
		// AccessLevel := 4 means User

		-------------------------------------------------
		-------------------------------------------------
		   Secure Routes: Priority Level || AccessLevel
		-------------------------------------------------
		-------------------------------------------------

		|----private
			|-super-admin : 1
			|--admin : 2
			|---employee : 3
			|----user : 4

	*/

	// /private is the secure root
	mux.Route("/private", func(mux chi.Router) {
		//secure routes that available only for Admin
		mux.Route("/admin", func(mux chi.Router) {
			mux.Use(AuthAdmin)

			//adminFileServer serves files for admin panel
			adminFileServer := http.FileServer(http.Dir("./static/private"))
			mux.Handle("/static/private/*", http.StripPrefix("/private/admin/static/private", adminFileServer))
			//admin secure routes
			mux.Get("/dashboard", handlers.Repo.AdminDashboard)
			mux.Get("/service/{mainID}/main", handlers.Repo.AdminSubItemList)        //--TODO
			mux.Get("/service/{mainID}/{subID}/sub", handlers.Repo.AdminSubItemList) //--TODO

			mux.Get("/category/add-new", handlers.Repo.AdminAddNewCategory)
			mux.Post("/category/add-new", handlers.Repo.PostAdminAddNewCategory)
			mux.Get("/category/{mainID}/update", handlers.Repo.AdminUpdateCategory)
			mux.Post("/category/{mainID}/update", handlers.Repo.PostAdminUpdateCategory)

			mux.Get("/sub-category/add-new", handlers.Repo.AdminAddNewSubCategory)
			mux.Post("/sub-category/add-new", handlers.Repo.PostAdminAddNewSubCategory)
			mux.Get("/sub-category/{mainID}/{subID}/update", handlers.Repo.AdminUpdateSubCategory)
			mux.Post("/sub-category/{mainID}/{subID}/update", handlers.Repo.PostAdminUpdateSubCategory)

			mux.Get("/service/add-new", handlers.Repo.AdminAddNewService)
			mux.Post("/service/add-new", handlers.Repo.PostAdminAddNewService)

			//Depper level to access
			//secure routes that available only for Super Admin
			mux.Route("/super", func(mux chi.Router) {
				mux.Use(AuthSuperAdmin)
				//privateFileServer serves files for super admin panel
				superAdminFileServer := http.FileServer(http.Dir("./static/private"))
				mux.Handle("/static/private/*", http.StripPrefix("/private/admin/super-admin/static/private", superAdminFileServer))

				mux.Get("/dashboard", handlers.Repo.SuperAdminDashboard)
				mux.Get("/service/{mainID}/main", handlers.Repo.AdminSubItemList)        //--TODO
				mux.Get("/service/{mainID}/{subID}/sub", handlers.Repo.AdminSubItemList) //--TODO

				mux.Get("/category/add-new", handlers.Repo.AdminAddNewCategory)
				mux.Post("/category/add-new", handlers.Repo.PostAdminAddNewCategory)
				// mux.Get("/category/update", handlers.Repo.AdminUpdateNewCategory)
				// mux.Post("/category/update", handlers.Repo.PostAdminUpdateNewCategory)

				mux.Get("/sub-category/add-new", handlers.Repo.AdminAddNewSubCategory)
				mux.Post("/sub-category/add-new", handlers.Repo.PostAdminAddNewSubCategory)

				mux.Get("/service/add-new", handlers.Repo.AdminAddNewService)
				mux.Post("/service/add-new", handlers.Repo.PostAdminAddNewService)
				// mux.Get("/service/{mainID}/{subID}/{serID}/update", handlers.Repo.AdminEditServiceDetails)
				// mux.Get("/sub/{mainID}/{subID}/update", handlers.Repo.PostAdminEditServiceCategoryMainDetails)
			})
		})
	})

	//secure routes that available only for Employee or Technician
	mux.Route("/employee", func(mux chi.Router) {
		mux.Use(AuthAdmin)

		//adminFileServer serves files for admin panel
		adminFileServer := http.FileServer(http.Dir("./static/private"))
		mux.Handle("/static/private/*", http.StripPrefix("/private/employee/static/private", adminFileServer))

		//admin secure routes
		// mux.Get("/dashboard", handlers.Repo.EmployeeDashboard)
	})

	//secure routes that available only for User or Client
	mux.Route("/user", func(mux chi.Router) {
		mux.Use(AuthAdmin)

		//adminFileServer serves files for admin panel
		adminFileServer := http.FileServer(http.Dir("./static/private"))
		mux.Handle("/static/private/*", http.StripPrefix("/private/user/static/private", adminFileServer))

		//admin secure routes
		// mux.Get("/dashboard", handlers.Repo.UserDashboard)
	})

	return mux
}

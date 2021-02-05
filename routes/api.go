package routes

import (
	"confetti-framework/app/http/controllers"
	"confetti-framework/app/http/middleware"
	"github.com/confetti-framework/foundation/http/routing"
)

/*
|---------------------------------------------------------------------------
| API routes
|---------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. By default
| this is loaded in a group. The group is assigned to the "Api" middleware
| and is placed with "/api" prefix. Feel free to remove the prefix if you
| are using a subdomain for your API (which is recommended).
| Enjoy building your API!
|
*/

var Api = routing.Group(
	routing.Get("/ping", controllers.Ping),
).Prefix("/api").Middleware(middleware.Api...)

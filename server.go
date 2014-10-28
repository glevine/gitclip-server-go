package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
)

func main() {
    server := martini.Classic()
    server.Use(render.Renderer())

    server.Get("/", func() string {
        return "Hello world!"
    })

    server.Group("/repos", func(router martini.Router) {
        router.Get("/:owner/:repo", func(params martini.Params, response render.Render) {
            repo := params["owner"] + "/" + params["repo"]

            payload := map[string]interface{}{
                "emberjs/ember.js": []interface{}{
                    "http://ember101.com/",
                },
                "go-martini/martini": []interface{}{
                    "http://0value.com/build-a-restful-API-with-Martini",
                    "http://blog.codegangsta.io/blog/2014/05/19/my-thoughts-on-martini/",
                },
            }

            response.JSON(200, payload[repo])
        })
    })

    server.Run()
}
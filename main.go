/*
 * todos
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
	"github.com/GIT_USER_ID/GIT_REPO_ID/infrastructure"
)

func main() {
	log.Printf("Server started")

	infrastructure.DbConnect()
	infrastructure.Db.AutoMigrate(&openapi.Task{})

	DefaultApiService := openapi.NewDefaultApiService()
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}

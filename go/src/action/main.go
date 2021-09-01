package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"oras.land/oras-go/pkg/remotes"
)

type Scope struct {
	Login       string `json:"scopeLogin"`
	Name        string `json:"scopeName"`
	Token       string `json:"scopeToken"`
	Permissions string `json:"scopePermissions"`
}

type ScopeMap struct {
	Scopes []Scope `json:"scopes"`
}

func main() {
	ctx := context.Background()

	ref := "jlteleport.azurecr.io/ubuntu:latest"

	file, err := os.Open("./secrets/auth.json")
	if err != nil {
		panic(err.Error())
	}

	scopeMap := &ScopeMap{}
	json.NewDecoder(file).Decode(scopeMap)

	scope := scopeMap.Scopes[0]

	reg := remotes.NewRegistryWithBasicAuthorization(
		ctx,
		ref,
		scope.Name,
		scope.Token,
		scope.Permissions)

	if reg == nil {
		panic("reg is nil")
	}

	resolver, err := reg.DiscoverFetch(ctx, ref)
	if err != nil {
		panic(err.Error())
	}

	ref, desc, err := resolver.Resolve(ctx, ref)
	if err != nil {
		panic(err.Error())
	}

	fetcher, err := resolver.Fetcher(ctx, ref)
	if err != nil {
		panic(err.Error())
	}

	content, err := fetcher.Fetch(ctx, desc)
	if err != nil {
		panic(err.Error())
	}

	defer content.Close()

	val, err := io.ReadAll(content)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(desc)
	fmt.Println(ref)
	fmt.Println(string(val))
}

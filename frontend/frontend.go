package frontend

import "embed"

//go:embed all:out/*
var BuildProd embed.FS

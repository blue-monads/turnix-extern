package output

import "embed"

//go:embed all:out/*
var BuildProd embed.FS

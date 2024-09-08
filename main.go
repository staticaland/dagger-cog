package main

import (
	"context"
)

type CogDagger struct{}

// RunCog installs Cog and runs it on the specified files
func (m *CogDagger) RunCog(ctx context.Context, source *Directory, files []string) (*Directory, error) {
	return m.installCog().
		WithDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(append([]string{"cog", "-r"}, files...)).
		Directory("/src", ctx)
}

// installCog creates a container with Cog installed
func (m *CogDagger) installCog() *Container {
	return dag.Container().
		From("python:3.9-slim").
		WithExec([]string{"pip", "install", "cogapp"})
}

// RunCogOnFile is a convenience method to run Cog on a single file
func (m *CogDagger) RunCogOnFile(ctx context.Context, source *Directory, file string) (*File, error) {
	dir, err := m.RunCog(ctx, source, []string{file})
	if err != nil {
		return nil, err
	}
	return dir.File(file)
}

// VerifyCog runs Cog in check mode to verify if files are up-to-date
func (m *CogDagger) VerifyCog(ctx context.Context, source *Directory, files []string) (bool, error) {
	_, err := m.installCog().
		WithDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(append([]string{"cog", "-r", "-c"}, files...)).
		Stdout(ctx)
	
	return err == nil, nil
}
// Cog CLI - generate content with inlined Python code.

package main

import (
	"context"
	"dagger/dagger-cog/internal/dagger"
	"fmt"
)

type DaggerCog struct{}

// Usage: dagger call cog --dir="." --file="README.md" --replace export --path="."

// Cog runs the cog command on the provided file within a Python container.
// It installs the cogapp package, copies the file into the container, and executes cog -r on the file.
//
// Parameters:
//
//	ctx  - The context for the operation.
//	file - The file to be processed by cog.
//
// Returns:
//
//	*dagger.File - The processed file.
func (m *DaggerCog) Cog(ctx context.Context,
	// Checksum the output to protect it against accidental change.
	// +optional
	checksum bool,
	// Delete the generator code from the output file.
	// +optional
	delete bool,
	// Warn if a file has no cog code in it.
	// +optional
	warnEmpty bool,
	// Add PATH to the list of directories for data files and modules.
	// +optional
	includePaths []string,
	// Use ENCODING when reading and writing files.
	// +optional
	encoding string,
	// Write the output to OUTNAME.
	// +optional
	outputName string,
	// Prepend the generator source with PROLOGUE.
	// +optional
	prologue string,
	// Use print() instead of cog.outl() for code output.
	// +optional
	usePrint bool,
	// Replace the input file with the output.
	// +optional
	replace bool,
	// Suffix all generated output lines with STRING.
	// +optional
	suffix string,
	// Write the output with Unix newlines (only LF line-endings).
	// +optional
	unixNewlines bool,
	// Use CMD if the output file needs to be made writable.
	// +optional
	writeCmd string,
	// Excise all the generated output without running the generators.
	// +optional
	excise bool,
	// The end-output marker can be omitted, and is assumed at eof.
	// +optional
	assumeOutput bool,
	// The patterns surrounding cog inline instructions.
	// +optional
	markers string,
	// Control the amount of output.
	// +optional
	verbosity int,
	// The input file to be processed by cog.
	file string,
	dir *dagger.Directory) *dagger.Directory {

	// Initialize the base command arguments
	execArgs := []string{
		"cog",
	}

	// Add options to execArgs
	if checksum {
		execArgs = append(execArgs, "-c")
	}
	if delete {
		execArgs = append(execArgs, "-d")
	}
	if warnEmpty {
		execArgs = append(execArgs, "-e")
	}
	for _, path := range includePaths {
		execArgs = append(execArgs, "-I", path)
	}
	if encoding != "" {
		execArgs = append(execArgs, "-n", encoding)
	}
	if outputName != "" {
		execArgs = append(execArgs, "-o", outputName)
	}
	if prologue != "" {
		execArgs = append(execArgs, "-p", prologue)
	}
	if usePrint {
		execArgs = append(execArgs, "-P")
	}
	if replace {
		execArgs = append(execArgs, "-r")
	}
	if suffix != "" {
		execArgs = append(execArgs, "-s", suffix)
	}
	if unixNewlines {
		execArgs = append(execArgs, "-U")
	}
	if writeCmd != "" {
		execArgs = append(execArgs, "-w", writeCmd)
	}
	if excise {
		execArgs = append(execArgs, "-x")
	}
	if assumeOutput {
		execArgs = append(execArgs, "-z")
	}
	if markers != "" {
		execArgs = append(execArgs, "--markers", markers)
	}
	if verbosity != 0 {
		execArgs = append(execArgs, "--verbosity", fmt.Sprintf("%d", verbosity))
	}

	execArgs = append(execArgs, file)

	return dag.Container().
		From("python:slim").
		WithExec([]string{"pip", "install", "cogapp"}).
		WithDirectory("/src", dir).
		WithWorkdir("/src").
		WithExec(execArgs).
		Directory("/src")
}

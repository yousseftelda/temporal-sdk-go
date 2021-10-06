package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing command")
	}

	_, thisFile, _, _ := runtime.Caller(0)
	thisDir := filepath.Dir(thisFile)

	switch os.Args[1] {
	case "build":
		if err := buildLibrary(thisDir, os.Args[2:]); err != nil {
			log.Fatalf("Failed building library: %v", err)
		}
	case "genprotos":
		if err := genProtos(thisDir); err != nil {
			log.Fatalf("Failed generating protos: %v", err)
		}
	default:
		log.Fatalf("Unknown command %v", os.Args[1])
	}
}

func buildLibrary(thisDir string, args []string) error {
	// TODO(cretz): This is a temporary tool that just runs cargo build and moves
	// libs. Find better ways to build and put things into their proper place
	// after cargo build.
	// TODO(cretz): Cross-compilation in Rust is hard w/ things like MSVC

	// Run cargo
	cmd := exec.Command("cargo", append([]string{"build"}, args...)...)
	cmd.Dir = thisDir
	cmd.Stdin = os.Stdin
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cargo build failed: %w", err)
	}

	// Now copy over the lib
	var outDir string
	switch osArch := runtime.GOOS + "/" + runtime.GOARCH; osArch {
	case "windows/amd64":
		outDir = "windows-x86_64"
	case "linux/amd64":
		outDir = "linux-x86_64"
	default:
		// TODO(cretz): more triples
		return fmt.Errorf("unrecognized OS/arch of %v", osArch)
	}
	// Do primitive read-all-write-all copy
	buildType := "debug"
	for _, arg := range os.Args[1:] {
		if arg == "--release" {
			buildType = "release"
			break
		}
	}
	inFile := filepath.Join(cmd.Dir, "target", buildType, "libtemporal_sdk_core_bridge.a")
	outFile := filepath.Join(cmd.Dir, "lib", outDir, "libtemporal_sdk_core_bridge.a")
	log.Printf("Copying %v to %v", inFile, outFile)
	if b, err := os.ReadFile(inFile); err != nil {
		return fmt.Errorf("failed reading %v: %w", inFile, err)
	} else if err := os.WriteFile(outFile, b, 0644); err != nil {
		return fmt.Errorf("failed writing %v: %w", outFile, err)
	}
	return nil
}

func genProtos(thisDir string) error {
	localProtoDir := filepath.Join(thisDir, "../sdk-core/protos/local")

	// Key is proto, value is dir/package
	protos := map[string]string{
		"activity_result.proto":     "corepb/activityresultpb",
		"activity_task.proto":       "corepb/activitytaskpb",
		"child_workflow.proto":      "corepb/childworkflowpb",
		"common.proto":              "corepb/commonpb",
		"core_interface.proto":      "corepb",
		"workflow_activation.proto": "corepb/workflowactivationpb",
		"workflow_commands.proto":   "corepb/workflowcommandspb",
		"workflow_completion.proto": "corepb/workflowcompletionpb",
	}
	// Create go package args
	goPackageArgs := make([]string, 0, len(protos))
	for file, dir := range protos {
		goPackageArgs = append(goPackageArgs, "--go_opt=M"+file+"="+"go.temporal.io/sdk/sdkcore/internal/ffi/"+dir)
	}

	// Generate protos
	for file, dir := range protos {
		// Create dir
		outDir := filepath.Join(thisDir, "..", dir)
		if err := os.MkdirAll(outDir, 0755); err != nil {
			return fmt.Errorf("failed creating dir %v: %w", outDir, err)
		}
		// Protoc args
		args := []string{
			"-I=" + localProtoDir,
			"-I=" + filepath.Join(thisDir, "../sdk-core/protos/api_upstream"),
			"--go_opt=paths=source_relative",
			"--go_out=" + outDir,
		}
		args = append(args, goPackageArgs...)
		args = append(args, filepath.Join(localProtoDir, file))
		log.Printf("Running protoc %v", strings.Join(args, " "))
		cmd := exec.Command("protoc", args...)
		cmd.Dir = thisDir
		cmd.Stdin = os.Stdin
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("protoc failed: %w", err)
		}
	}
	return nil
}

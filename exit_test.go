package errxit_test

import (
	"errors"
	"flag"
	"os"
	"os/exec"
	"testing"

	"git.sr.ht/~jamesponddotco/errxit-go"
)

func TestExit_NilError(t *testing.T) {
	t.Parallel()

	if os.Getenv("TEST_EXIT_NILERROR") == "1" {
		errxit.Exit(nil)

		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestExit_NilError") //nolint:gosec // this is not a security issue
	cmd.Env = append(os.Environ(), "TEST_EXIT_NILERROR=1")

	var exitErr *exec.ExitError
	if err := cmd.Run(); err != nil {
		if errors.As(err, &exitErr) && !exitErr.Success() {
			return
		}

		t.Fatalf("process ran with err = %v, want exit status 0", err)
	}
}

func TestExit_ErrHelp(t *testing.T) {
	t.Parallel()

	if os.Getenv("TEST_EXIT_ERRHELP") == "1" {
		errxit.Exit(flag.ErrHelp)

		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestExit_ErrHelp") //nolint:gosec // this is not a security issue
	cmd.Env = append(os.Environ(), "TEST_EXIT_ERRHELP=1")

	var exitErr *exec.ExitError
	if err := cmd.Run(); err != nil {
		if errors.As(err, &exitErr) && !exitErr.Success() {
			return
		}

		t.Fatalf("process ran with err = %v, want exit status 0", err)
	}
}

func TestExit_ErrxitError(t *testing.T) {
	t.Parallel()

	if os.Getenv("TEST_EXIT_ERRXITERROR") == "1" {
		testErr := errxit.New(errGeneric, 1)

		errxit.Exit(testErr)

		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestExit_ErrxitError") //nolint:gosec // this is not a security issue
	cmd.Env = append(os.Environ(), "TEST_EXIT_ERRXITERROR=1")

	var exitErr *exec.ExitError
	if err := cmd.Run(); err != nil {
		if errors.As(err, &exitErr) && !exitErr.Success() {
			return
		}

		t.Fatalf("process ran with err = %v, want exit status 0", err)
	}

	if exitErr.ExitCode() != 1 {
		t.Fatalf("process exited with code %d, want 1", exitErr.ExitCode())
	}
}

func TestExit_GenericError(t *testing.T) {
	t.Parallel()

	if os.Getenv("TEST_EXIT_GENERICERROR") == "1" {
		errxit.Exit(errGeneric)

		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestExit_GenericError") //nolint:gosec // this is not a security issue
	cmd.Env = append(os.Environ(), "TEST_EXIT_GENERICERROR=1")

	var exitErr *exec.ExitError
	if err := cmd.Run(); err != nil {
		if errors.As(err, &exitErr) && !exitErr.Success() {
			return
		}

		t.Fatalf("process ran with err = %v, want exit status 0", err)
	}

	if exitErr.ExitCode() != 1 {
		t.Fatalf("process exited with code %d, want 1", exitErr.ExitCode())
	}
}

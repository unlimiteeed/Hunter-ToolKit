package installation

import (
	"os"
	"os/exec"
)

func Clone(url, directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.MkdirAll(directory, 0755)
		if err != nil {
			return err
		}
	}

	err := os.Chdir(directory)
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "clone", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	err = os.Chdir("..")
	if err != nil {
		return err
	}

	return nil
}

func GoInstall(url string) error {
	cmd := exec.Command("go", "install", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

package serverworkspace

import (
	"os"
	"path"
	"strings"

	"github.com/arpanrec/netcli/internal/logger"
	"github.com/arpanrec/netcli/internal/utils"
)

func createVenv() {

	pythonVersion := whichPython()
	logger.Debug("Python version: ", pythonVersion)

	_, venvDirStatErr := os.Stat(venvDir)
	if venvDirStatErr != nil {
		if os.IsNotExist(venvDirStatErr) {
			logger.Info("Creating venv directory: ", venvDir)
			createEnvCmd := "python" + pythonVersion + " -m venv " + venvDir
			createEnv, createEnvErr := utils.BashExec(&createEnvCmd)
			if createEnvErr != nil {
				logger.Fatal("Failed to create venv directory: ", createEnv, createEnvErr)
			}
			logger.Debug("Venv directory created: ", createEnv)
		} else {
			logger.Fatal("Failed to get venv directory stat: ", venvDirStatErr)
		}
	}
	setEnvVars()
	installPipPackages()
	setPythonPath()
}

func whichPython() string {
	allVersions := []string{"3.13", "3.12", "3.11", "3.10", "3.9", "3.8", "3.7", "3.6"}
	for _, version := range allVersions {
		cmd := "python" + version + " --version"
		out, err := utils.BashExec(&cmd)
		if err != nil {
			if strings.Contains(out, "python"+version+": command not found") {
				logger.Debug("Python version: ", version, " not found", out, err)
				continue
			}
			logger.Fatal("Unable to get python version: ", version, "Out: ", out, "Error: ", err)
		} else {
			return version
		}
	}
	logger.Fatal("No python version found")
	return ""
}

func whereIsActivate() string {
	var binPath string
	possibleBinPaths := []string{path.Join(venvDir, "local/bin"), path.Join(venvDir, "bin")}
	for _, bP := range possibleBinPaths {
		if _, err := os.Stat(bP); err == nil {
			logger.Debug("Bin path found: ", bP)
			binPath = bP
			break
		}
	}
	if binPath == "" {
		logger.Fatal("Bin path not found")
	}
	logger.Info("Bin path: ", binPath)
	cmd := "find " + binPath + " -name activate"
	out, err := utils.BashExec(&cmd)
	if err != nil {
		logger.Fatal("Failed to find activate script: ", out, err)
	}
	outLines := strings.Split(out, "\n")
	activateFile := outLines[0]
	logger.Info("Activate script found at: ", activateFile)
	return activateFile
}

func setEnvVars() {
	cmd := "source " + whereIsActivate() + ">/dev/null && env"
	out, err := utils.BashExec(&cmd)
	if err != nil {
		logger.Fatal("Failed to get env vars: ", out, err)
	}
	envVars := strings.Split(out, "\n")
	for _, envVar := range envVars {
		venvEnvVars = append(venvEnvVars, envVar)
		logger.Debug("Env var: ", envVar)
	}
}

func installPipPackages() {
	cmd := `which python && \
pip3 install --upgrade pip && \
pip3 install setuptools-rust wheel setuptools --upgrade && \
pip3 install ansible cryptography requests hvac --upgrade && \
ansible-galaxy collection install git+https://github.com/arpanrec/arpanrec.nebula.git -f && \
ansible-galaxy collection install git+https://github.com/ansible-collections/community.general.git && \
ansible-galaxy collection install git+https://github.com/ansible-collections/community.crypto.git && \
ansible-galaxy collection install git+https://github.com/ansible-collections/amazon.aws.git && \
ansible-galaxy collection install git+https://github.com/ansible-collections/community.docker.git && \
ansible-galaxy collection install git+https://github.com/ansible-collections/ansible.posix.git && \
ansible-galaxy collection install git+https://github.com/kewlfft/ansible-aur.git && \
ansible-galaxy role install git+https://github.com/geerlingguy/ansible-role-docker.git
`
	out, err := utils.BashExecEnv(&cmd, &venvEnvVars)
	if err != nil {
		logger.Fatal("Failed to install pip packages: ", out, err)
	}
	logger.Info("Pip packages installed: ", out)
}

func setPythonPath() {
	cmd := "which python"
	out, err := utils.BashExecEnv(&cmd, &venvEnvVars)
	if err != nil {
		logger.Fatal("Failed to get python path: ", out, err)
	}
	basePythonPath = strings.TrimSuffix(out, "\n")

}

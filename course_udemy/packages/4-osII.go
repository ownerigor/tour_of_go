package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// 1 - Retornar a pasta atual
func getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório: ", err)
		return ""
	}

	return dir
}

// 2 - Listar arquivos e pastas
func listFilesAndDirectories() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Erro ao listar arquivos: ", err)
		return
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// 3 - Versão do OS
func getOSVersion() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "ver")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("uname", "-a")
	} else if runtime.GOOS == "darwin" {
		cmd = exec.Command("sw_vers")
	} else {
		fmt.Println("SO não suportado!")
		return
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao executar comando: ", err)
		return
	}
	fmt.Println(string(out))
}

// 4 - Configuração da máquina
func getSystemInfo() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "systeminfo")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("neofetch")
	} else {
		fmt.Println("SO não suportado!")
		return
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao executar comando: ", err)
		return
	}
	fmt.Println(string(out))
}

// 5 - Desligar o computador em uma hora
func shutdownInOneHour() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("shutdown", "/s", "/t", "3600")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("shutdown", "-h", "+60")
	} else {
		fmt.Println("SO não suportado!")
		return
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao agendar o desligamento!")
		return
	}
}

// 6 - Cancelar o desligamento
func cancelShutDown() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("shutdown", "/a")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("shutdown", "-c")
	} else {
		fmt.Println("SO não suportado!")
		return
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao cancelar o desligamento!")
		return
	}
}

func main() {
	fmt.Println("Pasta atual: ", getCurrentDirectory())
	fmt.Println("Arquivos e pastas:")
	listFilesAndDirectories()
	fmt.Println("Versão do sistema operacional: ")
	getOSVersion()
	fmt.Println("Configuração da máquina:")
	getSystemInfo()
	shutdownInOneHour()
	cancelShutDown()
}

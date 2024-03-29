package appScript

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

// AppStartScript 运行程序
func AppStartScript(path, logPath string, args []string) (int, error) {
	// 创建命令
	cmd := exec.Command("python", path)
	// 设置环境变量
	if runtime.GOOS == "windows" {
		cmd.Env = append(os.Environ(), "PATH="+os.Getenv("PATH")+";"+os.Getenv("PYTHONPATH"))
	}
	// 创建目录
	_ = os.MkdirAll(strings.Replace(logPath, "/out.log", "", 1), 0755)
	// 将空字节数组写入文件
	err := ioutil.WriteFile(logPath, []byte{}, 0644)
	if err != nil {
		return 0, err
	}
	// 将输出重定向到日志文件
	logFile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	defer logFile.Close()
	cmd.Stdout = logFile
	cmd.Stderr = logFile
	// 添加参数
	cmd.Args = append(cmd.Args, args...)
	fmt.Println(cmd.Args)
	// 启动命令
	err = cmd.Start()
	if err != nil {
		return 0, err
	}
	// 返回命令的进程 ID
	return cmd.Process.Pid, nil
}

// AppKillProcess 根据pid杀掉进程
func AppKillProcess(pid int) error {
	if runtime.GOOS == "windows" {
		// 构建 taskkill 命令
		cmd := exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid))
		// 执行命令
		err := cmd.Run()
		if err != nil {
			return err
		}
	} else {
		// 在 Linux 下，可以使用 `kill` 命令杀死进程
		cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

// AppIsProcessRunning 判断进程是否存在
func AppIsProcessRunning(pid int) bool {
	// 在 Linux 下，可以使用 `kill` 命令检查进程是否存在
	if runtime.GOOS == "linux" {
		cmd := exec.Command("kill", "-0", strconv.Itoa(pid))
		err := cmd.Run()
		if err != nil {
			// 进程不存在
			return false
		}
		// 进程存在
		return true
	} else if runtime.GOOS == "windows" {
		handle, err := syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION, false, uint32(pid))
		if err != nil {
			return false
		}
		defer syscall.CloseHandle(handle)
		return true
	}
	// 不支持的操作系统
	return false
}

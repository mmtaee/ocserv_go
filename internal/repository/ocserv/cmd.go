package ocserv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var occtlCmd = "sudo /usr/bin/occtl "

func (C *CMDRepository) ReloadService() error {
	command := occtlCmd + "reload"
	output, err := exec.Command(command).Output()
	if err != nil {
		return err
	}
	log.Println("reload result: ", output)
	return nil
}

func (C *CMDRepository) CreateOrUpdateUser(group, username, password string) error {
	command := fmt.Sprintf("/usr/bin/echo -e %s\n%s\n | sudo /usr/bin/ocpasswd", password, password)
	if group != "defaults" {
		command += fmt.Sprintf(" -g %s", group)
	}
	command += fmt.Sprintf("  -c /etc/ocserv/ocpasswd %s", username)
	output, err := exec.Command(command).Output()
	if err != nil {
		return err
	}
	log.Println("create user result: ", string(output))
	return nil
}

func (C *CMDRepository) ChangeGroup(group string, username string) error {
	command := fmt.Sprintf("sudo /usr/bin/ocpasswd")
	if group != "defaults" {
		command += fmt.Sprintf(" -g %s", group)
	}
	command += fmt.Sprintf("  -c /etc/ocserv/ocpasswd %s", username)
	output, err := exec.Command(command).Output()
	if err != nil {
		return err
	}
	log.Println("change user group result: ", string(output))
	return nil
}

func (C *CMDRepository) Lock(username string) error {
	command := fmt.Sprintf("sudo /usr/bin/ocpasswd -l -c /etc/ocserv/ocpasswd %s", username)
	output, err := exec.Command(command).Output()
	if err != nil {
		return err
	}
	log.Println("lock result: ", string(output))
	return nil
}

func (C *CMDRepository) Unlock(username string) error {
	command := fmt.Sprintf("sudo /usr/bin/ocpasswd -u -c /etc/ocserv/ocpasswd %s", username)
	output, err := exec.Command(command).Output()
	if err != nil {
		return err
	}
	log.Println("lock result: ", string(output))
	return nil
}

func (C *CMDRepository) DeleteUser(username string) error {
	command := fmt.Sprintf("sudo /usr/bin/ocpasswd  -c /etc/ocserv/ocpasswd -d %s", username)
	output, err := exec.Command(command).Output()
	if err != nil {
		return err
	}
	log.Println("delete user result: ", string(output))
	return nil
}

func (C *CMDRepository) Disconnect(username string) error {
	command := fmt.Sprintf("%s disconnect user %s", occtlCmd, username)
	output, err := exec.Command(command).Output()
	if err != nil {
		return err
	}
	log.Println("disconnect user result: ", string(output))
	return nil
}

func (C *CMDRepository) OnlineUsers(onlyUsername bool) (interface{}, error) {
	command := occtlCmd + "-j show users --output=json-pretty"
	output, err := exec.Command(command).Output()
	if err != nil {
		return nil, err
	}
	var result []map[string]interface{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		return nil, err
	}
	if onlyUsername {
		var username []string
		for _, item := range result {
			if name, ok := item["username"].(string); ok {
				username = append(username, name)
			}
		}
		return username, nil
	}
	return result, nil
}

func (C *CMDRepository) SyncUsers() ([][2]string, error) {
	var userList [][2]string

	file, err := os.Open("/etc/ocserv/ocpasswd")
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		userSplit := strings.Split(line, ":")
		if len(userSplit) >= 2 {
			username, group := userSplit[0], userSplit[1]
			userList = append(userList, [2]string{username, group})
		}
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return userList, nil
}

func (C *CMDRepository) ShowIPBans() interface{} {
	command := occtlCmd + "-j show ip bans"
	output, err := exec.Command(command).Output()
	if err != nil {
		return err
	}
	var result []map[string]interface{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		return err
	}
	return result
}

func (C *CMDRepository) ShowIPBansPoints() interface{} {
	command := occtlCmd + "-j show ip bans points"
	output, err := exec.Command(command).Output()
	if err != nil {
		return err
	}
	var result []map[string]interface{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		return err
	}
	return result
}

func (C *CMDRepository) UnBanIP(ip string) error {
	command := occtlCmd + "unban ip" + ip
	output, err := exec.Command(command).Output()
	if err != nil {
		return err
	}
	log.Println("unban ip result: ", string(output))
	return nil
}

func (C *CMDRepository) ShowStatus() (string, error) {
	command := occtlCmd + "show status"
	output, err := exec.Command(command).Output()
	if err != nil {
		return "", err
	}
	return string(output), nil

}

func (C *CMDRepository) ShowIRoutes() ([]map[string]interface{}, error) {
	command := occtlCmd + "-j show iroutes"
	output, err := exec.Command(command).Output()
	if err != nil {
		return nil, err
	}
	var result []map[string]interface{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

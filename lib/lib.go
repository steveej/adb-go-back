package lib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type App struct {
	Label string
}

type Apps map[string]App

func (apps Apps) toYaml() (string, error) {
	d, err := yaml.Marshal(apps)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	return "---\n" + string(d), err
}

func (apps Apps) DumpYaml(file string) error {
	yaml, err := apps.toYaml()
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	log.Printf("Writing resutls to %v\n", file)
	f, err := os.Create(file)
	defer f.Sync()
	defer f.Close()
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	_, err = f.WriteString(yaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return err
}

func LoadYaml(file string) (Apps, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}
	var apps Apps
	err = yaml.Unmarshal(buf, &apps)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}
	return apps, err
}

func GetUserApps() (Apps, error) {
	m := make(Apps, 0)
	cmd := exec.Command("adb", "shell", "pm list packages -3 -f | sed -e 's/package://'")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return m, err
	}

	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		re := regexp.MustCompile("(.*)=(.*)\r")
		lineRegex := re.FindStringSubmatch(line)
		if len(lineRegex) == 3 {
			apk := lineRegex[1]
			pkg := lineRegex[2]

			cmd = exec.Command("adb", "shell", fmt.Sprintf("aapt d badging %v", apk))
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				return m, err
			}

			re := regexp.MustCompile("application-label:'(.*)'")
			label := re.FindStringSubmatch(out.String())[1]

			s := App{
				Label: label,
			}
			m[pkg] = s
		}
	}
	return m, err
}

func (apps Apps) Backup(path string, args ...string) error {
	if len(apps) == 0 {
		return fmt.Errorf("Empty Packagelist")
	}
	if path == "" {
		return fmt.Errorf("Empty path")
	}

	cmd := exec.Command("adb", "backup", "-f", path)
	for _, s := range args {
		cmd.Args = append(cmd.Args, s)
	}
	for p := range apps {
		cmd.Args = append(cmd.Args, p)
	}
	log.Printf("Backing up to %v with command: %v\n", path, cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	err = apps.DumpYaml(path + ".yaml")
	return err
}

func Restore(path string) error {
	if path == "" {
		return fmt.Errorf("Empty path")
	}
	cmd := exec.Command("adb", "restore", path)
	log.Printf("Restoring from %v with command: %v\n", path, cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	return cmd.Run()
}

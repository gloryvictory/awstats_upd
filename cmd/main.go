package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	// perl "C:/Apps/awstats/wwwroot/cgi-bin/awstats.pl" -config="configName" -update -LogFile="C:\inetpub\logs\LogFiles\W3SVC1\u_ex220101.log"

	vp := viper.New()
	vp.SetConfigName("awstats_upd")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	vp.AddConfigPath(".cmd")

	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))

		// fmt.Println("Unable to read config file!!", " Err:", err)
	}
	log_folder := vp.GetString("log_folder")     //"C:\\inetpub\\logs\\LogFiles\\W3SVC1\\"
	path_cgi_bin := vp.GetString("path_cgi_bin") // "C:\\Apps\\awstats\\wwwroot\\cgi-bin\\"
	config_name := vp.GetString("config_name")   // "arcgis"
	fileName_out := vp.GetString("fileName_out") //"update_ststistics.bat"

	fmt.Println(log_folder, path_cgi_bin, config_name, fileName_out)

	// var log_folder string = "C:\\inetpub\\logs\\LogFiles\\W3SVC1\\"
	// var path_cgi_bin string = "C:\\Apps\\awstats\\wwwroot\\cgi-bin\\"
	// var config_name string = "arcgis"
	// fileName_out := "update_ststistics.bat"

	f, err := os.Create(fileName_out)

	if err != nil {
		fmt.Println("Unable to create file: ", fileName_out, " Err:", err)
		log.Fatal(err)
		os.Exit(1)
	}

	defer f.Close()

	deleteOldStatistics(path_cgi_bin, config_name, f)
	fmt.Println("List files: ")
	listDirByWalk(log_folder, f, path_cgi_bin, config_name)

}

func deleteOldStatistics(path string, config_name string, f *os.File) {
	// var files []string
	// files, err := getFilesStatistics(path, ext_awstats)

	var ext_awstats string = "." + config_name + ".txt"
	var files []string
	fmt.Println("Start deleting")

	files, err := listDirByIOReadDir(path, ext_awstats)

	if err != nil {
		fmt.Printf("Error: %v", err)
		log.Fatal(err)
	}

	for _, file := range files {

		fmt.Println(file)
		var str_bat_cmd string = "del \"" + file + "\""
		f.WriteString(str_bat_cmd + "\n")

		// fmt.Printf("[%s]\n", path)
	}

}

func listDirByIOReadDir(root string, ext string) ([]string, error) {
	var files []string
	// var err2 error
	fileInfo, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return files, err
	}
	for _, file := range fileInfo {
		reg, err2 := regexp.Compile(ext)

		if err2 != nil {
			fmt.Printf("Error: %v", err2)
			return nil, err2
		}

		if reg.MatchString(file.Name()) {
			full_path1 := path.Join(root, file.Name())
			full_path := strings.ReplaceAll(full_path1, "/", "")
			files = append(files, full_path)
		}
	}
	return files, nil
}

func listDirByWalk(path string, f *os.File, path_cgi_bin string, config_name string) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {

		var perl_name string = "perl"
		// var path_cgi_bin string = "C:/Apps/awstats/wwwroot/cgi-bin/"
		var path_file_script string = "\"" + path_cgi_bin + "awstats.pl\""
		var param_config string = "-config=\"" + config_name + "\""
		var param_name string = "-update"
		var param_logfile string = "-LogFile="

		// Выводится команда
		if wPath != path {
			var qq string = perl_name + " " + path_file_script + " " + param_config + " " + param_name + " " + param_logfile + "\"" + wPath + "\""
			f.WriteString(qq + "\n")
		}
		return nil
	})
}

// func FilePathWalkDir(root string) ([]string, error) {
// 	var files []string
// 	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
// 		if !info.IsDir() {
// 			files = append(files, path)
// 		}
// 		return nil
// 	})
// 	return files, err
// }

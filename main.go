package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// perl "C:/Apps/awstats/wwwroot/cgi-bin/awstats.pl" -config="arcgis" -update -LogFile="C:\inetpub\logs\LogFiles\W3SVC1\u_ex220101.log"

	var log_folder string = "C:\\inetpub\\logs\\LogFiles\\W3SVC1\\"

	fmt.Println("List files by Walk")
	listDirByWalk(log_folder)
}

func listDirByWalk(path string) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {

		var perl_name string = "perl"
		var file_script_path string = "\"C:/Apps/awstats/wwwroot/cgi-bin/awstats.pl\""
		var config_name string = "-config=\"arcgis\""
		var parametr_name string = "-update"

		// Обход директории без вывода
		// if wPath == path {
		// 	return nil
		// }

		// Если данный путь является директорией, то останавливаем рекурсивный обход
		// и возвращаем название папки
		// if info.IsDir() {
		// 	fmt.Printf("[%s]\n", wPath)
		// 	log.Println(wPath)
		// 	return filepath.SkipDir
		// }

		// Выводится название файла
		if wPath != path {
			// fmt.Println(wPath)
			var qq string = perl_name + " " + file_script_path + " " + config_name + " " + parametr_name + " " + "-LogFile=\"" + wPath + "\""
			fmt.Println(qq)
			// fmt.Printf("[%s]\n", info.Name())
			// fmt.Printf("[%d]\n", info.Size())
			// fmt.Printf("[%d]\n", info.ModTime().Day())
			// log.Println(wPath)
			// fmt.Printf("%s %v %v %v\n", perl_name, file_script_path, config_name, parametr_name)
		}
		return nil
	})
}

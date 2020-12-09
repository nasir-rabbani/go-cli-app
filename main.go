/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"mycart/app"
	"mycart/app/helpers/databasehelper"
	"mycart/app/helpers/loghelper"
	"mycart/app/utils"
)

func init() {

	var err error

	// to set application configs from appConfig.yaml file
	// err = setconfig()
	// if err != nil {
	// 	panic(err)
	// }

	// Initialize logger
	initLogger()

	// Initialize database
	err = databasehelper.InitDatabases()
	utils.HandelException(err, true)
}

func main() {
	app.Run()
}

// setconfig - map project config provided in configs/appconfig.yaml file
// func setconfig() error {
// 	err := confighelper.Init(models.ConfigFilePath, &models.AppConfig)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// Initialize Logger
func initLogger() {
	loghelper.Init()
}

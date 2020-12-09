package cmd

import (
	"fmt"
	"mycart/app/models"
	"mycart/app/services"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "To register a User or Admin",
	Long:  `To add users to the system register [-a] [Name] [Password]`,
	Run: func(cmd *cobra.Command, args []string) {
		isAdmin, err := cmd.Flags().GetBool("admin")
		if err != nil {
			fmt.Println(err)
			return
		}
		username, err := cmd.PersistentFlags().GetString("name")
		if err != nil {
			fmt.Println(err)
			return
		}
		password, err := cmd.PersistentFlags().GetString("password")
		if err != nil {
			fmt.Println(err)
			return
		}

		if isAdmin {
			user, err := services.RegisterUser(username, password, models.Roles.Admin)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("User Registered with userID :: %d", user.Model.ID)
		} else {
			user, err := services.RegisterUser(username, password, models.Roles.User)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("User Registered with userID :: %d", user.Model.ID)
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	registerCmd.PersistentFlags().String("name", "", "Username of the User")
	registerCmd.PersistentFlags().String("password", "", "Password of the User")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	registerCmd.Flags().BoolP("admin", "a", false, "Denotes an admin")
}

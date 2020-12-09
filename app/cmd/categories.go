package cmd

import (
	"fmt"
	"mycart/app/models"
	"mycart/app/services"

	"github.com/spf13/cobra"
)

// categoriesCmd represents the categories command
var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Displays all Categories",
	Long:  `This command displays all the categories defined by the admin`,
	Run: func(cmd *cobra.Command, args []string) {
		categories, err := services.GetAllCategories()
		if err != nil {
			fmt.Println("Error fetching categories ::", err)
			return
		}
		displayCategories(categories)
	},
}

func displayCategories(categories []models.Category) {
	if len(categories) == 0 {
		fmt.Println("No categories found")
		return
	}
	fmt.Println("List of available categories")
	fmt.Println("-----------------------------")
	for _, category := range categories {
		fmt.Printf("Category ID : %d \n", category.Model.ID)
		fmt.Printf("Category name : %s \n", category.Name)
		fmt.Println("-----------------------------")
	}
}

// addCategoryCmd represents the categories command
var addCategoryCmd = &cobra.Command{
	Use:   "addCategory",
	Short: "Adds a new category",
	Long:  `This commands is used to create new Category`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.PersistentFlags().GetString("name")
		if err != nil {
			fmt.Println("Name is required ::", err)
			return
		}

		category, err := services.AddCategory(name)
		if err != nil {
			fmt.Println("Failed while Adding Category ::", err)
			return
		}
		fmt.Printf("Category Added with categoryID :: %d", category.Model.ID)
	},
}

func init() {
	categoriesCmd.AddCommand(addCategoryCmd)
	rootCmd.AddCommand(categoriesCmd)

	addCategoryCmd.PersistentFlags().String("name", "", "Name of the category")
}

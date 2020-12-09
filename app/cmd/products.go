package cmd

import (
	"fmt"
	"mycart/app/models"
	"mycart/app/services"

	"github.com/spf13/cobra"
)

// productsCmd represents the products command
var productsCmd = &cobra.Command{
	Use:   "products",
	Short: "Displays all Products",
	Long:  `This command displays all the products added by the admin`,
	Run: func(cmd *cobra.Command, args []string) {
		var products []models.Product
		var err error
		byCategoryID, _ := cmd.Flags().GetBool("byCategoryID")
		byProductID, _ := cmd.Flags().GetBool("byProductID")

		if byCategoryID {
			id, _ := cmd.PersistentFlags().GetUint("categoryID")
			products, err = services.GetProductsByCategoryID(id)
			if err != nil {
				fmt.Println("Error fetching categories ::", err)
				return
			}
		} else if byProductID {
			id, _ := cmd.PersistentFlags().GetUint("categoryID")

			product, err := services.GetProductByID(id)
			products = append(products, product)
			if err != nil {
				fmt.Println("Error fetching categories ::", err)
				return
			}
		} else {
			products, err = services.GetAllProducts()
			if err != nil {
				fmt.Println("Error fetching categories ::", err)
				return
			}

		}

		displayProducts(products)
	},
}

func displayProducts(products []models.Product) {
	if len(products) == 0 {
		fmt.Println("No products found")
		return
	}
	fmt.Println("List of products")
	fmt.Println("-----------------------------")
	for _, product := range products {
		fmt.Printf("Product ID : %d \n", product.Model.ID)
		fmt.Printf("Product name : %s \n", product.Name)
		fmt.Printf("Product Price : %0.2f \n", product.Price)
		fmt.Printf("Category ID : %d \n", product.CategoryID)
		fmt.Println("-----------------------------")
	}
}

// addProductCmd represents the addProduct command
var addProductCmd = &cobra.Command{
	Use:   "addProduct",
	Short: "Adds a new product",
	Long: `This commands is used to add new Product
	eg. mycart products addProduct --name Nokia-1100 --price 1500 --categoryID 1
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.PersistentFlags().GetString("name")
		if err != nil {
			fmt.Println("Name is required ::", err)
			return
		}
		price, err := cmd.PersistentFlags().GetFloat32("price")
		if err != nil {
			fmt.Println("Price is required ::", err)
			return
		}
		categoryID, err := cmd.PersistentFlags().GetUint("categoryID")
		if err != nil {
			fmt.Println("CategoryID is required ::", err)
			return
		}

		product, err := services.AddProduct(name, price, categoryID)
		if err != nil {
			fmt.Println("Failed while Adding Product ::", err)
			return
		}
		fmt.Printf("Product Added with productID :: %d", product.Model.ID)
	},
}

func init() {
	productsCmd.AddCommand(addProductCmd)
	rootCmd.AddCommand(productsCmd)

	addProductCmd.PersistentFlags().String("name", "", "Name of the product")
	addProductCmd.PersistentFlags().Float32("price", 0, "Price of the product")
	addProductCmd.PersistentFlags().Uint("categoryID", 0, "ID of the category it belogs")

	productsCmd.PersistentFlags().Uint("categoryID", 0, "category ID to fetch products of")
	productsCmd.PersistentFlags().Uint("productID", 0, "product ID to fetch detail of")

	productsCmd.Flags().BoolP("byCategoryID", "c", false, "Get products by category ID")
	productsCmd.Flags().BoolP("byProductID", "p", false, "Get products details by product ID")
}

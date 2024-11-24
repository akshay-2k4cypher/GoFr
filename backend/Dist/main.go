package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// User struct represents the user data
type User struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Address       string `json:"address"`
	CompanyName   string `json:"companyName"`
	LicenseNumber string `json:"licenseNumber"`
	UserType      string `json:"userType"`
}

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Enable CORS to allow requests from other origins (e.g., React frontend)
	app.Use(cors.New())

	// Register endpoint
	app.Post("/register", func(c *fiber.Ctx) error {
		var user User

		// Parse the JSON request body into the User struct
		if err := c.BodyParser(&user); err != nil {
			log.Println("Body parsing error:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid input",
			})
		}

		// Open or create the users.txt file
		file, err := os.OpenFile("users.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("File opening error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to open file",
			})
		}
		defer file.Close()

		// Convert user struct to JSON and write to the file
		userData, err := json.Marshal(user)
		if err != nil {
			log.Println("JSON marshaling error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to process user data",
			})
		}

		if _, err := file.WriteString(string(userData) + "\n"); err != nil {
			log.Println("File writing error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save user data",
			})
		}

		log.Printf("User registered: %+v\n", user)
		return c.JSON(fiber.Map{"message": "User registered successfully"})
	})

	// Start the server
	port := ":5000"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

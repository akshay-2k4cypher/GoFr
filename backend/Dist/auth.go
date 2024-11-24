func main() {
    app := gofr.New()

    // Enable Basic Authentication
    app.EnableBasicAuth("admin", "secret_password")

    app.GET("/distributor/requests", func(ctx *gofr.Context) (interface{}, error) {
        // Handle access to secure data
        return "List of retailer requests", nil
    })

    app.Run()
}
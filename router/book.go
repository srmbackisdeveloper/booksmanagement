package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/srmbackisdeveloper/booksmanagement/common"
	"github.com/srmbackisdeveloper/booksmanagement/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddBookGroup(app *fiber.App) {
	bookGroup := app.Group("/books")

	bookGroup.Get("/", getBooks)
	bookGroup.Get("/:id", getBook)
	bookGroup.Post("/", createBook)
	bookGroup.Put("/:id", updateBook)
	bookGroup.Delete("/:id", deleteBook)
}

func getBooks(c *fiber.Ctx) error {
	coll := common.GetDbCollection("books")

	// find all books
	books := make([]models.Book, 0)

	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// iterate over the cursor
	for cursor.Next(c.Context()) {
		book := models.Book{}

		err := cursor.Decode(&book)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		books = append(books, book)
	}

	return c.Status(200).JSON(fiber.Map{"data": books})
}

func getBook(c *fiber.Ctx) error {
	coll := common.GetDbCollection("books")

	// find the book
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is required"})
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	book := models.Book{}

	err = coll.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"data": book})
}

type createDTO struct {
	Title  string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	Year   string `json:"year" bson:"year"`
}

func createBook(c *fiber.Ctx) error {
	// validate the body
	b := new(createDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	// create the book
	coll := common.GetDbCollection("books")
	result, err := coll.InsertOne(c.Context(), b)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "failed to create book",
			"message": err.Error(),
		})
	}

	// return the book
	return c.Status(201).JSON(fiber.Map{
		"data": result,
	})
}

type updateDTO struct {
	Title  string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	Year   string `json:"year" bson:"year"`
}

func updateBook(c *fiber.Ctx) error {
	// validate the body
	b := new(updateDTO)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	// get the id
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is required"})
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	// update the book
	coll := common.GetDbCollection("books")
	result, err := coll.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": b})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "failed to update book",
			"message": err.Error(),
		})
	}

	// return the book
	return c.Status(200).JSON(fiber.Map{
		"data": result,
	})
}

func deleteBook(c *fiber.Ctx) error {
	// get the id
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	// delete the book
	coll := common.GetDbCollection("books")
	result, err := coll.DeleteOne(c.Context(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete book",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}

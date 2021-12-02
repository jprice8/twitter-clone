package handler

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jprice8/twitter-clone/internal/model"
	"github.com/jprice8/twitter-clone/internal/shared/database"
)

type OrderDetailByOrder struct {
	OrderID      int64     `json:"order_id"`
	UserID       int64     `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	ProductID    int64     `json:"product_id"`
	Price        float64   `json:"price"`
	Quantity     int64     `json:"quantity"`
	ExtendedCost int64     `json:"extended_cost"`
}

type OrderDetailsByOrder struct {
	OrderDetailsByOrder []OrderDetailByOrder `json:"orderDetailsByOrder`
}

func GetOrdersByUser(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get user ID from url parameter
		userId := c.Params("userId")
		// Select all orders by userId
		rows, err := db.Query("SELECT * FROM orders WHERE user_id = $1", userId)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Query failed on ordersByUser", "data": err})
		}

		// Keep connection open until function completes
		defer rows.Close()
		result := model.Orders{}
		// Iterate through returned rows
		for rows.Next() {
			order := model.Order{}
			if err := rows.Scan(&order.ID, &order.UserID, &order.ExtendedCost, &order.CreatedAt); err != nil {
				return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error while scanning rows for ordersByUser", "data": err})
			}

			// Append Order to Orders
			result.Orders = append(result.Orders, order)
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Successfully returned orders by user", "data": result})
	}
}

// Get order by orderId
func GetOrderById(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get order ID from url param
		orderId := c.Params("orderId")
		// Select order from db using order ID
		var order model.Order
		row := db.QueryRow("SELECT * FROM orders WHERE id = $1", orderId)
		if err := row.Scan(&order.ID, &order.UserID, &order.ExtendedCost, &order.CreatedAt); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error while scanning row for orderByID", "data": err})
		}
		return c.JSON(fiber.Map{"status": "success", "message": "Successfully returned order by ID", "data": order})
	}
}

// Get all orderDetails by orderId
func GetOrderDetailsByOrderId(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get order ID from url param
		orderId := c.Params("orderId")
		// Select all orderDetails from db using order ID
		rows, err := db.Query(`
		SELECT orders.id AS order_id, 
				orders.user_id, 
				orders.created_at,
				order_details.product_id,
				order_details.price,
				order_details.quantity,
				orders.extended_cost
		FROM orders
		JOIN order_details ON orders.id = order_details.order_id
		WHERE order_id = $1
		`, orderId)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Query failed on orderDetailsByOrder", "data": err})
		}

		defer rows.Close()
		result := OrderDetailsByOrder{}

		for rows.Next() {
			orderDetailByOrder := OrderDetailByOrder{}
			if err := rows.Scan(
				&orderDetailByOrder.OrderID,
				&orderDetailByOrder.UserID,
				&orderDetailByOrder.CreatedAt,
				&orderDetailByOrder.ProductID,
				&orderDetailByOrder.Price,
				&orderDetailByOrder.Quantity,
				&orderDetailByOrder.ExtendedCost);
				err != nil {
				return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error while scanning rows for orderDetailsByOrder", "data": err})
			}
			// Append to OrderDetailByOrder
			result.OrderDetailsByOrder = append(result.OrderDetailsByOrder, orderDetailByOrder)
		}

		return c.JSON(fiber.Map{"status": "success", "message": "successfully returned order details by order", "data": result})
	}
}

// Create new order
func CreateOrder(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Allocate memory for new order
		order := new(model.Order)
		if err := c.BodyParser(order); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't parse new order input", "data": err})
		}
		// Insert new order into database
		res, err := db.Query("INSERT INTO orders (user_id, extended_cost, created_at) VALUES ($1, $2, $3)", order.UserID, order.ExtendedCost, time.Now())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error creating new order in db", "data": err})
		}
		// Output result to log and return successful JSON response
		log.Println(res)
		return c.JSON(fiber.Map{"status": "success", "message": "Successfully created new order", "data": order})
	}
}

// Create new order detail
func CreateOrderDetail(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Allocate memory for new order detail
		orderDetail := new(model.OrderDetail)
		if err := c.BodyParser(orderDetail); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't parse new orderDetail input", "data": err})
		}
		// Insert new orderDetail into database
		res, err := db.Query("INSERT INTO order_details (order_id, product_id, price, quantity) VALUES ($1, $2, $3, $4)", orderDetail.OrderID, orderDetail.ProductID, orderDetail.Price, orderDetail.Quantity)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error creating new orderDetail in db", "data": err})
		}
		// Output result to log and return successful JSON response
		log.Println(res)
		return c.JSON(fiber.Map{"status": "success", "message": "Successfully created new orderDetail", "data": orderDetail})
	}
}

// Cancel existing order
func CancelOrder(db database.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get order ID from url param
		orderId := c.Params("orderId")
		res, err := db.Query("DELETE FROM orders WHERE id = $1", orderId)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error deleting order from database", "data": err})
		}
		log.Println(res)
		return c.JSON(fiber.Map{"status": "success", "message": "Successfully deleted order from database", "data": orderId})
	}
}

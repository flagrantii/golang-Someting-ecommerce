package models

type User struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	PasswordHash    string `json:"password_hash"`
	ShippingAddress string `json:"shipping_address"`
	PaymentInfo     string `json:"payment_info"`
	Role            string `json:"role"` // or use an enum
}

type Product struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Quantity    int      `json:"quantity"`
	Category    string   `json:"category"`
	Images      []string `json:"images"`
	SellerID    int      `json:"seller_id"`
}

type Order struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	Products        []Product `json:"products"`
	TotalPrice      float64   `json:"total_price"`
	Status          string    `json:"status"`
	ShippingAddress string    `json:"shipping_address"`
	PaymentStatus   string    `json:"payment_status"`
	Date            string    `json:"date"`
}

type Cart struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Products   []Product `json:"products"`
	TotalPrice float64   `json:"total_price"`
	Date       string    `json:"date"`
}

type Payment struct {
	ID      int     `json:"id"`
	OrderID int     `json:"order_id"`
	Method  string  `json:"method"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
	Date    string  `json:"date"`
}

type Shipping struct {
	ID         int    `json:"id"`
	OrderID    int    `json:"order_id"`
	Carrier    string `json:"carrier"`
	TrackingNo string `json:"tracking_no"`
	Status     string `json:"status"`
	Date       string `json:"date"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Review struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	UserID    int    `json:"user_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	Date      string `json:"date"`
}

type Discount struct {
	ID         int     `json:"id"`
	Code       string  `json:"code"`
	Discount   float64 `json:"discount"`
	ExpiryDate string  `json:"expiry_date"`
	UsageLimit int     `json:"usage_limit"`
}

type Shop struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     int       `json:"owner_id"`
	Products    []Product `json:"products"`
}

type Subscription struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	Plan          string `json:"plan"`
	Status        string `json:"status"`
	PaymentMethod string `json:"payment_method"`
}

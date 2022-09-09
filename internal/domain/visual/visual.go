package visual

type Visual struct {
	ID            int64
	OrderID       *int64
	InventoryID   *int64
	ProductID     *int64
	PartID        *int64
	PathID        *int64
	User          *string
	ConnectionID  int64
	HostProcessID *int64
	HostName      *string
	SiteName      *string
}

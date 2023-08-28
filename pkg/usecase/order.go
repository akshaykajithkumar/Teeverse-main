package usecase

import (
	domain "Teeverse/pkg/domain"
	"Teeverse/pkg/helper"
	internal "Teeverse/pkg/helper/pdf"
	interfaces "Teeverse/pkg/repository/interface"
	services "Teeverse/pkg/usecase/interface"
	"Teeverse/pkg/utils/models"
	"errors"
	"fmt"
	"time"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
	userUseCase     services.UserUseCase
	walletRepo      interfaces.WalletRepository
	couponRepo      interfaces.CouponRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository, userUseCase services.UserUseCase, walletRepository interfaces.WalletRepository, couponRepository interfaces.CouponRepository) *orderUseCase {
	return &orderUseCase{
		orderRepository: repo,
		userUseCase:     userUseCase,
		walletRepo:      walletRepository,
		couponRepo:      couponRepository,
	}
}

func (i *orderUseCase) GetOrders(id, page, limit int) ([]domain.Order, error) {

	orders, err := i.orderRepository.GetOrders(id, page, limit)
	if err != nil {
		return []domain.Order{}, err
	}

	return orders, nil

}

func (i *orderUseCase) OrderItemsFromCart(userid int, order models.Order) (string, error) {
	//cartchange
	cart, err := i.userUseCase.GetCart(userid, 0, 0)
	if err != nil {
		return "", err
	}

	var total float64
	for _, v := range cart {
		total = total + v.Total
	}

	//finding discount if any
	DiscountRate := i.couponRepo.FindCouponDiscount(order.CouponID)
	if DiscountRate > 0 {
		totalDiscount := (total * float64(DiscountRate)) / 100
		total = total - totalDiscount
	} else {
		totalDiscount := 0.0
		total = total - totalDiscount
	}

	var invoiceItems []*internal.InvoiceData
	for _, v := range cart {
		inventory, err := internal.NewInvoiceData(v.ProductName, int(v.Quantity), v.Total)
		if err != nil {
			panic(err)
		}
		invoiceItems = append(invoiceItems, inventory)
	}
	// Create single invoice
	invoice := internal.CreateInvoice("Teeverse", "www.Teeverse.store", invoiceItems)
	internal.GenerateInvoicePdf(*invoice)
	fmt.Printf("The Total Invoice Amount is: %f", invoice.CalculateInvoiceTotalAmount())

	//COD
	if order.PaymentID == 1 {
		order_id, err := i.orderRepository.OrderItems(userid, order, total)
		if err != nil {
			return "", err
		}

		if err := i.orderRepository.AddOrderProducts(order_id, cart); err != nil {
			return "", err
		}
		cartID, _ := i.userUseCase.GetCartID(userid)
		if err := i.userUseCase.ClearCart(cartID); err != nil {
			return "", err
		}

	} else if order.PaymentID == 2 {
		// razorpay
		order_id, err := i.orderRepository.OrderItems(userid, order, total)
		if err != nil {
			return "", err
		}

		if err := i.orderRepository.AddOrderProducts(order_id, cart); err != nil {
			return "", err
		}
		link := fmt.Sprintf("http://localhost:1243/users/payment/razorpay?id=%d", order_id)
		return link, err
	}

	//wallet

	return "", nil

}

func (i *orderUseCase) CancelOrder(id int) error {

	err := i.orderRepository.CancelOrder(id)
	if err != nil {
		return err
	}
	return nil

}

func (i *orderUseCase) EditOrderStatus(status string, id int) error {

	err := i.orderRepository.EditOrderStatus(status, id)
	if err != nil {
		return err
	}
	return nil

}

func (i *orderUseCase) AdminOrders(page, limit int) (domain.AdminOrdersResponse, error) {

	var response domain.AdminOrdersResponse

	pending, err := i.orderRepository.AdminOrders("PENDING")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	shipped, err := i.orderRepository.AdminOrders("SHIPPED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	delivered, err := i.orderRepository.AdminOrders("DELIVERED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	canceled, err := i.orderRepository.AdminOrders("CANCELED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	returned, err := i.orderRepository.AdminOrders("RETURNED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	response.Canceled = canceled
	response.Pending = pending
	response.Shipped = shipped
	response.Delivered = delivered
	response.Returned = returned
	return response, nil

}

func (i *orderUseCase) DailyOrders() (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := time.Now()
	startDate := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, time.UTC)
	SalesReport.Orders, _ = i.orderRepository.GetOrdersInRange(startDate, endDate)
	SalesReport.TotalOrders = len(SalesReport.Orders)
	total := 0.0
	for _, v := range SalesReport.Orders {
		total += v.Price
	}
	SalesReport.TotalRevenue = total

	products, err := i.orderRepository.GetProductsQuantity()
	if err != nil {
		return domain.SalesReport{}, err
	}
	bestSellerIDs := helper.FindMostBoughtProduct(products)
	var bestSellers []string
	for _, v := range bestSellerIDs {
		product, err := i.orderRepository.GetProductNameFromID(v)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}

func (i *orderUseCase) WeeklyOrders() (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := time.Now()
	startDate := endDate.Add(-time.Duration(endDate.Weekday()) * 24 * time.Hour)
	SalesReport.Orders, _ = i.orderRepository.GetOrdersInRange(startDate, endDate)
	SalesReport.TotalOrders = len(SalesReport.Orders)
	total := 0.0
	for _, v := range SalesReport.Orders {
		total += v.Price
	}
	SalesReport.TotalRevenue = total

	products, err := i.orderRepository.GetProductsQuantity()
	if err != nil {
		return domain.SalesReport{}, err
	}
	bestSellerIDs := helper.FindMostBoughtProduct(products)
	var bestSellers []string
	for _, v := range bestSellerIDs {
		product, err := i.orderRepository.GetProductNameFromID(v)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}

func (i *orderUseCase) MonthlyOrders() (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := time.Now()
	startDate := time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, 0, 0, time.UTC)
	SalesReport.Orders, _ = i.orderRepository.GetOrdersInRange(startDate, endDate)
	SalesReport.TotalOrders = len(SalesReport.Orders)
	total := 0.0
	for _, v := range SalesReport.Orders {
		total += v.Price
	}
	SalesReport.TotalRevenue = total

	products, err := i.orderRepository.GetProductsQuantity()
	if err != nil {
		return domain.SalesReport{}, err
	}
	bestSellerIDs := helper.FindMostBoughtProduct(products)
	var bestSellers []string
	for _, v := range bestSellerIDs {
		product, err := i.orderRepository.GetProductNameFromID(v)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers
	return SalesReport, nil
}

func (i *orderUseCase) AnnualOrders() (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := time.Now()
	startDate := time.Date(endDate.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	SalesReport.Orders, _ = i.orderRepository.GetOrdersInRange(startDate, endDate)
	SalesReport.TotalOrders = len(SalesReport.Orders)
	total := 0.0
	for _, v := range SalesReport.Orders {
		total += v.Price
	}
	SalesReport.TotalRevenue = total

	products, err := i.orderRepository.GetProductsQuantity()
	if err != nil {
		return domain.SalesReport{}, err
	}
	bestSellerIDs := helper.FindMostBoughtProduct(products)
	var bestSellers []string
	for _, v := range bestSellerIDs {
		product, err := i.orderRepository.GetProductNameFromID(v)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}

func (i *orderUseCase) CustomDateOrders(dates models.CustomDates) (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := dates.EndDate
	startDate := dates.StartingDate
	SalesReport.Orders, _ = i.orderRepository.GetOrdersInRange(startDate, endDate)
	SalesReport.TotalOrders = len(SalesReport.Orders)
	total := 0.0
	for _, v := range SalesReport.Orders {
		total += v.Price
	}
	SalesReport.TotalRevenue = total

	products, err := i.orderRepository.GetProductsQuantity()
	if err != nil {
		return domain.SalesReport{}, err
	}
	bestSellerIDs := helper.FindMostBoughtProduct(products)
	var bestSellers []string
	for _, v := range bestSellerIDs {
		product, err := i.orderRepository.GetProductNameFromID(v)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}

func (i *orderUseCase) ReturnOrder(id int) error {

	//should check if the order is already returned peoples will misuse this security breach
	// and will get  unlimited money into their wallet
	status, err := i.orderRepository.CheckIfTheOrderIsAlreadyReturned(id)
	if err != nil {
		return err
	}

	if status == "RETURNED" {
		return errors.New("order already returned")
	}

	//should also check if the order is already returned
	//or users will also earn money by returning pending orders by opting COD

	if status != "DELIVERED" {
		return errors.New("user is trying to return an order which is still not delivered")
	}

	//make order as returned order
	if err := i.orderRepository.ReturnOrder(id); err != nil {
		return err
	}

	//find amount to be credited to user
	amount, err := i.orderRepository.FindAmountFromOrderID(id)
	fmt.Println(amount)
	if err != nil {
		return err
	}

	//find the user
	userID, err := i.orderRepository.FindUserIdFromOrderID(id)
	fmt.Println(userID)
	if err != nil {
		return err
	}
	//find if the user having a wallet
	walletID, err := i.walletRepo.FindWalletIdFromUserID(userID)
	fmt.Println(walletID)
	if err != nil {
		return err
	}
	//if no wallet create new one
	if walletID == 0 {
		walletID, err = i.walletRepo.CreateNewWallet(userID)
		if err != nil {
			return err
		}
	}
	//credit the amount into users wallet
	if err := i.walletRepo.CreditToUserWallet(amount, walletID); err != nil {
		return err
	}

	return nil

}

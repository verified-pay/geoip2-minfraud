package geoip2

import "github.com/verified-pay/geoip2-minfraud/model"

func (mf *MinFraud) WithAccount(data *model.AccountReq) *MinFraud {
	mf.Req.Account = data
	return mf
}

func (mf *MinFraud) WithBilling(data *model.BillingReq) *MinFraud {
	mf.Req.Billing = data
	return mf
}

func (mf *MinFraud) WithShipping(data *model.ShippingReq) *MinFraud {
	mf.Req.Shipping = data
	return mf
}

func (mf *MinFraud) WithCreditCard(data *model.CreditCardReq) *MinFraud {
	mf.Req.CreditCard = data
	return mf
}

func (mf *MinFraud) WithCustomInputs(data *model.CustomInputsReq) *MinFraud {
	mf.Req.CustomInputs = data
	return mf
}

func (mf *MinFraud) WithDevice(data *model.DeviceReq) *MinFraud {
	mf.Req.Device = data
	return mf
}

func (mf *MinFraud) WithEmail(data *model.EmailReq) *MinFraud {
	mf.Req.Email = data
	return mf
}

func (mf *MinFraud) WithEvent(data *model.EventReq) *MinFraud {
	mf.Req.Event = data
	return mf
}

func (mf *MinFraud) WithOrder(data *model.OrderReq) *MinFraud {
	mf.Req.Order = data
	return mf
}

func (mf *MinFraud) WithPayment(data *model.PaymentReq) *MinFraud {
	mf.Req.Payment = data
	return mf
}

func (mf *MinFraud) WithShoppingCartItem(data *model.ShoppingCartReq) *MinFraud {
	mf.Req.ShoppingCart = append(mf.Req.ShoppingCart, data)
	return mf
}

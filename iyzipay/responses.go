package iyzipay

type (
	Resource struct {
		Status         string `json:"status"`
		ErrorCode      string `json:"errorCode"`
		ErrorMessage   string `json:"errorMessage"`
		ErrorGroup     string `json:"errorGroup"`
		Locale         string `json:"locale"`
		SystemTime     int64  `json:"systemTime"`
		ConversationID string `json:"conversationId"`
	}

	SubMerchantResponse struct {
		Resource
		Name                  string `json:"name"`
		Email                 string `json:"email"`
		GsmNumber             string `json:"gsmNumber"`
		Address               string `json:"address"`
		Iban                  string `json:"iban"`
		SwiftCode             string `json:"swiftCode"`
		Currency              string `json:"currency"`
		TaxOffice             string `json:"taxOffice"`
		ContactName           string `json:"contactName"`
		ContactSurname        string `json:"contactSurname"`
		LegalCompanyTitle     string `json:"legalCompanyTitle"`
		SubMerchantExternalID string `json:"subMerchantExternalId"`
		IdentityNumber        string `json:"identityNumber"`
		TaxNumber             string `json:"taxNumber"`
		SubMerchantType       string `json:"subMerchantType"`
		SubMerchantKey        string `json:"subMerchantKey"`
	}

	CheckoutFormInitializeResource struct {
		Resource
		Token               string `json:"token"`
		CheckoutFormContent string `json:"checkoutFormContent"`
		TokenExpireTime     int64  `json:"tokenExpireTime"`
		PaymentPageURL      string `json:"paymentPageUrl"`
	}

	RefundResponse struct {
		Resource
		PaymentID            string  `json:"paymentId"`
		PaymentTransactionID string  `json:"paymentTransactionId"`
		Price                float64 `json:"price"`
		Currency             string  `json:"currency"`
		ConnectorName        string  `json:"connectorName"`
	}

	CheckoutFormResponse struct {
		Status                       string  `json:"status"`
		Locale                       string  `json:"locale"`
		SystemTime                   int64   `json:"systemTime"`
		Price                        float64 `json:"price"`
		PaidPrice                    float64 `json:"paidPrice"`
		Installment                  int     `json:"installment"`
		PaymentID                    string  `json:"paymentId"`
		FraudStatus                  int     `json:"fraudStatus"`
		MerchantCommissionRate       float64 `json:"merchantCommissionRate"`
		MerchantCommissionRateAmount float64 `json:"merchantCommissionRateAmount"`
		IyziCommissionRateAmount     float64 `json:"iyziCommissionRateAmount"`
		IyziCommissionFee            float64 `json:"iyziCommissionFee"`
		CardType                     string  `json:"cardType"`
		CardAssociation              string  `json:"cardAssociation"`
		CardFamily                   string  `json:"cardFamily"`
		BinNumber                    string  `json:"binNumber"`
		LastFourDigits               string  `json:"lastFourDigits"`
		BasketID                     string  `json:"basketId"`
		Currency                     string  `json:"currency"`
		ItemTransactions             []struct {
			ItemID                        string  `json:"itemId"`
			PaymentTransactionID          string  `json:"paymentTransactionId"`
			TransactionStatus             int     `json:"transactionStatus"`
			Price                         float64 `json:"price"`
			PaidPrice                     float64 `json:"paidPrice"`
			MerchantCommissionRate        float64 `json:"merchantCommissionRate"`
			MerchantCommissionRateAmount  float64 `json:"merchantCommissionRateAmount"`
			IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
			IyziCommissionFee             float64 `json:"iyziCommissionFee"`
			BlockageRate                  float64 `json:"blockageRate"`
			BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
			BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`
			BlockageResolvedDate          string  `json:"blockageResolvedDate"`
			SubMerchantKey                string  `json:"subMerchantKey"`
			SubMerchantPrice              float64 `json:"subMerchantPrice"`
			SubMerchantPayoutRate         float64 `json:"subMerchantPayoutRate"`
			SubMerchantPayoutAmount       float64 `json:"subMerchantPayoutAmount"`
			MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
			ConvertedPayout               struct {
				PaidPrice                     float64 `json:"paidPrice"`
				IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
				IyziCommissionFee             float64 `json:"iyziCommissionFee"`
				BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
				BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`
				SubMerchantPayoutAmount       float64 `json:"subMerchantPayoutAmount"`
				MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
				IyziConversionRate            float64 `json:"iyziConversionRate"`
				IyziConversionRateAmount      float64 `json:"iyziConversionRateAmount"`
				Currency                      string  `json:"currency"`
			} `json:"convertedPayout"`
		} `json:"itemTransactions"`
		AuthCode      string `json:"authCode"`
		Phase         string `json:"phase"`
		MdStatus      int    `json:"mdStatus"`
		HostReference string `json:"hostReference"`
		Token         string `json:"token"`
		CallbackURL   string `json:"callbackUrl"`
		PaymentStatus string `json:"paymentStatus"`
	}
)

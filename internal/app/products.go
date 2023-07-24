package app

/*
func (b *Base) GetStripeProducts() []models.Product {
	products := []models.Product{}
	list := b.Stripe.Products.List(nil)
	for _, val := range list.ProductList().Data {
		price, _ := b.Stripe.Prices.Get(val.DefaultPrice.ID, nil)

		product := models.Product{
			ID:          val.ID,
			Name:        val.Name,
			Description: val.Description,
			Price: models.Price{
				ID:       val.DefaultPrice.ID,
				Currency: string(price.Currency),
				Amount:   int(price.UnitAmount),
			},
			Images:     val.Images,
			URL:        val.URL,
			Metadata:   val.Metadata,
			Attributes: val.Attributes,
			Created:    val.Created,
			Updated:    val.Updated,
		}

		products = append(products, product)
	}

	return products
}

func (b *Base) AddStripeProducts() *stripe.Product {
	newProduct := &stripe.ProductParams{
		Active:      stripe.Bool(true),
		Name:        stripe.String("Test Name"),
		Description: stripe.String("This is a description"),
		URL:         stripe.String("http://example.com"),
		Images: stripe.StringSlice([]string{
			"https://werbot.com/assets/img/steps-1.png",
			"https://werbot.com/assets/img/steps-2.png",
			"https://werbot.com/assets/img/steps-3.png",
		}),
		Shippable: stripe.Bool(true),
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			Currency:   stripe.String(string(stripe.CurrencyEUR)),
			UnitAmount: stripe.Int64(5000),
		},
	}
	newProduct.AddMetadata("key", "value")

	recponse, err := b.Stripe.Products.New(newProduct)
	if err != nil {
		b.Log.Fatal().Err(err).Send()
	}

	return recponse
}
*/

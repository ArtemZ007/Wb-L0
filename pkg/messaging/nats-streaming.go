// Настройка подключения к NATS
nc, err := nats.Connect(natsURL)
if err != nil {
	log.Fatal(err)
}
defer nc.Close()

// Основной цикл прослушивания сообщений
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
go func() {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg := <-nc.Subscribe("new_order")
			if msg == nil {
				continue
			}
			order := &Order{}
			err := json.Unmarshal(msg.Data, order)
			if err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}
			// TODO: Process order
			log.Printf("Received new order: %v", order)
			err = nc.Publish("processed_order", []byte(fmt.Sprintf("Processed order: %v", order)))
			if err != nil {
				log.Printf("Error publishing processed order: %v", err)
			}
		}
	}
}()
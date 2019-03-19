package consent

func Test() {
	conf := ConsentStoreConfig{
		Addr:     "localhost:1234",
		Password: "12345",
		DB:       0,
	}
	store := Init(conf)

	store.SetConsent(ConsentMessage{id: "Maki", status: REVOKED})
	store.GetConsent("Maki")
}

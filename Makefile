db_connect:
	@echo "Connecting to database..."
	@docker compose exec db mysql -umy_food_app -pmy_food_app my_food_app
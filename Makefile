# These commands will only be used during development (usernames and passwords that might be used here are not secure and will be changed in production)
db-connect:
	echo "Connecting to database..."
	@docker compose exec db mysql -umy_food_app -pmy_food_app my_food_app

build-api-bin:
	@echo "Building API binary..."
	@start_time=$$(date +%s%N); \
	cd api && go build -o ./build/api .; \
	end_time=$$(date +%s%N); \
	elapsed_ns=$$((end_time - start_time)); \
	elapsed_ms=$$((elapsed_ns / 1000000)); \
	minutes=$$((elapsed_ms / 60000)); \
	seconds=$$(( (elapsed_ms % 60000) / 1000 )); \
	milliseconds=$$(( (elapsed_ms % 1000) )); \
	microseconds=$$(( (elapsed_ns % 1000000) / 1000 )); \
	if [ $$minutes -gt 0 ]; then \
		if [ $$seconds -gt 0 ]; then \
			printf "API binary built successfully in %dmin %ds %dms %dµs\n" $$minutes $$seconds $$milliseconds $$microseconds; \
		else \
			printf "API binary built successfully in %dmin %dms %dµs\n" $$minutes $$milliseconds $$microseconds; \
		fi; \
	elif [ $$seconds -gt 0 ]; then \
		printf "API binary built successfully in %ds %dms %dµs\n" $$seconds $$milliseconds $$microseconds; \
	else \
		printf "API binary built successfully in %dms %dµs\n" $$milliseconds $$microseconds; \
	fi

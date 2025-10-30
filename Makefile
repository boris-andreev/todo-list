MOCKS_DIR := ./internal/mocks
INTERFACES := \
    ./internal/model/repository.go \

.PHONY: mocks clean

mocks: $(MOCKS_DIR)
	@echo "Generating mocks..."
	@for file in $(INTERFACES); do \
		filename=$$(basename $$file); \
		mockgen -source=$$file -destination=$(MOCKS_DIR)/$$filename; \
	done
	@echo "Mocks generated in $(MOCKS_DIR)"

$(MOCKS_DIR):
	@mkdir -p $(MOCKS_DIR)

clean:
	@echo "Cleaning up mocks..."
	@rm -rf $(MOCKS_DIR)
	@echo "Mocks directory removed."

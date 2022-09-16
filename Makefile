BIN_DIR := $(CURDIR)/bin
DRIVERS_DIR := $(BIN_DIR)/drivers
OPENAPI_GENERATOR_JAR_LOCATION := $(BIN_DIR)/openapi-generator-cli.jar

OPENAPI_DIR := $(CURDIR)/openapi
OPENAPI_SPEC_LOCATION := $(OPENAPI_DIR)/spec.yaml
OPENAPI_CUSTOM_GENERATOR_DIR := $(OPENAPI_DIR)/codegen

OPENAPI_OUTPUT_DIR := $(CURDIR)/api

.PHONY: generate
generate: compile-custom-generator
	rm -r api/
	java -cp \
		$(OPENAPI_CUSTOM_GENERATOR_DIR)/target/go-custom-openapi-generator-1.0.0.jar:$(OPENAPI_GENERATOR_JAR_LOCATION) \
		org.openapitools.codegen.OpenAPIGenerator generate \
		-i $(OPENAPI_SPEC_LOCATION) \
		-g go-custom \
    	-o $(OPENAPI_OUTPUT_DIR) \
		--git-repo-id webdriver \
		--git-user-id dsmontoya \
		-p packageName=api,generateInterfaces=true,enumClassPrefix=true,structPrefix=true

.PHONY: --ensure-bin-dir
--ensure-bin-dir:
	mkdir -p $(BIN_DIR)

.PHONY: --ensure-chrome-driver-dir
--ensure-chrome-driver-dir:
	mkdir -p $(CHROME_DRIVER_DIR)

--set-chrome-driver-filename-%:
	mkdir -p $(BIN_DIR)

.PHONY: compile-custom-generator
compile-custom-generator:
	cd $(OPENAPI_CUSTOM_GENERATOR_DIR) && mvn package

.PHONY: setup-common
setup-common:
	asdf install
	wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/6.0.1/openapi-generator-cli-6.0.1.jar -O $(OPENAPI_GENERATOR_JAR_LOCATION)

.PHONY: update-openapi-spec
update-openapi-spec:
	wget https://raw.githubusercontent.com/aerokube/selenium-openapi/master/selenium.yaml -O $(OPENAPI_SPEC_LOCATION)

### Drivers

CHROME_DRIVER_VERSION := 104.0.5112.79
CHROME_DRIVER_DIR:=$(DRIVERS_DIR)/chrome
download-chrome-%: export FILE_NAME=$(CHROME_DRIVER_DIR)/$(*F)
download-chrome-%: --ensure-bin-dir --ensure-chrome-driver-dir
	wget https://chromedriver.storage.googleapis.com/$(CHROME_DRIVER_VERSION)/chromedriver_$(*F).zip -O $(FILE_NAME).zip
	unzip -o $(FILE_NAME).zip -d $(FILE_NAME) && rm $(FILE_NAME).zip

INTEGRATION_TEST_BUILD_TAGS := integration_test
INTEGRATION_TEST_FLAGS := -tags=$(INTEGRATION_TEST_BUILD_TAGS)
test-integration-chrome-%: #TODO: target to run integration tests for all the platforms.
test-integration-chrome-%: export CHROME_DRIVER_PATH=$(CHROME_DRIVER_DIR)/$(*F)/chromedriver
test-integration-chrome-%:
	ls $(CHROME_DRIVER_PATH) || $(MAKE) download-chrome-$(*F)
	
	go test $(INTEGRATION_TEST_FLAGS) -v ./...
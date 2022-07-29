XSD_DIR=./res/xsd
XSD_FILES=BPMN20.xsd BPMNDI.xsd Semantic.xsd DC.xsd DI.xsd
JSON_SCHEMA_DIR=./res/json
JSON_SCHEMA_FILES=bpmn.json bpmndi.json dc.json di.json

define load_xsd_file
	curl "https://www.omg.org/spec/BPMN/20100501/$(1)" -o "$(XSD_DIR)/$(1)"

endef

define load_json_schema_file
	curl "https://raw.githubusercontent.com/bpmn-io/bpmn-moddle/master/resources/bpmn/json/$(1)" -o "$(JSON_SCHEMA_DIR)/$(1)"

endef

load_bpmn_xsd:
	$(foreach XSD_FILE, $(XSD_FILES), $(call load_xsd_file,$(XSD_FILE)))

load_bpmn_json_schema_files:
	$(foreach JSON_FILE, $(JSON_SCHEMA_FILES), $(call load_json_schema_file,$(JSON_FILE)))

install_zek:
	go install github.com/miku/zek/cmd/zek@latest

gen_zek:
	cat ./res/bpmn/complex.bpmn | zek -P gen > ./pkg/gen/code.go

build:
	go install ./...

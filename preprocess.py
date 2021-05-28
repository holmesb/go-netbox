#!/usr/bin/python3
import json
import logging
import os

# logging configuration
LOGLEVEL = os.environ.get("LOGLEVEL", default="INFO").upper()
logging.basicConfig(level=getattr(logging, LOGLEVEL))
logger = logging.getLogger(__name__)

source_file = "swagger.json"

# First, add a default response api call. If these do not exists, we get no error messages from the API.
with open(source_file, "r") as readfile:
    data = json.load(readfile)

default_response = {
    "description": "",
    "schema": {"type": "object", "additionalProperties": True},
}

for path, path_spec in data["paths"].items():
    logging.debug("checking path: " + path)
    for verb, verb_spec in path_spec.items():
        # ignore the parameters "verb"
        if verb == "parameters":
            continue
        if "responses" in verb_spec and "default" not in verb_spec["responses"]:
            logging.info("Adding default response to " + path + " " + verb)
            verb_spec["responses"]["default"] = default_response

# Second, fix the config contexts and local context data to our needs
# This implements https://github.com/fbreckle/go-netbox/commit/987669a91daf2d04a155feaf032a24ed684169b7 and https://github.com/fbreckle/go-netbox/commit/030637bab4bb25cec173035cd2d4a78fb3f47053
for definition, definition_spec in data["definitions"].items():
    for property, property_spec in definition_spec["properties"].items():
        # for config_context, change additionalProperties.type from string to object
        if property == "config_context":
            if property_spec["additionalProperties"]["type"] == "string":
                logging.debug(
                    f"{definition}.{property}.additionalProperties.type == string"
                )
                logging.debug("before: " + str(property_spec))
                property_spec["additionalProperties"] = {"type": "object"}
                logging.debug("after: " + str(property_spec))
                logging.info(f"fixed {definition}.{property}")

        # for local_context_data, change type from string to object
        # and add additionalProperties with type object
        if property == "local_context_data":
            if property_spec["type"] == "string":
                logging.debug(f"{definition}.{property}.type == string")
                logging.debug("before: " + str(property_spec))
                property_spec["type"] = "object"
                property_spec["additionalProperties"] = {"type": "object"}
                logging.debug("after: " + str(property_spec))
                logging.info(f"fixed {definition}.{property}")


# Third, allow unsetting some optional attributes
# This implements https://github.com/fbreckle/go-netbox/commit/23f80e06cb3da6f0fd9974a2d986c09826e518b5
for property, property_spec in data["definitions"]["WritableIPAddress"][
    "properties"
].items():
    if (
        "x-nullable" in property_spec
        and property_spec["x-nullable"] == True
        and property_spec["type"] == "integer"
    ):
        property_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritableIPAddress.{property}")
    if property == "tags":
        property_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritableIPAddress.{property}")

# This implements https://github.com/fbreckle/go-netbox/commit/1363e14cfc7bce4bd3d5ee93c09ca70543c51279
for property, property_spec in data["definitions"][
    "WritableVirtualMachineWithConfigContext"
]["properties"].items():
    if "x-nullable" in property_spec and property_spec["x-nullable"] == True:
        property_spec["x-omitempty"] = False
        logging.info(
            f"set x-omitempty = false on WritableVirtualMachineWithConfigContext.{property}"
        )
    if property == "tags":
        property_spec["x-omitempty"] = False
        logging.info(
            f"set x-omitempty = false on WritableVirtualMachineWithConfigContext.{property}"
        )

# todo: check if we actually need https://github.com/fbreckle/go-netbox/commit/7a076f034fc50bb913d5496c193c267e55eb2d2c


# Write output file
with open("swagger.processed.json", "w") as writefile:
    json.dump(data, writefile, indent=2)

#!/bin/bash

scriptPos=${0%/*}

docker run --rm \
  --user $(id -u):$(id -g) \
  -v "$(cd $scriptPos/.. && pwd)/docs/puml":/data \
  plantuml/plantuml \
  -tsvg badginator_model.puml -tsvg config_model.puml

# docker run --rm \
#   --user $(id -u):$(id -g) \
#   -v "$(cd $scriptPos/.. && pwd)/docs/puml":/data \
#   plantuml/plantuml \
#   -tsvg config_model.puml

stupidOutput="$scriptPos/../docs/puml/?"
if [ -d "$stupidOutput" ]; then
  rm -rf "$stupidOutput"
fi

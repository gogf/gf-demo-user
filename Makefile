ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "gf-demo-user"
DOCKER_NAME = "gf-demo-user"

include ./hack/hack.mk

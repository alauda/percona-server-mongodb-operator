# Copyright 2018 Oracle and/or its affiliates. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SOURCE_DIR="./pkg"
BUNDLE_IMAGE ?= build-harbor.alauda.cn/middleware/percona-server-mongodb-operator-bundle
OPERATOR_IMAGE ?= build-harbor.alauda.cn/middleware/percona-server-mongodb-operator
INDEX_IMAGE ?= build-harbor.alauda.cn/middleware/percona-server-mongodb-operator-index
BUNDLE_VERSION ?= $(shell gitversion /output json /showvariable MajorMinorPatch)-$(shell gitversion /output json /showvariable ShortSha)-$(GITNAME)
CONTROLLER_VERSION ?= $(BUNDLE_VERSION)

lint:
	@golint $(SOURCE_DIR)

.PHONY: opm
# find or download opm
# download opm if necessary
opm: ## 安装 opm
ifeq (, $(shell which opm))
	@{ \
	set -e ;\
	OPM_TMP_DIR=$$(mktemp -d) ;\
	cd $$OPM_TMP_DIR ;\
	git clone --branch v1.15.2 --depth 1 https://github.com/operator-framework/operator-registry.git && cd operator-registry ;\
	export CGO_ENABLED=1 ;\
	make build ;\
	cp ./bin/* $(GOBIN)/ ;\
	rm -rf $$OPM_TMP_DIR ;\
	}
OPM=$(GOBIN)/opm
else
OPM=$(shell which opm)
endif

.PHONY: kustomize
kustomize: ## 安装 kustomize
ifeq (, $(shell which kustomize))
	@{ \
	set -e ;\
	KUSTOMIZE_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$KUSTOMIZE_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/kustomize/kustomize/v3@v3.5.4 ;\
	rm -rf $$KUSTOMIZE_GEN_TMP_DIR ;\
	}
KUSTOMIZE=$(GOPATH)/bin/kustomize
else
KUSTOMIZE=$(shell which kustomize)
endif

.PHONY: operator-sdk
operator-sdk: ## 安装 operator-sdk
ifeq (, $(shell which operator-sdk))
	@{ \
	set -e ;\
	OPERATOR_SDK_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$OPERATOR_SDK_GEN_TMP_DIR ;\
	git clone --branch v1.2.0 --depth 1 https://github.com/operator-framework/operator-sdk.git && cd operator-sdk ;\
	make install ;\
	rm -rf $$OPERATOR_SDK_GEN_TMP_DIR ;\
	}
OPERATOR_SDK=$(GOPATH)/bin/operator-sdk
else
OPERATOR_SDK=$(shell which operator-sdk)
endif

.PHONY: bundle
bundle: operator-sdk kustomize ## 构建bundle
	@cd config/manager && $(KUSTOMIZE) edit set image controller=$(OPERATOR_IMAGE):$(CONTROLLER_VERSION)
	@cd config/manager && $(KUSTOMIZE) edit add annotation -f operatorversion:$(CONTROLLER_VERSION)
	@$(KUSTOMIZE) build config/manifests | $(OPERATOR_SDK) generate bundle --version $(BUNDLE_VERSION) --channels stable
	@$(OPERATOR_SDK) bundle validate ./bundle

.PHONY: bundle_image
bundle_image: bundle ## 构建bundle镜像
	@docker build -f bundle.Dockerfile -t $(BUNDLE_IMAGE):$(BUNDLE_VERSION) .

.PHONY: push_bundle_image
push_bundle_image: bundle_image ## 推送bundle镜像
	@docker push $(BUNDLE_IMAGE):$(BUNDLE_VERSION)

.PHONY: index_image
index_image: opm ## 构建operator index镜像
	@$(OPM) index add --bundles $(BUNDLE_IMAGE):$(BUNDLE_VERSION) --tag $(INDEX_IMAGE):$(BUNDLE_VERSION) -u docker

.PHONY: push_index_image
push_index_image: index_image ## 推送operator index镜像
	@docker push $(INDEX_IMAGE):$(BUNDLE_VERSION)

.PHONY: clean
clean: ## 清理
	@-rm -rf bin bundle bundle.Dockerfile index_build_* index_tmp_*

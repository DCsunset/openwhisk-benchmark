#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

OW_USER?=openwhisk
OW_VER?=v1.15
OW_RUNTIME?=$(OW_USER)/action-golang-$(OW_VER)
OW_COMPILER?=$(OW_USER)/action-golang-$(OW_VER)
WSK?=wsk
SRC=main.go go.mod
NAME=hello-go

deploy: main.go
	zip -r src.zip $(SRC)
	docker run -i $(OW_COMPILER) -compile main <src.zip >exec.zip
	$(WSK) -i action update $(NAME) exec.zip --kind go:1.15

clean:
	$(WSK) -i action delete counter
	rm src.zip exec.zip

.PHONY: clean

#!/bin/bash

# 設定基礎路徑
BASE_PATH="/Users/ed/EDGO/twelveGo"

# 建立目錄結構
mkdir -p "$BASE_PATH/cmd/api"
mkdir -p "$BASE_PATH/pkg/api"
mkdir -p "$BASE_PATH/pkg/model"
mkdir -p "$BASE_PATH/internal/database"
mkdir -p "$BASE_PATH/web/static"
mkdir -p "$BASE_PATH/web/templates"
mkdir -p "$BASE_PATH/scripts"
mkdir -p "$BASE_PATH/build"
mkdir -p "$BASE_PATH/deployments"
mkdir -p "$BASE_PATH/api/api-spec"
mkdir -p "$BASE_PATH/test"
mkdir -p "$BASE_PATH/docs"

# 建立檔案
touch "$BASE_PATH/cmd/api/main.go"
touch "$BASE_PATH/pkg/api/authRouter.go"
touch "$BASE_PATH/pkg/api/shoppingCartRouter.go"
touch "$BASE_PATH/pkg/api/memberRouter.go"
touch "$BASE_PATH/pkg/model/product.go"
touch "$BASE_PATH/pkg/model/order.go"
touch "$BASE_PATH/pkg/model/user.go"
touch "$BASE_PATH/internal/database/db.go"
touch "$BASE_PATH/README.md"

# 為空目錄建立.gitkeep檔案
touch "$BASE_PATH/web/static/.gitkeep"
touch "$BASE_PATH/web/templates/.gitkeep"
touch "$BASE_PATH/scripts/.gitkeep"
touch "$BASE_PATH/build/.gitkeep"
touch "$BASE_PATH/deployments/.gitkeep"
touch "$BASE_PATH/api/api-spec/.gitkeep"
touch "$BASE_PATH/test/.gitkeep"
touch "$BASE_PATH/docs/.gitkeep"

echo "Project directories and files created."

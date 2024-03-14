#!/bin/bash
echo "Building Lambda functions - Email"
FOLDER_EMAIL=email-process
GO_FILE_EMAIL="main.go"
OUTPUT_NAME_EMAIL="main"
OUTPUT_ZIP_NAME_EMAIL="main-email"
STATIC_DIR="templates"

rm -f ${OUTPUT_NAME_EMAIL}.zip #Clean up the previous build

cd $FOLDER_EMAIL
rm -f $OUTPUT_NAME_EMAIL #Clean up the previous build

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $OUTPUT_NAME_EMAIL -ldflags '-w' $GO_FILE_EMAIL

zip -j ../${OUTPUT_ZIP_NAME_EMAIL}.zip "$OUTPUT_NAME_EMAIL"
zip -r ../${OUTPUT_ZIP_NAME_EMAIL}.zip "$STATIC_DIR"

cd ..

echo "Email Lambda built successfully"

########################################################################################


echo "Building Lambda functions - Transaction"
FOLDER_TRANSACTION=transaction-process
GO_FILE_TRANSACTION="main.go"
OUTPUT_NAME_TRANSACTION="main"
OUTPUT_ZIP_NAME_TRANSACTION="main-transaction"

rm -f ${OUTPUT_NAME_TRANSACTION}.zip #Clean up the previous build
cd $FOLDER_TRANSACTION 
rm -f $OUTPUT_NAME_EMAIL #Clean up the previous build

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $OUTPUT_NAME_TRANSACTION -ldflags '-w' $GO_FILE_TRANSACTION
cd ..

zip -j ${OUTPUT_ZIP_NAME_TRANSACTION}.zip "$FOLDER_TRANSACTION/$OUTPUT_NAME_TRANSACTION"
echo "Transaction Lambda built successfully"

#-----------------------------------------------------
# bucket S3
#-----------------------------------------------------

resource "aws_s3_bucket" "lambda_bucket" {
  bucket = "car-smile-go"
}

resource "aws_s3_bucket_acl" "access_control_list" {
  bucket = aws_s3_bucket.lambda_bucket.id
  acl    = "private"
}


data "archive_file" "lambda_zip" {
  type = "zip"

  source_dir  = "${path.module}/target/"
  output_path = "${path.module}/target/golambda.zip"
}

resource "aws_s3_object" "lambda_car_smile" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = "golambda.zip"
  source = data.archive_file.lambda_zip.output_path

  etag = filemd5(data.archive_file.lambda_zip.output_path)
}
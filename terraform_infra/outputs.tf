# Output value definitions

output "lambda_bucket_name" {
  description = "Name of the S3 bucket used to store function code."

  value = aws_s3_bucket.lambda_bucket.id
}

output "function_name" {
  description = "Name of the Lambda function."

  value = aws_lambda_function.car-smile-mngr.function_name
}

output "base_url" {
  value = aws_api_gateway_deployment.rest_api_car_deploy.invoke_url
}
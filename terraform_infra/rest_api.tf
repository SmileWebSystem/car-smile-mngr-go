resource "aws_api_gateway_rest_api" "rest_api_car" {
  name = var.rest_api_name
  body = data.template_file.specification.rendered
}

data "template_file" "specification" {
  template = file("api-specification.json")

  vars = {
    lambda_identity_arn = aws_lambda_function.car-smile-mngr.arn
    aws_region          = var.AWS_REGION
  }

}

resource "aws_api_gateway_deployment" "rest_api_car_deploy" {
  rest_api_id = aws_api_gateway_rest_api.rest_api_car.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.rest_api_car.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "example" {
  deployment_id = aws_api_gateway_deployment.rest_api_car_deploy.id
  rest_api_id   = aws_api_gateway_rest_api.rest_api_car.id
  stage_name    = "dev"

}

resource "aws_lambda_permission" "api-gateway-invoke-lambda" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.car-smile-mngr.function_name
  principal     = "apigateway.amazonaws.com"

  # The /*/* portion grants access from any method on any resource
  # within the specified API Gateway.
  source_arn = "${aws_api_gateway_rest_api.rest_api_car.execution_arn}/*/*"
}
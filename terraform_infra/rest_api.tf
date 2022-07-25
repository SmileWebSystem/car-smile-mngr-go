resource "aws_api_gateway_rest_api" "rest_api_car" {
  name = var.rest_api_name
  body = jsonencode({
    "openapi" : "3.0.1",
    "info" : {
      "title" : "car-smile-api",
      "version" : "1.0.0"
    },
    "servers" : [{
      "url" : "https://ix6encemed.execute-api.us-east-1.amazonaws.com/{basePath}",
      "variables" : {
        "basePath" : {
          "default" : "/dev"
        }
      }
    }],
    "paths" : {
      "/hello/{name}" : {
        "get" : {
          "parameters" : [{
            "name" : "name",
            "in" : "path",
            "required" : true,
            "schema" : {
              "type" : "string"
            }
          }],
          "responses" : {
            "200" : {
              "description" : "200 response",
              "content" : {
                "application/json" : {
                  "schema" : {
                    "$ref" : "#/components/schemas/saludo"
                  }
                }
              }
            }
          },
          "x-amazon-apigateway-integration" : {
            "httpMethod" : "POST",
            "uri" : "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:225716232501:function:car-smile-mngr-go/invocations",
            "responses" : {
              "default" : {
                "statusCode" : "200"
              }
            },
            "passthroughBehavior" : "when_no_match",
            "contentHandling" : "CONVERT_TO_TEXT",
            "type" : "aws_proxy"
          }
        }
      }
    },
    "components" : {
      "schemas" : {
        "saludo" : {
          "required" : ["valor"],
          "type" : "object",
          "properties" : {
            "valor" : {
              "type" : "string"
            }
          }
        }
      }
    }
    }
  )
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
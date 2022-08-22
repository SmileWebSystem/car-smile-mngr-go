resource "aws_api_gateway_rest_api" "rest_api_car" {
  name = var.rest_api_name
  body = jsonencode({
    "openapi" : "3.0.1",
    "info" : {
      "title" : "car-smile-api",
      "version" : "1.0.0"
    },
    "servers" : [
      {
        "url" : "https://x2eyenqnl5.execute-api.us-east-1.amazonaws.com/{basePath}",
        "variables" : {
          "basePath" : {
            "default" : "/dev"
          }
        }
    }],
    "paths" : {
      "/smile/v2/car/{licensePlate}" : {
        "get" : {
          "parameters" : [
            {
              "name" : "licensePlate",
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
                    "$ref" : "#/components/schemas/CarInfo"
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
        "CarInfo" : {
          "required" : [
          "LicensePlate"],
          "type" : "object",
          "properties" : {
            "Claims" : {
              "type" : "array",
              "items" : {
                "$ref" : "#/components/schemas/CarClaim"
              }
            },
            "Guia" : {
              "$ref" : "#/components/schemas/CarGuia"
            },
            "History" : {
              "type" : "array",
              "items" : {
                "$ref" : "#/components/schemas/CarHistory"
              }
            },
            "LicensePlate" : {
              "type" : "string"
            },
            "Owners" : {
              "type" : "array",
              "items" : {
                "$ref" : "#/components/schemas/CarOwner"
              }
            },
            "Score" : {
              "$ref" : "#/components/schemas/CarScore"
            },
            "Simis" : {
              "type" : "array",
              "items" : {
                "$ref" : "#/components/schemas/CarSimi"
              }
            }
          }
        },
        "CarSimi" : {
          "type" : "object",
          "properties" : {
            "Code" : {
              "type" : "string"
            },
            "Date" : {
              "type" : "string"
            },
            "Name" : {
              "type" : "string"
            },
            "Number" : {
              "type" : "string"
            },
            "Status" : {
              "type" : "string"
            },
            "Value" : {
              "type" : "string"
            }
          }
        },
        "CarHistory" : {
          "type" : "object",
          "properties" : {
            "Active" : {
              "type" : "string"
            },
            "Beneficiary" : {
              "type" : "string"
            },
            "Chassis" : {
              "type" : "string"
            },
            "CodeCompany" : {
              "type" : "string"
            },
            "DateFin" : {
              "type" : "string"
            },
            "DateIni" : {
              "type" : "string"
            },
            "Engine" : {
              "type" : "string"
            },
            "InsuredAmount" : {
              "type" : "string"
            },
            "LicensePlate" : {
              "type" : "string"
            },
            "NameCompany" : {
              "type" : "string"
            },
            "PolicyHolder" : {
              "type" : "string"
            },
            "PolicyHolderId" : {
              "type" : "string"
            },
            "PolicyNumber" : {
              "type" : "string"
            },
            "Service" : {
              "type" : "string"
            }
          }
        },
        "CarGuia" : {
          "type" : "object",
          "properties" : {
            "ActualValue" : {
              "type" : "number",
              "format" : "float"
            },
            "Brand" : {
              "type" : "string"
            },
            "Class" : {
              "type" : "string"
            },
            "Country" : {
              "type" : "string"
            },
            "Maker" : {
              "type" : "string"
            },
            "Model" : {
              "type" : "string"
            },
            "Type" : {
              "type" : "string"
            }
          }
        },
        "CarShelter" : {
          "type" : "object",
          "properties" : {
            "AmountPaid" : {
              "type" : "string"
            },
            "Date" : {
              "type" : "string"
            },
            "Name" : {
              "type" : "string"
            },
            "ReclaimedAmount" : {
              "type" : "number",
              "format" : "float"
            },
            "Status" : {
              "type" : "string"
            }
          }
        },
        "CarOwner" : {
          "type" : "object",
          "properties" : {
            "Id" : {
              "type" : "string"
            },
            "Name" : {
              "type" : "string"
            }
          }
        },
        "CarScore" : {
          "type" : "object",
          "properties" : {
            "ScoreClaims" : {
              "type" : "number",
              "format" : "float"
            },
            "ScoreCloseSimis" : {
              "type" : "number",
              "format" : "float"
            },
            "ScoreOpenSimis" : {
              "type" : "number",
              "format" : "float"
            },
            "ScoreOwners" : {
              "type" : "number",
              "format" : "float"
            },
            "ScoreTotal" : {
              "type" : "number",
              "format" : "float"
            }
          }
        },
        "CarClaim" : {
          "type" : "object",
          "properties" : {
            "Company" : {
              "type" : "string"
            },
            "Date" : {
              "type" : "string"
            },
            "NumberClaim" : {
              "type" : "string"
            },
            "NumberPolicy" : {
              "type" : "string"
            },
            "Shelters" : {
              "type" : "array",
              "items" : {
                "$ref" : "#/components/schemas/CarShelter"
              }
            }
          }
        }
      }
    }
  })
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
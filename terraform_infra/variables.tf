# Input variable definitions

variable "aws_region" {
  description = "AWS region for all resources."

  type    = string
  default = "us-east-1"
}

variable "rest_api_name" {
  type        = string
  description = "Name of the API Gateway created"
  default     = "car-smile-api"
}

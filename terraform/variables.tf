variable "aws_region" {
  default = "us-east-1"
}

variable "instance_type" {
  default = "t2.micro"
}

variable "key_name" {
  type        = string
  default     = "andrew-projects-api"
}

variable "public_key_path" {
  type        = string
  default     = "~/.ssh/andrew-projects-api.pem"
}

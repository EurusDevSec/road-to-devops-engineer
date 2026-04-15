provider "aws" {
  region = "ap-southeast-1"
}

resource "aws_s3_bucket" "my_test_bucket" {
  bucket = "eurus-learning-tf-2304"

  tags = {
    Name        = "My First Terraform Bucket"
    Environment = "Dev"
  }
}

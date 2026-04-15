# 1. Lấy thông tin VPC mặc định của Lab
data "aws_vpc" "default" {
  default = true
}

# 2. Lấy danh sách các Subnet trong VPC đó
data "aws_subnets" "default" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }
}

# 3. Lấy AMI Amazon Linux 2
data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]
  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-x86_64-gp2"]
  }
}

# 4. Tạo 4 Server
resource "aws_instance" "my_servers" {
  count         = var.instance_count
  ami           = data.aws_ami.amazon_linux.id
  instance_type = var.instance_type

  # Chọn đại subnet đầu tiên trong danh sách tìm được
  subnet_id = data.aws_subnets.default.ids[0]

  tags = {
    Name = "Lab-Server-${count.index + 1}"
  }
}

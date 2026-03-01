locals {
  local16 = var.test16
}

variable "test18" {
  type = bool
}

variable "test17" {
  type = bool
}

terraform {
  required_version = ">=1.6.2"
}

output "output16" {
  value = local.local16
}

output "output17" {
  value = var.test17
}

output "output18" {
  value = var.test18
}

variable "test19" {
  type = bool
}

variable "test20" {
  type = bool
}

output "output20" {
  value = var.test20
}

variable "test21" {
  type = bool
}

output "output21" {
  value = var.test21
}

variable "tag" {
  type = string
}

locals {
  name = "proxy-test"
  repo = "cgetzen/proxy-test"
}

provider "kubernetes" {
  config_path = "~/.kube/config"
}

resource "kubernetes_deployment" "this" {
  metadata {
    name = local.name
  }
  spec {
    replicas = 2
    selector {
      match_labels = {
        name = local.name
      }
    }
    template {
      metadata {
        labels = {
          name = local.name
        }
      }
      spec {
        container {
          image = "${local.repo}:${var.tag}"
          name  = local.name
        }
      }
    }
  }
}

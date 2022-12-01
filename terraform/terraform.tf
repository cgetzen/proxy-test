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

          env_from {
            config_map_ref {
              name = "environment"
            }
          }
        }
      }
    }
  }
}

// Enhancement 3
resource "kubernetes_config_map" "environment" {
  metadata {
    name = "environment"
  }

  data = {
    TEMPLATE_DATA = "Hello World."
  }
}

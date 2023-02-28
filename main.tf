terraform {
  required_providers {
    order = {
      version = "~> 1.0.0"
      source  = "sirena.com/tutorial/order"
    }
  }
}


resource "order_request" "order_01" {
        style = "latte"
        size  = "large"
        place = "togo"
}

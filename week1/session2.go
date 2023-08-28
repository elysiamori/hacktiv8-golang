package main

import "fmt"

func calculateShippingCost(weight float64, destination string) float64 {
    var shippingCost float64

    if weight > 0 {
        if destination == "USA" {
            if weight <= 5 {
                shippingCost = 10.0
            } else {
                shippingCost = 15.0
            }
        } else {
            shippingCost = 20.0
        }
    } else {
        shippingCost = 0.0
    }

    return shippingCost
}

func calculateShippingCostWithoutNestedIfs(weight float64, destination string) float64 {
    var shippingCost float64

    if weight <= 0 {
        shippingCost = 0.0
        return shippingCost
    }

    if destination != "USA" {
        shippingCost = 20.0
        return shippingCost
    }

    if weight <= 5 {
        shippingCost = 10.0
        return shippingCost
    }

    shippingCost = 15.0
    return shippingCost

}

func main() {
    weight := 4.5
    destination := "USA"

    cost := calculateShippingCost(weight, destination)
    fmt.Printf("Shipping Cost: $%.2f\n", cost)
}

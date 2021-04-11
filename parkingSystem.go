type ParkingSystem struct {
    big int
    medium int
    small int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    var p ParkingSystem
    p.big = big
    p.medium = medium
    p.small = small
    return p
}


func (this *ParkingSystem) AddCar(carType int) bool {
    var res bool
    if carType == 1 {
        if this.big > 0 {
            this.big = this.big - 1
            res = true
        } else {
            res = false
        }
    }
    if carType == 2 {
        if this.medium > 0 {
            this.medium = this.medium - 1
            res = true
        } else {
            res = false
        }
    }
    if carType == 3 {
        if this.small > 0 {
            this.small = this.small - 1
            res = true
        } else {
            res = false
        }
    }
    return res
}
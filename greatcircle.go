package greatcircle

import (
    "math"
)

//  Radius of the Earth (meters)
const R = 6371000.0


// Coordinate is a datatype used to represent a latitude-longitude coordinate
type Coordinate struct {
    Latitude    float64
    Longitude   float64
}

// Convert decimal coordinate value to a radian
func radians(decimal float64) (radians float64) {
    return ( (decimal * math.Pi) / 180)
}

// Distance calculates the distance between two points (in meters) using the Great Circle method
func Distance(c0 *Coordinate, c1 *Coordinate) (distance float64) {

    theta0 := radians(c0.Latitude)
    theta1 := radians(c1.Latitude)
    dTheta := theta1 - theta0
    dLambda := radians(c1.Longitude) - radians(c0.Longitude)

    a := (math.Pow(math.Sin(dTheta / 2.0), 2.0)) + math.Cos(theta0) * math.Cos(theta1) * (math.Pow(math.Sin(dLambda / 2.0), 2.0))
    c := 2.0 * math.Atan2(math.Sqrt(a), math.Sqrt(1.0 - a))

    return R * c
}

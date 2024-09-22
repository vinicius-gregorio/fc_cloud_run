package failures

import "errors"

/// Celsius Errors

// Invalid Temperature, absolute values that could not be greater than 100
var ErrInvalidTemperatureCelsius_GreaterThan100 = errors.New("invalid temperature Celsius - Celsius could not be greater than 100")

// Invalid Temperature, absolute values that could not be less than -100
var ErrInvalidTemperatureCelsius_LessThan100 = errors.New("invalid temperature Celsius - Celsius could not be less than -100")

// Empty Temperature, the temperature could not be empty
var ErrEmptyTemperatureCelsius = errors.New("empty temperature Celsius - Celsius could not be empty")

/// Fahrenheit Errors

// Invalid Temperature, absolute values that could not be greater than 212
var ErrInvalidTemperatureFahrenheit_GreaterThan212 = errors.New("invalid temperature Fahrenheit - Fahrenheit could not be greater than 212")

// Invalid Temperature, absolute values that could not be less than -148
var ErrInvalidTemperatureFahrenheit_LessThan148 = errors.New("invalid temperature Fahrenheit - Fahrenheit could not be less than -148")

// Empty Temperature, the temperature could not be empty
var ErrEmptyTemperatureFahrenheit = errors.New("empty temperature Fahrenheit - Fahrenheit could not be empty")

/// Kelvin Errors

// Invalid Temperature, absolute values that could not be greater than 373
var ErrInvalidTemperatureKelvin_GreaterThan373 = errors.New("invalid temperature Kelvin - Kelvin could not be greater than 373")

// Invalid Temperature, absolute values that could not be less than 0
var ErrInvalidTemperatureKelvin_LessThan0 = errors.New("invalid temperature Kelvin - Kelvin could not be less than 0")

// ErrNameIsRequired is a error that represents a Name is required
var ErrNameIsRequired = errors.New("name is required")

// ErrRegionIsRequired is a error that represents a Region is required
var ErrRegionIsRequired = errors.New("region is required")

// ErrCountryIsRequired is a error that represents a Country is required
var ErrCountryIsRequired = errors.New("country is required")

// ErrLocalTimeIsRequired is a error that represents a LocalTime is required
var ErrLocalTimeIsRequired = errors.New("local time is required")

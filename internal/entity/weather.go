package entity

type Weather struct {
	Location LocationWeather `json:"location"`
	Current  CurrentWeather  `json:"current"`
}

func NewWeather(location LocationWeather, current CurrentWeather) (*Weather, error) {
	w := &Weather{
		Location: location,
		Current:  current,
	}

	if err := w.validate(); err != nil {
		return nil, err
	}

	return w, nil
}

func (w *Weather) validate() error {
	if err := w.Location.validate(); err != nil {
		return err
	}
	if err := w.Current.validate(); err != nil {
		return err
	}

	return nil
}

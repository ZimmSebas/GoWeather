package weatherapi

type WeatherAPIRequest struct {
	location WeatherAPILocation
	current  WeatherAPICurrent
}

type WeatherAPILocation struct {
	name            string
	region          string
	country         string
	lat             float32
	lon             float32
	tz_id           string
	localtime_epoch int16
	localtime       string
}

type WeatherAPICurrent struct {
}

type weatherRes struct {
	temp_min      float32
	temp_max      float32
	pressure      float32
	precipitation float32
}

func request() weatherRes {
	var res weatherRes

	return res
}

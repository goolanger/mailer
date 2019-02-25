package utils

func CleanParameterString(p *string, def string) {
	if len(*p) == 0 {
		*p = def
	}
}

func CleanParameterInt(p *int, def int) {
	if *p == 0 {
		*p = def
	}
}

func GenerateToken() string {
	return "qwlehfowiuq4hnir2qc342y34o8tu2n34ihmto35utc8924y5g9ot4u7t24y5"
}

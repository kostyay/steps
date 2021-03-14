module github.com/stackpulse/public-steps/public-steps/elastic/info

go 1.14

replace (
	github.com/stackpulse/public-steps/common v0.0.0 => ../../common
	github.com/stackpulse/public-steps/elastic/base v0.0.0 => ../base
)

require (
	github.com/caarlos0/env/v6 v6.3.0
	github.com/stackpulse/public-steps/elastic/base v0.0.0
)

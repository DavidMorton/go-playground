module example/main

go 1.21.5

replace example/greetings => ../greetings

require example/greetings v0.0.0-00010101000000-000000000000
require webserver v0.0.0-00010101000000-000000000000
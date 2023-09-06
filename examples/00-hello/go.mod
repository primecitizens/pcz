module example.com/hello

go 1.21

replace (
  github.com/primecitizens/std => ../src
)

require (
  github.com/primecitizens/std v0.0.0
)

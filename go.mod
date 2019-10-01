module github.com/tarotaro0/exp-mm-repo

go 1.13

require github.com/tarotaro0/exp-mm-repo/services/foo v1.0.0

replace github.com/tarotaro0/exp-mm-repo/services/foo v1.0.0 => ./services/foo

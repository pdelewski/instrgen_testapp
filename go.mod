module github.com/pdelewski/instrgen_testapp

require github.com/pdelewski/instrgen_test_module v0.1.0
require github.com/pdelewski/instrgen_testapp/component v0.1.0

replace github.com/pdelewski/instrgen_testapp/component => ./component

go 1.19

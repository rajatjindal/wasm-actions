#!/bin/bash

tinygo build -x -target=wasip2 --wit-package ./wit --wit-world cli -o main.wasm -x -work main.go
# Snyky - Alternative Snyk IaC implementation

## What is this?

An experiment to see what an alternative Snyk IaC experience could look like

## Usage

`$ snyky test -s <filename>` will test the current directory of files against all of the policy packs in the `packs` folder

`$ snyky test -s <filename> -p snyk,user` will only use the Snyk & User provided policies found in `packs/snyk` and `packs/user` respectively

## Installation

First make sure you have conftest installed, following the instructions [here](https://www.conftest.dev/install/)

Next run `go get github.com/benlaplanche/snyky` to instlal this onto your path

Run `$ snyky test --help` to check it's all working correctly.

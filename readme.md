# Snyky - Conftest plugin

## What is this?

An experiment to see what an alternative Snyk IaC experience could look like

## Usage

`$ conftest snyky` will test the current directory of files against all of the policy packs in the `packs` folder

`$ conftest snyky --packs=snyk,user` will only use the Snyk & User provided policies found in `packs/snyk` and `packs/user` respectively

## Installation

First make sure you have conftest installed, following the instructions [here](https://www.conftest.dev/install/)

Next install the plugin with `conftest plugin install github.com/benlaplanche/snyky`

Test it works with `$ conftest snyky --verify` and you should see a `SUCCESS` message

## TODOS

- can conftest display success messages as well as failures in the output

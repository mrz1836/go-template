# Common makefile commands & variables between projects
include .make/common.mk

# Common Golang makefile commands & variables between projects
include .make/go.mk

# Temporary makefile for initializing a new Go project template (remove this!)
include .make/temp.mk

## Set default repository details if not provided
REPO_NAME  ?= go-template
REPO_OWNER ?= mrz1836

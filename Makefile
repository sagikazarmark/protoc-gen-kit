# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

# Add the ability to override some variables
# Use with care
-include override.mk

# Main targets
include main.mk

# Add custom targets here
-include custom.mk

.PHONY: deploy invoke

# invoke served model - deprecated
invoke:
	python3 scripts/invoke.py -i $(shell minikube ip) -p 32000

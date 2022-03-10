validate-env:
ifndef VERIFICATION_TOKEN
	echo 'VERIFICATION_TOKEN を環境変数に設定して下さい。'
	exit 1
endif

.PHONY: deploy
deploy: validate-env
	gcloud functions deploy takopi --entry-point TakopiCommand --runtime go116 --set-env-vars VERIFICATION_TOKEN=$(VERIFICATION_TOKEN) --trigger-http
run_dev_frontend:
    cd frontend && npm run dev
run_dev_backend:
    export TURNIX_DEV_EXTERNTEST_SERVER="http://localhost:3000" && export TURNIX_DEV_MODE="true" && go run -v main.go
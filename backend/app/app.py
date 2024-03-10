import uvicorn
from suggest import router as router_suggest
from search import router as route_search
from metrics import router as router_metrics
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware


app = FastAPI(debug=True)


app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


app.include_router(router_metrics)
app.include_router(router_suggest)
app.include_router(route_search)


if __name__ == "__main__":
    uvicorn.run(app, host="localhost", port=8001)
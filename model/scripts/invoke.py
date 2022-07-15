"""Utility script to invoke hosted model."""
import logging
import sys

import pandas as pd
import requests
import click

logging.basicConfig(stream=sys.stdout, level=logging.INFO)

PROXIES = {"http": None, "https": None}
HEADERS = {"Content-Type": "application/json"}


def load_data(path: str) -> pd.DataFrame:
    return pd.read_csv(path, parse_dates=["date_recorded"], index_col="id")


@click.command()
@click.option("-i", "--ip", default="127.0.0.1", help="ip address of API call")
@click.option("-p", "--port", default=5000, help="port of API call")
def invoke_model(ip, port) -> None:
    """Utilty mehtod to invoke hosted model"""
    data = load_data("files/water_pump_set.csv").reset_index().sample(20)
    url = f"http://{ip}:{port}/invocations"
    # on http://platform/invocations the model can be invoked
    # url =     "http://platform/model/invocations"  # this one is redirected to the model
    logging.info(f"URL: {url}")
    logging.info(f"Headers: {HEADERS}")
    logging.info(f"Data: {data.head()}")
    r = requests.post(
        url,
        data=data.to_json(orient="split"),
        headers=HEADERS,
        verify=False,
        proxies=PROXIES,
    )
    logging.info(f"Status code: {r.status_code}")
    logging.info(f"Response text: {r.text}")

    response_df = pd.read_json(r.text, orient="records")
    logging.info(f"Response df: {response_df}")


if __name__ == "__main__":
    invoke_model()

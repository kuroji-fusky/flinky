from bs4 import BeautifulSoup
import requests

__all__ = [
    "req_wrapper",
    "req_soup"
]

rs = requests.Session()

HEADERS = {
    "User-Agent": "kurobot-python/1.0"
}


def req_wrapper(url: str,
                referer=None,
                timeout: int = 15,
                retries: int = 3):
    """
    Args:
        url (str): The URL to request from
        referer (str, optional): Adds a `Referer` header to the request
        timeout (int, optional): Timeout that bitch
        retries (int, optional): How many retries before it dies

    Returns:
        Response: lmao
    """
    retry_count = 0

    while retry_count < retries:
        try:
            return rs.get(url,
                          headers={**HEADERS, "Referer": referer},
                          timeout=timeout)

        except requests.exceptions.Timeout:
            retry_count += 1

            if retry_count == retries:
                raise

        except ConnectionError:
            raise ConnectionError("Bruh you not connected to the internet")


def req_soup(url: str):
    req = req_wrapper(url)

    if req is None:
        raise ValueError("Request failed and returned nothing, just like your soul")  # noqa

    return BeautifulSoup(req.text, "html.parser")

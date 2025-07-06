import bs4
import requests
from typing import Union

__all__ = [
    "req_wrapper",
    "req_soup",
    "SoupMe"
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
    return bs4.BeautifulSoup(req.text, "html.parser")


class SoupMe:
    @staticmethod
    def extract_links(link_el: Union[bs4.ResultSet[bs4.Tag], bs4.Tag], *, prefix):
        if isinstance(link_el, bs4.ResultSet):
            return [l.get("href") for l in link_el]
        else:
            return link_el.get("href")

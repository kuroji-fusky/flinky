import bs4
import requests  # type:ignore
from typing import Union, final
import time
from functools import wraps

__all__ = [
    "req_wrapper",
    "req_soup",
    "soup_utils"
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
        Response: The response
    """
    retry_count = 0

    while retry_count < retries:
        try:
            headers = dict(HEADERS)
            if referer is not None:
                headers["Referer"] = referer

            return rs.get(url,
                          headers=headers,
                          timeout=timeout)

        except (requests.exceptions.Timeout, ConnectionError) as e:
            if isinstance(e, requests.exceptions.Timeout):
                retry_count += 1

                if retry_count == retries:
                    raise
            else:  # ConnectionError
                raise ConnectionError("You got no internet fam")


def req_soup(url: str):
    req = req_wrapper(url)
    return bs4.BeautifulSoup(req.text, "html.parser")


FlexiTag = Union[bs4.ResultSet[bs4.Tag], bs4.Tag]

# TODO add a constructor for flexible poosay


def bs4_instance_check(func):
    def wrapper():
        func()

    return wrapper


@final
class soup_utils:
    @staticmethod
    def extract_links(link_el: FlexiTag, *, prefix=""):
        if not isinstance(link_el, (bs4.ResultSet, bs4.Tag)):
            raise TypeError(
                "Item(s) provided must be a BeautifulSoup `ResultSet` or `Tag`")

        if isinstance(link_el, bs4.ResultSet):
            if not all(isinstance(el, bs4.Tag) and el.name == "a" for el in link_el):
                raise ValueError("link_el must contain only `<a>` tags")

        elif not (isinstance(link_el, bs4.Tag) and link_el.name == "a"):
            raise ValueError("link_el must be an `<a>` tag or a type of ResultSet of `<a>` tags")  # noqa

        if isinstance(link_el, bs4.ResultSet):
            return [f'{prefix}{l.get("href")}' for l in link_el]
        else:
            return f'{prefix}{link_el.get("href")}'

    @staticmethod
    def extract_text_content(el: FlexiTag):
        if not isinstance(el, (bs4.ResultSet, bs4.Tag)):
            raise TypeError(
                "Item(s) provided must be a BeautifulSoup `ResultSet` or `Tag`")

        if isinstance(el, bs4.ResultSet):
            return [e.get_text("", True) for e in el]
        else:
            return el.get_text("", True)

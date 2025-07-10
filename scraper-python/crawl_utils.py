import bs4
import requests  # type:ignore
from typing import Union, Optional, Any, final

import json


__all__ = [
    "RequestWrapper",
    "BadgerCache",
    "soup_utils",
]

rs = requests.Session()

HEADERS = {
    "User-Agent": "kurobot-python/1.0"
}


class RequestWrapper:
    def __init__(self,
                 url: str,
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

                self.req_content = rs.get(url,
                                          headers=headers,
                                          timeout=timeout)

            except (requests.exceptions.Timeout, ConnectionError) as e:
                if isinstance(e, requests.exceptions.Timeout):
                    retry_count += 1

                    if retry_count == retries:
                        raise
                else:  # ConnectionError
                    raise ConnectionError("You got no internet fam")

    @property
    def soup(self):
        req = self.req_content
        return bs4.BeautifulSoup(req.text, "html.parser")


_CompatTag = Union[bs4.ResultSet[bs4.Tag], Optional[bs4.Tag]]


def bs4_scope_tags(func):
    def wrapper():
        func()

    return wrapper


@final
class soup_utils:
    @staticmethod
    def extract_links(link_el: _CompatTag, *, prefix=""):
        if not isinstance(link_el, (bs4.ResultSet, bs4.Tag)):
            raise TypeError(
                "Item(s) provided must be a BeautifulSoup `ResultSet` or `Tag`")

        if isinstance(link_el, bs4.ResultSet):
            if not all(isinstance(el, bs4.Tag) and el.name == "a" for el in link_el):
                raise ValueError("link_el must contain only `<a>` tags")

        elif not (isinstance(link_el, bs4.Tag) and link_el.name == "a"):
            raise ValueError("link_el must be an `<a>` tag or a type of ResultSet of `<a>` tags")  # noqa

        return [f'{prefix}{l.get("href")}'
                for l in (link_el if isinstance(link_el, bs4.ResultSet) else [link_el])]

    @staticmethod
    def extract_text_content(el: _CompatTag):
        if not isinstance(el, (bs4.ResultSet, bs4.Tag)):
            raise TypeError(
                "Item(s) provided must be a BeautifulSoup `ResultSet` or `Tag`")

        return [e.get_text(strip=True)
                for e in (el if isinstance(el, bs4.ResultSet) else [el])]


class BadgerCache:
    def __init__(self, filename: str, initial_content: Any, *, json_mode=False):
        self.filename = filename
        self.json_mode = json_mode
        self.contents = None

        if self.json_mode and not isinstance(initial_content, (dict, list)):
            raise TypeError("You can't do that shit in JSON mode lmao")

        try:
            with open(filename, "r", encoding="utf-8") as f:
                if self.json_mode:
                    self.contents = json.load(f)
                else:
                    self.contents = f.read()

        except FileNotFoundError:
            with open(filename, "w", encoding="utf-8") as f:
                if self.json_mode:
                    json.dump(initial_content, f, ensure_ascii=True, indent=2)
                else:
                    f.write(str(initial_content))
            self.contents = initial_content

    @property
    def get_dat_juicy_data(self):
        return self.contents

    def append(self, new_content: Any):
        if not self.json_mode:
            with open(self.filename, "a", encoding="utf-8") as f:
                f.write(new_content)

            self.contents += new_content
            return

        if not isinstance(self.contents, list):
            raise TypeError("JSON append only works if the file contains a list.")  # noqa

        self.contents.append(new_content)
        with open(self.filename, "w", encoding="utf-8") as f:
            json.dump(self.contents, f, ensure_ascii=True, indent=2)

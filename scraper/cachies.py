import json
from typing import Any


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

__all__ = [
    "VILLAINS_BASE_URL",
    "HEROES_BASE_URL",
    "_WHITELISTED_KEYWORDS"
]

VILLAINS_BASE_URL = "villains.fandom.com"
HEROES_BASE_URL = "hero.fandom.com"

# This is a really scuffed implementation but useful for
# tracking words from the "Appearance" sections or tracking if no infobox exists
_WHITELISTED_KEYWORDS = ["fox", "wolf", "werewolf"]

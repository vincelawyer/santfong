#!/usr/bin/env python
# -*- coding: utf-8 -*- #
from __future__ import unicode_literals

# This file is only used if you use `make publish` or
# explicitly specify it as your config file.

import os
import sys
sys.path.append(os.curdir)
from pelicanconf import *

# no trailing /
SITEURL = 'https://sunteron.com'
RELATIVE_URLS = False

FEED_ALL_ATOM = None
CATEGORY_FEED_ATOM = None
TRANSLATION_FEED_ATOM = None
AUTHOR_FEED_ATOM = None
AUTHOR_FEED_RSS = None

DELETE_OUTPUT_DIRECTORY = True

# Following items are often useful when publishing

#DISQUS_SITENAME = ""
GOOGLE_ANALYTICS = "UA-130100894-1"

# Google Custom Search
GOOGLE_CSE = '015740835703749574172:xcjpa2d0eqw'

# Google Adsense
#   Remember to modify
#   theme/templates/layout/includes/adsense.html
#   theme/templates/layout/includes/adsense_page_level.html
#GOOGLE_ADSENSE = True

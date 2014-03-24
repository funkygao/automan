#!/usr/bin/env python
#encoding=utf-8
'''Prerequirement:
sudo easy_install confluence '''

import sys
import json
from confluence import Confluence

URL='http://wiki.localhost.cn/'
USER='peng.gao'
PASS='peng.gao'
SPACE='DG'

def wikitag(line):
    js = json.loads(line)
    return '{code:JavaScript}' + json.dumps(js, indent=4) + '{code}'

def wiki_content(fn):
    n = 0
    content = 'The following content is genereated by automan\n\n'
    for line in open(fn):
        if n == 0:
            # title
            content += line
        elif n == 1:
            # input json
            content += 'Request:\n'
            content += wikitag(line) + '\n'
        elif n == 2:
            # ouptut json
            content += 'Response:\n'
            content += wikitag(line) + '\n'
        n += 1
        n %= 3

    return content

def main():
    cf = Confluence(url=URL, username=USER, password=PASS)
    ok = cf.storePageContent('test auto wiki', SPACE, wiki_content(sys.argv[1]))
    if ok:
        print 'feed wiki ok'
    else:
        print 'feed wiki fail'

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print 'Usage: %s api_filename' % sys.argv[0]
        sys.exit(0)

    main()

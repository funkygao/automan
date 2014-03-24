#!/usr/bin/env python
#encoding=utf-8
'''Prerequirement:
sudo easy_install confluence '''

import os
from confluence import Confluence

URL='http://wiki.localhost.cn/'
USER='test'
PASS='test'
SPACE='DG'

def main():
    cf = Confluence(url=URL, username=USER, password=PASS)
    ok = cf.storePageContent('test auto wiki', SPACE, 'hello world, blah')
    if ok:
        print 'feed ok'
    else:
        print 'feed fail'


if __name__ == '__main__':
    main()

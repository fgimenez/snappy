# -*- Mode: Makefile; indent-tabs-mode:t; tab-width: 4 -*-

all:
install:
	mkdir -p $(DESTDIR)/bin
	cp -a bin/get $(DESTDIR)/bin/get
	cp -a bin/set $(DESTDIR)/bin/set
	cp -a $(SNAPCRAFT_STAGE)/usr/share/glib-2.0 $(DESTDIR)/usr/share/glib-2.0
	glib-compile-schemas $(DESTDIR)/usr/share/glib-2.0/schemas
